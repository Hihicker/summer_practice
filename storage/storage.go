package storage

import "github.com/jmoiron/sqlx"

type Storage struct {
	Dict DictionaryStorage
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Dict: NewDictionaryStorage(db),
	}
}
