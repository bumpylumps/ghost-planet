package data

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"ghostplanet.bumpsites.com/internal/validator"
)

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
	Title     string    `json:"title"`
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

	textNotesJSON, err := json.Marshal(evidence.TextNotes)
	if err != nil {
		return err
	}

	audioNotesJSON, err := json.Marshal(evidence.AudioNotes)
	if err != nil {
		return err
	}

	photosJSON, err := json.Marshal(evidence.Photos)
	if err != nil {
		return err
	}

	evpsJSON, err := json.Marshal(evidence.EVPS)
	if err != nil {
		return err
	}

	args := []interface{}{
		evidence.Visibility,
		textNotesJSON,
		audioNotesJSON,
		photosJSON,
		evpsJSON,
	}

	return e.DB.QueryRow(query, args...).Scan(&evidence.ID, &evidence.CreatedAt, &evidence.Version)
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
