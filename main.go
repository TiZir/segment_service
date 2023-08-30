package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TiZir/segment_service/db"
	"github.com/TiZir/segment_service/helper"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

type Data struct {
	Added   []string `json:"added"`
	Deleted []string `json:"deleted"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("App started successfully")
	helper.MakeRespons(w, "App started successfully", http.StatusOK, nil)
}

func createSegment(w http.ResponseWriter, r *http.Request) {
	var segment db.Segment
	segment.Name = mux.Vars(r)["name"]
	err := db.InsertSegment(segment)
	if err != nil {
		log.Printf("Error adding to segment: %v", err)
		helper.MakeRespons(w, "Error adding to segment", http.StatusInternalServerError, err)
		return
	}
	helper.MakeRespons(w, "Successful add", http.StatusOK, nil)
}

func deleteSegment(w http.ResponseWriter, r *http.Request) {
	var segment db.Segment
	segment.Name = mux.Vars(r)["name"]
	err := db.DeleteSegment(segment)
	if err != nil {
		log.Printf("Error deleting from segment: %v", err)
		helper.MakeRespons(w, "Error deleting from segment", http.StatusInternalServerError, err)
		return
	}
	helper.MakeRespons(w, "Successful del", http.StatusOK, nil)
}

func getSegments(w http.ResponseWriter, r *http.Request) {
	segments, err := db.SelectSegmentTest()
	if err != nil {
		log.Printf("Selection error from segment: %v", err)
		helper.MakeRespons(w, "Selection error from segment", http.StatusInternalServerError, err)
		return
	}
	jsonData, err := json.Marshal(segments)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		helper.MakeRespons(w, "Failed to marshal data", http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getCompliance(w http.ResponseWriter, r *http.Request) {
	compliance, err := db.SelectCompliance()
	if err != nil {
		log.Printf("Selection error from compliance: %v", err)
		helper.MakeRespons(w, "Selection error from compliance", http.StatusInternalServerError, err)
		return
	}
	jsonData, err := json.Marshal(compliance)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		helper.MakeRespons(w, "Failed to marshal data", http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getComplianceById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id_user"])
	if err != nil {
		log.Printf("Incorrect id_user for get compliance: %v", err)
		helper.MakeRespons(w, "Incorrect id_user for get compliance", http.StatusBadRequest, err)
		return
	}
	compliance, err := db.SelectComplianceById(id)
	if err != nil {
		log.Printf("Selection error from compliance: %v", err)
		helper.MakeRespons(w, "Selection error from compliance", http.StatusInternalServerError, err)
		return
	}
	jsonData, err := json.Marshal(compliance)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		helper.MakeRespons(w, "Failed to marshal data", http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func addUserToSegment(w http.ResponseWriter, r *http.Request) {

	var data Data
	id, err := strconv.Atoi(mux.Vars(r)["id_user"])
	if err != nil {
		log.Printf("Incorrect id_user for add segment: %v", err)
		helper.MakeRespons(w, "Incorrect id_user for add segment", http.StatusBadRequest, err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Incorrect data in body for add segment: %v", err)
		helper.MakeRespons(w, "Incorrect data in body for add segment", http.StatusBadRequest, err)
		return
	}
	if len(data.Added) >= 1 {
		_, err = db.SelectSegment(data.Added)
		if err != nil {
			log.Printf("Selection error from segment: %v", err)
			helper.MakeRespons(w, "Selection error from segment", http.StatusInternalServerError, err)
			return
		} else {
			err = db.InsertCompliance(data.Added, id)
			if err != nil {
				log.Printf("Error adding to segment: %v", err)
				helper.MakeRespons(w, "Error adding to segment", http.StatusInternalServerError, err)
				return
			}
		}
	}
	if len(data.Deleted) >= 1 {
		_, err = db.SelectSegment(data.Deleted)
		if err != nil {
			log.Printf("Selection error from segment: %v", err)
			helper.MakeRespons(w, "Selection error from segment", http.StatusInternalServerError, err)
			return
		} else {
			err = db.DeleteCompliance(data.Deleted, id)
			if err != nil {
				log.Printf("Error deleting from segment: %v", err)
				helper.MakeRespons(w, "Error deleting from segment", http.StatusInternalServerError, err)
				return
			}
		}
	}
	helper.MakeRespons(w, "Successful add", http.StatusCreated, nil)
}

func main() {
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")

	router.HandleFunc("/segments", getSegments).Methods("GET")
	router.HandleFunc("/segments/create/{name}", createSegment)
	router.HandleFunc("/segments/delete/{name}", deleteSegment)

	router.HandleFunc("/compliances", getCompliance).Methods("GET")
	router.HandleFunc("/compliances/{id_user}", getComplianceById).Methods("GET")
	router.HandleFunc("/compliances/{id_user}/segments", addUserToSegment).Methods("PUT")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
