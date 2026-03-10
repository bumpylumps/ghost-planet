package main

import (
	"fmt"
	"net/http"

	"time"

	"ghostplanet.bumpsites.com/internal/data"
	"ghostplanet.bumpsites.com/internal/validator"
)

func (app *application) createEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID              int64     `json:"id"`
		InvestigationID int64     `json:"investigation_id"`
		LocationID      int64     `json:"location_id"`
		CreatedByUserID int64     `json:"created_by_user_id"`
		CreatedAt       time.Time `json:"created_at"`
		Visibility      *bool     `json:"visibility"`
		Version         int32     `json:"version"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	evidence := &data.Evidence{
		ID:              input.ID,
		InvestigationID: input.InvestigationID,
		LocationID:      input.LocationID,
		CreatedByUserID: input.CreatedByUserID,
		CreatedAt:       input.CreatedAt,
		Visibility:      input.Visibility,
		Version:         input.Version,
	}

	v := validator.New()

	if data.ValidateEvidence(v, evidence); !v.Valid() {

	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	TestVisibility := true // need memory address for pointer bool Visibility
	evidence := envelope{"evidence": data.Evidence{
		ID:              id,
		InvestigationID: 123,
		LocationID:      666,
		CreatedByUserID: 435,
		CreatedAt:       time.Now(),
		Visibility:      &TestVisibility,
		Version:         666,
	}}

	err = app.writeJSON(w, http.StatusOK, evidence, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
