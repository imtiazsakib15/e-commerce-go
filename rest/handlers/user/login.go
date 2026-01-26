package user

import (
	"ecommerce/config"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	user, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if(err != nil) {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, err)
		return
	}
	if(user == nil) {
		util.SendError(w, http.StatusBadRequest, "Invalid email or password")
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
		util.SendError(w, http.StatusBadRequest, "An error occurred")
	}

	util.SendData(w, http.StatusOK, accessToken)
}