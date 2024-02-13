package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"sure/metafy/pkg/utils"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Connect(w http.ResponseWriter, r *http.Request) {
	utils.LogInboundRequest(w, r)

	headersJson, err := json.Marshal(r.Header)
	if err != nil {
		log.Printf("error marshalling headers: %v", err)
		return
	}

	log.Printf("headers: [%s]", string(headersJson))

	log.Printf("proto: [%s]", r.Proto)

	trailersJson, err := json.Marshal(r.Trailer)
	if err != nil {
		log.Printf("error marshalling trailers: %v", err)
		return
	}

	log.Printf("trailers: [%s]", string(trailersJson))

	log.Print(r.TransferEncoding)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading: %v", err)

		w.WriteHeader(500)
		w.Write([]byte("500 internal server error"))
		return
	}
	defer conn.Close()

	connId, err := uuid.NewUUID()
	if err != nil {
		log.Printf("error generating connection uuid: %v", err)
		return
	}

	log.Printf("inbound connection [%s]", connId)

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				log.Printf("outbound connection [%s]", connId)
				return
			}

			log.Printf("error reading message: %v", err)
			break
		}

		log.Printf("[%s]: %s", connId, string(message))

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Printf("error writing: %v", err)
			break
		}
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: login
	w.Write([]byte("ok"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: register
	w.Write([]byte("ok"))
}
