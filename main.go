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
	fmt.Println("Endpoint Hit: compliancePage")
	json.NewEncoder(w).Encode(compliance)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	err := db.OpenEnv()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/compliance", compliancePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
