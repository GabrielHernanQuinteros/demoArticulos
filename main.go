package main

import (
	"log"
	"net/http"
	"time"

	myvars "github.com/GabrielHernanQuinteros/demoArticulos/vars"
	mytools "github.com/GabrielHernanQuinteros/prueba/video"
	myroute "github.com/GabrielHernanQuinteros/prueba/routes"
	"github.com/gorilla/mux"
)

func main() {

	mytools.Hola()

	auxBaseDatos, err := mytools.ConectarDB(myvars.ConnectionString) // Ping database

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

	mytools.EnableCORS(auxRouter)

	auxRouter.HandleFunc("/videogames", TraerRegistros).Methods(http.MethodGet)          //Modificar
	auxRouter.HandleFunc("/videogames/{id}", TraerRegistroPorId).Methods(http.MethodGet) //Modificar
	auxRouter.HandleFunc("/videogames", CrearRegistro).Methods(http.MethodPost)          //Modificar
	auxRouter.HandleFunc("/videogames", ModificarRegistro).Methods(http.MethodPut)       //Modificar
	auxRouter.HandleFunc("/videogames/{id}", BorrarRegistro).Methods(http.MethodDelete)  //Modificar

	//===================================================================================================
	// Setup and start server

	server := &http.Server{
		Handler: auxRouter,
		Addr:    myvars.Port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", myvars.Port)
	log.Fatal(server.ListenAndServe())
}
