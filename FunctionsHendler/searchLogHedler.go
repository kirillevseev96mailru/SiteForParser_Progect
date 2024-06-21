package functionshendler

import (
	"fmt"
	"html/template"
	"net/http"
)

func searchLogHedler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/search.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "search", nil)
}
