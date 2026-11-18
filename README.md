# Fortgo

Fortgo is in active pre-release development and constant breaking changes between commit versions are to be expected.

A library for making requests to Epic Games APIs--primarily targeting fortnite related endpoints.

Installation: `go get github.com/8h9x/fortgo`

Simple client creation example:

#### `/examples/authorization_code/main.go`

```go
package main

import (
    "fmt"
    "log"

    "github.com/8h9x/fortgo"
    "github.com/8h9x/fortgo/account"
)

func main() {
    var code string

    fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", account.FortniteNewIOSClient.ID)
    _, err := fmt.Scan(&code)
    if err != nil {
        log.Fatal(err)
    }

    client := fortgo.NewClient(
        fortgo.WithAuthClient(account.FortniteNewIOSClient),
        fortgo.WithTokenGrant(account.TokenGrantAuthorizationCode(code)),
    )

    err = client.Connect()
    if err != nil {
        log.Fatal("Failed to connect to client: ", err)
    }

    log.Println("Fortgo client connected!")
}
```

For a more complete example, see [Blast](https://github.com/8h9x/BlastBot)
