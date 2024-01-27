package handler

import (
	"around/model"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved one post from request")

	decoder := json.NewDecoder(r.Body)

	var p model.Post

	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved one request for search")
	w.Header().Set("Content_type", "application/json")

	user := r.URL.Query().Get("user")
	keywords := r.URL.Query().Get("keywords")

	var posts []model.Post
	var err error

	if user != "" {
		posts, err = service.searchPostsByUser(user)
	} else {
		posts, err = service.searchPostsByKeyWords(keywords)
	}

}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/upload",
		http.HandlerFunc(uploadHandler)).Methods("POST")
	return router
}
