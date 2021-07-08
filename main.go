package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Run Ekyc")
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{user_id}", getUser).Methods("GET")
	myRouter.HandleFunc("/user", createUser).Methods("POST")
	myRouter.HandleFunc("/user/{user_id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{user_id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/notif", getNotif).Methods("POST")
	log.Fatal(http.ListenAndServe(":3030", myRouter))
}
