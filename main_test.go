package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/saiprasadkrishnamurthy/web-api/handlers"
)

func Test_Echo(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handlers.EchoHandler(res, req, nil)
	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}
