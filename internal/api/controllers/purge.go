package controllers

import "net/http"

func Purge(w http.ResponseWriter, r *http.Request) {
	// Logic for purging data will go here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data purged successfully"))
}
