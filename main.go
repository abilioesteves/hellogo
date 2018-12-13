package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")

	logrus.Info("Initialized Hello World")
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		logrus.Errorf("Error initializing the DNS Manager Webhook: %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	hostname, _ := os.Hostname()
	json.NewEncoder(w).Encode(fmt.Sprintf("Hello World: %s", hostname))
}
