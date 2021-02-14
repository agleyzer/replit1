package main

import (
	"fmt"
	"net/http"
  "strconv"
  
  "github.com/replit/database-go"
)

const key = "default"

func count(w http.ResponseWriter, req *http.Request) {
  value, err := database.Get(key)
  if err == database.ErrNotFound {
    value = "0"
    err = nil
  }
  if err != nil {
    http.Error(w, fmt.Sprintf("Database get error: %v", err), http.StatusInternalServerError)
    return
  }

  count, err := strconv.Atoi(value)
  if err != nil {
    http.Error(w, fmt.Sprintf("Atoi error: %v", err), http.StatusInternalServerError)
    return
  }

  count = count + 1 

  if err := database.Set(key, strconv.Itoa(count)); err != nil {
    http.Error(w, fmt.Sprintf("Database set error: %v", err), http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, `{"count": %d}`, count)
}

func main() {
	http.HandleFunc("/", count)  
	http.ListenAndServe(":8080", nil)
}
