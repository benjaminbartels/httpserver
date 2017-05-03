package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/*name", index)

	log.Fatal(http.ListenAndServe(":3333", router))

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Printf(r.RequestURI)
}
