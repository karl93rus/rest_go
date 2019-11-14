package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ReqData struct {
	CodeId        string `json:"code_id"`
	ValidateText  string `json:"validate_text"`
	ValidatePhone string `json:"validate_phone"`
}

type ErrorResp struct {
	ErrorCode int `json:"errorCode"`
	Timeout   int `json:"timeout,omitempty"`
}

func validateCode(w http.ResponseWriter, r *http.Request) {
	cid := "xdSGMwHje3bYnNtUwEHwel0fPU4HXn"
	var reqData ReqData
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
	}

	jerr := json.Unmarshal(body, &reqData)
	if jerr != nil {
		log.Println("Error unmarshaling json:", jerr)
	}

	if reqData.CodeId == cid && reqData.ValidateText == "123456789" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json, text/plain")
		w.WriteHeader(200)
	} else {
		e := ErrorResp{
			ErrorCode: 1002,
		}
		resp, merr := json.Marshal(e)
		if merr != nil {
			log.Println("Error marshalling json:", merr)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, PUT, POST, OPTIONS")
		w.WriteHeader(409)
		w.Write(resp)
	}

}
