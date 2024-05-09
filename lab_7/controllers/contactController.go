package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}
	
	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var EditContact = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	
	contact := &models.Contact{}
	
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}
	
	res := contact.EditContact(id)
 	
	if !res {
		u.Respond(w, u.Message(false, "Unknown user."))
	} else {
		u.Respond(w, u.Message(true, "success"))
	}
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	
	contact := &models.Contact{}
	
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}
	
	res := contact.DeleteContact(id)
	
	if !res {
		u.Respond(w, u.Message(false, "Unknown user."))
	} else {
		u.Respond(w, u.Message(true, "success"))
	}
	
	fmt.Println(id)
}
