package main

import (
	"fmt"
	"net/http"
	"time"

	"ghostplanet.bumpsites.com/internal/data"
	"ghostplanet.bumpsites.com/internal/validator"
)

func (app *application) createInvestigationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new investigation")
	// TODO: make this once location and evidence is created and stored in DB
}

func (app *application) showInvestigationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	investigation := app.mockInvestigation(id) // data.Investigation for non tests when ready

	err = app.writeJSON(w, http.StatusOK, envelope{"investigation": investigation}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) mockInvestigation(id int64) data.Investigation {
	return data.Investigation{
		ID: id,
		User: &data.User{
			ID:                 1,
			CreatedAt:          time.Now().AddDate(0, -6, 0), // 6 months ago
			Status:             "Active",
			ProfilePage:        "https://example.com/profile/testuser",
			Firstname:          "John",
			Lastname:           "Doe",
			Username:           "ghosthunter123",
			Investigations:     []data.Investigation{}, // empty to avoid circular reference
			Evidence:           []data.Evidence{},
			PrivateLocations:   []data.Location{},
			CommunityLocations: []data.Location{},
		},
		Location: &data.Location{
			ID:                          1,
			Name:                        "Haunted Manor",
			Address:                     "123 spooky st, Albuquerque, NM, 08712",
			Lore:                        "Local legend says a woman in white roams the halls every full moon",
			LatLong:                     []string{"40.7128", "-74.0060"},
			PastInvestigationsUser:      []data.Investigation{},
			PastInvestigationsCommunity: []data.Investigation{},
			Popularity:                  42,
			Visibility:                  true,
		},
		Evidence: []data.Evidence{
			{
				ID:         1,
				TextNotes:  []string{"Cold spot near staircase", "Strange whispers heard at 2:30 AM"},
				AudioNotes: []string{"https://example.com/audio/session1.mp3"},
				Photos:     []string{"https://example.com/photos/orb1.jpg", "https://example.com/photos/shadow.jpg"},
				EVPS:       []string{"https://example.com/evp/voice1.mp3"},
				Visibility: true,
			},
		},
	}
}

// TODO: figure out how to actually validate this info
// create structure for text notes
func (app *application) createEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TextNotes  []string `json:"textnotes,omitempty"`
		AudioNotes []string `json:"audionotes,omitempty"`
		Photos     []string `json:"photos,omitempty"`
		EVPS       []string `json:"evps,omitempty"`
		Visibility *bool    `json:"visibility"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Visibility != nil, "visibility", "must be provided")

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	evidence := data.Evidence{
		ID:         id,
		TextNotes:  []string{"very dark place"},
		AudioNotes: []string{"screeching"},
		Photos:     []string{"white lady"},
		EVPS:       []string{"scary noises"},
		Visibility: true,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"evidence": evidence}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name                        string               `json:"name"`
		Address                     string               `json:"address"`
		Lore                        string               `json:"lore"`
		LatLong                     []string             `json:"lat_long_coordinates"`
		PastInvestigationsUser      []data.Investigation `json:"past_investigations_user,omitempty"`
		PastInvestigationsCommunity []data.Investigation `json:"past_investigations_community,omitempty"`
		Popularity                  data.Popularity      `json:"popularity,omitempty"`
		Visibility                  bool                 `json:"visibility"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Name != "", "name", "must be provided")
	v.Check(validator.Unique(input.LatLong), "lat_long_coordinates", "coordinates can not be identical")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showLocationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	location := data.Location{
		ID:                          id,
		Name:                        "Dork Knight",
		Address:                     "123 Apple lane",
		Lore:                        "Ghost",
		LatLong:                     []string{"1235", "-93949"},
		PastInvestigationsUser:      []data.Investigation{},
		PastInvestigationsCommunity: []data.Investigation{},
		Popularity:                  32,
		Visibility:                  true,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"location": location}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
