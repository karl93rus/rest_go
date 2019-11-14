package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var mail_counter int = 0

func validateEmail(w http.ResponseWriter, r *http.Request) {
	mail_counter += 1
	if mail_counter > 3 {
		fmt.Println("COUNTER > 3")
		e := ErrorResp{
			ErrorCode: 1010,
		}
		r, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling json:", err)
		}
		mail_counter = 0
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		w.Write(r)
		return
	}
	param := r.URL.Path[len("/validateEmail/"):]
	fmt.Println(param)
	emailok1 := "azaza@azaz.za"
	emailok2 := "azaza@azaz.zu"
	emailok3 := "azaza@azaz.zam"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json, text/plain")

	if param == emailok3 {
		w.WriteHeader(401)
		return
	}

	if param == emailok1 || param == emailok2 {
		w.WriteHeader(200)
	} else {
		e := ErrorResp{
			ErrorCode: 1015,
		}
		res, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling json:", err)
		}

		w.WriteHeader(409)
		w.Write(res)
	}
}
