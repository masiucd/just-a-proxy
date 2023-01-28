package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi/v5"
)

// Post is the response from the API
type Post struct {
  UserID int    `json:"userId"`
  ID     int    `json:"id"`
  Title  string `json:"title"`
  Body   string `json:"body"`
}

// Me is the response from the API
type Me struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
}

func main() {

  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/users", getUsers)
  r.Get("/me", getMe)

  http.ListenAndServe(":8080", r)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    http.Error(w, "Error while fetching posts", resp.StatusCode)
    return
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  var post []Post
  if err := json.Unmarshal(body, &post); err != nil {
    log.Println(err)
  }

  // send Post as json
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(post)

}

func getMe(w http.ResponseWriter, r *http.Request) {
  me := Me{Name: "Ibra", Age: 42}

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(me)
}
