package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sure/metafy/pkg/models"
	endpoints "sure/metafy/pkg/modules"
	"sure/metafy/pkg/utils"
)

var routes = []models.ApiRoute{
	{
		Path: "/api/health",
		Get:  endpoints.Health,
	},
	{
		Path: "/api/login",
		Post: endpoints.Login,
	},
	{
		Path: "/api/register",
		Post: endpoints.Register,
	},
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../public")))

	for _, route := range routes {
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			utils.LogInboundRequest(w, r)

			switch r.Method {
			case "DELETE":
				if route.Delete != nil {
					route.Delete(w, r)
				}
			case "GET":
				if route.Get != nil {
					route.Get(w, r)
				}
			case "PATCH":
				if route.Patch != nil {
					route.Patch(w, r)
				}
			case "POST":
				if route.Post != nil {
					route.Post(w, r)
				}
			case "PUT":
				if route.Put != nil {
					route.Put(w, r)
				}
			}

			utils.SendMessage(w, 405, "method not allowed")
		})
	}

	http.HandleFunc("/api/connect", endpoints.Connect)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
