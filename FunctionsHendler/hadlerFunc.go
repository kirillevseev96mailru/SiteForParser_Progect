package functionshendler

import (
	"net/http"
)

func HadlerFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/viewAllLogs/", viewAllLogsHendler)
	http.HandleFunc("/viewOneLog/", viewOneLogHendler)
	http.HandleFunc("/search/", searchLogHedler)
	http.HandleFunc("/moreInfo/", moreInfoHedler)
	http.HandleFunc("/view_specific_logs/", viewSpecificLogsOrLog)
	http.ListenAndServe(":8080", nil)

}
