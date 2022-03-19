package middleware

import (
	"log"
	"net/http"
)

func validJwtToken(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("jwt_token")
		if err != nil {
			log.Println(err, "err jwt token")
			return
		}
		f.ServeHTTP(w, r)
	}
}
