package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// Set up a global string for our secret
var mySigningKey = []byte("secret")

// SecurityController main
type SecurityController struct{}

// NewSecurityController represents the controller for operating on the JWT resource
func NewSecurityController() *SecurityController {
	return &SecurityController{}
}

// GetTokenHandler generate a new token
func (sc *SecurityController) GetTokenHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set token claims
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign the token with our secret
	tokenString, _ := token.SignedString(mySigningKey)

	// Finally, write the token to the browser window
	w.Write([]byte(tokenString))
}
