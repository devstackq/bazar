package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/devstackq/bazar/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//tok, err := ExtractTokenMetadata(c.Request)
		token, err := VerifyToken(c.Request, "accessx")//check accesToken
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
				Status: "Info",
				Message: "access token expired or incorrect",
				Data: nil,
			})
			return
		}
		//tokenValid - ok; compare Uuid - db; setContext(user_id);  - else -> send 401; /refresh 
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok  && token.Valid {
		// check uuid Db
		// val, ok := redis.get(claims["access_uuid"])

		c.Set("user_id",  claims["user_id"].(float64))//set context user_id
		c.Next()

		} else {
			// if access_uuid = equal(clinet access_uuid) -> update- refresh/access
			// acs_uuid := claims["access_uuid"]
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
					Status: "Info",
					Message: "access token expired or incorrect",
					Data: nil,
				})
			}
		}
}
}
func AuthorizeRefreshJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
	if 	token, err := VerifyToken(c.Request, "refreshx"); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
			Status: "Info",
			Message: "refreshx token expired or incorrect1",
			Data: nil,
		})
		return
	}else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("user_id",  claims["user_id"].(float64))//set context user_id
			c.Next()
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
				Status: "Info",
				Message: "refreshx token expired or incorrect",
				Data: nil,
			})
		}
	}
}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request, secretKey string) (*jwt.Token, error) {

	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := VerifyToken(r, "accessx")
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// log.Println(claims, "cla")

	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func ValidateToken(encodedToken, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
