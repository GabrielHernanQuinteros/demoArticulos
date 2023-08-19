package main

import (
	"log"
	"net/http"
	"time"

	mytools "github.com/GabrielHernanQuinteros/prueba/video"
	"github.com/gorilla/mux"
)

type VideoGame struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int64  `json:"year"`
}

func main() {

	mytools.Hola()

	// Ping database
	//auxBaseDatos, err := ConectarDB()
	auxBaseDatos, err := mytools.ConectarDB(ConnectionString)

	if err != nil {

		log.Printf("Error con la base de datos" + err.Error())
		return

	} else {

		err = auxBaseDatos.Ping()

		if err != nil {
			log.Printf("Error estableciendo la conexión con la base de datos. Por favor verifique sus credenciales. El error es: " + err.Error())
			return
		}
	}

	//===================================================================================================
	// Define routes

	auxRouter := mux.NewRouter()

	//setupRoutesForVideoGames(router)

	mytools.EnableCORS(auxRouter)

	auxRouter.HandleFunc("/videogames", TraerVideogames).Methods(http.MethodGet)
	auxRouter.HandleFunc("/videogames/{id}", TraerVideogamePorId).Methods(http.MethodGet)
	auxRouter.HandleFunc("/videogames", CrearVideogame).Methods(http.MethodPost)
	auxRouter.HandleFunc("/videogames", ModificarVideogame).Methods(http.MethodPut)
	auxRouter.HandleFunc("/videogames/{id}", BorrarVideogame).Methods(http.MethodDelete)

	//===================================================================================================
	// Setup and start server

	port := ":8000"

	server := &http.Server{
		Handler: auxRouter,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}
