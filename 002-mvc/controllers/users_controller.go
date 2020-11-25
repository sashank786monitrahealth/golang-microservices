package controllers

import (
	"log"
	"net/http"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	var userID string = req.URL.Query().Get("user_id")
	log.Printf("About to process user_id %v\n", userID)
	log.Printf("About to process user_id %T\n", userID)
}
