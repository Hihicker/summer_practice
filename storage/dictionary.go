package storage

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jmoiron/sqlx"
)

type DictionaryStorage struct {
	db *sqlx.DB
}

func NewDictionaryStorage(db *sqlx.DB) DictionaryStorage {
	return DictionaryStorage{
		db: db,
	}
}

func (d *DictionaryStorage) GetCompleteWord(word *string) ([]string, error) {
	var words []string
	decoded, err := url.QueryUnescape(*word)
	if err != nil {
		return words, err
	}
	query := fmt.Sprintf("SELECT rw.word as words FROM russian_words as rw WHERE rw.word LIKE '%s%%' UNION SELECT ow.word as words FROM osetian_words as ow WHERE ow.word LIKE '%s%%'", decoded, decoded)
	log.Print(query)
	err = d.db.Select(&words, query)
	if err != nil {
		return words, err
	}

	return words, nil
}
