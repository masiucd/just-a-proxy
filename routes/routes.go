package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Route struct
type Route func(http.ResponseWriter, *http.Request)


func start (w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w,"hello")
}


func about (w http.ResponseWriter, r * http.Request) {
	t,err:=template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		log.Print("templating parsing error",err)
	}

	err = t.Execute(w,nil)

}

// HandlerFn function
func HandlerFn() []Route {
	routes := []Route{start,about}

	return routes
}