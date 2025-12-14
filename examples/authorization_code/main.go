package main

import (
	"fmt"
	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/consts"
	"log"
	"net/http"
)

func main() {
	httpClient := &http.Client{}
	authClientNewIOS := &auth.AuthClient{consts.FortniteNewIOSClientID,consts.FortniteNewIOSClientSecret}

	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", authClientNewIOS.ID)
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}

	authCodePayload := auth.PayloadAuthorizationCode{
		Code: code,
	}

	credentials, err := auth.Authenticate(httpClient, authClientNewIOS, authCodePayload, true)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fortgo.NewClient(httpClient, credentials)
	if err != nil {
		log.Fatal("Failed to construct client", err)
	}

	log.Println("Fortgo client successfully created")
}
