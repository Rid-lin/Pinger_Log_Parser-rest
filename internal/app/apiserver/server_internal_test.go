package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HadleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
