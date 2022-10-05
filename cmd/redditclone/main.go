package main

import (
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/middleware"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"github.com/Julia1505/RedditCloneBack/pkg/user"
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
	userStorage := user.NewUsersStorage()
	userHandlers := &handlers.UserHandler{
		UserStorage: userStorage,
	}

	postStorage := post.NewPostsStorage()
	post := &post.Post{Author: post.Author{Id: "2", Username: "dfs"}, Title: "dfghjk"}
	post.Type = "text"
	post.Category = "music"
	post.Text = "sdfghjkl"
	postStorage.AddPost(post)
	postHandlers := &handlers.PostHandler{
		PostStorage: postStorage,
	}

	mux := mux.NewRouter().StrictSlash(true)
	mux.PathPrefix("/static/").Handler(handlers.StaticHandler)
	mux.HandleFunc("/", indexHandler)

	apiRouter := mux.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/register", userHandlers.SignUp).Methods("POST")
	apiRouter.HandleFunc("/login", userHandlers.SignIn).Methods("POST")

	apiRouter.HandleFunc("/posts/", postHandlers.List).Methods("GET")
	apiRouter.HandleFunc("/posts", postHandlers.AddPost).Methods("POST")
	apiRouter.HandleFunc("/posts/{category_name}", postHandlers.CategoryList).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}", postHandlers.Post).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}", postHandlers.AddComment).Methods("POST")
	apiRouter.HandleFunc("/post/{post_id}/{comment_id}", postHandlers.DeleteComment).Methods("DELETE")
	apiRouter.HandleFunc("/post/{post_id}/upvote", postHandlers.UpVote).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}/downvote", postHandlers.DownVote).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}/unvote", postHandlers.UnVote).Methods("GET")
	apiRouter.HandleFunc("/post/{post_id}", postHandlers.DeletePost).Methods("DELETE")
	apiRouter.HandleFunc("/user/{user_login}", postHandlers.UserList).Methods("GET")

	allSiteMux := middleware.IsAuthorized(userStorage, mux)
	allSiteMux = middleware.Logging(allSiteMux)
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
