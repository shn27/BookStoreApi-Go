package Controller

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"os"
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

var TokenAuth *jwtauth.JWTAuth

func Jwt() string {
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{os.Getenv("NAME"): os.Getenv("PASSWORD")})
	return tokenString
}
