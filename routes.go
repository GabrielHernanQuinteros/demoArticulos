package main

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielHernanQuinteros/demoArticulos/other"
	"github.com/gorilla/mux"

	mytools "github.com/GabrielHernanQuinteros/prueba/video"
)

//===================================================================================================
// Funciones de ROUTERS

func TraerVideogames(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxRegistros, err := TraerVideogamesSQL()

	if err == nil {
		mytools.RespondWithSuccess(auxRegistros, parWriter)
	} else {
		mytools.RespondWithError(err, parWriter)
	}

}

func TraerVideogamePorId(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := other.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	auxRegistro, err := TraerVideogamePorIdSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(auxRegistro, parWriter)
	}

}

func CrearVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := CrearVideogameSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func ModificarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := ModificarVideogameSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func BorrarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := other.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	err = BorrarVideogameSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(true, parWriter)
	}

}
