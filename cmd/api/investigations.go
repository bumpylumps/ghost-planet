package main

import (
	"fmt"
	"net/http"
)

func (app *application) createInvestigationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new investigation")
}

func (app *application) showInvestigationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of Paranormal Investigation %d\n", id)
}
