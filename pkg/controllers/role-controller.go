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

func CreateRole(w http.ResponseWriter, r *http.Request) {
	newRole := &models.Role{}
	utils.ParseBody(r, newRole)
	createdRole := newRole.CreateRole()
	res, _ := json.Marshal(createdRole)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	roles := models.GetAllRoles()
	res, _ := json.Marshal(roles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID := vars["roleId"]
	id, err := strconv.ParseInt(roleID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	role := models.GetRoleById(id)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	updatedRole := &models.Role{}
	utils.ParseBody(r, updatedRole)
	vars := mux.Vars(r)
	roleId := vars["roleId"]
	id, err := strconv.ParseInt(roleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	existingRole := models.GetRoleById(id)
	if updatedRole.Name != "" {
		existingRole.Name = updatedRole.Name
	}
	models.UpdateRole(existingRole)
	res, _ := json.Marshal(existingRole)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId := vars["roleId"]
	id, err := strconv.ParseInt(roleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	role := models.DeleteRole(id)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
