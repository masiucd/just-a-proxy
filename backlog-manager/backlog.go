package backlogmanager

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/lithammer/shortuuid"
)

func start(htmlFile string) func(http.ResponseWriter, *http.Request) {
	// fmt.Fprintf(wr, "Start!")

	return func(wr http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(htmlFile)

		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error", err)
		}

		err = t.Execute(wr, nil)
	}
}

// PageVariables struct
type PageVariables struct {
	Title    string
	TodoList []Todo
}


func getTodos() []Todo  {
	var todos, err = ReadFile("todos.json")

	if err != nil {
		log.Fatal("could not read file from main.go", err)
	}

	return todos
}


func backlogList(htmlFile string) func(http.ResponseWriter, *http.Request) {
	
		todos := getTodos()

	pageVariables := PageVariables{
		Title:    "Backlog App With GO",
		TodoList: todos,
	}
	

	return func(wr http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(htmlFile)

		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error: ", err)
		}

		err = t.Execute(wr, pageVariables)
	}
}

func addTodo(htmlFile string) func(http.ResponseWriter, *http.Request)  {


	return func(wr http.ResponseWriter, r *http.Request) {
		todos := getTodos()

		err := r.ParseForm()
		
		if err != nil {
			http.Error(wr,err.Error(),http.StatusBadRequest)
			log.Print("Request parsing error",err)
		}


		todo := Todo{
			ID:  shortuuid.New(),
			Text: r.FormValue("title"),
			About: r.FormValue("text"),
			Completed: false,
		}

		todos = append(todos, todo)

		
		WriteFile(todos)

		http.Redirect(wr,r,"/backlog-list",http.StatusSeeOther)

	}

}

// Handlers func
func Handlers(htmlList map[string]string) {

	http.HandleFunc("/", start(htmlList["start"]))
	http.HandleFunc("/backlog-list", backlogList(htmlList["backlog-list"]))
	http.HandleFunc("/add-todo", addTodo(htmlList["add-todo"]) )

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
