package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	store.User().Create(u)
	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handlerFake := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handlerFake).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HadleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name        string
		payload     interface{}
		excpectCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			excpectCode: http.StatusCreated,
		},
		{
			name:        "invalid payload",
			payload:     "invalid",
			excpectCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			excpectCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.excpectCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name        string
		payload     interface{}
		excpectCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			excpectCode: http.StatusOK,
		},
		{
			name:        "invalid payload",
			payload:     "invalid",
			excpectCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			excpectCode: http.StatusUnauthorized,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			excpectCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.excpectCode, rec.Code)
		})
	}

}
