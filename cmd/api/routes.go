package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)
	router.HandlerFunc(http.MethodPost, "/v1/records", app.createRecordHandler)
	router.HandlerFunc(http.MethodGet, "/v1/records/:id", app.showRecordHandler)
	router.HandlerFunc(http.MethodGet, "/v1/records", app.showRecordsHandler)

	return router
}
