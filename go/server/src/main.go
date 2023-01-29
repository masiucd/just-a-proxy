package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi/v5"
  "github.com/masiucd/just-a-proxy/src/utils"
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

// StarWarsUser is the response from the API
type StarWarsUser struct {
  Name      string   `json:"name"`
  Height    string   `json:"height"`
  Mass      string   `json:"mass"`
  HairColor string   `json:"hair_color"`
  SkinColor string   `json:"skin_color"`
  EyeColor  string   `json:"eye_color"`
  BirthYear string   `json:"birth_year"`
  Gender    string   `json:"gender"`
  HomeWorld string   `json:"homeworld"`
  Films     []string `json:"films"`
  Species   []string `json:"species"`
  Vehicles  []string `json:"vehicles"`
  Starships []string `json:"starships"`
  Created   string   `json:"created"`
  Edited    string   `json:"edited"`
  URL       string   `json:"url"`
}

func main() {

  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/api/posts", getUsers)
  r.Get("/api/me", getMe)
  r.Get("/api/starwars/users/{user}", getStarWarsUser)

  http.ListenAndServe(":8080", r)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
  if err != nil {
    log.Fatalln(err)
  }
  // close the body when the function returns
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

func getStarWarsUser(w http.ResponseWriter, r *http.Request) {
  slug := chi.URLParam(r, "user")
  userSlugID := utils.GetUserSlugID(slug)

  // resp, err := http.Get("https://swapi.dev/api/people/1")
  resp, err := http.Get(fmt.Sprintf("https://swapi.dev/api/people/%s", userSlugID))
  if err != nil {
    log.Fatalln(err)
  }
  // close the body when the function returns
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  var starWarsUser StarWarsUser
  if err := json.Unmarshal(body, &starWarsUser); err != nil {
    log.Println(err)
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(starWarsUser)
}
