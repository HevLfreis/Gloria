package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Dot struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string `json:"name"`
	Location  Location `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

type Location struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}

func NewDot(name string, cords [2]float32, timestamp time.Time) *Dot {
	loc := Location{"Point", cords}
	dot := Dot{Name: name, Location: loc, Timestamp: timestamp}
	return &dot
}

func DotHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		DotGetHandler(w, r)
	case "POST":
		DotPostHandler(w, r)
	case "DELETE":
		DotDeleteHandler(w, r)
	default:
		// Todo: change to unsupported api handler
		DotGetHandler(w, r)
	}
}

func DotGetHandler(w http.ResponseWriter, r *http.Request) {
	SetResponse(w, "json")

	var dots []Dot

	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")
	c.Find(nil).All(&dots)

	results := make([]interface{}, len(dots))
	for i, dot := range dots {
		results[i] = dot
	}

	json.NewEncoder(w).Encode(Response{0, "get dots ok", results})
}

func DotPostHandler(w http.ResponseWriter, r *http.Request) {
	SetResponse(w, "json")

	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")

	r.ParseForm()

	lat, err1 := strconv.ParseFloat(r.PostFormValue("lat"), 32)
	lng, err2 := strconv.ParseFloat(r.PostFormValue("lng"), 32)
	if err1 != nil || err2 != nil {
		json.NewEncoder(w).Encode(Response{Status: 1, Message: "post dots failed, invalid value"})
		return
	}

	name := r.PostFormValue("name")
	if name == "" {
		name = "Unknown"
	}

	loc := [2]float32{float32(lat), float32(lng)}
	dot := NewDot(name, loc, time.Now())
	err := c.Insert(&dot)

	if err != nil {
		return
	}

	var dots []Dot
	c.Find(nil).All(&dots)

	results := make([]interface{}, len(dots))
	for i, dot := range dots {
		results[i] = dot
	}

	json.NewEncoder(w).Encode(Response{0, "post dots ok", results})
}

func DotDeleteHandler(w http.ResponseWriter, r *http.Request) {
	SetResponse(w, "json")

	var dot Dot
	
	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")
	c.Find(nil).Sort("-_id").Limit(1).One(&dot)
	c.Remove(bson.M{"_id": dot.Id})

	results := make([]interface{}, 0)
	results = append(results, dot)

	json.NewEncoder(w).Encode(Response{0, "delete dot ok", results})
}
