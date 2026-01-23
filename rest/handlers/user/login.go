package user

import (
	"ecommerce/config"
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	user := database.Find(loginReq.Email, loginReq.Password)
	if(user == nil) {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	
	cnf := config.GetConfig()
	accessToken, err := util.CreateJwt(util.Payload{
		Sub: user.ID,
		Name: user.Name,
		Email: user.Email,
		IsShopOwner: user.IsShopOwner,
	}, cnf.JwtSecret)
	if(err != nil) {
		http.Error(w, "An error occurred", http.StatusBadRequest)
	}

	util.SendData(w, accessToken, http.StatusOK)
}