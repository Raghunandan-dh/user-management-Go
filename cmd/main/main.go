package main

import (
	"log"
	"net/http"
	"simple-REST/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	routes.RegisterRoleRoutes(r)
	routes.RegisterUserRoleRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
