package main

import (
	backlogmanager "github.com/masiuciszek/Gophers/backlog-manager"
)

func main() {
	// champs := champions.Champs("champs.json")

	// routes := routes.HandlerFn()

	// for i, v := range routes {
	// 	switch i {
	// 	case 0:
	// 		http.HandleFunc("/", v)
	// 	case 1:
	// 		http.HandleFunc("/about", v)
	// 	case 2:
	// 		http.HandleFunc("/test", v)
	// 	}
	// }

	// fmt.Println("Serer is on port :9000")
	// in this case Log.fatal is a like a backup wrapper, os if something goes wrong handle the error
	// log.Fatal(http.ListenAndServe(":9000", nil))
	// htmlList := []string{"./public/backlog-list.html", "/add-todo" }
	htmlList := make(map[string]string)
		
	htmlList["start"]	= "./public/index.html"
	htmlList["backlog-list"]	= "./public/backlog-list.html"
	htmlList["add-todo"]	= "./public/add-todo.html"

	backlogmanager.Handlers(htmlList)

}
