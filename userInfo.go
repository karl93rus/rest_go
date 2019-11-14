package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	id := "jh681ui4s66dti1386" // phone and mail
	// id := "8ty34756hg93jvfj" // phone
	// id := "asd" // mail
	log.Println(r.RemoteAddr)
	var phone string

	user := User{
		Attributes: Attribs{
			HashType: []string{"sha1"},
		},
	}

	rows, err := db.Query("SELECT * from users WHERE id=($1)", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	user.Attributes.Phone = append(user.Attributes.Phone, phone)

	resp, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
