package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
    "os"

	"github.com/rs/cors"
)

type Post struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Email string `json:"email"`
    Description string `json:"description"`
}

var posts = []Post{}

func newPost(w http.ResponseWriter, r *http.Request) {
    var p Post

    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    idx := slices.IndexFunc(posts, func(nwP Post) bool {return nwP.Email == p.Email})
    w.Header().Set("Content-Type", "application/json")
    if idx == -1 {
        posts = append(posts, p)
        json.NewEncoder(w).Encode(p)
    } else {
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode("email already exits")
    }
}

func getPosts(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(posts)
}

func main() {
    mux := http.NewServeMux()
    port := os.Getenv("POL_BACK_PORT")

    if port == "" {
        panic("Couldn't find env variable POL_BACK_PORT")
    }

    mux.HandleFunc("/post", newPost)
    mux.HandleFunc("/posts", getPosts)
    
    handler := cors.Default().Handler(mux)
    fmt.Printf("Listening on port %s...", port)
    http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
