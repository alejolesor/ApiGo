package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	Session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return Session
}

var collection = getSession().DB("curso-go")
var comicsdb = collection.C("ComicsMarvel")
var peliculas = collection.C("movies")

//var collectionComics = getSession().DB("curso-go").C("Comics")

//Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to server")
}

//movieList Listado de peliculas ...
func movieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := peliculas.Find(nil).Sort("-_id").All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Listado de Peliculas ", results)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

//movieShow ...
func movieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	if !bson.IsObjectIdHex(movieID) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(404)

		Error := Errores{
			Errorctrl{1, "El Id no Existe en la DB"},
		}
		// er := Errorctrl{}
		// er.idPelicula = 1
		// er.Description = "pruebas2"
		json.NewEncoder(w).Encode(Error)
		return
	}
	oid := bson.ObjectIdHex(movieID)

	results := Movie{}
	err := peliculas.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

}

func movieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = peliculas.Insert(movieData)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movieData)
}

//getComics ...
func getComics(w http.ResponseWriter, r *http.Request) {
	var methodComics = "public/comics?"
	var responsgeneral ResponseGeneral
	//Services consumo de Api
	response := Services(methodComics)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.Unmarshal(databyte, &responsgeneral)
	if err1 != nil {
		fmt.Printf("Error decodificando: %v\n", err1)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responsgeneral)
	}
}

//getComicsDetails ...
func getComicsDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	comicID := params["id"]
	var methodComics = "public/comics/" + comicID + "?"
	var responsgeneral ResponseGeneral
	//Services consumo de Api
	response := Services(methodComics)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.Unmarshal(databyte, &responsgeneral)
	if err1 != nil {
		fmt.Printf("Error decodificando: %v\n", err1)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responsgeneral)
	}
}

//InsertComics ...
func InsertComics(w http.ResponseWriter, r *http.Request) {

	var methodComics = "public/comics?"
	var responsgeneral ResponseGeneral
	//Services consumo de Api
	response := Services(methodComics)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.Unmarshal(databyte, &responsgeneral)
	if err1 != nil {
		fmt.Printf("Error decodificando: %v\n", err1)
	} else {
		fmt.Println(responsgeneral.Data.Results)
		err = comicsdb.Insert(responsgeneral.Data)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Insert OK")
	}
}

//getComicsDb Lista de Comics en la base de datos
func getComicsDb(w http.ResponseWriter, r *http.Request) {
	var comics []Data
	err := comicsdb.Find(nil).Sort("-_id").All(&comics)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Listado de Peliculas en la Tabla de MongoDB ", comics)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comics)
}

func getComicxName(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//title := params["title"]
	var comicsfilter []Data

	//err := comicsdb.Find(nil).Sort("-_id").All(&comicsfilter)
	err := comicsdb.Find(nil).Sort("-_id").All(&comicsfilter)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("Listado de Comics", comicsfilter)

	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comicsfilter)

}
