package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// User type
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Todo struct type
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Album struct
type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

// Photo struct
type Photo struct {
	AlbumID int    `json:"albumId"`
	ID      int    `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url`
}

// Comment struct
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// URLConst string
const URLConst = "https://jsonplaceholder.typicode.com/"

// App to init the application
func App() {

	endPointXs := []string{"users", "todos", "images", "albums", "comments"}

	fmt.Println("enter a endpoint like for example")

	for _, v := range endPointXs {
		fmt.Println(v)
	}

	var endPoint string
	fmt.Scanf("%s", &endPoint)

	resp, err := http.Get(URLConst + endPoint)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var users []User
	var todos []Todo
	var albums []Album
	var photos []Photo
	var comments []Comment

	switch strings.TrimSpace(endPoint) {
	case "users":
		err = json.Unmarshal(body, &users)
		fmt.Println(users)
	case "todos":
		err = json.Unmarshal(body, &todos)
		fmt.Println(todos)
	case "albums":
		err = json.Unmarshal(body, &albums)
		fmt.Println(albums)
	case "photos":
		err = json.Unmarshal(body, &photos)
		fmt.Println(photos)
	case "comments":
		err = json.Unmarshal(body, &comments)
		fmt.Println(comments)
	default:
		fmt.Println("there was no match!")

	}

}
