package functionshendler

import (
	"fmt"
	"html/template"
	"net/http"
)

func viewOneLogHendler(w http.ResponseWriter, r *http.Request) {
	users := r.URL.Query().Get("User")
	types := r.URL.Query().Get("Type")
	dates := r.URL.Query().Get("Dates")
	command := r.URL.Query().Get("Command")
	var OurLogs = []Log{
		{
			Date:    dates,
			User:    users,
			Command: command,
			Type:    types,
		},
	}

	web_data := Show{OurLogs}

	tmpl, err := template.ParseFiles("templates/viewOneLog.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "viewOneLog", web_data)
}
