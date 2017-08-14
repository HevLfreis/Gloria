package main

import (
	"net/http"
	"flag"
	"github.com/inconshreveable/log15"
)

var (
	MONGO *MongoStore
	log log15.Logger
	SUPPORTED_APIS []Api
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
	router := InitRouter()

	// init logger
	InitLogger()

	// start server
	log.Info("start api server", "port", *port)
	http.ListenAndServe(":"+*port, router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	results := make([]interface{}, len(SUPPORTED_APIS))
	for i, api := range SUPPORTED_APIS {
		results[i] = api
	}

	SendJsonResponse(w, STATUS_SUCCESS, "welcome to seeleit api", results)
}

func UnsupportedHandler(w http.ResponseWriter, r *http.Request) {
	results := make([]interface{}, len(SUPPORTED_APIS))
	for i, api := range SUPPORTED_APIS {
		results[i] = api
	}

	SendJsonResponse(w, STATUS_UNSUPPORTED, "unsupported api", results)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ParseArgs(r, map[string]interface{}{"float": 0.0, "string": "string", "bool": false})

	if err != nil {
		log.Warn("parse arg failed", "err", err)
		SendJsonResponse(w, STATUS_ERR, "receive params failed", nil)
		return
	}

	result := map[string]interface{}{
		"float": data["float"].(float64),
		"string": data["string"].(string),
		"bool": data["bool"].(bool),
	}

	SendJsonResponse(w, STATUS_SUCCESS, "receive params", []interface{}{result})
}