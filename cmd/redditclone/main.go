package main

import (
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/middleware"
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

	mux := mux.NewRouter()
	mux.PathPrefix("/static/").Handler(handlers.StaticHandler)
	mux.HandleFunc("/", indexHandler)

	apiRouter := mux.PathPrefix("/api").Subrouter()

	//siteMux.HandleFunc("/register", )
	//siteMux.HandleFunc("/login",)
	apiRouter.HandleFunc("/posts/", postHandlers.List).Methods("GET")
	apiRouter.HandleFunc("/posts/", postHandlers.AddPost).Methods("POST")
	apiRouter.HandleFunc("/posts/{category_name}", postHandlers.CategoryList).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}", postHandlers.Post).Methods("GET")
	//siteMux.HandleFunc("/post/{post_id}", postHandlers.).Methods("POST")
	//siteMux.HandleFunc("/post/{post_id}/{comment_id}").Methods("DELETE")
	//siteMux.HandleFunc("/post/{post_id}/upvote", postHandlers.).Methods("GET")
	//siteMux.HandleFunc("/post/{post_id}/downvote", postHandlers.).Methods("GET")
	//siteMux.HandleFunc("/post/{post_id}", postHandlers.).Methods("DELETE")
	//siteMux.HandleFunc("/user/{user_login}", postHandlers.).Methods("GET")

	allSiteMux := middleware.Logging(mux)
	allSiteMux = middleware.PanicRecovery(allSiteMux)

	return http.Server{
		Addr:         port,
		Handler:      allSiteMux,
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
