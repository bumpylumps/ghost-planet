package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createInvestigationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new investigation")
}

func (app *application) showInvestigationHandler(w http.ResponseWriter, r *http.Request) {

	// slice request parameter names/values
	params := httprouter.ParamsFromContext(r.Context())

	//if id param can't convert, err
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of Paranormal Investigation %d\n", id)
}
