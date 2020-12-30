package backlogmanager

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func start(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "Start!")
}

func handleRoute(htmlFile string) func(http.ResponseWriter, *http.Request) {

	return func(wr http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(htmlFile)

		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error: ", err)
		}

		err = t.Execute(wr, nil)
	}
}

// func backlog(wr http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("backlog-list.html")

// 	if err != nil {
// 		http.Error(wr, err.Error(), http.StatusBadRequest)
// 		log.Print("Template parsing error: ", err)
// 	}

// 	err = t.Execute(wr, nil)
// }

// Handlers func
func Handlers(htmlFile string) {
	http.HandleFunc("/", start)
	http.HandleFunc("/backlog-list", handleRoute(htmlFile))

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
