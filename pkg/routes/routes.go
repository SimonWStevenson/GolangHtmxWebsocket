package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Home Page")
//}

func reader(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println((err))
			return
		}

		/*
		log.Println("is this working here")
		message := []byte("This is a periodic message.")
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println(err)
		return
		}		
		*/
	}
}

func sendWSMessage(conn *websocket.Conn){
	log.Println("is this working")
	for i:= 0; i<10; i++ {
		message := []byte("This is a periodic message.")
		time.Sleep(5 * time.Second)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println(err)
			return
		}		
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {log.Println(err)}
	log.Println("Client successfully connected...")
	sendWSMessage(ws)
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "C:\\Users\\simon\\GitHub\\Four\\pkg\\website\\pets.html")
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "C:\\Users\\simon\\GitHub\\Four\\pkg\\website\\home.html")
	})
	clickCounter := 0
	http.HandleFunc("/ClickCount", func(w http.ResponseWriter, r *http.Request){
		clickCounter++
		fmt.Fprint(w, clickCounter)
	})
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home Page")
	})
}

func Routes() {
	fmt.Println("Go Websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}