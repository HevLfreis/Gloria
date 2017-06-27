package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
	"strings"
)

var (
	MONGO *MongoStore
)

func main() {
	MONGO = NewMongoStore()
	defer MONGO.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/bus", BusHandler).Methods("GET")
	router.HandleFunc("/dott", DotHandler).Methods("GET", "POST", "DELETE")
	router.HandleFunc("/bg", BgHandler).Methods("GET")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(m, ","), t)
		return nil
	})
	http.ListenAndServe(":8080", router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	SetResponse(w, "json")

	json.NewEncoder(w).Encode(Response{Message: "welcome to seeleit api"})
}

func BgHandler(w http.ResponseWriter, r *http.Request) {
	f, err := exec.Command("/root/anaconda2/bin/python", "/home/hevlfreis/projects/iBot/src/main.py").CombinedOutput()
	if err != nil {

	}
	fmt.Fprintf(w, string(f))
}
