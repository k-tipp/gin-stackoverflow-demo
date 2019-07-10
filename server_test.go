package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testBody = `{
	"resourceType": "Organization",
	"id": "f001",
	"name": "Burgers University Medical Center",
	"telecom": [
	  {
		"system": "phone",
		"value": "022-655 2300",
		"use": "work"
	  }
	],
	"partOf": {
	  "reference": "Organization/f001"
	},
	"endpoint": [
	  {
		"reference": "Endpoint/example"
	  }
	]
  }`

func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Successful
func TestSuccessEmptyJsonBody(t *testing.T) {
	log.Println("Running TestSuccessEmptyJsonBody...")

	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/json")

	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

// Successful
func TestSuccessJsonBody(t *testing.T) {
	log.Println("Running TestSuccessJsonBody...")
	// Grab our router
	router := SetupRouter()
	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "application/json")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestEmptyBodyApplicationJavascript(t *testing.T) {
	log.Println("Running TestEmptyBodyJavascript...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/javascript")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestEmptyBodyTextPlain(t *testing.T) {
	log.Println("Running TestEmptyBodyTextPlain...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "text/plain")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

// Successful
func TestEmptyBodyApplicationXml(t *testing.T) {
	log.Println("Running TestEmptyBodyApplicationXml...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/xml")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

// Successful
func TestEmptyBodyTextXml(t *testing.T) {
	log.Println("Running TestEmptyBodyTextXml...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "text/xml")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestEmptyBodyTextHTML(t *testing.T) {
	log.Println("Running TestEmptyBodyTextHTML...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(""))
	req.Header.Add("Content-Type", "text/html")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestBodyApplicationJavascript(t *testing.T) {
	log.Println("Running TestBodyJavascript...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "application/javascript")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestBodyTextPlain(t *testing.T) {
	log.Println("Running TestBodyTextPlain...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "text/plain")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

// Successful
func TestBodyApplicationXml(t *testing.T) {
	log.Println("Running TestBodyApplicationXml...")
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "application/xml")
	w := performRequest(router, req)

	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

// Successful
func TestBodyTextXml(t *testing.T) {
	log.Println("Running TestBodyTextXml...")
	// Grab our router
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "text/xml")
	w := performRequest(router, req)

	// the request gives a 200
	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}

func TestBodyTextHTML(t *testing.T) {
	log.Println("Running TestBodyTextHTML...")
	// Grab our router
	router := SetupRouter()

	req, _ := http.NewRequest(http.MethodPost, "/Organization", strings.NewReader(testBody))
	req.Header.Add("Content-Type", "text/html")
	w := performRequest(router, req)

	// the request gives a 200
	if w.Code != http.StatusOK {
		t.Errorf("HTTP status was %v", w.Code)
	}
}
