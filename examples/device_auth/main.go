package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/8h9x/fortgo"
	"github.com/8h9x/fortgo/account"
)

const deviceAuthFile = "./deviceAuth.json"

func main() {
	var tokenGrant account.TokenGrant
	shouldGenerate := false

	deviceAuth, err := readDeviceAuthData(deviceAuthFile)
	if err != nil {
		var code string

		fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", account.FortniteNewIOSClient.ID)
		_, err := fmt.Scan(&code)
		if err != nil {
			log.Fatal("Failed to read auth code input: ", err)
		}

		tokenGrant = account.TokenGrantAuthorizationCode(code)
		shouldGenerate = true
	} else {
		tokenGrant = account.TokenGrantDeviceAuth(
			deviceAuth.AccountID,
			deviceAuth.DeviceID,
			deviceAuth.Secret,
		)
	}

	client := fortgo.NewClient(
		fortgo.WithAuthClient(account.FortniteNewIOSClient),
		fortgo.WithTokenGrant(tokenGrant),
	)

	err = client.Connect()
	if err != nil {
		log.Fatal("Failed to connect to client: ", err)
	}

	log.Println("Fortgo client connected!")

	if shouldGenerate {
		log.Println("Generating new deviceAuth...")

		deviceAuth, err := client.AccountService.CreateDeviceAuth(client.CurrentCredentials().AccountID)
		if err != nil {
			log.Fatal("Unable to create deviceAuth: ", err)
		}

		err = writeDeviceAuthData(deviceAuth)
		if err != nil {
			log.Fatal("Unable to save deviceAuth to file: ", err)
		}

		log.Println(fmt.Sprintf("New deviceAuth saved to '%s'", deviceAuthFile))
	}
}

func readDeviceAuthData(filepath string) (account.DeviceAuth, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return account.DeviceAuth{}, err
	}

	deviceAuth := account.DeviceAuth{}
	err = json.Unmarshal(data, &deviceAuth)
	if err != nil {
		return account.DeviceAuth{}, err
	}

	return deviceAuth, nil
}

func writeDeviceAuthData(deviceAuth account.DeviceAuth) error {
	deviceAuthBytes, err := json.Marshal(deviceAuth)
	if err != nil {
		return err
	}

	err = os.WriteFile(deviceAuthFile, deviceAuthBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
