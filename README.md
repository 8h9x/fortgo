# Fortgo

A library for making requests to Epic Games APIs--primarily targeting fortnite related endpoints.

Installation: `go get github.com/8h9x/fortgo`

Simple client creation example: 
#### `/examples/authorization_code/main.go`
```go
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

	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", consts.FortniteNewIOSClientID)
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}

	authCodePayload := auth.PayloadAuthorizationCode{
		Code: code,
	}

	credentials, err := auth.Authenticate(httpClient, auth.FortnitePS4USClient, authCodePayload, true)
	if err != nil {
		log.Println(err)
	}

	_, err = fortgo.NewClient(httpClient, credentials)
	if err != nil {
		log.Println("Failed to construct client", err)
	}

	log.Println("Fortgo client successfully created")
}
```