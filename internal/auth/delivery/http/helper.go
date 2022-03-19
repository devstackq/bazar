package v1

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/devstackq/bazar/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func responseWithStatus(c *gin.Context, status int, message, text string, data interface{}) {
	c.AbortWithStatusJSON(status, models.Response{
		Status:  text,
		Message: message,
		Data:    data,
	})
}

type AccessDetails struct {
	AccessUuid string
	UserId   int64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
	SubTimeRefresh time.Duration
	SubTimeAccess time.Duration
}

func CreateToken(userID int, accessSecret, refreshSecret string)( *TokenDetails,  error) {
	td := &TokenDetails{}
	var err error
	
	td.AtExpires = time.Now().Add(time.Minute * 20).Unix() // set mlsec token 20 min
	td.AccessUuid = uuid.New().String() //set access uuid
	
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix() //set refreshTokenExpires - 7day
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(userID)// generated uuid (separator) userID

	//set jwt data; key:value
	atClaims := jwt.MapClaims{}
	atClaims["authorization"]= true
	atClaims["access_token"]=td.AccessToken
	atClaims["user_id"]=userID
	atClaims["expired"]=td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// token := jwt.New(jwt.SigningMethodHS256)

	log.Println(td.AccessToken, "before signed", accessSecret)
	td.AccessToken, err = at.SignedString([]byte(accessSecret)) //currentUuid + accessSecret
	if err != nil {
		return nil, err
	}
	log.Println(td.AccessToken, "after signed")

	//set refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_token"]=td.RefreshToken
	rtClaims["user_id"]=userID
	rtClaims["expired"]=td.RtExpires
	
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret)) //update refreshToken, + secretRefresh

	td.SubTimeAccess=  time.Unix(td.AtExpires, 0).Sub(time.Now()) //converting Unix to UTC(to Time object)
	// td.SubTimeRefresh=  time.Unix(td.RtExpires, 0).Sub(time.Now()) //converting Unix to UTC(to Time object)

	log.Println(td.SubTimeAccess, "remove token from sec")
	go func (t time.Duration) {
		time.AfterFunc(t, func() { 
			log.Println("remove access token from db, and update token")
		 })
	}(td.SubTimeAccess)

	return td, nil

}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//signin -> createToken/setExpire -> save db; - return client;
//middleware -> /handler -> checkTokenFromClient;
//refreshToken -> when: timeExpire;
//removeToken - logout || expire;