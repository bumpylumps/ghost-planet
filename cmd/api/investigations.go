package main

import (
	"fmt"
	"net/http"
	"time"

	"ghostplanet.bumpsites.com/internal/data"
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

	investigation := app.mockInvestigation(id)

	err = app.writeJSON(w, http.StatusOK, investigation, nil)
	if err != nil {
		app.logger.Println(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) mockInvestigation(id int64) data.Investigation {
	return data.Investigation{
		ID: id,
		User: data.User{
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
		Location: data.Location{
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
