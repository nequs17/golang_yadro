package models

import (
	u "contactsBook/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// add function for email validation
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	// теперь проверка на email полноценна
	if !isEmailValid(account.Email) {
		return u.Message(false, "Email address is not valid!"), false
	}

	if len(account.Password) < 4 {
		return u.Message(false, "Password must be longer then 4 symbols"), false
	}

	acc := &Account{}
	
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error!"), false
	}
	if acc.Email != "" {
		return u.Message(false, "The email is already occupied by another user!"), false
	}

	return u.Message(false, "Check is passed!"), true
}

func (account *Account) CreateAccount() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	account.Password = ""

	response := u.Message(true, "Account has been created!")
	response["account"] = account
	return response
}

func LoginAccount(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password does not match!")
	}
	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
