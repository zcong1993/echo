package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(c *gin.Context) {
	cc, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer cc.Close()
	for {
		mt, message, err := cc.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = cc.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func helloHandler(c *gin.Context) {
	res := ""
	for k, v := range c.Request.Header {
		res += fmt.Sprintf("%s:%s\n", k, v)
	}
	c.String(http.StatusOK, res)
}

func echoHandler(c *gin.Context) {
	io.Copy(c.Writer, c.Request.Body)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", helloHandler)
	r.GET("/hello", helloHandler)
	r.GET("/ws", wsHandler)
	r.POST("/echo", echoHandler)

	r.Run()
}
