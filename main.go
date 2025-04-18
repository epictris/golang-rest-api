package main

import (
	"net/http"

	"tris.sh/go/api/base"
	"tris.sh/go/api/routes"
	"tris.sh/go/db"
)

func main() {
	database := db.Init()
	defer database.Close()
	base.RegisterRoute("/test", database, routes.GetUser)
	base.RegisterRoute("/api/create_user", database, routes.CreateUser)
	http.ListenAndServe(":8080", nil)
}
