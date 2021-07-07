package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UserData struct {
	UserID      string `json:"trx_id"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	FullName    string `json:"fullName"`
	DateOfBirth string `json:"dateOfBirth"`
	GovId       string `json:"govId"`
}

type allUsers []UserData

var users = allUsers{
	{
		UserID: "IL150695b7d0b82b1739RI",
		Email : "dartikadara@gmail.com",
		Mobile: "+6282381135788",
		FullName: "Dartika Anie Marian",
		DateOfBirth: "1997-08-26",
		GovId: "1374026608970021",
	},
}



func returnAllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(users)
}


func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createUser")
	var newUser UserData
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: getUser")
	UserID := mux.Vars(r)["user_id"]

	fmt.Println("UserID:" + UserID)
	for _, singleEvent := range users {
		if singleEvent.UserID == UserID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUser")
	vars := mux.Vars(r)
	id := vars["user_id"]

	for index, user := range users {
		if user.UserID  == id {
			users = append(users[:index], users[index+1:]...)
		}
	}

}
