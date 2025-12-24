package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
    "flag"
	"fmt"
	"log"
	"net/http"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/auth"
)

func main() {
	keyType := flag.String("key-type", "ed25519", "Key algorithm (ed25519 or rsa)")
	flag.Parse()

	switch *keyType {
	case "ed25519", "rsa":
		fmt.Printf("Using %s key type\n", *keyType)
	default:
		log.Fatalln(fmt.Sprintf("Error: invalid key type %q. Must be 'ed25519' or 'rsa'", *keyType))
	}

	httpClient := &http.Client{}

	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", auth.FortniteNewIOSClient.ID)
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}

	authCodePayload := auth.PayloadAuthorizationCode{
		Code: code,
	}

	credentials, err := auth.Authenticate(httpClient, auth.FortniteNewIOSClient, authCodePayload, true)
	if err != nil {
		log.Fatal(err)
	}

	client := fortgo.NewClient(httpClient, credentials)

	err = client.Connect()
	if err != nil {
		log.Fatal("Failed to connect to client", err)
	}

	log.Println("Fortgo client connected!")

	if *keyType == "rsa" {
		rsaKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		res, err := client.PublicKeyService.RegisterKeyRSA(&rsaKey.PublicKey)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("RSA public key registered\nPublic Key: %s\nPrivate Key: %s\nJWT: %s", res.Key, base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(rsaKey)), res.JWT))
	} else {
		pub, priv, _ := ed25519.GenerateKey(rand.Reader)
		res, err := client.PublicKeyService.RegisterKeyED25519(pub)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(fmt.Sprintf("ED25519 public key registered\nPublic Key: %s\nPrivate Key: %s\nJWT: %s", base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(priv), res.JWT))
	}
}
