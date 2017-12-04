package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const port = ":19933"

func checkOriginFunc(r *http.Request) bool {
	// ignore all origin check
	return true
}

var upgrader = websocket.Upgrader{CheckOrigin: checkOriginFunc}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read err : %s\n", err.Error())
			return
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Printf("write err : %s\n", err.Error())
			return
		}
	}
}

func main() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(port, nil))
}
