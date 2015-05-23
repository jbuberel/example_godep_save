package main


import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func HomeHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, godep save user")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.ListenAndServe(":8080", r)

}
