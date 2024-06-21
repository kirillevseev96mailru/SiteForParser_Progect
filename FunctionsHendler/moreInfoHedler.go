package functionshendler

import (
	"fmt"
	"html/template"
	"net/http"
)

func moreInfoHedler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/moreInfo.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "moreInfo", nil)
}
