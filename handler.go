package goauthlib

import (
	"log"
	"net/http"
)

type Handler func(hut Http) interface{}

func (handleFunc Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hut := NewHttp(w, r)
	response := handleFunc(hut)
	switch response.(type) {
	case error:
		if Config.ErrorHandler != nil {
			Config.ErrorHandler(response.(error), hut)
			break
		}
		http.Error(w, "Error", 500)
		log.Print(response.(error).Error())
	case string:
		w.Write([]byte(response.(string)))
	default:
	}
}
