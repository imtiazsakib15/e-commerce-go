package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	createdUser := newUser.Store(newUser)

	util.SendData(w, createdUser, http.StatusCreated)
}