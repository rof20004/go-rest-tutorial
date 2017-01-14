package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/julienschmidt/httprouter"
	"github.com/rof20004/go-rest-tutorial/api/controllers"
	"gopkg.in/mgo.v2"
)

// Database connection
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}

// AuthRequest middleware
func AuthRequest(handleFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if token == nil {
			http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
			return
		}

		if token.Valid {
			handleFunc(w, r, p)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				http.Error(w, "Error: Malformed auth token.", http.StatusBadRequest)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				w.Header().Add("WWW-Authenticate", "Bearer")
				http.Error(w, "Error: Auth token expired.", http.StatusUnauthorized)
			} else {
				http.Error(w, "Error: "+http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}
			return
		} else {
			http.Error(w, "Error: "+http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
	}
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Security controller
	sc := controllers.NewSecurityController()
	r.GET("/get-token", sc.GetTokenHandler)

	// Websocket controller
	wsc := controllers.NewSocketController()
	r.GET("/enviar", wsc.Socket)
	r.GET("/enviarReceber", wsc.Echo)

	// User controller
	uc := controllers.NewUserController(getSession())
	r.GET("/users", AuthRequest(uc.ListUsers))
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	log.Fatal(http.ListenAndServe(":3000", r))
}
