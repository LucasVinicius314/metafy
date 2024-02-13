package utils

import (
	"log"
	"net/http"
)

func LogInboundRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
}
