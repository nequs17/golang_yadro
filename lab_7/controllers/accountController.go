package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request!"))
		return
	}
	
	resp := account.CreateAccount()
	u.Respond(w, resp)
}

var LoginAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request!"))
		return
	}
	resp := models.LoginAccount(account.Email, account.Password)
	u.Respond(w, resp)
}
