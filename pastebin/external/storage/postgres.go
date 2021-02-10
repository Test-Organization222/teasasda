package storage

import (
	"database/sql"

	"github.com/stdapps/pastebin/pastebin"
)

// Postgres is postgres implementation of storage interface
type Postgres struct {
	db *sql.DB
}

var _ pastebin.Storage = Postgres{}

// NewPostgresStorage is a custructor for Postgres
func NewPostgresStorage(db *sql.DB) Postgres {
	return Postgres{
		db: db,
	}
}

// GetPaste retrieves a Paste from database by Key
func (p Postgres) GetPaste(key string) (pastebin.Paste, error) {
	paste := pastebin.Paste{}
	err := p.db.QueryRow("SELECT id, body FROM pastes WHERE id = $1", key).Scan(&paste.Key, &paste.Body)

	return paste, err
}

// StorePaste persists a Paste to database
func (p Postgres) StorePaste(paste pastebin.Paste) error {
	_, err := p.db.Exec("INSERT INTO pastes (id, body) VALUES($1, $2);", paste.Key, paste.Body)
	return err
}
