package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Evidences interface {
		Insert(*Evidence) error
		Get(int64) (*Evidence, error)
		Update(*Evidence) error
		Delete(int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Evidences: EvidenceModel{DB: db},
	}
}

// testing
func NewMockModels() Models {
	return Models{
		Evidences: MockEvidenceModel{},
	}
}
