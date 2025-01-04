package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const DB_PATH = "../project.db"

type Project struct {
	Project_name string `json:"project_name"`
	Description  string `json:"description"`
}

func Root(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CreateProject(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Println("Error on reading the body")
	}

	var project Project
	err = json.Unmarshal(body, &project)

	if err != nil {
		log.Println("Error on unmarshaling ", err)
	}

	db.Exec("INSERT INTO projects(project_name, description) VALUES (?, ?)", project.Project_name, project.Description)
	w.WriteHeader(http.StatusCreated)
}

func ServerInit() {
	var err error
	db, err = sql.Open("sqlite3", DB_PATH)

	if err != nil {
		log.Fatalln("Error on opening sqlite3 database: ", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(Root))
	mux.Handle("/create", http.HandlerFunc(CreateProject))

	http.ListenAndServe(":3001", mux)
}
