package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/account"
)

func main() {
	httpClient := &http.Client{}
	deviceAuthPath := "./deviceAuth.json"

	deviceAuth, err := fetchDeviceAuth(httpClient, account.FortniteNewIOSClient, deviceAuthPath)
	if err != nil {
		log.Fatal(err)
	}

	deviceAuthPayload := account.TokenPayloadDeviceAuth{
		AccountID: deviceAuth.AccountID,
		DeviceID:  deviceAuth.DeviceID,
		Secret:    deviceAuth.Secret,
	}

	credentials, err := account.Authenticate(httpClient, account.FortniteNewIOSClient, deviceAuthPayload, true)
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

func fetchDeviceAuth(httpClient *http.Client, authClient *account.AuthClient, filepath string) (account.DeviceAuthResponse, error) {
	data, err := os.ReadFile(filepath)
	if os.IsNotExist(err) || len(data) == 0 {
		deviceAuth, err := generateDeviceAuth(httpClient, authClient)
		if err != nil {
			return account.DeviceAuthResponse{}, err
		}

		buf, err := json.Marshal(deviceAuth)
		if err != nil {
			return account.DeviceAuthResponse{}, err
		}

		err = os.WriteFile(filepath, buf, 0600)
		if err != nil {
			return account.DeviceAuthResponse{}, err
		}

		return deviceAuth, nil
	} else if err != nil {
		return account.DeviceAuthResponse{}, err
	}

	deviceAuth := &account.DeviceAuthResponse{}
	err = json.Unmarshal(data, deviceAuth)
	if err != nil {
		return account.DeviceAuthResponse{}, err
	}

	return *deviceAuth, nil
}

func generateDeviceAuth(httpClient *http.Client, authClient *account.AuthClient) (account.DeviceAuthResponse, error) {
	credentials, err := authCodeFlow(httpClient, authClient)
	if err != nil {
		return account.DeviceAuthResponse{}, err
	}

	deviceAuth, err := account.CreateDeviceAuth(httpClient, &credentials)
	if err != nil {
		return account.DeviceAuthResponse{}, err
	}

	return deviceAuth, nil
}

func authCodeFlow(httpClient *http.Client, authClient *account.AuthClient) (account.TokenResponse, error) {
	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", authClient.ID)
	_, err := fmt.Scan(&code)
	if err != nil {
		return account.TokenResponse{}, err
	}

	credentials, err := account.Authenticate(httpClient, authClient, account.TokenPayloadAuthorizationCode{Code: code}, true)
	if err != nil {
		return account.TokenResponse{}, err
	}

	return credentials, nil
}
