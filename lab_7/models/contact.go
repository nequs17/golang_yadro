package models

import (
	u "contactsBook/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"regexp"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty!"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number cannot be empty!"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User not found!"), false
	}

	phoneRegex := regexp.MustCompile(`^7([0-9]{3})([0-9]{3})([0-9]{2})([0-9]{2})$`)
	if !phoneRegex.MatchString(contact.Phone) {
		return u.Message(false, "Invalid telephoe number format"), false
	}
	
	return u.Message(true, "success"), true
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}
	
	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}

func (contact *Contact) EditContact(user uint) bool {
	err := GetDB().Table("contacts").Where("user_id = ? AND phone = ?", user, contact.Phone).Update(map[string]interface{}{"name": contact.Name}).Error
	
	if err != nil {
		return false
	}
	
	return true
}

func (contact *Contact) DeleteContact(user uint) bool {
	err := GetDB().Where("user_id = ? AND phone = ?", user, contact.Phone).Delete(&Contact{}).Error
	
	if err != nil {
		return false
	}
	
	return true
}