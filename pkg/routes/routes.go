package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"
	"net/http" // could change this to Echo https://echo.labstack.com/
	"github.com/gorilla/websocket" // could change this to Echo https://echo.labstack.com/

)

type ChatMessage struct {
    ChatMessage string `json:"chatMessage"`
    Headers     struct {
        HXRequest      string `json:"HX-Request"`
        HXTrigger      string `json:"HX-Trigger"`
        HXTriggerName  string `json:"HX-Trigger-Name"`
        HXTarget       string `json:"HX-Target"`
        HXCurrentURL   string `json:"HX-Current-URL"`
    } `json:"HEADERS"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	log.Println("Starting reader")
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading the websocket message:", err)
			return
		}
		log.Println(messageType)
		log.Println(string(p))

		// Unmarshal the JSON message into the ChatMessage struct
		var chatMsg ChatMessage
        if err := json.Unmarshal(p, &chatMsg); err != nil {
            log.Println("Error unmarshaling JSON:", err)
            continue
        }

        // Return the chat message back to the UI
        log.Println("Received chat message:", chatMsg.ChatMessage)
		message := "You just typed '" + chatMsg.ChatMessage + "'"
		messageAppend := []byte(`<div id="idChatroomAppend" hx-swap-oob="beforeend"><p>` + message + `</p></div>`)
		if err := conn.WriteMessage(websocket.TextMessage, messageAppend); err != nil {
			log.Println("Error writing message back to client:", err)
			return
		}
	}
}

func sendWSMessage(conn *websocket.Conn) {
	log.Println("Starting sendWSMessage")
	for i := 0; i < 10; i++ {
		message := "This is message update number " + strconv.Itoa(i+1) + "; "
		messageSwap := []byte(`<div id="idMessageSwap" hx-swap-oob="true">` + message + `</div>`)
		messageAppend := []byte(`<div id="idMessageAppend" hx-swap-oob="beforeend">` + message + `</div>`)
		log.Println(i + 1)
		if err := conn.WriteMessage(websocket.TextMessage, messageSwap); err != nil {
			log.Println("Error writing swap message", err)
			return
		}
		if err := conn.WriteMessage(websocket.TextMessage, messageAppend); err != nil {
			log.Println("Error writing append message", err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading the endpoint to a websocket", err)
	}
	log.Println("Starting wsEndpoint")
	go sendWSMessage(ws)
	go reader(ws)
}

func setupRoutes() {
	absPath, err := filepath.Abs(".")
	if err != nil {panic(err)}
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, absPath + "\\pkg\\website\\home.html")
		x := http.Dir("website")
		log.Println(x)
	})
	clickCounter := 0
	http.HandleFunc("/ClickCount", func(w http.ResponseWriter, r *http.Request) {
		clickCounter++
		fmt.Fprint(w, clickCounter)
	})
}

func Routes() {
	fmt.Println("Go Websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}