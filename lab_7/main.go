package main

import (
	app "contactsBook/authentication"
	"contactsBook/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	
	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginAccount).Methods("POST")
	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/me/contacts", controllers.GetContacts).Methods("GET")
	
	// my edtis
	router.HandleFunc("/me/contacts", controllers.EditContact).Methods("PUT")
	router.HandleFunc("/me/contacts", controllers.DeleteContact).Methods("DELETE")
	
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// ready print
	fmt.Println("Server is ready. Listening port: " + port)
	
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
