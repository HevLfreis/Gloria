package main

import (
	"fmt"
	"net/http"
	"strings"
	"flag"
	"github.com/gorilla/mux"
	"github.com/inconshreveable/log15"
	"path/filepath"
	//"strconv"
)

var (
	MONGO *MongoStore
	log log15.Logger
	API []Api
)

func main() {

	// parse args
	port := flag.String("port", "8080", "http listen port")
	mongoUrl := flag.String("mongo", "mongodb://127.0.0.1", "mongodb host")
	flag.Parse()

	// init mongo
	MONGO = NewMongoStore(*mongoUrl)
	defer MONGO.Close()

	// config routers
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET").Name("index")
	router.HandleFunc("/bus", BusHandler).Methods("GET").Name("bus")
	router.HandleFunc("/dott", DotHandler).Methods("GET", "POST", "DELETE").Name("dot")
	//router.HandleFunc("/bg", BgHandler).Methods("GET")
	router.HandleFunc("/test", TestHandler).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(UnsupportedHandler)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		n := route.GetName()

		fmt.Println(strings.Join(m, ","), t)

		// init available apis
		API = append(API, Api{n, t, m})
		return nil
	})

	// set log15
	logPath := "log"
	logFile := "gloria.log"

	log = log15.New("module", "api/server")
	if err := DirExistedOrCreate(logPath); err != nil {
		log.Warn("log path create failed")
		return
	}
	handler, _ := log15.FileHandler(filepath.Join(logPath, logFile), log15.TerminalFormat())
	log.SetHandler(handler)
	log.Info("start logging", "path", filepath.Join(logPath, logFile))

	// start server
	log.Info("start api server", "port", *port)
	http.ListenAndServe(":"+*port, router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	results := make([]interface{}, len(API))
	for i, api := range API {
		results[i] = api
	}

	SendJsonResponse(w, STATUS_SUCCESS, "welcome to seeleit api", results)
}

func UnsupportedHandler(w http.ResponseWriter, r *http.Request) {
	results := make([]interface{}, len(API))
	for i, api := range API {
		results[i] = api
	}

	SendJsonResponse(w, STATUS_UNSUPPORTED, "unsupported api", results)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ParseForm(r, map[string]interface{}{"float1": 0.0, "string1": "string", "bool1": nil})

	if err != nil {
		log.Warn("parse arg failed", "err", err)
		return
	}

	fmt.Fprintf(w, "hello %s", "world")
}

//func BgHandler(w http.ResponseWriter, r *http.Request) {
//	f, err := exec.Command("/root/anaconda2/bin/python", "/home/hevlfreis/projects/iBot/src/main.py").CombinedOutput()
//	if err != nil {
//
//	}
//	fmt.Fprintf(w, string(f))
//}
