package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rof20004/go-rest-tutorial/api/controllers"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/users", uc.ListUser)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
