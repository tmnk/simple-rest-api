package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmnk/simple-rest-api/internal/app/store/teststore"
)

func TestServer_handleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		expectedCode int
		payload      interface{}
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "ex@ggg.ru",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "ivalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "exggg.ru",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
	// rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	// s := newServer(teststore.New())
	// s.ServeHTTP(rec, req)
	// assert.Equal(t, rec.Code, http.StatusOK)
}
