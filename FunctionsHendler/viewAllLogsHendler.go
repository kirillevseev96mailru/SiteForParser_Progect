package functionshendler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func outPutAllLogs() ([]Log, error) {
	file, err := ioutil.ReadFile("New-logs.json")
	if err != nil {
		return nil, err
	}

	var logs []Log
	err = json.Unmarshal(file, &logs)
	if err != nil {
		log.Fatal(err)
	}

	return logs, nil
}

type Log struct {
	Date    string `json:"date"`
	User    string `json:"user"`
	Command string `json:"command"`
	Type    string `json:"type"`
}

type Show struct {
	Pages []Log
}

func viewAllLogsHendler(w http.ResponseWriter, r *http.Request) {
	logs, err := outPutAllLogs()
	if err != nil {
		fmt.Println(err)
	}

	web_data := Show{logs}

	tmpl, err := template.ParseFiles("templates/viewAllLogs.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "viewAllLogs", web_data)
}
