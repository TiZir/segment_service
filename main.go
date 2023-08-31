package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TiZir/segment_service/handler"
	"github.com/TiZir/segment_service/helper"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("App started successfully")
	helper.MakeRespons(w, "App started successfully", http.StatusOK, nil)
}

func main() {
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")

	router.HandleFunc("/segments", handler.GetSegments).Methods("GET")
	router.HandleFunc("/segments/create/{name}", handler.CreateSegment)
	router.HandleFunc("/segments/delete/{name}", handler.DeleteSegment)

	router.HandleFunc("/compliances", handler.GetCompliance).Methods("GET")
	router.HandleFunc("/compliances/{id_user}", handler.GetComplianceById).Methods("GET")
	router.HandleFunc("/compliances/{id_user}/segments", handler.AddUserToCompliance).Methods("PUT")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
