package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

type User struct {
	Id            string  `json:"id"`
	Username      string  `json:"username"`
	Enabled       bool    `json:"enabled"`
	EmailVerified bool    `json:"emailVerified"`
	Email         string  `json:"email"`
	Self          string  `json:"self"`
	Attributes    Attribs `json:"attributes"`
}

type Attribs struct {
	Phone                 []string `json:"phone"`
	Description           []string `json:"description"`
	AuthSmsVerified       []string `json;"AUTH_SMS_VERIFIED"`
	LastAuthorizationDate []string `json:"last_authorization_date"`
	CreateDate            []string `json:"create_date"`
	HashType              []string `json:"hash_type"`
	AuthEmailVerified     []string `json:"AUTH_EMAIL_VERIFIED"`
}

func init() {
	connStr := "user=*********** password=************ host=localhost dbname=testdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	defer db.Close()

	http.HandleFunc("/getUserInfo", userInfo)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/sendCode", checkMethod("POST", sendCode))
	http.HandleFunc("/validateCode", checkMethod("POST", validateCode))
	http.HandleFunc("/validatePhone/", validatePhone)
	http.HandleFunc("/validateEmail/", validateEmail)
	http.HandleFunc("/updUserPhone", checkMethod("PUT", updateUserPhone))
	http.HandleFunc("/sendCodeToEmail", checkMethod("POST", sendCodeToEmail))
	http.HandleFunc("/updUser", checkMethod("PUT", updateUser))

	http.ListenAndServe(":3000", nil)
}
