package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type VideoGame struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Year  int64  `json:"year"`
}

//===================================================================================================
// Funciones de ROUTERS

func fnTraerVideogames(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxRegistros, err := fnTraerVideogamesSQL()

	if err == nil {
		fnRespondWithSuccess(auxRegistros, parWriter)
	} else {
		fnRespondWithError(err, parWriter)
	}

}

func fnTraerVideogamePorId(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := fnStringToInt64(auxIdString)

	if err != nil {
		fnRespondWithError(err, parWriter)
		return
	}

	auxRegistro, err := fnTraerVideogamePorIdSQL(auxId)

	if err != nil {
		fnRespondWithError(err, parWriter)
	} else {
		fnRespondWithSuccess(auxRegistro, parWriter)
	}

}

func fnCrearVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		fnRespondWithError(err, parWriter)
	} else {
		err := fnCrearVideogameSQL(auxRegistro)

		if err != nil {
			fnRespondWithError(err, parWriter)
		} else {
			fnRespondWithSuccess(true, parWriter)
		}

	}

}

func fnModificarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro VideoGame
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		fnRespondWithError(err, parWriter)
	} else {
		err := fnModificarVideogameSQL(auxRegistro)

		if err != nil {
			fnRespondWithError(err, parWriter)
		} else {
			fnRespondWithSuccess(true, parWriter)
		}

	}

}

func fnBorrarVideogame(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := fnStringToInt64(auxIdString)

	if err != nil {
		fnRespondWithError(err, parWriter)
		return
	}

	err = fnBorrarVideogameSQL(auxId)

	if err != nil {
		fnRespondWithError(err, parWriter)
	} else {
		fnRespondWithSuccess(true, parWriter)
	}

}

//===================================================================================================
// Funciones de CORS

func fnEnableCORS(parRouter *mux.Router) {

	parRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	parRouter.Use(fnMiddlewareCors)

}

func fnMiddlewareCors(next http.Handler) http.Handler {
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

func fnRespondWithError(parError error, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(parWriter).Encode(parError.Error())

}

func fnRespondWithSuccess(parDato interface{}, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(parWriter).Encode(parDato)

}
