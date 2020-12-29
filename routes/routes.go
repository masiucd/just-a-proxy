package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/masiuciszek/Gophers/champions"
)

type ReqType struct {
	Test string
}

func getUsers(fileName string)[]champions.Champion {
	champs := champions.Champs(fileName)

	return champs
}

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

func parseGhPost(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var t ReqType
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Test)
}



// HandlerFn function
func HandlerFn() []Route {
	routes := []Route{start,about,parseGhPost}

	return routes
}