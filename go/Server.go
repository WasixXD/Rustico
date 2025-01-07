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

type ProjectRequest struct {
	Project_name string `json:"project_name"`
	Description  string `json:"description"`
}

type ProjectResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"project_name"`
	Desc    string `json:"description"`
	Created string `json:"created_at"`
}

func Root(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func CreateProject(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Println("Error on reading the body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var project ProjectRequest
	err = json.Unmarshal(body, &project)

	if err != nil {
		log.Println("Error on unmarshaling ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Exec("INSERT INTO projects(project_name, description) VALUES (?, ?)", project.Project_name, project.Description)
	w.WriteHeader(http.StatusCreated)
}

func GetRandomProject(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	result := db.QueryRow("select id, project_name, description, created_at from projects order by random() limit 1;")

	var response ProjectResponse

	err := result.Scan(&response.Id, &response.Name, &response.Desc, &response.Created)

	if err != nil {
		log.Println("Error on scanning the row: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
	mux.Handle("/random", http.HandlerFunc(GetRandomProject))

	http.ListenAndServe(":3001", mux)
}
