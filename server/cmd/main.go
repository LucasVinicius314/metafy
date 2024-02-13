package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type ApiRoute struct {
	Path string
	Get  func(w http.ResponseWriter, r *http.Request)
	Post func(w http.ResponseWriter, r *http.Request)
}

var upgrader = websocket.Upgrader{}

func main() {
	routes := []ApiRoute{
		{
			Path: "/api/health",
			Get: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("ok"))
			},
		},
		{
			Path: "/api/login",
			Post: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("ok"))
			},
		},
		{
			Path: "/api/register",
			Post: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("ok"))
			},
		},
	}

	http.Handle("/", http.FileServer(http.Dir("../public")))

	for _, route := range routes {
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			logInboundRequest(w, r)

			switch r.Method {
			case "GET":
				if route.Get != nil {
					route.Get(w, r)
				}
			case "POST":
				if route.Post != nil {
					route.Post(w, r)
				}
			}

			w.WriteHeader(405)
			w.Write([]byte("405 method not allowed"))
		})
	}

	http.HandleFunc("/api/connect", func(w http.ResponseWriter, r *http.Request) {
		logInboundRequest(w, r)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade failed: ", err)

			w.WriteHeader(500)
			w.Write([]byte("500 internal server error"))
			return
		}
		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}

			log.Println(string(message))

			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	log.Println(err)
}

func logInboundRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
}
