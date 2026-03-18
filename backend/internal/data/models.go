package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Evidence interface {
		Get(id int64) (*Evidence, error)
		Update(evidence *Evidence) error
		Delete(id int64) error
		FullSync(evidence *Evidence /*, audios []AudioNote, texts []TextNote, photos []Photo*/) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Evidence: EvidenceModel{DB: db},
	}
}

// testing
func NewMockModels() Models {
	return Models{
		Evidence: MockEvidenceModel{},
	}
}
