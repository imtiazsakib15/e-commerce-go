package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser ReqCreateUser
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	createdUser, err := h.userRepo.Create(repo.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, err)
		return
	}

	util.SendData(w, http.StatusCreated, createdUser)
}