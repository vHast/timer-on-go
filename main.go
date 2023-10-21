package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Create a struct that holds information to be dispalyed in our HTML file

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	// We tell Go where we can find our html file, we ask Go to parse the html file, we wrap it in a call to template.Must() which handles any errors
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	// Here we tell Go to create a handle that looks in the static directory, go then uses the /static/ as a url that our html can refer to when looking for our css and other files

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) // Go looks in the relative static directory first, then matches it to a url of our choice as shown in httt.Handle("/static")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Takes the name from the URL query ?name=Martin, will set welcome.Name = Martin
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		// * You can change the name in the url
		// localhost:8080/?name=Martin

		//If errors show an internal server message

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Server on port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
