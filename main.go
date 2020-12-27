package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/masiuciszek/Gophers/routes"
)




func main() {
	// champs := champions.Champs("champs.json")
	
	routes := routes.HandlerFn()

	for i, v := range routes {
		switch i {
		case 0:
			http.HandleFunc("/",v)
		case 1:
			http.HandleFunc("/about",v)
		}
	}

	
	
	fmt.Println("Serer is on port :9000")
	// in this case Log.fatal is a like a backup wrapper, os if something goes wrong handle the error
	log.Fatal(http.ListenAndServe(":9000",nil))

}
