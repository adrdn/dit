package home

import (
	"net/http"
	"text/template"
	
	"adrdn/dit/credential"
)

var tmpl = template.Must(template.ParseGlob("forms/home/*"))

// Home revokes the home page
func Home(w http.ResponseWriter, r *http.Request) {
	ok, user := credential.CheckAuthentication(w, r)
	if !ok {
		return
	}
	tmpl.ExecuteTemplate(w, "Home", user.Username)
}
