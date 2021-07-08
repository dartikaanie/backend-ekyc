package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type UserData struct {
	UserID      string `json:"trx_id"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	FullName    string `json:"fullName"`
	DateOfBirth string `json:"dateOfBirth"`
	GovId       string `json:"govId"`
	Status      string `json:"status"`
}

type ResponseUserStatus struct {
	Status string `json:"status"`
	Desc   string `json:"description"`
	UserId string `json:"user_id"`
}

type allUsers []UserData

var users = allUsers{
	{
		UserID:      "IL150695b7d0b82b1739RI",
		Email:       "dartikadara@gmail.com",
		Mobile:      "+6282381135788",
		FullName:    "Dartika Anie Marian",
		DateOfBirth: "1997-08-26",
		GovId:       "1374026608970021",
		Status:      "-",
	},
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser UserData
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			fmt.Println(msg)
			http.Error(w, msg, http.StatusUnsupportedMediaType)

			return
		}
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	fmt.Println("Endpoint Hit: createUser : " + newUser.UserID)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	UserID := mux.Vars(r)["user_id"]
	fmt.Println("Endpoint Hit: getUser id : " + UserID)
	for _, singleEvent := range users {
		if singleEvent.UserID == UserID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getNotif(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit getNotif")

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			fmt.Println(msg)
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var responseUserStatus ResponseUserStatus
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(" getNotif request " + string(reqBody))
	json.Unmarshal(reqBody, &responseUserStatus)

	if err != nil {
		log.Fatalln(err)
	}
	responseUserStatus.Desc = "User Not Found"
	responseUserStatus.Status = "-"
	if responseUserStatus.UserId != "" {
		for _, singleUser := range users {
			if singleUser.UserID == responseUserStatus.UserId {
				responseUserStatus.Desc = "notify success"
				responseUserStatus.Status = "success"
			}
		}
	}

	fmt.Println("getNotif response: " + responseString(responseUserStatus))
	json.NewEncoder(w).Encode(responseUserStatus)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]
	fmt.Println("Endpoint Hit: deleteUser : " + id)
	for index, user := range users {
		if user.UserID == id {
			users = append(users[:index], users[index+1:]...)
		}
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println("Endpoint Hit update  userID: " + userID)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the status to update")
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			fmt.Println(msg)
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var updatedUser UserData
	json.Unmarshal(reqBody, &updatedUser)
	fmt.Println(" updateUser  " + string(reqBody))
	var responseUserStatus ResponseUserStatus
	responseUserStatus.UserId = userID
	responseUserStatus.Desc = "failed to update status"
	responseUserStatus.Status = "-"

	if updatedUser.Status != "" {
		for i, singleUser := range users {
			if singleUser.UserID == userID {
				singleUser.Status = updatedUser.Status
				users = append(users[:i], singleUser)
				responseUserStatus.Status = updatedUser.Status
				responseUserStatus.Desc = "success update status "
			}
		}
	}

	fmt.Println("update : " + responseString(responseUserStatus))
	json.NewEncoder(w).Encode(responseUserStatus)
}

func responseString(status ResponseUserStatus) string {
	return "{ 'user_id' :" + status.UserId + ", status :" + status.Status + ", description :" + status.Desc + " }"
}
