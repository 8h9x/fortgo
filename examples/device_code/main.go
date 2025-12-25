package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/auth"
)

const (
	USE_CUSTOM_CHECK_TIMEOUT = false

	CHECK_INTERVAL = time.Second * 10
	CUSTOM_CHECK_TIMEOUT = time.Minute * 2
)

func main() {
	httpClient := &http.Client{}

	clientCredentials, err := auth.Authenticate(httpClient, auth.FortnitePS4USClient, auth.PayloadClientCredentials{}, false)
	if err != nil {
		log.Fatal("Unable to generate client credentials: ", err)
	}

	deviceAuthorization, err := auth.CreateDeviceCode(httpClient, clientCredentials)
	if err != nil {
		log.Fatal("Unable to start device code flow: ", err)
	}

	var timeout time.Duration
	if USE_CUSTOM_CHECK_TIMEOUT {
		timeout = CUSTOM_CHECK_TIMEOUT
	} else {
		timeout = time.Second * time.Duration(deviceAuthorization.ExpiresIn)
	}

	fmt.Printf("Visit this URL: %s\nthen press 'Confirm' on the epic games login page.\n\nThis login flow will time out in %v minutes\n", deviceAuthorization.VerificationURIComplete, timeout.Minutes())

	credentials, err := waitForDeviceCodeConfirm(httpClient, deviceAuthorization.DeviceCode, CHECK_INTERVAL, timeout)
	if err != nil {
		log.Fatal("Unable to complete device code flow: ", err)
	}

	client := fortgo.NewClient(httpClient, credentials)

	err = client.Connect()
	if err != nil {
		log.Fatal("Failed to connect to client", err)
	}

	log.Println("Fortgo client connected!")
}

func waitForDeviceCodeConfirm(httpClient *http.Client, deviceCode string, interval, timeout time.Duration) (auth.TokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	payload := auth.PayloadDeviceCode{
		DeviceCode: deviceCode,
	}

	for {
		select {
		case <-ctx.Done():
			return auth.TokenResponse{}, ctx.Err()
		case <-ticker.C:
			credentials, err := auth.Authenticate(httpClient, auth.FortnitePS4USClient, payload, true)
			if err != nil {
				continue
			}
			return credentials, nil
		}
	}
}
