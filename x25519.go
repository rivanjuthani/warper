package main

import (
	"encoding/base64"

	"github.com/keys-pub/keys"
)

type KeyPairResponse struct {
	PublicKey   []byte
	PrivateKey  []byte
	RegisterKey string
}

func NewX25519KeyPair() KeyPairResponse {
	x25519Key := keys.GenerateX25519Key()

	publicKey := x25519Key.Public()
	privateKey := x25519Key.Private()

	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)

	return KeyPairResponse{
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
		RegisterKey: publicKeyBase64,
	}
}
