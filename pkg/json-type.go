package pkg

import (
	ma "github.com/multiformats/go-multiaddr"
)

//A container for the encapsulated struct.
//
//File containts a json array of provider records.
//[{ProviderRecord1},{ProviderRecord2},{ProviderRecord3}]
type ProviderRecords struct {
	EncapsulatedJSONProviderRecords []EncapsulatedJSONProviderRecord `json:"ProviderRecords"`
}

//This struct will be used to create,read and store the encapsulated data necessary for reading the
//provider records.
type EncapsulatedJSONProviderRecord struct {
	ID        string         `json:"PeerID"`
	CID       string         `json:"ContentID"`
	Addresses []ma.Multiaddr `json:"PeerMultiaddress"`
}

//Creates a new:
//	EncapsulatedCidProvider struct {
//		ID      string
//		CID     string
//		Address ma.Multiaddr
//	}
func NewEncapsulatedJSONCidProvider(id string, cid string, addresses []ma.Multiaddr) EncapsulatedJSONProviderRecord {
	return EncapsulatedJSONProviderRecord{
		ID:        id,
		CID:       cid,
		Addresses: addresses,
	}
}

const filename = "C:\\Users\\fotis\\GolandProjects\\retrieval-success-rate\\go-libp2p-kad-dht\\providers.json"
