package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/account"
)

const checkInterval = time.Second * 10

func main() {
	client := fortgo.NewClient()

	clientToken, err := client.AccountService.GetToken(
		account.FortnitePS4USClient,
		account.TokenGrantClientCredentials(),
		account.TokenTypeDefault,
	)
	if err != nil {
		log.Fatal("Unable to generate client credentials: ", err)
	}

	deviceAuthorization, err := client.AccountService.CreateDeviceCode(clientToken)
	if err != nil {
		log.Fatal("Unable to generate client credentials: ", err)
	}

	log.Println(deviceAuthorization)

	timeout := time.Second * time.Duration(deviceAuthorization.ExpiresIn)

	fmt.Printf("Visit this URL: %s\nthen press 'Confirm' on the epic games login page.\n\nThis login flow will time out in %v minutes\n", deviceAuthorization.VerificationURIComplete, timeout.Minutes())

	token, err := waitForConfirmAndGetToken(client, deviceAuthorization.DeviceCode, checkInterval, timeout)
	if err != nil {
		log.Fatal("Failed to login: ", err)
	}

	err = client.ConnectWithToken(token)
	if err != nil {
		log.Fatal("Failed to connect to client", err)
	}

	log.Println("Fortgo client connected!")
}

func waitForConfirmAndGetToken(client *fortgo.Client, deviceCode string, interval, timeout time.Duration) (account.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return account.Token{}, ctx.Err()
		case <-ticker.C:
			token, err := client.AccountService.GetToken(
				account.FortnitePS4USClient,
				account.TokenGrantDeviceCode(deviceCode),
				account.TokenTypeEG1,
			)
			if err != nil {
				log.Println(err)
				continue
			}

			return token, nil
		}
	}
}
