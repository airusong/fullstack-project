package handler

import (
	"around/model"
	"around/service"

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
		posts, err = service.SearchPostsByUser(user)
	} else {
		posts, err = service.SearchPostsByKeywords(keywords)
	}

	if err != nil {
		http.Error(w, "Failed to read post from backend",
			http.StatusInternalServerError)
		fmt.Printf("Failed to read post from backend %v.\n", err)

		return
	}
	js, err := json.Marshal(posts)
	if err != nil {

		http.Error(w, "Failed to parse posts into JSON format",
			http.StatusInternalServerError)

		fmt.Printf("Failed to parse posts into JSON format %v.\n", err)

		return
	}
	w.Write(js)

}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/upload",
		http.HandlerFunc(uploadHandler)).Methods("POST")
	router.Handle("/search",
		http.HandlerFunc(searchHandler)).Methods("GET")
	return router
}
