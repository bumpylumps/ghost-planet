package data

import (
	"database/sql"
	// "encoding/json"
	"strings"
	"time"

	"ghostplanet.bumpsites.com/internal/validator"
)

type Evidence struct {
	ID              int64     `json:"id"`
	InvestigationID int64     `json:"investigation_id"`
	LocationID      int64     `json:"location_id"`
	CreatedByUserID int64     `json:"created_by_user_id"`
	CreatedAt       time.Time `json:"created_at"`
	Visibility      bool      `json:"visibility"`
	Version         int32     `json:"version"`
}

type TextNote struct {
	ID         int64     `json:"id"`
	EvidenceID int64     `json:"evidence_id"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	LocationID int64     `json:"location_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type AudioNote struct {
	ID            int64         `json:"id"`
	EvidenceID    int64         `json:"evidence_id"`
	Title         string        `json:"title"`
	SourceURL     string        `json:"source_url"`
	Duration      time.Duration `json:"duration"`
	FileSizeBytes int64         `json:"file_size_bytes"`
	IsEVP         bool          `json:"is_evp"`
	CreatedAt     time.Time     `json:"created_at"`
}

type Photo struct {
	ID            int64     `json:"id"`
	EvidenceID    int64     `json:"evidence_id"`
	SourceURL     string    `json:"source_url"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Caption       string    `json:"caption"`
	FileType      string    `json:"file_type"`
	FileSizeBytes int64     `json:"file_size_bytes"`
	CreatedAt     time.Time `json:"created_at"`
}

func ValidateTextNote(v *validator.Validator, textNote *TextNote) {
	v.Check(textNote.Subject != "", "subject", "must be provided")
	v.Check(len(textNote.Subject) <= 500, "subject", "must not be more than 500 bytes long")

	v.Check(textNote.Body != "", "body", "must be provided")
	v.Check(len(textNote.Body) <= 10000, "body", "must not be more than 10,000 bytes long")

	// location, err := GetLocation(textNote.LocationID)
	// v.Check(err == nil, "locationID", "location not found")
	// if err == nil {
	// 	ValidateLocation(v, location)
	// }

}

func ValidateAudioNote(v *validator.Validator, audioNote *AudioNote) {
	v.Check(audioNote.SourceURL != "", "source_url", "must be provided")
	v.Check(len(audioNote.SourceURL) <= 2000, "source_url", "must not be more than 2000 bytes long")

	const maxSize = 5 * 1024 * 1024
	v.Check(audioNote.FileSizeBytes > 0 && audioNote.FileSizeBytes <= maxSize, "size", "must be between 1 byte and 5 MB")
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
	v.Check(photo.FileSizeBytes > 0 && photo.FileSizeBytes <= maxSize, "size", "must be between 1 byte and 5 MB")

	// check that caption is a valid string, if provided
	v.Check(photo.Caption != "", "caption", "must be provided")
	v.Check(len(photo.Caption) <= 500, "caption", "must not be more than 500 bytes")
	// thumbnail is programmatically generated? Still a url

}

func ValidateEvidence(v *validator.Validator, evidence *Evidence) {

}

type EvidenceModel struct {
	DB *sql.DB
}

/*
	 TODO: write insert to handle new evidence flow:
		1) Parse JSON from client into Request struct that includes evidence slices
		2) Insert evidence into parent evidence table, return generated ID
		3) Loop through evidence slices, assign returned EvidenceID to each item and insert
		4) Commit/Rollback based on success/failure of steps
*/
func (e EvidenceModel) Insert(evidence *Evidence) error {
	return nil
}

func (e EvidenceModel) Get(id int64) (*Evidence, error) {
	return nil, nil
}

func (e EvidenceModel) Update(evidence *Evidence) error {
	return nil
}

func (e EvidenceModel) Delete(id int64) error {
	return nil
}

// testing
type MockEvidenceModel struct{}

func (e MockEvidenceModel) Insert(evidence *Evidence) error {
	return nil
}

func (e MockEvidenceModel) Get(id int64) (*Evidence, error) {
	return nil, nil
}

func (e MockEvidenceModel) Update(evidence *Evidence) error {
	return nil
}

func (e MockEvidenceModel) Delete(id int64) error {
	return nil
}
