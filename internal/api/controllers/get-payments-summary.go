package controllers

import (
	"net/http"
)

func GetPaymentsSummary(w http.ResponseWriter, r *http.Request) {
	// Logic for getting payments summary will go here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payments summary"))
}
