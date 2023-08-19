package main

import (
	"encoding/json"
	"net/http"

	mytools "github.com/GabrielHernanQuinteros/prueba/video"
	"github.com/gorilla/mux"
)

//===================================================================================================
// Funciones de ROUTERS

func TraerRegistros(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxRegistros, err := TraerRegistrosSQL()

	if err == nil {
		mytools.RespondWithSuccess(auxRegistros, parWriter)
	} else {
		mytools.RespondWithError(err, parWriter)
	}

}

func TraerRegistroPorId(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := mytools.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	auxRegistro, err := TraerRegistroPorIdSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(auxRegistro, parWriter)
	}

}

func CrearRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro EstrucReg
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := CrearRegistroSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func ModificarRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro EstrucReg
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := ModificarRegistroSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func BorrarRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := mytools.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	err = BorrarRegistroSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(true, parWriter)
	}

}
