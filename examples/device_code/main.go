package main

import (
	"context"
	"fmt"
	"github.com/8h9x/vinderman"
	"github.com/8h9x/vinderman/auth"
	"github.com/8h9x/vinderman/consts"
	"log"
	"net/http"
	"time"
)

const (
	CHECK_INTERVAL = time.Second * 10
	CHECK_TIMEOUT  = time.Minute * 2
)

func main() {
	httpClient := &http.Client{}

	clientCredentials, err := auth.Authenticate(httpClient, consts.FortnitePS4USClientID, consts.FortnitePS4USClientSecret, auth.PayloadClientCredentials{}, false)
	if err != nil {
		log.Fatal("Unable to generate client credentials: ", err)
	}

	deviceAuthorization, err := auth.GetDeviceCode(httpClient, clientCredentials)
	if err != nil {
		log.Fatal("Unable to start device code flow: ", err)
	}

	fmt.Printf("Visit this URL: %s\nThen press 'Confirm' on the epic games login page.\n", deviceAuthorization.VerificationURIComplete)

	credentials, err := waitForDeviceCodeConfirm(httpClient, deviceAuthorization.DeviceCode, CHECK_INTERVAL, CHECK_TIMEOUT)
	if err != nil {
		log.Fatal("Unable to start device code flow: ", err)
	}

	_, err = vinderman.NewClient(httpClient, credentials)
	if err != nil {
		log.Println("Failed to construct client", err)
	}

	log.Println("Vinderman client successfully created")
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
			credentials, err := auth.Authenticate(httpClient, consts.FortnitePS4USClientID, consts.FortnitePS4USClientSecret, payload, true)
			if err != nil {
				continue
			}
			return credentials, nil
		}
	}
}
