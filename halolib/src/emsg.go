package halo

import (
	"crypto/aes"
	"crypto/cipher"
)

var (
	nonce = [20]byte{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
)

type Session struct {
	SessKey [32]byte
	SessId uint64
}

func EncryptData(indata []byte, key []byte) ([]byte, error) {
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

func DecryptData(indata []byte, key []byte) ([]byte, error) {
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
