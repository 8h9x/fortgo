package fortgo

import "github.com/8h9x/fortgo/links"

type Mnemonic struct {
	client *Client
	links.MnemonicData
}

func (m *Mnemonic) GetRelated() (links.GetRelatedMnemonicsResponse, error) {
	return m.client.LinkService.GetRelatedMnemonics("fn", m.Mnemonic, m.Version)
}