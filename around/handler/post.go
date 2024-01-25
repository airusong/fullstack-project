package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"around/model"

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

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/upload",
		http.HandlerFunc(uploadHandler)).Methods("POST")
	return router
}
