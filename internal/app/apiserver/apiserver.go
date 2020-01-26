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
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlitestore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)
	LogPatch = config.LogPatch

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

// func (s *APIServer) configureLogger() error {
// 	level, err := logrus.ParseLevel(s.config.LogLevel)
// 	if err != nil {
// 		return err
// 	}
// 	s.logger.SetLevel(level)
// 	return nil
// }

// func (s *APIServer) configureRouter() {
// 	s.router.HandleFunc("/hello", s.handleHello())
// }

// func (s *APIServer) configureStore() error {
// 	st := store.New(s.config.Store)
// 	if err := st.Open(); err != nil {
// 		return err
// 	}

// 	s.store = st

// 	return nil
// }

// func (s *APIServer) handleHello() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello")
// 	}
// }
