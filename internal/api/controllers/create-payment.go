package controllers

import (
	"net/http"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	// Logic for creating a payment will go here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment created successfully"))
}
