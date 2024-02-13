package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func LogInboundRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
}

func SendMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"message\":%s}", strconv.Quote(message))))
}
