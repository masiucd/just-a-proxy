package main

import (
  "net/http"

  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi/v5"
)

func main() {

  r := chi.NewRouter()
  r.Use(middleware.Logger)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
  })

  http.ListenAndServe(":8080", r)

}
