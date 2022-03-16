package controllers

import (
	// "encoding/json"
	// "fmt"
	// "time"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootPage(t *testing.T) {
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

func TestFetchAllItems(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "http://localhost:8080/items", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fetchAllItems)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Check the response body is what we ecpect
	//I don't know how to test the return value of NewRecorder
	// It was unclear if the *bytes.Buffer type is just a string, a dictionary type, or if the value can be retrieved using an address.
}
