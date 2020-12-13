package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-REST/pkg/models"
	"simple-REST/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	createdUser := newUser.CreateUser()
	res, _ := json.Marshal(createdUser)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	id, err := strconv.ParseInt(userID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.GetUserById(id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updatedUser := &models.User{}
	utils.ParseBody(r, updatedUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	existingUser := models.GetUserById(id)
	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}
	models.UpdateUser(existingUser)
	res, _ := json.Marshal(existingUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
