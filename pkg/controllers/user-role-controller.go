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

func CreateUserRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userID, _ := strconv.ParseInt(userId, 0, 0)

	newUserRoleMapping := &models.UserRoleMapping{UserId: userID}
	utils.ParseBody(r, newUserRoleMapping)

	role := models.GetRoleById(newUserRoleMapping.RoleId)
	var res []byte
	if role.Name == "" {
		errorMsg := "Invalid Role"
		w.WriteHeader(http.StatusBadRequest)
		res = []byte(errorMsg)

	} else {
		newUserRoleMapping.CreateUserRole()
		w.WriteHeader(http.StatusOK)
		res, _ = json.Marshal(newUserRoleMapping)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	id, err := strconv.ParseInt(userID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	roles := models.GetUserRoles(id)
	res, _ := json.Marshal(roles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
