package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gridfs-upload/backend/golang/controller"
)

//var controller = &Controller{Repository: repository.Repository{}}
var controller = &controllers.Controller{}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{"FileUpload", "POST", "/upload", controller.Upload},
	Route{"FileDownload", "POST", "/download", controller.Download},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//serving curl as requirment
	//router.HandleFunc("/", controller.GetNewJoke)
	//adding a ui
	//router.HandleFunc("/", Home)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
