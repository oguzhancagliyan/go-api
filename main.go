package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"yemek-sepeti-case-api/src/configuration"
	"yemek-sepeti-case-api/src/controllers"
	"yemek-sepeti-case-api/src/models"
	"yemek-sepeti-case-api/src/services"
)

type Handler struct {
	keyValueController *controllers.KeyValueController
}

func NewHandler(keyValueController *controllers.KeyValueController) *Handler {
	return &Handler{keyValueController: keyValueController}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/api"):
		keys := strings.Split(r.URL.Path, "/api/")
		h.keyValueController.Get(w, keys[1])
		return
	case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/api"):
		h.keyValueController.Set(w, r)
		return
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/api"):
		h.keyValueController.Flush(w)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	db := configuration.NewDatabase()
	configuration.RestoreDatabase(db)
	keyValueService := services.NewKeyValueService(db)
	keyValueController := controllers.NewKeyValueController(keyValueService)

	go func() {
		for true {

			if len(db.Db) < 1 {
				time.Sleep(1 * time.Minute)
				continue
			}

			var dataArray []models.Data
			storedDb := db.Db
			for key, value := range storedDb {
				data := models.Data{
					Key:   key,
					Value: value,
				}
				dataArray = append(dataArray, data)
			}

			serializedArray, err := json.Marshal(dataArray)

			if err != nil {
				panic(err.Error())
			}

			currentTime := time.Now().Unix()
			fileName := fmt.Sprintf("./src/backups/%d-data.json", currentTime)
			file, err := os.Create(fileName)

			fmt.Printf("Saving data to %s", fileName)
			file.Write(serializedArray)
			fmt.Printf("writed data to file. FileName : %s\n", fileName)

			file.Close()

			time.Sleep(1 * time.Minute)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("/api/", NewHandler(keyValueController))
	err := http.ListenAndServe(":8081", mux)

	if err != nil {
		panic(fmt.Errorf("error on listenandService %s", err.Error()))
	}
}
