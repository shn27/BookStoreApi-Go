package Middleware

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"strings"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := r.Header.Get("Authorization")
		if ai == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Get the token from the header.
		eAuthToken := strings.Split(r.Header.Get("Authorization"), " ")
		uAuthInfo, err := base64.StdEncoding.DecodeString(eAuthToken[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		cx := bytes.Index(uAuthInfo, []byte(":"))
		username := string(uAuthInfo[:cx])
		password := string(uAuthInfo[cx+1:])

		if username != "admin" || password != "1234" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var tokenAuth *jwtauth.JWTAuth

func Init() string {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	return tokenString
}

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("jwt")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Error(res, "cookie not found.May be expires", http.StatusBadRequest)
			default:
				log.Println(err)
				http.Error(res, "server error", http.StatusInternalServerError)
			}
			return
		}

		if Init() != cookie.Value {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}

		next.ServeHTTP(res, req)
	})
}
