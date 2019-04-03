package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	addr = flag.String("addr", "127.0.0.1:1941", "HTTP service address")
	sessions []Session
	nonce = [...]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, }
)

const (
	writeWait        = 10 * time.Second
	maxMsgSize       = 8192
	pongWait         = 60 * time.Second
	pingPeriod       = (pongWait * 9) / 10
	closeGracePeriod = 10 * time.Second
)

func handleMsg(data []byte) error {
	magic := string(data[0:5])
	if magic == "EMSG" {
		var sessKey [32]byte
		found := false
		sessId_, err := strconv.Atoi(string(data[5:13]))
		if err != nil {
			return errors.New("Invalid session ID sent in EMSG: " +
				string(sessId_))
		}
		sessId := uint64(sessId_)
		for _, elem := range sessions {
			if elem.SessId == sessId {
				sessKey = elem.SessKey
				found = true
				break
			}
		}
		if !found {
			return errors.New("Bad session ID sent in EMSG, session ID: " +
				string(sessId))
		}
		pload, err := DecryptData(data, sessKey[:])
		if err != nil {
			return err
		}
		parseEMSG(pload)
	} else if magic == "HMSG" {
		var ver [4]byte
		copy(ver[:], data[5:9])
		version := decodeNetUint32(ver)
		if version > 0 {
			return errors.New("Version requested is unsupported: " +
				string(version))
		}
		var fprint [20]byte
		copy(fprint[:], data[9:29])
		fp := hex.EncodeToString(fprint[:])
		req := "https://pgp.key-server.io/pks/lookup?op=get&search=0x" + fp
		resp, err := http.Get(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		log.Println(body)
	}
	return nil
}

func parseEMSG(data []byte) error {
	return nil
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade HTTP to WebSocket: %v", err)
		return
	}
	log.Println("Upgraded HTTP to WebSocket.")
	defer conn.Close()
	for {
		_, d, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Println("Conn.ReadMessage() failed:", err)
			}
			return
		}
		handleMsg(d)
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
