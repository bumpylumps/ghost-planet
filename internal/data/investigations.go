package data

import (
	"math"
	"time"

	"ghostplanet.bumpsites.com/internal/validator"
)

type Investigation struct {
	ID               int64      `json:"id"`
	UserID           int64      `json:"user,omitempty"`
	LocationID       int64      `json:"location,omitempty"`
	Phenomena        string     `json:"phenomena"`
	CreatedAt        time.Time  `json:"created_at"` // hidden input
	Evidence         []Evidence `json:"evidence,omitempty"`
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

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Firstname != "", "firstname", "must be provided")
	v.Check(len(user.Firstname) <= 500, "firstname", "must not be more than 500 bytes long")

	v.Check(user.Lastname != "", "lastname", "must be provided")
	v.Check(len(user.Lastname) <= 500, "lastname", "must not be more than 500 bytes")

	v.Check(user.Username != "", "username", "must be provided")
	v.Check(len(user.Username) <= 500, "username", "must not be more than 500 bytes long")

	v.Check(len(user.Status) <= 500, "status", "must not be more than 500 bytes long")
}

type Location struct {
	ID                          int64           `json:"id"`
	Name                        string          `json:"name"`
	Address                     string          `json:"address"`
	State                       string          `json:"state"` // dropdown options
	City                        string          `json:"city"`
	Zip                         string          `json:"zip"`
	Lore                        string          `json:"lore"` // TODO figure out structure for lore
	Latitude                    float64         `json:"latitude"`
	Longitude                   float64         `json:"longitude"`
	PastInvestigationsUser      []Investigation `json:"past_investigations_user"`
	PastInvestigationsCommunity []Investigation `json:"past_investigations_community"`
	Popularity                  Popularity      `json:"popularity"` // customize to add "stars" for now
	Visibility                  bool            `json:"visibility"` // Public/Private Location
	CreatedByUserID             int64           `json:"created_by_user_id"`
	OwnerUserID                 int64           `json:"owner_user_id"`
}

func ValidateLocation(v *validator.Validator, location *Location) {
	v.Check(location.Name != "", "name", "must be provided")
	v.Check(len(location.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(location.Address != "", "address", "must be provided")
	// TODO: check address for valid address format
	// split string into parts
	// check that first part is a number
	// check that second part is a string
	// check that last part is a string?

	// check that state is valid string
	// check that city is valid string
	// check that zip is valid numbers

	v.Check(location.Lore != "", "lore", "must be provided")
	v.Check(len(location.Lore) <= 500, "lore", "must not be more than 500 bytes long")

	// needs check for data type - generic malformed data error for when log/lat is not a number
	v.Check(!math.IsNaN(location.Latitude), "latitude", "must be a valid number")
	v.Check(!math.IsInf(location.Latitude, 0), "latitude", "must be a finite number")
	v.Check(location.Latitude > -90 && location.Latitude < 90, "latitude", "must be between -90 and 90")

	v.Check(!math.IsNaN(location.Longitude), "longitude", "must be a valid number")
	v.Check(!math.IsInf(location.Longitude, 0), "longitude", "must be a finite number")
	v.Check(location.Longitude > -180 && location.Longitude < 180, "longitude", "must be between -180 and 180")
}

type Evidence struct {
	ID              int64       `json:"id"`
	TextNotes       []TextNote  `json:"text_notes"`
	AudioNotes      []AudioNote `json:"audio_notes"`
	Photos          []Photo     `json:"photos"`
	EVPS            []AudioNote `json:"evps"`
	Visibility      bool        `json:"visibility"`
	CreatedByUserID int64       `json:"created_by_user_id"`
}

type TextNote struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Subject    string    `json:"subject"`
	LocationID int64     `json:"locationid"`
	Body       string    `json:"body"`
}

type AudioNote struct {
	ID        int64     `json:"id"`
	SourceURL string    `json:"source_url"`
	CreatedAt time.Time `json:"created_at"`
	Length    string    `json:"length"`
	Size      string    `json:"size"`
}

type Photo struct {
	ID        int64  `json:"id"`
	SourceURL string `json:"source_url"`
	FileType  string `json:"fileType"`
	Size      string `json:"size"`
	Caption   string `json:"caption"`
	Thumbnail string `json:"thumbnail"`
}

// TODO: Validator for Evidences
