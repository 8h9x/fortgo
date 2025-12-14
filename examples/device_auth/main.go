package main

import (
	"encoding/json"
	"fmt"
	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/auth"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	deviceAuthPath := "./deviceAuth.json"

	deviceAuth, err := fetchDeviceAuth(httpClient, auth.FortniteNewIOSClient, deviceAuthPath)
	if err != nil {
		log.Fatal(err)
	}

	deviceAuthPayload := auth.PayloadDeviceAuth{
		AccountID: deviceAuth.AccountID,
		DeviceID: deviceAuth.DeviceID,
		Secret: deviceAuth.Secret,
	}

	credentials, err := auth.Authenticate(httpClient, auth.FortniteNewIOSClient, deviceAuthPayload, true)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fortgo.NewClient(httpClient, credentials)
	if err != nil {
		log.Fatal("Failed to construct client", err)
	}

	log.Println("Fortgo client successfully created")
}

func fetchDeviceAuth(httpClient *http.Client, authClient *auth.AuthClient, filepath string) (auth.DeviceAuthResponse, error) {
	data, err := os.ReadFile(filepath)
	if os.IsNotExist(err) || len(data) == 0 {
		deviceAuth, err := generateDeviceAuth(httpClient, authClient)
		if err != nil {
			return auth.DeviceAuthResponse{}, err
		}

		buf, err := json.Marshal(deviceAuth)
		if err != nil {
			return auth.DeviceAuthResponse{}, err
		}

		err = os.WriteFile(filepath, buf, 0600)
		if err != nil {
			return auth.DeviceAuthResponse{}, err
		}

		return deviceAuth, nil
	} else if err != nil {
		return auth.DeviceAuthResponse{}, err
	}

	deviceAuth := &auth.DeviceAuthResponse{}
	err = json.Unmarshal(data, deviceAuth)
	if err != nil {
		return auth.DeviceAuthResponse{}, err
	}

	return *deviceAuth, nil
}

func generateDeviceAuth(httpClient *http.Client, authClient *auth.AuthClient) (auth.DeviceAuthResponse, error) {
	credentials, err := authCodeFlow(httpClient, authClient)
	if err != nil {
		return auth.DeviceAuthResponse{}, err
	}

	deviceAuth, err := auth.CreateDeviceAuth(httpClient, credentials)
	if err != nil {
		return auth.DeviceAuthResponse{}, err
	}

	return deviceAuth, nil
}

func authCodeFlow(httpClient *http.Client, authClient *auth.AuthClient) (auth.TokenResponse, error) {
	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", authClient.ID)
	_, err := fmt.Scan(&code)
	if err != nil {
		return auth.TokenResponse{}, err
	}

	credentials, err := auth.Authenticate(httpClient, authClient,  auth.PayloadAuthorizationCode{code}, true)
	if err != nil {
		return auth.TokenResponse{}, err
	}

	return credentials, nil
}