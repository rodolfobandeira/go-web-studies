package main

import (
	"net/http"

	"github.com/rodolfobandeira/go-web-studies/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
