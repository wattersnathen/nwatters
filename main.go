package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
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
	http.HandleFunc("/contact-confirmation", contactConfirmation)
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
	if req.Method == http.MethodGet {
		err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
		if err != nil {
			log.Println(err)
		}
	} else if req.Method == http.MethodPost {
		req.ParseForm()

		name := req.FormValue("name")
		from := req.FormValue("email")
		subject := req.FormValue("subject")
		body := req.FormValue("body")

		type EmailConfig struct {
			Username string
			Password string
		}

		file, _ := os.Open("./config/smtp.json")
		decoder := json.NewDecoder(file)
		emailConfig := EmailConfig{}
		err := decoder.Decode(&emailConfig)
		if err != nil {
			log.Println(err)
		}
		auth := smtp.PlainAuth(
			"",
			emailConfig.Username,
			emailConfig.Password,
			"smtp.gmail.com",
		)

		msg := "Reply-To: " + name + " <" + from + ">" + "\r\n" +
			"To: " + emailConfig.Username + "\r\n" +
			"Subject: " + subject + "\r\n" +
			body

		err = smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			from,
			[]string{emailConfig.Username},
			[]byte(msg),
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/contact-confirmation", http.StatusSeeOther)
	}
}
func contactConfirmation(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact_confirmation.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
func resume(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/resume.pdf", http.StatusPermanentRedirect)
}
