package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	addr = flag.String("addr", "127.0.0.1:1941", "HTTP service address")
)

const (
	writeWait = 10 * time.Second
	maxMsgSize = 8192
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	closeGracePeriod = 10 * time.Second
)

func handleMsg(data []byte) {

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade HTTP to WebSocket: %v", err)
		return
	}
	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("ReadMessage() failed: %v", err)
			return
		}
		if msgType == websocket.BinMessage {
			return
		}
		handleMsg(data)
	}
}

func httpSetup(addr string) {
	http.HandleFunc("/", wsHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Must specify listening address")
	}
	httpSetup(*addr)
}
