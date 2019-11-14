package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func logout(w http.ResponseWriter, r *http.Request) {
	var urlParams url.Values
	r.ParseForm()
	urlParams = r.Form
	// for k, v := range urlParams {
	// 	fmt.Println("param:", k, "value:", v)
	// }
	redirectUri := urlParams["redirect_uri"][0]
	fmt.Println(redirectUri)

	w.WriteHeader(302)
	w.Header().Set("Location", redirectUri)
}
