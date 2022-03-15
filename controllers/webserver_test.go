package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestrootPage(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	res := httptest.NewRecorder()

	rootPage(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}

	if res.Body.String() != "Welcome to the Go Api Server" {
		t.Errorf("invalid response: %#v", res)
	}

	t.Logf("%#v", res)
}
