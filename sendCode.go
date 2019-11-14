package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CodeResponse struct {
	Code_id string `json:"code_id"`
	Code    int    `json:"code"`
	Timeout int    `json:"timeout"`
}

var counter int = 0

func sendCode(w http.ResponseWriter, r *http.Request) {
	counter += 1
	fmt.Println("COUNTER /sendCode:", counter)

	if counter == 2 {
		e := ErrorResp{
			ErrorCode: 1008,
			Timeout:   2433,
		}
		r, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling json:", err)
		}
		w.WriteHeader(409)
		w.Write(r)
		return
	}

	if counter > 3 {
		fmt.Println("COUNTER > 3")
		e := ErrorResp{
			ErrorCode: 1006,
		}
		r, err := json.Marshal(e)
		if err != nil {
			log.Println("Error marshalling json:", err)
		}
		counter = 0
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		w.Write(r)
		return
	}

	resp := CodeResponse{
		Code_id: "xdSGMwHje3bYnNtUwEHwel0fPU4HXn",
		Code:    123456789,
		Timeout: 3000,
	}

	rows, dberr := db.Query(`
		INSERT INTO smscodes (code_id, code)
		VALUES ($1, $2)
	`, resp.Code_id, resp.Code)
	if dberr != nil {
		log.Println("Error writing into DB:", dberr)
	}
	defer rows.Close()
	res, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error marshalling json:", err)
	}

	time.Sleep(1500 * time.Millisecond)

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
