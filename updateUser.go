package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestData struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handler", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))

	id := "jh681ui4s66dti1386"
	var data RequestData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}

	rows, dberr := db.Query(`
		UPDATE users
		SET email=($1), phone=($2)
		WHERE id=($3)
	`, data.Email, data.Phone, id)
	if dberr != nil {
		log.Println("Error updating user:", dberr)
	}
	rows.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
	json.NewEncoder(w).Encode(data)
}
