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

type Session struct {
	SessKey [32]byte
	SessId uint64
}

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

func encryptData(indata []byte, key []byte) ([]byte, error) {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(ciph)
	if err != nil {
		panic(err.Error())
	}
	outdata := aesgcm.Seal(nil, nonce[:], indata, nil)
	return outdata, nil
}

func decryptData(indata []byte, key []byte) ([]byte, error) {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(ciph)
	if err != nil {
		panic(err.Error())
	}
	outdata, err := aesgcm.Open(nil, nonce[:], indata, nil)
	return outdata, nil
}

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
		pload, err := decryptData(data, sessKey[:])
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

// incoming, TLS only (no encryption)
type HandshakeIn struct {
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
