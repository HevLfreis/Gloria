package main

import (
	"net/http"
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
		UnsupportedHandler(w, r)
	}
}

func DotGetHandler(w http.ResponseWriter, r *http.Request) {

	var dots []Dot

	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")
	c.Find(nil).All(&dots)

	results := make([]interface{}, len(dots))
	for i, dot := range dots {
		results[i] = dot
	}

	SendJsonResponse(w, STATUS_SUCCESS, "get dots ok", results)
}

func DotPostHandler(w http.ResponseWriter, r *http.Request) {

	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")

	data, err := ParseForm(r, map[string]interface{}{"lat": 0.0, "lng": 0.0, "name": "Unknown"})

	if err != nil {
		SendJsonResponse(w, STATUS_ERR, "post dots failed", nil)
		log.Warn("parse args failed", "err", err)
		return
	}

	lat := data["lat"].(float64)
	lng := data["lng"].(float64)
	name := data["name"].(string)

	loc := [2]float32{float32(lat), float32(lng)}
	dot := NewDot(name, loc, time.Now())
	if err := c.Insert(&dot); err != nil {
		SendJsonResponse(w, STATUS_ERR, "post dots failed", nil)
		return
	}

	var dots []Dot
	c.Find(nil).All(&dots)

	results := make([]interface{}, len(dots))
	for i, dot := range dots {
		results[i] = dot
	}

	SendJsonResponse(w, STATUS_SUCCESS, "post dots ok", results)
}

func DotDeleteHandler(w http.ResponseWriter, r *http.Request) {

	var dot Dot
	
	sess := MONGO.Session()
	defer sess.Close()
	c := sess.DB(MONGO_DBNAME).C("dot")
	c.Find(nil).Sort("-_id").Limit(1).One(&dot)
	c.Remove(bson.M{"_id": dot.Id})

	results := make([]interface{}, 0)
	results = append(results, dot)

	SendJsonResponse(w, STATUS_SUCCESS, "delete dot ok", results)
}
