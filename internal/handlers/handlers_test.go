package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileHandler_MissingFileParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/file", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FileHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	expected := `{"code":400,"message":"File parameter is missing"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestCustomDataHandler_MissingKeyParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CustomDataHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	expected := `{"code":400,"message":"Key parameter is missing"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestCustomDataPostHandler_InvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/data-post", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CustomDataPostHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	expected := `{"code":405,"message":"Invalid request method"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

