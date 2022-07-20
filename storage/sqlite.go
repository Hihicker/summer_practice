package storage
import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConn(dbPath string) (*sqlx.DB, error) {
	if _, err := os.Stat(dbPath); err != nil {
		return nil, errors.New("the database on the route `"+dbPath+"` does not exist")
	}

	conn, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}