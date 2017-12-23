package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/portfolio", portfolio)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/resume", resume)
	http.HandleFunc("/resume.pdf", resumePDF)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index")
}
func about(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "about")
}
func portfolio(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "portfolio")
}
func contact(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "contact")
}
func resume(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "resume")
}
func resumePDF(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "resume.pdf")
}
