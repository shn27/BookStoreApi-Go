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
		Expires: time.Now().Add(10 * time.Minute),
		Path:    "/",
	}
	cookie.Value = Jwt()
	http.SetCookie(res, &cookie)

	res.Write([]byte("Cookie set!\n"))
	fmt.Fprintf(res, "Successfully Logged In ")
	// Write a response
}

var tokenAuth *jwtauth.JWTAuth
var Secret string

func Jwt() string {
	tokenAuth = jwtauth.New("HS256", []byte(Secret), nil) // replace with secret key
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"admin": 1234})
	return tokenString
}
