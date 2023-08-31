package handler

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/TiZir/segment_service/db"
	"github.com/TiZir/segment_service/helper"
)

func WriteCSV(w http.ResponseWriter, r *http.Request) {
	file, err := os.CreateTemp("", "output_*.csv")
	if err != nil {
		log.Printf("Unable to create csv file: %v", err)
		helper.MakeRespons(w, "Unable to create csv file", http.StatusInternalServerError, err)
		return
	}
	log.Printf("Temporary file path: %v", file.Name()) // Добавьте эту строку для отладки
	defer file.Close()
	defer os.Remove(file.Name())

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data, err := db.SelectHistory()
	if err != nil {
		log.Printf("Selection error from history: %v", err)
		helper.MakeRespons(w, "Selection error from history", http.StatusInternalServerError, err)
		return
	}
	for _, history := range data {
		var timeAction string
		var action string

		if history.TSdel != nil {
			action = "del"
			tsDel, err := time.Parse("2006-01-02 15:04:05", string(history.TSdel))
			if err != nil {
				log.Printf("Error parse time: %v", err)
				helper.MakeRespons(w, "Error parse time", http.StatusInternalServerError, err)
				return
			}
			timeAction = tsDel.String()
		} else {
			action = "add"
			tsAdd, err := time.Parse("2006-01-02 15:04:05", string(history.TSadd))
			if err != nil {
				log.Printf("Error parse time: %v", err)
				helper.MakeRespons(w, "Error parse time", http.StatusInternalServerError, err)
				return
			}
			timeAction = tsAdd.String()
		}
		record := []string{
			strconv.Itoa(history.IdUser),
			history.NameSegment,
			action,
			timeAction,
		}
		err := writer.Write(record)
		if err != nil {
			log.Printf("Csv write error: %v", err)
			helper.MakeRespons(w, "Csv write error", http.StatusInternalServerError, err)
			return
		}
		writer.Flush()
	}
	file.Seek(0, 0)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=output.csv")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}
