package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "rest services routing with go httprouter\n")
}

func Save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "save , %s!\n", ps.ByName("id"))
}

func Query(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "query , %s!\n", ps.ByName("qp"))
}

func Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "find , %s!\n", ps.ByName("id"))
}

func main() {
	router := httprouter.New()
	// router.GET("/", Index)

	router.POST("/user/save/:id", Save)
	router.GET("/user/query/:qp", Query)
	router.GET("/user/find/:id", Find)

	router.NotFound = http.FileServer(http.Dir("public"))

	log.Fatal(http.ListenAndServe(":3001", router))
}
