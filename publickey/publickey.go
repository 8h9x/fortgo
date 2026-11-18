package publickey

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) RegisterKeyED25519(key ed25519.PublicKey) (RegisterPublicKeyResponse, error) {
	return c.registerKey("ed25519", base64.StdEncoding.EncodeToString(key))
}

func (c *Client) RegisterKeyRSA(key *rsa.PublicKey) (RegisterPublicKeyResponse, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return RegisterPublicKeyResponse{}, err
	}
	return c.registerKey("rsa", base64.StdEncoding.EncodeToString(pubKeyBytes))
}

func (c *Client) registerKey(algorithm, publicKey string) (RegisterPublicKeyResponse, error) {
	payload := registerKeyPayload{
		Algorithm: algorithm,
		PublicKey: publicKey,
	}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.PublicKeyService,
		"publickey/v2/publickey",
		rest.WithJSONBody(payload),
	)
	if err != nil {
		return RegisterPublicKeyResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return RegisterPublicKeyResponse{}, err
	}

	var data RegisterPublicKeyResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return RegisterPublicKeyResponse{}, err
	}

	return data, err
}
