package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/portfolio", portfolio)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/resume", resume)
	http.HandleFunc("/resume.pdf", resume)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func portfolio(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "portfolio.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func blog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "bloghome.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func resume(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/resume.pdf", http.StatusPermanentRedirect)
}
