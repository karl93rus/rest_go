package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type EmailRequest struct {
	Address string
	Type    string
}

func sendCodeToEmail(w http.ResponseWriter, r *http.Request) {
	uid := "jh681ui4s66dti1386"
	var reqObj EmailRequest
	req, _ := ioutil.ReadAll(r.Body)
	uerr := json.Unmarshal(req, &reqObj)
	if uerr != nil {
		log.Println("Error unmarshalling json:", uerr)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json, text/plain")
	rows, dberr := db.Query(`
			UPDATE users
			SET email=($1)
			WHERE id=($2)
	`, reqObj.Address, uid)
	if dberr != nil {
		log.Println("Error updating user:", dberr)
	}
	rows.Close()

	w.WriteHeader(201)
}
