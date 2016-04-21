package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
)

func main() {
	// Simple Server
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	var r *http.Request
	// Fetch Cookie Monster
	cookie, _ := r.Cookie("username")

	fmt.Println(cookie)

	var w http.ResponseWriter
	// Another way to fetch cookie
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Example for setting Cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: "bunchhieng", Expires: expiration}

	http.SetCookie(w, &cookie)
}
