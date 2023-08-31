package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TiZir/segment_service/db"
	"github.com/TiZir/segment_service/helper"
	"github.com/gorilla/mux"
)

type Data struct {
	Added   []string `json:"added"`
	Deleted []string `json:"deleted"`
}

func GetCompliance(w http.ResponseWriter, r *http.Request) {
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

func GetComplianceById(w http.ResponseWriter, r *http.Request) {
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

func AddUserToCompliance(w http.ResponseWriter, r *http.Request) {
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
	flagAdd := false
	flagDel := false
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
		flagAdd = true
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
		flagDel = true
	}
	if flagAdd && flagDel {
		helper.MakeRespons(w, "Successful add and del", http.StatusCreated, nil)
	} else if flagAdd {
		helper.MakeRespons(w, "Successful add", http.StatusOK, nil)
	} else if flagDel {
		helper.MakeRespons(w, "Successful del", http.StatusOK, nil)
	} else {
		helper.MakeRespons(w, "Error removing or adding", http.StatusInternalServerError, nil)
	}
}
