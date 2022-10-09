package main

import (
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/server"
	"log"
)

//collection.Count()
//collection.Insert
//

func main() {

	myServer := server.NewServer(":8080")

	fmt.Println("Server is listening 8080")
	err := myServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
