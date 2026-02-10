package handlers

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search results"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login success"))
}
