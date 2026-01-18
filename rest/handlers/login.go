package handler

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginReq struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	user := database.Find(loginReq.Email, loginReq.Password)
	if(user != nil) {
		util.SendData(w, user, http.StatusOK)
		return
	}
	
	http.Error(w, "Invalid credentials", http.StatusBadRequest)	
}