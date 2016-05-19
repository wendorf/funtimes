package main

import (
	"net/http"

	"github.com/tedsuo/rata"
)

func getIndex(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello from a rata!"))
}

func main() {
	routes := rata.Routes{
		{Name: "get_index", Method: "GET", Path: "/"},
	}

	handlers := rata.Handlers{
		"get_index": http.HandlerFunc(getIndex),
	}

	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", router)
}
