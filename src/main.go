package main

import (
	"log"
	"net/http"
	"io/ioutil"
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

func main() {
	var fileServer = http.FileServer(http.Dir("templates"))
	http.Handle("/", fileServer)

	log.Println("Server up...")
	http.ListenAndServe(":8080", nil)
}

// RootHandler is the base url router handler 
type RootHandler struct { }

func (rootHandler *RootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var urlPath = req.URL.Path[1:]
	log.Println(urlPath)

	var data, err = ioutil.ReadFile(string(urlPath))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 !" + http.StatusText(404)))
	}
}

