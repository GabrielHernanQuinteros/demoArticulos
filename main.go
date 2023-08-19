package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/GabrielHernanQuinteros/demoArticulos/video"
	
)

type VideoGame struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int64  `json:"year"`
}

func main() {

	// Ping database
	auxBaseDatos, err := fnConectarDB()

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

	fnEnableCORS(auxRouter)

	auxRouter.HandleFunc("/videogames", fnTraerVideogames).Methods(http.MethodGet)
	auxRouter.HandleFunc("/videogames/{id}", fnTraerVideogamePorId).Methods(http.MethodGet)
	auxRouter.HandleFunc("/videogames", fnCrearVideogame).Methods(http.MethodPost)
	auxRouter.HandleFunc("/videogames", fnModificarVideogame).Methods(http.MethodPut)
	auxRouter.HandleFunc("/videogames/{id}", fnBorrarVideogame).Methods(http.MethodDelete)

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
