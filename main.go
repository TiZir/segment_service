package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TiZir/segment_service/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func compliancePage(w http.ResponseWriter, r *http.Request) {
	compliance, err := db.SelectCompliance()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(compliance)
	if err != nil {
		log.Fatal("Error marshaling data to JSON:", err)
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// return
	}
	fmt.Println("Endpoint Hit: compliancePage")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/compliance", compliancePage)
	// err := db.OpenEnv()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
