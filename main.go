package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

var rooms = make(map[string]*melody.Melody)

func main() {
	r := gin.Default()
	m := melody.New()
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws/:room", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			fmt.Println("q.Request.URL.Path", q.Request.URL.Path)
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	r.Run(":5000")
}
