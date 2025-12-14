package fortgo

import "github.com/8h9x/fortgo/links"

type Mnemonic struct {
	client *Client
	links.MnemonicData
}

func (m *Mnemonic) FetchRelated() (links.FetchRelatedMnemonicsResponse, error) {
	return links.FetchRelatedMnemonics(m.client.HttpClient, m.client.CurrentCredentials(), "fn", m.Mnemonic, m.Version)
}