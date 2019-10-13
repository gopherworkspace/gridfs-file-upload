package main

import (
	"github.com/gridfs-upload/backend/golang/router"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
	"log"
)

//entrypoint
func main() {

	// if changing port change in homepage to serve the ui also(optional)
	router := router.NewRouter() // create routes

	router.Methods("GET", "POST")

	serv := &http.Server {
		Handler : router,
		Addr: "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serv.ListenAndServe())
}


/*intialize MongoDB session*/
func intiDB(){

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

}
