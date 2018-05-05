package servicehandlers

import (
	"log"
	"net/http"
	"time"
)

type HttpServiceHandler interface {
	Get(*http.Request) SrvcRes
	Put(*http.Request) SrvcRes
	Post(*http.Request) SrvcRes
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) interface{} {
	var response interface{}

	if r.Method == "GET" {
		startTime := time.Now()
		response = p.Get(r)
		duration := time.Now().Sub(startTime)
		log.Println(duration)
	} else if r.Method == "PUT" {
		startTime := time.Now()
		response = p.Put(r)
		duration := time.Now().Sub(startTime)
		log.Println(duration)
	} else if r.Method == "POST" {
		startTime := time.Now()
		response = p.Post(r)
		duration := time.Now().Sub(startTime)
		log.Println(duration)
	}
	return response
}
