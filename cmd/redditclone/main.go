package main

import (
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/Julia1505/RedditCloneBack/pkg/handlers"
)

var tpl = template.Must(template.ParseFiles("./static/html/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Println("err: %w", err)
	}
	log.Println("OK")
}

func NewServer(port string) http.Server {
	postStorage := post.NewPostsStorage()
	post := post.NewPost("ds", "music", "dsadsad")
	post.Type = "text"
	post.Text = "sdfghjkl"
	postStorage.AddPost(post)
	postHandlers := &handlers.PostHandler{
		PostStorage: postStorage,
	}

	siteMux := mux.NewRouter()
	siteMux.PathPrefix("/static/").Handler(handlers.StaticHandler)
	siteMux.HandleFunc("/", indexHandler)

	//siteMux.HandleFunc("/api/register", )
	//siteMux.HandleFunc("/api/login",)
	siteMux.HandleFunc("/api/posts/", postHandlers.List).Methods("GET")
	siteMux.HandleFunc("/api/posts/", postHandlers.AddPost).Methods("POST")
	siteMux.HandleFunc("/api/posts/{category_name}", postHandlers.CategoryList).Methods("GET")
	siteMux.HandleFunc("/api/post/{post_id}", postHandlers.Post).Methods("GET")
	//siteMux.HandleFunc("/api/post/{post_id}", postHandlers.).Methods("POST")
	//siteMux.HandleFunc("/api/post/{post_id}/{comment_id}").Methods("DELETE")
	//siteMux.HandleFunc("/api/post/{post_id}/upvote", postHandlers.).Methods("GET")
	//siteMux.HandleFunc("/api/post/{post_id}/downvote", postHandlers.).Methods("GET")
	//siteMux.HandleFunc("/api/post/{post_id}", postHandlers.).Methods("DELETE")
	//siteMux.HandleFunc("/api/user/{user_login}", postHandlers.).Methods("GET")

	return http.Server{
		Addr:         port,
		Handler:      siteMux,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}
}

func main() {

	server := NewServer(":8080")

	fmt.Println("Server is listening 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
