package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi/v5"
)

// UserResponse is the response from the API
type UserResponse struct {
  UserID int    `json:"userId"`
  ID     int    `json:"id"`
  Title  string `json:"title"`
  Body   string `json:"body"`
}

func main() {

  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/users", getUsers)

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

  var userResponse []UserResponse
  if err := json.Unmarshal(body, &userResponse); err != nil {
    log.Println(err)
  }

  // send userResponse as json
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(userResponse)

}
