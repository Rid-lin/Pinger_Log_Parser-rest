package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"net/http"
)

const (
	sessionName        = "plpsessions"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect E-mail or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST") // Create Users
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	s.router.HandleFunc("/getdevices", s.handleGetDevices()).Methods("GET")       // Get All Devices
	s.router.HandleFunc("/updatedevices", s.handleUpdateDevices()).Methods("PUT") // Run check status all Devices

	s.router.HandleFunc("/editdevice", s.handleGetDevice()).Methods("GET")       // Get information on one device
	s.router.HandleFunc("/editdevice", s.handleCreateDevice()).Methods("POST")   // Create Device and him's info
	s.router.HandleFunc("/editdevice", s.handleUpdateDevice()).Methods("PUT")    // Update info one Device
	s.router.HandleFunc("/editdevice", s.handleDeleteDevice()).Methods("DELETE") // Delete Device

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start))
	})
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(tos.ServersList)
		// TODO тут нужно отдать index.html как статичную страницу, в которой будет логика по дальнейшей работе приложения
	}
}

//handleGetDevices Return list devices as JSON and/or Error
func (s *server) handleGetDevices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var devices [](*model.Device)
		devices, err := s.store.Device().GetAllAsList() // FIXME исправить GetAllAsList() на GetStatusOfAllDevices()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, devices) // FIXME Выводить не просто список всех устройств, а список со статусами
	}
}

func (s *server) handleUpdateDevices() http.HandlerFunc {
	// опрашивает все devices и и обновляет информацию о доступности (статус)
	devicesList, _ := s.store.Device().GetAllAsList()
	for _, device := range devicesList {
		device.CheckNLogStatus(getFullPatchFile(LogPatch))
	}
	return s.handleGetDevices()
}

func (s *server) handleGetDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO Написать функцию которая отдаёт информацию о конкретном устройстве в JSON

		// var devices [](*model.Device)
		// devices, err := s.store.Device().GetAllAsList()
		// if err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// s.respond(w, r, http.StatusOK, devices)
	}
}

func (s *server) handleCreateDevice() http.HandlerFunc {
	type request struct {
		ID          int    `json:"id"`
		IP          string `json:"ip"`
		Place       string `json:"place"`
		Description string `json:"description"`
		MethodCheck string `json:"methodcheck"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		d := &model.Device{
			ID:          req.ID,
			IP:          req.IP,
			Place:       req.Place,
			Description: req.Description,
			MethodCheck: req.MethodCheck,
		}

		if err := s.store.Device().Create(d); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		d.CheckNLogStatus(getFullPatchFile(LogPatch))

		s.respond(w, r, http.StatusCreated, d)
	}

	// return func(w http.ResponseWriter, r *http.Request) {
	// 	var device (*model.Device)
	// 	// TODO Написать функцию которая создаёт устройство при получении информации в JSON

	// 	// var devices [](*model.Device)
	// 	// devices, err := s.store.Device().GetAllAsList()
	// 	// if err != nil {
	// 	// 	s.error(w, r, http.StatusInternalServerError, err)
	// 	// 	return
	// 	// }
	// 	// s.respond(w, r, http.StatusOK, devices)
	// 	getFullPatchFile(LogPatch)
	// 	device.CheckNLogStatus(LogPatch)
	// }
}

func (s *server) handleUpdateDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var device (*model.Device)
		// TODO Написать функцию которая обновляет информацию об устройстве при получении информации в JSON

		// var devices [](*model.Device)
		// devices, err := s.store.Device().GetAllAsList()
		// if err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// s.respond(w, r, http.StatusOK, devices)
		getFullPatchFile(LogPatch)
		device.CheckNLogStatus(LogPatch)
	}
}

func (s *server) handleDeleteDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO Написать функцию которая удаляет устройство при получении информации в JSON

		// var devices [](*model.Device)
		// devices, err := s.store.Device().GetAllAsList()
		// if err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// s.respond(w, r, http.StatusOK, devices)
	}
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func getFullPatchFile(logPatch string) string {
	dateNow := time.Now().Format("2006_01_02")
	filename := fmt.Sprintf("%s/%s.csv", logPatch, dateNow)
	// fmt.Printf("%s", filename) //DEBUG
	return filename
}
