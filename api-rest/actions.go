package main

import (
	"encoding/json"
	"fmt"
	"mux"
	"net/http"

	"mgo"
)

//conexion a mongo
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

var collection = getSession().DB("go_movies").C("movies")

var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Juan Antonio"},
}

//Respuesta a la ruta /
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

//Respuesta a la ruta /peliculas
func MovieList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Listado de peliculas")
	//Responder en formato json
	json.NewEncoder(w).Encode(movies)
}

//Respuesta a la ruta /pelicula{id}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	//Metodo para obtener las variables
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

//Respuesta a la ruta /pelicula
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	//log.Println(movie_data)

	//Mongo
	err := collection.Insert(movie_data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(movie_data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	//movies = append(movies, movie_data)
}
