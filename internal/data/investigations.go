package data

import (
	"database/sql"
	"math"
	"strings"
	"time"

	"ghostplanet.bumpsites.com/internal/validator"

	"github.com/lib/pq"
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

func GetLocation(locationID int64) (*Location, error) {
	//db location lookup
	return &Location{}, nil
}

type Evidence struct {
	ID              int64       `json:"id"`
	TextNotes       []TextNote  `json:"text_notes"`
	AudioNotes      []AudioNote `json:"audio_notes"`
	Photos          []Photo     `json:"photos"`
	EVPS            []AudioNote `json:"evps"`
	Visibility      bool        `json:"visibility"`
	CreatedByUserID int64       `json:"created_by_user_id"`
	CreatedAt       time.Time   `json:"created_at"`
	Version         int64       `json:"version"`
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
	Size      int64     `json:"size"`
}

type Photo struct {
	ID           int64  `json:"id"`
	SourceURL    string `json:"source_url"`
	FileType     string `json:"fileType"`
	Size         int64  `json:"size"`
	Caption      string `json:"caption"`
	ThumbnailURL string `json:"thumbnail"`
}

// TODO: test ValidateAudioNote, ValidateTextNote, ValidatePhoto
func ValidateTextNote(v *validator.Validator, textNote *TextNote) {
	v.Check(textNote.Subject != "", "subject", "must be provided")
	v.Check(len(textNote.Subject) <= 500, "subject", "must not be more than 500 bytes long")

	v.Check(textNote.Body != "", "body", "must be provided")
	v.Check(len(textNote.Body) <= 10000, "body", "must not be more than 10,000 bytes long")

	location, err := GetLocation(textNote.LocationID)
	v.Check(err == nil, "locationID", "location not found")
	if err == nil {
		ValidateLocation(v, location)
	}

}

func ValidateAudioNote(v *validator.Validator, audioNote *AudioNote) {
	v.Check(audioNote.SourceURL != "", "source_url", "must be provided")
	v.Check(len(audioNote.SourceURL) <= 2000, "source_url", "must not be more than 2000 bytes long")

	const maxSize = 5 * 1024 * 1024
	v.Check(audioNote.Size > 0 && audioNote.Size <= maxSize, "size", "must be between 1 byte and 5 MB")
	// add size check
	// the rest of the info is programattically generated
}

func ValidatePhoto(v *validator.Validator, photo *Photo) {
	// check that sourceURL is valid string
	v.Check(photo.SourceURL != "", "source_url", "must be provided")
	v.Check(len(photo.SourceURL) <= 2000, "source_url", "must not be more than 2000 bytes long")

	allowedTypes := map[string]bool{
		"jpeg": true, "jpg": true, "png": true, "webp": true,
	}

	v.Check(allowedTypes[strings.ToLower(photo.FileType)], "fileType", "must be a valid image type")

	const maxSize = 5 * 1024 * 1024
	v.Check(photo.Size > 0 && photo.Size <= maxSize, "size", "must be between 1 byte and 5 MB")

	// check that caption is a valid string, if provided
	v.Check(photo.Caption != "", "caption", "must be provided")
	v.Check(len(photo.Caption) <= 500, "caption", "must not be more than 500 bytes")
	// thumbnail is programmatically generated? Still a url

}

func ValidateEvidence(v *validator.Validator, evidence *Evidence) {
	hasEvidence := len(evidence.TextNotes) > 0 || len(evidence.AudioNotes) > 0 || len(evidence.Photos) > 0 || len(evidence.EVPS) > 0
	v.Check(hasEvidence, "evidence", "must contain at least one note, photo, or EVP")

	for _, note := range evidence.TextNotes {
		ValidateTextNote(v, &note)
	}

	for _, note := range evidence.AudioNotes {
		ValidateAudioNote(v, &note)
	}

	for _, photo := range evidence.Photos {
		ValidatePhoto(v, &photo)
	}

	for _, evp := range evidence.EVPS {
		ValidateAudioNote(v, &evp)
	}
}

type EvidenceModel struct {
	DB *sql.DB
}

func (e EvidenceModel) Insert(evidence *Evidence) error {
	query := `
		INSERT INTO evidence (visibility, text_notes, audio_notes, photos, evps)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	args := []interface{}{evidence.Visibility, pq.Array(evidence.TextNotes), pq.Array(evidence.AudioNotes), pq.Array(evidence.Photos), pq.Array(evidence.EVPS)}

	return e.DB.QueryRow(query, args...).Scan(&evidence.ID, &evidence.CreatedAt, &evidence.Version)
}

func (e EvidenceModel) Get(id int64) (*Evidence, error) {
	return &Evidence{
		ID:              id,
		TextNotes:       []TextNote{},
		AudioNotes:      []AudioNote{},
		Photos:          []Photo{},
		EVPS:            []AudioNote{},
		Visibility:      true,
		CreatedByUserID: 1,
	}, nil
}

func (e EvidenceModel) Update(evidence *Evidence) error {
	return nil
}

func (e EvidenceModel) Delete(id int64) error {
	return nil
}

// Testing
type MockEvidenceModel struct{}

func (e MockEvidenceModel) Insert(evidence *Evidence) error {
	evidence.ID = 1
	evidence.CreatedAt = time.Now()
	evidence.Version = 1
	return nil
}

func (e MockEvidenceModel) Get(id int64) (*Evidence, error) {
	return &Evidence{}, nil
}

func (e MockEvidenceModel) Update(evidence *Evidence) error {
	return nil
}

func (e MockEvidenceModel) Delete(id int64) error {
	return nil
}
