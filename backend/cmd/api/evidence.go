package main

import (
	"net/http"

	"ghostplanet.bumpsites.com/internal/data"
	// "time"
)

func (app *application) createEvidenceHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Define the INPUT struct (The "Filter")
	// Use pointers (*bool, *int64) for mandatory fields to distinguish
	// between a "zero value" and a missing field.
	// keep the input struct here instead of in data/evidence to keep data layer clean
	// var input struct {
	// Add InvestigationID and LocationID here (mandatory in DB)
	// Add slices (TextNotes, Photos, etc.)
	// Add Visibility (*bool) and CreatedByUserId
	// }
	var input struct {
		InvestigationID *int64           `json:"investigation_id"`
		LocationID      *int64           `json:"location_id"`
		TextNotes       []data.TextNote  `json:"text_notes"`
		AudioNotes      []data.AudioNote `json:"audio_notes"`
		Photos          []data.Photo     `json:"photos"`
		Visibility      *bool            `json:"visibility"`
		CreatedByUserID *int64           `json:"created_by_user_id"`
	}

	// 2. Decode the incoming JSON into the input struct
	// Use app.readJSON. Handle errors with app.badRequestResponse.
	return nil
	// 3. Start Validation
	// Initialize your validator.

	// 4. Perform "Existence" checks (The Gatekeeper)
	// Use v.Check to ensure mandatory pointers (like Visibility) are not nil.

	// 5. Check "Gatekeeper" validation status
	// If !v.Valid(), call app.failedValidationResponse and RETURN.
	// (This prevents the server from panicking in the next step!)

	// 6. Map Input to Data Model (The "Gold Standard")
	// Now that you know pointers aren't nil, dereference them (*)
	// and copy values into a &data.Evidence{} struct.

	// 7. Perform "Business Logic" validation
	// Use your data.ValidateEvidence(v, evidence) function.
	// If !v.Valid(), return the validation response.

	// 8. Execute the Database Insert
	// Call app.models.Evidence.Insert(evidence).
	// Note: This is where you'll eventually handle the transaction for the slices!

	// 9. Send Success Response
	// Set the "Location" header for the new resource.
	// Use app.writeJSON to send a 201 Created status and the evidence data.
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
