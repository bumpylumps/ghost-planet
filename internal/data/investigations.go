package data

import (
	"time"
)

type Investigation struct {
	ID               int64      `json:"id"`
	User             User       `json:"user"`
	Location         Location   `json:"location"`
	Phenomena        string     `json:"phenomena"`
	CreatedAt        time.Time  `json:"created_at"`
	Evidence         []Evidence `json:"evidence"`
	EmergencyContact string     `json:"emergency_contact"`
	Visibility       bool       `json:"visibility"` // Public or Private Investigation
}

type User struct {
	ID                 int64           `json:"id"`
	CreatedAt          time.Time       `json:"created_at"`
	Status             string          `json:"status"`
	ProfilePage        string          `json:"profile_page"`
	Firstname          string          `json:"firstname"`
	Lastname           string          `json:"lastname"`
	Username           string          `json:"username"`
	Investigations     []Investigation `json:"investigations"`
	Evidence           []Evidence      `json:"evidence"`
	PrivateLocations   []Location      `json:"private_locations"`
	CommunityLocations []Location      `json:"community_locations"` // locations that have been contributed to the Public locations list
}

type Location struct {
	ID                          int64           `json:"id"`
	Name                        string          `json:"name"`
	Address                     string          `json:"address"`
	Lore                        string          `json:"lore"`                 // TODO figure out structure for lore
	LatLong                     []string        `json:"lat_long_coordinates"` // [Lattitude, Longitude]
	PastInvestigationsUser      []Investigation `json:"past_investigations_user"`
	PastInvestigationsCommunity []Investigation `json:"past_investigations_community"`
	Popularity                  Popularity      `json:"popularity"` // customize to add "stars" for now
	Visibility                  bool            `json:"visibility"` // Public/Private Location
}

type Evidence struct {
	ID         int64    `json:"id"`
	TextNotes  []string `json:"text_notes"`  // TODO: Flesh these out with their own types
	AudioNotes []string `json:"audio_notes"` // slice of audio urls
	Photos     []string `json:"photos"`      // slice of photo urls
	EVPS       []string `json:"evps"`        // slice of audio urls
	Visibility bool
}
