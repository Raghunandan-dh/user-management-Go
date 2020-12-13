package routes

import (
	"simple-REST/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/v1/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/v1/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/users/{userId}", controllers.DeleteUser).Methods("DELETE")
}

var RegisterRoleRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/roles/", controllers.CreateRole).Methods("POST")
	router.HandleFunc("/v1/roles/", controllers.GetRoles).Methods("GET")
	router.HandleFunc("/v1/roles/{roleId}", controllers.UpdateRole).Methods("PUT")
	router.HandleFunc("/v1/roles/{roleId}", controllers.DeleteRole).Methods("DELETE")
}

var RegisterUserRoleRoutes = func(router *mux.Router) {
	router.HandleFunc("/v1/users/{userId}/roles", controllers.CreateUserRole).Methods("POST")
	router.HandleFunc("/v1/users/{userId}/roles", controllers.GetUserRoles).Methods("GET")
}
