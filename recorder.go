package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type RecordedDevice struct {
	Id          string
	Humidity    float64
	Temperature float64
	LastReport  time.Time
}

type Recorder struct {
	Store map[string]*RecordedDevice
}

func NewRecorder() *Recorder {
	return &Recorder{
		Store: make(map[string]*RecordedDevice),
	}
}

func (r *Recorder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	id := q["id"][0]
	humidity, _ := strconv.ParseFloat(q["hum"][0], 64)
	temperature, _ := strconv.ParseFloat(q["temp"][0], 64)
	record := &RecordedDevice{
		Id:          id,
		Humidity:    humidity,
		Temperature: ftoc(temperature),
		LastReport:  time.Now(),
	}
	log.Print(id, record)
	r.Store[id] = record
}

func ftoc(f float64) float64 {
	return (f - 32.0) * (5.0 / 9.0)
}
