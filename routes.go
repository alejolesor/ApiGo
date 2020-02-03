package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route ...
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

//Routes ...
type Routes []Route

//newRouter ...
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"movieList",
		"GET",
		"/peliculas",
		movieList,
	},
	Route{
		"movieShow",
		"GET",
		"/pelicula/{id}",
		movieShow,
	},
	Route{
		"movieAdd",
		"POST",
		"/peliculaAdd",
	    movieAdd,
	},
	Route{
		"getComics",
		"GET",
		"/comics",
		getComics,
	},
	Route{
		"getComicsDetails",
		"GET",
		"/comicsDetails/{id}",
		getComicsDetails,
	},
	Route{
		"InsertComics",
		"GET",
		"/insertComics",
		InsertComics,
	},
	Route{
		"getComicsDb",
		"GET",
		"/Comicsdb",
		getComicsDb,
	},
	Route{
		"getComicxName",
		"GET",
		"/ComicsName/{title}",
		getComicxName,
	},

	
}
