package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMainFunc(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Main panicked ??")
		}
	}()

	go main()
	time.Sleep(1 * time.Second)
}

func TestPanic(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/panic/", nil)
	if err != nil {
		t.Fatalf("Unable to request %v", err)
	}
	ResponcerRecord := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	panicHandler(ResponcerRecord, req)
	response := ResponcerRecord.Result()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("No error in /panic/ %v", response.StatusCode)
	}
}

func TestRecoveryHandle(t *testing.T) {
	panichandler := http.HandlerFunc(panicHandler)
	makeRequest("Get", "/panic/", recoveryHandle(panichandler))
}

func makeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	Request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	ResponcerRecord := httptest.NewRecorder()
	ResponcerRecord.Result()
	handler.ServeHTTP(ResponcerRecord, Request)
	return ResponcerRecord, err
}

func makeTestHandler(t *testing.T) http.HandlerFunc {
	fun := func(w http.ResponseWriter, r *http.Request) {
		log.Print("I should not execute")
	}
	return http.HandlerFunc(fun)
}

func TestDebugHandler(t *testing.T) {
	testserver := httptest.NewServer(getHandles())
	defer testserver.Close()

	makeNewRequest := func(method string, url string, body io.Reader) *http.Request {
		request, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Errorf("Unable to request %v", err)
		}
		return request
	}
	testSuit := []struct {
		Name       string
		request    *http.Request
		Statuscode int
	}{
		// TODO: Add test cases.
		{Name: "Testcase 1", request: makeNewRequest("GET", testserver.URL+"/debug/?path=/usr/local/go/src/balaji-dongare/gophercises/recover/main.go&line=72", nil), Statuscode: 200},
		{Name: "Testcase 2", request: makeNewRequest("GET", testserver.URL+"/debug/?path=/usr/local/go/src/balaji-dongare/gophercises/recover/main.go&line=", nil), Statuscode: 500},
	}
	for _, testcase := range testSuit {
		t.Run(testcase.Name, func(t *testing.T) {
			response, err := http.DefaultClient.Do(testcase.request)
			if err != nil {
				t.Fatal(err)
			}
			defer response.Body.Close()
			if response.StatusCode != testcase.Statuscode {
				t.Error("Error in /debug/")
			}
		})
	}
}
