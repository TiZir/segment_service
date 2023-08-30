package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TiZir/segment_service/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

type Data struct {
	Added   []string `json:"added"`
	Deleted []string `json:"deleted"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the App! Server started successfully.")
}

func createSegment(w http.ResponseWriter, r *http.Request) {
	var segment db.Segment
	segment.Name = mux.Vars(r)["name"]
	err := db.InsertSegment(segment)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteSegment(w http.ResponseWriter, r *http.Request) {
	var segment db.Segment
	segment.Name = mux.Vars(r)["name"]
	err := db.DeleteSegment(segment)
	if err != nil {
		log.Fatal(err)
	}
}

func getSegments(w http.ResponseWriter, r *http.Request) {
	segments, err := db.SelectSegmentTest()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(segments)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getCompliance(w http.ResponseWriter, r *http.Request) {
	compliance, err := db.SelectCompliance()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(compliance)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getComplianceById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id_user"])
	if err != nil {
		log.Fatal(err)
	}
	compliance, err := db.SelectComplianceById(id)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(compliance)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func addUserToSegment(w http.ResponseWriter, r *http.Request) {

	var data Data
	id, err := strconv.Atoi(mux.Vars(r)["id_user"])
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	if len(data.Added) >= 1 {
		_, err = db.SelectSegment(data.Added)
		if err != nil {
			log.Fatal(err)
		} else {
			err = db.InsertCompliance(data.Added, id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if len(data.Deleted) >= 1 {
		_, err = db.SelectSegment(data.Deleted)
		if err != nil {
			log.Fatal(err)
		} else {
			err = db.DeleteCompliance(data.Deleted, id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/segments", getSegments).Methods("GET")
	router.HandleFunc("/segments/create/{name}", createSegment)
	router.HandleFunc("/segments/delete/{name}", deleteSegment)
	router.HandleFunc("/compliance", getCompliance).Methods("GET")
	router.HandleFunc("/compliance/{id_user}", getComplianceById).Methods("GET")
	router.HandleFunc("/compliance/{id_user}/segments", addUserToSegment).Methods("PUT")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
