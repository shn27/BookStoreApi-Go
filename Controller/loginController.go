package Controller

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"time"
)

var secretKey []byte

func Login(res http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(60 * time.Second),
		Path:    "/",
	}
	cookie.Value = Init()

	http.SetCookie(res, &cookie)

	res.Write([]byte("Cookie set!"))
	fmt.Fprintf(res, "Successfully Logged In ")

	// Write a response
}

var tokenAuth *jwtauth.JWTAuth

func Init() string {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	return tokenString
}
