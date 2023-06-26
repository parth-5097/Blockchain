// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the application
	app := App{}
	app.Initialize()

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// app.go

package main

import (
	"html/template"
	"log"
	"net/http"
)

type App struct {
	Router *http.ServeMux
}

func (app *App) Initialize() {
	app.Router = http.NewServeMux()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/", app.homeHandler)
	app.Router.HandleFunc("/submit", app.submitHandler)
	// Add more routes as needed
}

func (app *App) renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/"+templateName)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the home page template
		app.renderTemplate(w, r, "home.html", nil)
	} else {
		// Handle invalid request method
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (app *App) submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Process form submission
		r.ParseForm()
		name := r.Form.Get("name")
		// Perform KYC/AML checks and reporting requirements
		// ...
		// Handle the response
		app.renderTemplate(w, r, "success.html", name)
	} else {
		// Handle invalid request method
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
