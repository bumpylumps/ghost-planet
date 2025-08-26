package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// api routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/investigations", app.createInvestigationHandler)
	router.HandlerFunc(http.MethodGet, "/v1/investigations/:id", app.showInvestigationHandler)

	//return that dang ol httprouter instance
	return router
}
