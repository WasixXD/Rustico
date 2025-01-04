package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestRoot(t *testing.T) {
	expected := "Hello, World!"
	res, err := http.Get("http://localhost:3001/")

	if err != nil {
		t.Fatal("Error on making GET request: ", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal("Error on reading the body: ", err)
	}

	if res.StatusCode != 200 || string(body) != expected {
		t.Fatalf("StatusCode Expected: %d Got %d\nBody Expected: %s Got %s", http.StatusOK, res.StatusCode, expected, string(body))
	}
}

func TestCreate(t *testing.T) {

	payload := []byte(`{
		"project_name": "Project One",
		"description": "This is a good Description"	
	}`)

	expected := http.StatusCreated

	res, err := http.Post("http://localhost:3001/create", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal("Error on making POST request: ", err)
	}

	defer res.Body.Close()

	if err != nil {
		t.Fatal("Error on reading response: ", err)
	}

	if res.StatusCode != expected {
		log.Fatalf("StatusCode expected: %d Got: %d", expected, res.StatusCode)
	}

}

func TestMain(m *testing.M) {
	go ServerInit()
	m.Run()

	os.Exit(1)
}
