package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	//local modules
	"../domain"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	var userIDStr string = req.URL.Query().Get("user_id")
	var userID int64
	var err error
	userID, err = strconv.ParseInt(userIDStr, 10, 64)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number."))
		return
	}

	var user domain.User
	user, err = services.GetUser(userID)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))

		//handle the err and return to the client
		return
	}

	// return user to client
	jsonValue, _ = json.Marshal(user)
	resp.Write(jsonValue)
}
