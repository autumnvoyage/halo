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
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	addr = flag.String("addr", "127.0.0.1:1941", "HTTP service address")
)

const (
	writeWait        = 10 * time.Second
	maxMsgSize       = 8192
	pongWait         = 60 * time.Second
	pingPeriod       = (pongWait * 9) / 10
	closeGracePeriod = 10 * time.Second
)

func handleMsg(data []byte) {

}

// For messages, we assume we have TLS

// incoming, TLS only (no encryption)
type HandshakeIn struct {
	Magic [4]byte // ASCII no-NUL "HMSG"
	Version uint32 // 0
	PubkeyFprint [20]byte // Look up the client’s PGP
}

// outgoing, RSA enc, client’s key
type HandshakeOut struct {
	Magic [4]byte // ASCII no-NUL "HMSG"
	Version uint32 // 0
	TTL uint32 // Lifetime of session
	SessKey [32]byte // AES-256, future msgs use this instead of RSA
	SessId uint64 // Tag future comms
}

type MessageIn struct {
	Magic [4]byte // ASCII no-NUL "EMSG"
	SessId uint64
	Payload []byte // decrypted with AES-256 sesskey
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade HTTP to WebSocket: %v", err)
		return
	}
	log.Println("Upgraded HTTP to WebSocket.")
	for {
		msgType, data, err := conn.ReadJSON()
		if err != nil {
			log.Fatalf("Conn.ReadJSON() failed: %v", err)
			return
		}
		handleMsg(data)
	}
}

func httpSetup(addr string) {
	http.HandleFunc("/", wsHandler)
	log.Println("Set HTTP server handler")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	flag.Parse()
	log.Printf("Service address: %v", *addr)
	httpSetup(*addr)
}
