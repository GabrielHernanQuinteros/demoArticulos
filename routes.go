package main

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielHernanQuinteros/demoArticulos/other"
	"github.com/gorilla/mux"
)

//===================================================================================================
// Funciones de ROUTERS

func TraerVideogames(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxRegistros, err := TraerVideogamesSQL()

	if err == nil {
		RespondWithSuccess(auxRegistros, parWriter)
	} else {
		RespondWithError(err, parWriter)
	}

}

func TraerVideogamePorId(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := other.StringToInt64(auxIdString)

	if err != nil {
		RespondWithError(err, parWriter)
		return
	}

	auxRegistro, err := TraerVideogamePorIdSQL(auxId)

	if err != nil {
		RespondWithError(err, parWriter)
	} else {
		RespondWithSuccess(auxRegistro, parWriter)
	}

}

func CrearVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		RespondWithError(err, parWriter)
	} else {
		err := CrearVideogameSQL(auxRegistro)

		if err != nil {
			RespondWithError(err, parWriter)
		} else {
			RespondWithSuccess(true, parWriter)
		}

	}

}

func ModificarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		RespondWithError(err, parWriter)
	} else {
		err := ModificarVideogameSQL(auxRegistro)

		if err != nil {
			RespondWithError(err, parWriter)
		} else {
			RespondWithSuccess(true, parWriter)
		}

	}

}

func BorrarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := other.StringToInt64(auxIdString)

	if err != nil {
		RespondWithError(err, parWriter)
		return
	}

	err = BorrarVideogameSQL(auxId)

	if err != nil {
		RespondWithError(err, parWriter)
	} else {
		RespondWithSuccess(true, parWriter)
	}

}

//===================================================================================================
// Funciones de CORS

func EnableCORS(parRouter *mux.Router) {

	parRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	parRouter.Use(MiddlewareCors)

}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})

}

//===================================================================================================
// Funciones de respuesta

func RespondWithError(parError error, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(parWriter).Encode(parError.Error())

}

func RespondWithSuccess(parDato interface{}, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(parWriter).Encode(parDato)

}
