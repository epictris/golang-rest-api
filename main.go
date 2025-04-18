package main

import (
	"net/http"

	"tris.sh/go/api/base"
	"tris.sh/go/api/routes"
)

func main() {
	base.RegisterRoute("/test", routes.GetUser)
	http.ListenAndServe(":8080", nil)
}
