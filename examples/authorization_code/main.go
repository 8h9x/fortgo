package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/account"
)

func main() {
	httpClient := &http.Client{}

	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", account.FortniteNewIOSClient.ID)
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}

	authCodePayload := account.TokenPayloadAuthorizationCode{
		Code: code,
	}

	credentials, err := account.Authenticate(httpClient, account.FortniteNewIOSClient, authCodePayload, true)
	if err != nil {
		log.Fatal(err)
	}

	client := fortgo.NewClient(httpClient, credentials)

	err = client.Connect()
	if err != nil {
		log.Fatal("Failed to connect to client", err)
	}

	log.Println("Fortgo client connected!")
}
