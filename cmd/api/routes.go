package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/investigations", app.createInvestigationHandler)
	router.HandlerFunc(http.MethodGet, "/v1/investigations/:id", app.showInvestigationHandler)

	router.HandlerFunc(http.MethodPost, "/v1/evidence", app.createEvidenceHandler)
	router.HandlerFunc(http.MethodGet, "/v1/evidence/:id", app.showEvidenceHandler)

	router.HandlerFunc(http.MethodPost, "/v1/locations", app.createLocationHandler)
	router.HandlerFunc(http.MethodGet, "/v1/locations/:id", app.showLocationHandler)

	//return that dang ol httprouter instance
	return router
}
