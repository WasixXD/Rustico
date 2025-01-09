package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const DB_PATH = "./project.db"

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
	result := db.QueryRow("select id, project_name, description, created_at from projects where github_url is null order by random() limit 1;")

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

func AllProjects(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	result, err := db.Query("select id, project_name, description, strftime('%Y.%m.%d', created_at) as created_at, github_url from projects where github_url is not null;")

	if err != nil {
		log.Println("Error on querying: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response []ProjectResponse
	var project ProjectResponse

	for {
		result.Next()
		err := result.Scan(&project.Id, &project.Name, &project.Desc, &project.Created, &project.GithubUrl)
		if err != nil {
			break
		}
		response = append(response, project)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func DeleteProject(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	query := req.URL.Query().Get("id")

	if query == "" {
		log.Println("Deleting doesn't have a query")
		http.Error(w, "Please provide the id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(query)

	if err != nil {
		log.Println("ID is not a number: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("delete from projects where id = ?", id)

	if err != nil {
		log.Println("Something went wrong when querying: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()

	json.NewEncoder(w).Encode(rows)
}

// should probably put this in a middleware
func CompleteProject(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	query := req.URL.Query().Get("id")
	url := req.URL.Query().Get("url")

	if query == "" {
		log.Println("Completing doesn't have a query")
		http.Error(w, "Please provide the id", http.StatusBadRequest)
		return
	} else if url == "" {
		log.Println("Completing doesn't have a url")
		http.Error(w, "Please provide the url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(query)

	if err != nil {
		log.Println("ID is not a number: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("update projects set github_url = ? where id = ?", url, id)

	if err != nil {
		log.Println("Something went wrong when querying: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()

	json.NewEncoder(w).Encode(rows)
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
	mux.Handle("/delete", http.HandlerFunc(DeleteProject))
	mux.Handle("/complete", http.HandlerFunc(CompleteProject))
	mux.Handle("/all", http.HandlerFunc(AllProjects))

	http.ListenAndServe("0.0.0.0:3001", mux)
}
