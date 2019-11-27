package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

/*
 Simple server to host my portfolio site
 Capabilities:
		- host simple static pages with images
		- design in context for 127 class

 Stretch Goals:
		- Implement "contact me" functionality
		- Track who visits site, collect / harvest
			 relevant data
		- Design an API for other side projects
		- Put current RELEVANT projects on site
		- (VERY STRETCH) Short, quick blog posts
			(like things I learn about golang / swift along the way)

	Outstanding Problems:
		-
*/

// Needs to be publically accessible for now?
// Investigate further if this is good practice
var templates *template.Template
var workingDir string

func main() {
	initTemplates()

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/projects", projects)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8000", nil)
}

// Handlers
func home(w http.ResponseWriter, r *http.Request) {
	var finalPath = workingDir + "index.html"
	var err = templates.ExecuteTemplate(w, finalPath, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	var err = templates.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	var err = templates.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return 
	}

}
func projects(w http.ResponseWriter, r *http.Request) {
	var err = templates.ExecuteTemplate(w, "projects.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

// Helpers
func initTemplates() {
	var workingDir, err = os.Getwd()
	if err != nil {
		log.Println("Got directory Error:", err)
	}
	var finalDir = workingDir + "/public/templates/*.html"

	templates = template.Must(template.ParseGlob(finalDir))	
}
