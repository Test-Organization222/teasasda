package storage

import (
	"database/sql"
	"net/url"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stdapps/pastebin/pastebin/specs"
	"github.com/stretchr/testify/require"
)

func TestPostgres(t *testing.T) {
	connstr := GetTestDbConnStr()
	db, err := sql.Open("postgres", connstr)
	require.Nil(t, err)
	postgresStorage := NewPostgresStorage(db)

	specs.TestStorage(t, postgresStorage)
}

func GetTestDbConnStr() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("pastebin", "pastebin"),
		Host:     "localhost:5444",
		Path:     "pastebin",
		RawQuery: "sslmode=disable",
	}
	connStr := u.String()
	return connStr
}
