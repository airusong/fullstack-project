package handler

import (
	"around/model"
	"around/service"

	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

var (
	mediaTypes = map[string]string{
		".jpeg": "image",
		".jpg":  "image",
		".gif":  "image",
		".png":  "image",
		".mov":  "video",
		".mp4":  "video",
		".avi":  "video",
		".flv":  "video",
		".wmv":  "video",
	}
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved one post from request")

	//decoder := json.NewDecoder(r.Body)

	p := model.Post{
		Id:      uuid.New(),
		User:    r.FormValue("user"),
		Message: r.FormValue("message"),
	}
	file, header, err := r.FormFile("media_file")
	if err != nil {
		http.Error(w, "Media file is not available", http.StatusBadRequest)
		fmt.Printf("Media file is not available %v\n", err)
		return
	}
	suffix := filepath.Ext(header.Filename)

	if t, ok := mediaTypes[suffix]; ok {
		p.Type = t
	} else {
		p.Type = "Unknown"
	}

	err = service.SavePost(&p, file)

	if err != nil {
		http.Error(w, "Failed to save post to backend", http.StatusBadRequest)
		fmt.Printf("Failed to save post to backend %v\n", err)
		return
	}

	fmt.Println("post is save successfully")

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
