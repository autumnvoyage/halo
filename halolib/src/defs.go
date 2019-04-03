package halo

/**
 * This file contains network transmission structures for most communications
 * involving Halo software. It exists more for reference purposes, although it
 * may be used sometimes in the software.
 */

// incoming handshake (unencrypted)
// Client initiated
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
