package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var payload []byte = []byte(`{
		"project_name": "Project One",
		"description": "This is a good Description"	
}`)

func makeGetRequest(url string, t *testing.T) (*http.Response, []byte) {
	res, err := http.Get(url)

	if err != nil {
		t.Fatal("Couldn't make GET request: ", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal("Couldn't parse body")
	}

	return res, body
}

func TestRoot(t *testing.T) {
	expected := "Hello, World!"
	res, body := makeGetRequest("http://localhost:3001/", t)

	if res.StatusCode != 200 || string(body) != expected {
		t.Fatalf("StatusCode Expected: %d Got %d\nBody Expected: %s Got %s", http.StatusOK, res.StatusCode, expected, string(body))
	}
}

func TestCreate(t *testing.T) {

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

func TestRandomProject(t *testing.T) {

	_, body := makeGetRequest("http://localhost:3001/random", t)

	var project ProjectResponse
	err := json.Unmarshal(body, &project)

	if err != nil {
		t.Fatalf("Expected %+v GOT: %v", project, err)
	}
}

func TestAll(t *testing.T) {

	res, body := makeGetRequest("http://localhost:3001/all", t)

	if res.StatusCode != 200 {
		t.Fatalf("StatusCode Expected %d, GOT: %d", http.StatusOK, res.StatusCode)
	}

	var projects []ProjectResponse
	err := json.Unmarshal(body, &projects)

	if err != nil {
		t.Fatalf("Expected %+v GOT: %v", projects, err)
	}
}

func TestDelete(t *testing.T) {

	c1 := make(chan bool)
	go func(c chan bool) {
		res, _ := makeGetRequest("http://localhost:3001/delete", t)

		if res.StatusCode == http.StatusBadRequest {
			c <- true
			return
		}
		c <- false

	}(c1)

	c2 := make(chan bool)
	go func(c chan bool) {
		res, _ := makeGetRequest("http://localhost:3001/delete?id=abc", t)

		if res.StatusCode == http.StatusBadRequest {
			c <- true
			return
		}
		c <- false

	}(c2)

	c3 := make(chan bool)
	go func(c chan bool) {
		_, body := makeGetRequest("http://localhost:3001/delete?id=1000000", t)
		expected := "0\n"
		parsed := string(body)

		if expected == parsed {
			c <- true
			return
		}
		c <- false

	}(c3)

	c4 := make(chan bool)
	go func(c chan bool) {
		_, body := makeGetRequest("http://localhost:3001/delete?id=1", t)
		expected := "1\n"
		parsed := string(body)

		if expected == parsed {
			c <- true
			return
		}
		c <- false

	}(c4)

	ok := <-c1
	if !ok {
		t.Fatalf("Error on GET request")
	}
	ok = <-c2
	if !ok {
		t.Fatalf("Passed id=abc but server didn't catch it")
	}

	ok = <-c3
	if !ok {
		t.Fatalf("Passed id=1000000 but server didn't catch it")
	}
	ok = <-c4
	if !ok {
		t.Fatalf("Expected '1' GOT: Something else")
	}

}

func TestMain(m *testing.M) {
	go ServerInit()
	time.Sleep(time.Millisecond)
	m.Run()

	os.Exit(1)
}
