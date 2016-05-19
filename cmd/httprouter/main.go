package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getIndex(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	resp.Write([]byte("Hello from an httprouter!"))
}

func main() {
	router := httprouter.New()
	router.GET("/", getIndex)

	http.ListenAndServe(":8081", router)
}
