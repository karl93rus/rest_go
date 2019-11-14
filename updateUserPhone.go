package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type PhoneInfo struct {
	PhoneValidateCodeId string `json:"phone_validate_code_id"`
	Phone               string `json:"phone"`
}

func updateUserPhone(w http.ResponseWriter, r *http.Request) {
	id := "xdSGMwHje3bYnNtUwEHwel0fPU4HXn"
	uid := "jh681ui4s66dti1386"
	var reqBody PhoneInfo
	res, _ := ioutil.ReadAll(r.Body)
	herr := json.Unmarshal(res, &reqBody)
	if herr != nil {
		log.Println("Error unmarshalling json:", herr)
	}

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// w.Header().Set("Content-Type", "application/json, text/plain")

	if reqBody.PhoneValidateCodeId == id {
		rows, dberr := db.Query(`
			UPDATE users
			SET phone=($1)
			WHERE id=($2)
	`, reqBody.Phone, uid)
		if dberr != nil {
			log.Println("Error updating user:", dberr)
			w.WriteHeader(500)
			return
		}
		rows.Close()
		w.WriteHeader(204)
	} else if reqBody.PhoneValidateCodeId != id {
		e := ErrorResp{
			ErrorCode: 1013,
		}
		resp, merr := json.Marshal(e)
		if merr != nil {
			log.Println("Error marshalling json:", merr)
		}
		w.WriteHeader(409)
		w.Write(resp)
	}
}
