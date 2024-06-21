package functionshendler

import (
	"fmt"
	"html/template"
	"net/http"
)

func viewSpecificLogsOrLog(w http.ResponseWriter, r *http.Request) {
	dates := r.FormValue("logsdate")
	types := r.FormValue("logstype")
	names := r.FormValue("usersname")
	logs, err := outPutAllLogs()
	if err != nil {
		fmt.Println(err)
	}

	var NewLogs []Log
	if dates != "" && types != "" && names != "" {
		for _, log := range logs {
			if log.User == names && log.Date == dates && log.Type == types {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if dates != "" && types != "" {
		for _, log := range logs {
			if log.Date == dates && log.Type == types {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if types != "" && names != "" {
		for _, log := range logs {
			if log.User == names && log.Type == types {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if dates != "" && names != "" {
		for _, log := range logs {
			if log.User == names && log.Date == dates {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if dates != "" {
		for _, log := range logs {
			if log.Date == dates {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if types != "" {
		for _, log := range logs {
			if log.Type == types {
				NewLogs = append(NewLogs, log)
			}
		}
	} else if names != "" {
		for _, log := range logs {
			if log.User == names {
				NewLogs = append(NewLogs, log)
			}
		}
	} else {
		for _, log := range logs {
			NewLogs = append(NewLogs, log)
		}
	}

	if len(NewLogs) == 0 {
		tmpl, err := template.ParseFiles("templates/notFoundLogs.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		tmpl.ExecuteTemplate(w, "notFoundLogs", nil)
	} else {
		tmpl, err := template.ParseFiles("templates/viewSpecificLogs.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		web_data := Show{NewLogs}
		tmpl.ExecuteTemplate(w, "viewSpecificLogs", web_data)
	}

}
