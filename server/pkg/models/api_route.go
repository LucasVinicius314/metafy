package models

import "net/http"

type ApiRoute struct {
	Path   string
	Delete func(w http.ResponseWriter, r *http.Request)
	Get    func(w http.ResponseWriter, r *http.Request)
	Patch  func(w http.ResponseWriter, r *http.Request)
	Post   func(w http.ResponseWriter, r *http.Request)
	Put    func(w http.ResponseWriter, r *http.Request)
}
