package main

import (
	"fmt"
	"net/http"

	// "time"

	"ghostplanet.bumpsites.com/internal/data"
	"ghostplanet.bumpsites.com/internal/validator"
)

func (app *application) createEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TextNotes  []data.TextNote  `json:"text_notes,omitempty"`
		AudioNotes []data.AudioNote `json:"audio_notes,omitempty"`
		Photos     []data.Photo     `json:"photos,omitempty"`
		EVPS       []data.AudioNote `json:"evps,omitempty"`
		Visibility *bool            `json:"visibility"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	evidence := &data.Evidence{
		TextNotes:  input.TextNotes,
		AudioNotes: input.AudioNotes,
		Photos:     input.Photos,
		EVPS:       input.EVPS,
		Visibility: *input.Visibility,
	}

	v := validator.New()

	v.Check(input.Visibility != nil, "visibility", "must be provided")

	if data.ValidateEvidence(v, evidence); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Evidence.Insert(evidence)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/evidences/%d", evidence.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"evidence": evidence}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	// id, err := app.readIDParam(r)
	// if err != nil {
	// 	app.notFoundResponse(w, r)
	// 	return
	// }

	// evidence := data.Evidence{
	// 	ID:         id,
	// 	AudioNotes: []string{"screeching"},
	// 	Photos:     []string{"white lady"},
	// 	EVPS:       []string{"scary noises"},
	// 	Visibility: true,
	// }

	// err = app.writeJSON(w, http.StatusOK, envelope{"evidence": evidence}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}
