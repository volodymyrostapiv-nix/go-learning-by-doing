package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const INDEX = "No keyword specified"

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/:target", Target)

	log.Fatal(http.ListenAndServe(":1234", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(INDEX))
}

func Target(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	target := ps.ByName("target")
	http.Redirect(w, r, "https://www.google.com/search?q="+target, http.StatusSeeOther)
}
