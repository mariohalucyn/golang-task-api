package initializers

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

var EcdsaPublicKey *ecdsa.PublicKey

func LoadAndParsePublicKey() {
	var ok bool
	der, err := os.ReadFile("ec-prime256v1-pub-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("err")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("err")
	}

	EcdsaPublicKey, ok = pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("err")
	}
}
