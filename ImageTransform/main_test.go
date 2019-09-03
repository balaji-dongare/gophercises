package main

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestModifyAPI(t *testing.T) {
	testserver := httptest.NewServer(getHandlers())
	defer testserver.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: test get", r: newreq("GET", testserver.URL+"/modify/img/input.png?mode=2", nil), status: 200},
		{name: "2: test get", r: newreq("GET", testserver.URL+"/modify/img/input.png?mode=2&number=100", nil), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Error("error in response")
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestSingleModeAPI(t *testing.T) {
	testserver := httptest.NewServer(getHandlers())
	defer testserver.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1:/localhost:3000", r: newreq("GET", testserver.URL+"/", nil), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error")
			}
		})
	}
}

func TestUploadAPI(t *testing.T) {
	testserver := httptest.NewServer(getHandlers())
	defer testserver.Close()

	request := func(method, url string) *http.Request {
		file, err := os.Open("./img/input.png")
		if err != nil {
			t.Error(err)
		}
		body := &bytes.Buffer{}
		writter := multipart.NewWriter(body)
		part, err := writter.CreateFormFile("image", file.Name())
		if err != nil {
			t.Error("error in copy")
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("error in copy")
		}
		err = writter.Close()
		if err != nil {
			t.Error("error in close writer")
		}
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Set("Content-Type", writter.FormDataContentType())
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: /Upload", r: request("POST", testserver.URL+"/upload"), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}
func TestErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	err := errors.New("error")
	errorResponse(w, err)
}

func TestUploadAPIError(t *testing.T) {
	testserver := httptest.NewServer(getHandlers())
	defer testserver.Close()

	request := func(method, url string) *http.Request {
		file, err := os.Open("./img/input.png")
		if err != nil {
			t.Error(err)
		}
		body := &bytes.Buffer{}
		writter := multipart.NewWriter(body)
		part, err := writter.CreateFormFile("ima", file.Name())
		if err != nil {
			t.Error("error in copy")
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("error in copy")
		}
		err = writter.Close()
		if err != nil {
			t.Error("error in close writer")
		}
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Set("Content-Type", writter.FormDataContentType())
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: /Upload", r: request("POST", testserver.URL+"/upload"), status: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestUploadAPIError1(t *testing.T) {
	testserver := httptest.NewServer(getHandlers())
	defer testserver.Close()

	request := func(method, url string) *http.Request {
		file, err := os.Open("./img/input.png")
		if err != nil {
			t.Error(err)
		}
		body := &bytes.Buffer{}
		writter := multipart.NewWriter(body)
		part, err := writter.CreateFormFile("image", "input.pngg")
		if err != nil {
			t.Error("error in copy")
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("error in copy")
		 }
		// err = writter.Close()
		// if err != nil {
		// 	t.Error("error in close writer")
		// }
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Set("Content-Type", writter.FormDataContentType())
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: /Upload", r: request("POST", testserver.URL+"/upload"), status: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}
