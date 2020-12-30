package backlogmanager

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
func start(htmlFile string) func(http.ResponseWriter, *http.Request) {
	// fmt.Fprintf(wr, "Start!")

	return func(wr http.ResponseWriter, r *http.Request) {
		t,err := template.ParseFiles(htmlFile)

		if err != nil {
			http.Error(wr,err.Error(),http.StatusBadRequest)
			log.Print("Template parsing error", err)
		}

		err = t.Execute(wr,nil)
	}
}



func backlogList(htmlFile string) func(http.ResponseWriter, *http.Request) {
	return func(wr http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(htmlFile)

		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error: ", err)
		}

		err = t.Execute(wr, nil)
	}
}


// Handlers func
func Handlers(htmlList map[string]string) {
	http.HandleFunc("/", start(htmlList["start"]))
	http.HandleFunc("/backlog-list", backlogList(htmlList["backlog-list"]))

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
