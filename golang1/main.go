// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START cloudrun_helloworld_service]

// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"log"
    "html/template"
	"net/http"
	"os"
)

type Begin struct {
	Success bool
	LinkGithub string
}

type ContactDetails struct {
	Success bool
    Name    string
    State   string
    Message string
	LinkGithub string
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		link_github := os.Getenv("LINK_GITHUB")
		if link_github == "" {
			link_github = "https://github.com/italo2sanfer/"
		}
		begin := Begin{
			Success: false,
			LinkGithub: link_github,
		}		
		if r.Method != http.MethodPost {
			tmpl.Execute(w, begin)
			return
		}
		details := ContactDetails{
			Success: true,
			Name:   r.FormValue("name"),
			State: r.FormValue("state"),
			Message: r.FormValue("message"),
			LinkGithub: link_github,
		}
		tmpl.Execute(w, details)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}