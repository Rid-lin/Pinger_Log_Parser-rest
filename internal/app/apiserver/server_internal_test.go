package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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
			name: "invalid params - password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			excpectCode: http.StatusUnauthorized,
		},
		{
			name: "invalid params - Email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
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

func TestServer_HandleGetDevices(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	t.Run("GetDevices", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/getdevices", nil)
		s.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

}

func TestServer_HandleGetDevice(t *testing.T) {
	store := teststore.New()

	d := model.TestDevice(t)
	store.Device().Create(d)
	srtID := strconv.Itoa(d.ID)

	d2 := model.TestDevice2(t)
	store.Device().Create(d2)

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name        string
		payload     interface{}
		excpectCode int
	}{
		{
			name: "valid - two parameter",
			payload: map[string]string{
				"ip": d.IP,
				"id": srtID,
			},
			excpectCode: http.StatusOK,
		},
		{
			name: "valid - two parameter with diferent device",
			payload: map[string]string{
				"ip": d2.IP,
				"id": srtID,
			},
			excpectCode: http.StatusOK,
		},
		{
			name: "valid - only ip",
			payload: map[string]string{
				"ip": d.IP,
			},
			excpectCode: http.StatusOK,
		},
		{
			name: "valid - only id",
			payload: map[string]string{
				"id": srtID,
			},
			excpectCode: http.StatusOK,
		},
		{
			name:        "invalid params - empty headers",
			payload:     map[string]string{},
			excpectCode: http.StatusBadRequest,
		},
		{
			name:        "invalid params - invalid headers",
			payload:     map[string]string{"invalid": "invalid"},
			excpectCode: http.StatusBadRequest,
		},
		{
			name: "invalid params - not existed ID",
			payload: map[string]string{
				"id": "99999999999999999999"},
			excpectCode: http.StatusBadRequest,
		},
		{
			name: "invalid params - invalid IP",
			payload: map[string]string{
				"ip": "invalid"},
			excpectCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid params - invalid ID",
			payload: map[string]string{
				"id": "invalid"},
			excpectCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			headers := make(map[string]string)
			headers = tc.payload.(map[string]string)
			req, _ := http.NewRequest(http.MethodGet, "/editdevice", nil)
			for key, val := range headers {
				req.Header.Set(key, val)
			}
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.excpectCode, rec.Code)
		})
	}

}

func TestServer_HandleUpdateDevices(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	t.Run("UpdateDevices", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/updatedevices", nil)
		s.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

}

func TestServer_HandleCreateDevice(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name        string
		payload     interface{}
		excpectCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"ip":          "127.0.0.1",
				"place":       "ЧНГКМ",
				"description": "Описание",
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
			req, _ := http.NewRequest(http.MethodPost, "/editdevice", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.excpectCode, rec.Code)
		})
	}
}

// func TestServer_HandleDeleteDevice(t *testing.T) {
// 	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
// 	d := model.TestDevice(t)

// 	testCases := []struct {
// 		name        string
// 		payload     interface{}
// 		excpectCode int
// 	}{
// 		{
// 			name: "valid",
// 			payload: func() map[string]string {
// 				return map[string]string{
// 					"id":          string(d.ID),
// 					"ip":          d.IP,
// 					"place":       d.Place,
// 					"description": d.Description,
// 					"methodcheck": d.MethodCheck,
// 				}
// 			},
// 			excpectCode: http.StatusOK,
// 		},
// 		{
// 			name:        "invalid payload",
// 			payload:     "invalid",
// 			excpectCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "invalid params",
// 			payload: map[string]string{
// 				"ip":          "invalid",
// 				"place":       "invalid",
// 				"description": "invalid",
// 				"methodcheck": "invalid",
// 			},
// 			excpectCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			_ = s.store.Device().Create(d)
// 			rec := httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tc.payload)
// 			fmt.Printf("%v\n", string(b.Bytes()))
// 			req, _ := http.NewRequest(http.MethodDelete, "/editdevice", b)
// 			s.ServeHTTP(rec, req)
// 			assert.Equal(t, tc.excpectCode, rec.Code)
// 		})
// 	}
// }
