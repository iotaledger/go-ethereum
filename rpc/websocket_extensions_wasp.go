package rpc

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebSocketCodec(conn *websocket.Conn, host string, req http.Header, readLimit int64) ServerCodec {
	return newWebsocketCodec(conn, host, req, readLimit)
}
