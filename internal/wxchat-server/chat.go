package wxchatserver

import (
	"encoding/json"
	"log"
	"wx-server/internal/httpserver"
	"wx-server/internal/logging"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

func (s *Server) chat(c *gin.Context) {
	var m Message
	err := c.Bind(&m)
	if err != nil {
		httpserver.ResponseBadRequest(c, err)
		return
	}
	content, err := s.qwen.Chat(m.Text)
	if err != nil {
		httpserver.ResponseServerError(c, err)
		return
	}
	httpserver.ResponseOK(c, Message{
		Text: content,
	})
}

func (s *Server) wschat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.WithContext(c).With("err", err).Warn("failed to upgrade websocket")
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logging.WithContext(c).With("err", err).Warn("failed to read message")
			return
		}

		switch messageType {
		case websocket.TextMessage:
		case websocket.BinaryMessage:
			var m Message
			err := json.Unmarshal(p, &m)
			if err != nil {
				logging.WithContext(c).With("err", err).Warn("failed to unmarshal message")
				log.Println(err)
				return
			}
			ch, err := s.qwen.StreamChat(m.Text)
			if err != nil {
				log.Println(err)
				return
			}
			go func() {
				for r := range ch {
					rm := Message{
						User: "qwen",
						Text: r,
					}
					buf, err := json.Marshal(rm)
					if err != nil {
						logging.WithContext(c).With("err", err).Warn("failed to marshal message")
						continue
					}
					conn.WriteMessage(websocket.TextMessage, buf)
				}
			}()
		case websocket.CloseMessage:
		case websocket.PongMessage:
			return
		case websocket.PingMessage:
			conn.WriteMessage(websocket.PongMessage, p)
		}
	}
}
