package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tsaridoor/gpio"
	"tsaridoor/middleware"
)

func main() {
	log.Println("starting web service")

	err := gpio.Setup()
	if err != nil {
		log.Fatal(fmt.Sprint("unable to setup gpio", err.Error()))
	}

	defer func() {
		gpio.Close()
	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/unlock", middleware.LogRequestWrapper(middleware.BasicAuth(unlockHandler)))
	http.Handle("/", middleware.LogRequestWrapper(middleware.BasicAuth(homeHandler)))
	log.Fatal(http.ListenAndServe(":80", nil))
}

func unlockHandler(w http.ResponseWriter, _ *http.Request) {
	gpio.Unlock()
	_, err := fmt.Fprint(w, "OK")
	if err != nil {
		log.Printf("Fprint err %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "home", &Page{Title: "Tsaridoor"})
}

var templates = template.Must(template.ParseFiles("./home.html"))

type Page struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
