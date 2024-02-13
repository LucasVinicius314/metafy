package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sure/metafy/pkg/models"
	"sure/metafy/pkg/utils"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Connect(w http.ResponseWriter, r *http.Request) {
	utils.LogInboundRequest(w, r)

	token := ""
	for _, value := range r.Header.Values("Sec-Websocket-Protocol") {
		if strings.HasPrefix(value, "token-") {
			token = value[5:]
			break
		}
	}

	if token == "" {
		utils.SendMessage(w, 400, "missing token")
		return
	}

	log.Printf("token: [%s]", token)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading: %v", err)
		utils.SendMessage(w, 500, "internal server error")
		return
	}
	defer conn.Close()

	connId, err := uuid.NewUUID()
	if err != nil {
		log.Printf("error generating connection uuid: %v", err)
		utils.SendMessage(w, 500, "internal server error")
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
	utils.SendMessage(w, 200, "ok")
}

func Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		utils.SendMessage(w, 400, "invalid body")
		return
	}
	defer r.Body.Close()

	request := &models.LoginRequest{}
	err = json.Unmarshal(bodyBytes, request)
	if err != nil {
		log.Printf("error reading body json: %v", err)
		utils.SendMessage(w, 400, "invalid body json")
		return
	}

	// TODO: login
	if request.Email == "lorem" && request.Password == "ipsum" {
		utils.SendMessage(w, 200, "sample jwt token")
		return
	}

	utils.SendMessage(w, 400, "user not found")
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: register
	utils.SendMessage(w, 200, "ok")
}
