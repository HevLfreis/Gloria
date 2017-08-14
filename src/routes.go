package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"strings"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler).Methods("GET").Name("index")
	router.HandleFunc("/bus", BusHandler).Methods("GET").Name("bus")
	router.HandleFunc("/dott", DotHandler).Methods("GET", "POST", "DELETE").Name("dot")
	router.HandleFunc("/test", TestHandler).Methods("GET", "POST").Name("test")
	router.NotFoundHandler = http.HandlerFunc(UnsupportedHandler)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		n := route.GetName()

		fmt.Println(strings.Join(m, ","), t)

		// init available apis
		SUPPORTED_APIS = append(SUPPORTED_APIS, Api{n, t, m})
		return nil
	})
	return router
}
