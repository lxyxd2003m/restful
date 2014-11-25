package main

import (
	"fmt"
	"github.com/drone/routes"
	"log"
	"github.com/lxyxd2003m/restful"
	"net/http"
)

func main() {
	fmt.Println("hello restful")
	mux := routes.New()
	mux.Get("/user/:uid", restful.Getuser)
	mux.Post("/user/:uid", restful.Modifyuser)
	mux.Del("/user/:uid", restful.Deleteuser)
	mux.Put("/user/", restful.Adduser)
	http.Handle("/", mux)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal("ListenServer error", nil)
	}

}
