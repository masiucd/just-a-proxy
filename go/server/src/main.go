package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"

  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi/v5"
  "github.com/masiucd/just-a-proxy/src/models"
  "github.com/masiucd/just-a-proxy/src/utils"
)

// AllUsers is the response from the API
type AllUsers struct {
  Success  bool                  `json:"success"` // custom field
  Count    int                   `json:"count"`
  Next     string                `json:"next"`
  Previous string                `json:"previous"`
  Users    []models.StarWarsUser `json:"results"`
}

func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/api/me", me)
  r.Get("/api/starwars/users", starWarsUsers)
  r.Get("/api/starwars/users/{user}", starWarsUser)

  fmt.Println("Server is running on port 8080 🚀")
  http.ListenAndServe(":8080", r)

}

func me(w http.ResponseWriter, r *http.Request) {
  me := models.Me{Name: "Ibra", Age: 42}
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(me)
}

func starWarsUsers(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("https://swapi.dev/api/people")

  // get query params
  // q := r.URL.Query()
  // page := q.Get("page")
  // if page != "" {
  //   resp, err = http.Get(fmt.Sprintf("https://swapi.dev/api/people?page=%s", page))
  // }

  if err != nil {
    log.Fatalln(err)
  }

  // close the body when the function returns
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  var allUsers AllUsers
  if err := json.Unmarshal(body, &allUsers); err != nil {
    log.Println(err)
  }
  allUsers.Success = true
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(allUsers)
}

func starWarsUser(w http.ResponseWriter, r *http.Request) {
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

  var starWarsUser models.StarWarsUser
  if err := json.Unmarshal(body, &starWarsUser); err != nil {
    log.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(starWarsUser)
}
