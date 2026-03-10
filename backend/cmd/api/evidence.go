package main

import (
	"fmt"
	"net/http"

	"time"

	"ghostplanet.bumpsites.com/internal/data"
)

func (app *application) createEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new parent evidence")
}

func (app *application) showEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	evidence := data.Evidence{
		ID:              id,
		InvestigationID: 123,
		LocationID:      666,
		CreatedByUserID: 435,
		CreatedAt:       time.Now(),
		Visibility:      true,
		Version:         666,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"evidence": evidence}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
