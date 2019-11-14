package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorCode struct {
	ErrorCode int `json:"errorCode"`
}

func validatePhone(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path[len("/validatePhone/"):]
	okphone1 := "79183954442"
	okphone2 := "79183954441"

	if param == okphone1 || param == okphone2 {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json, text/plain")
		w.WriteHeader(200)
	} else {
		e := ErrorCode{
			ErrorCode: 1009,
		}
		resp, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling json:", err)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json, text/plain")
		w.WriteHeader(409)
		w.Write(resp)
	}
}
