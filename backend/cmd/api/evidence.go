package main

import (
	"errors"
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
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Evidence.FullSync(evidence)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/evidence/%d", evidence.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"evidence": evidence}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	evidence, err := app.models.Evidence.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"evidence": evidence}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	evidence, err := app.models.Evidence.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		InvestigationID int64 `json:"investigation_id"`
		LocationID      int64 `json:"location_id"`
		CreatedByUserID int64 `json:"created_by_user_id"`
		Visibility      *bool `json:"visibility"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	evidence.InvestigationID = input.InvestigationID
	evidence.LocationID = input.LocationID
	evidence.CreatedByUserID = input.CreatedByUserID
	evidence.Visibility = input.Visibility

	v := validator.New()

	if data.ValidateEvidence(v, evidence); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Evidence.Update(evidence)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"evidence": evidence}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Evidence.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "evidence successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
