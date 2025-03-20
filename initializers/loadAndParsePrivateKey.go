package initializers

import (
	"crypto/ecdsa"
	"crypto/x509"
	"log"
	"os"
)

var EcdsaPrivateKey *ecdsa.PrivateKey

func LoadAndParsePrivateKey() {
	der, err := os.ReadFile("ec-prime256v1-priv-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	EcdsaPrivateKey, err = x509.ParseECPrivateKey(der)
	if err != nil {
		log.Fatal(err)
	}
}
