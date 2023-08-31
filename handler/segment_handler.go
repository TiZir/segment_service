package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TiZir/segment_service/db"
	"github.com/TiZir/segment_service/helper"
	"github.com/gorilla/mux"
)

func CreateSegment(w http.ResponseWriter, r *http.Request) {
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

func DeleteSegment(w http.ResponseWriter, r *http.Request) {
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

func GetSegments(w http.ResponseWriter, r *http.Request) {
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
