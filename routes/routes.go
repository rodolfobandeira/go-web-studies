package routes

import (
	"net/http"

	"github.com/rodolfobandeira/go-web-studies/controllers"
)

// LoadRoutes - Loads all Routes
func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
