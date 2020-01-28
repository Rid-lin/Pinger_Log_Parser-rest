package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store/sqlitestore"
	"github.com/gorilla/sessions"
)

//LogPatch - patch of logs checks status devices
var LogPatch string

// Start ...
func Start(config *Config) error {
	LogPatch = config.LogPatch
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlitestore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	go srv.periodicCheck(config.TimeoutCheck)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
