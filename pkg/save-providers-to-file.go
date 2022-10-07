package pkg

import (
	"encoding/json"
	"io"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//Saves the providers along with the CIDs in a json format. In an error occurs it returns the error or else
//it returns nil.
//
//Because we want to add a new provider record in the file for each new provider record
//we need to read the contents and add the new provider record to the already existing array.
func saveProvidersToFile(filename string, contentID string, addressInfos []*peer.AddrInfo) error {
	log.Debug("starting to save providers to file")
	log.Debugf("cid is: %s", contentID)
	log.Debugf("address infos: %v", addressInfos)
	jsonFile, err := os.Open(filename)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Errorf("error %s while closing down providers file", err)
		}
	}(jsonFile)
	if err != nil {
		return errors.Wrap(err, "whiel trying to open json file")
	}
	//create a new instance of ProviderRecords struct which is a container for the encapsulated struct
	var records ProviderRecords

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrap(err, "while trying to read json file")
	}

	if len(bytes) != 0 {
		//read the existing data. Will throw error if they are not of type EncapsulatedJSONproviderRecord
		err = json.Unmarshal(bytes, &records)
		if err != nil {
			return errors.Wrap(err, "while unmarshalling json")
		}
	}

	for _, addressInfo := range addressInfos {

		if addressInfo == nil {
			continue
		}

		//create a new encapsulated struct
		NewEncapsulatedJSONProviderRecord := NewEncapsulatedJSONCidProvider(contentID, addressInfo.ID.String(), addressInfo.Addrs)
		log.Debugf("Created new encapsulated JSON provider record: ID:%s,CID:%s,Addresses:%v", NewEncapsulatedJSONProviderRecord.ID, NewEncapsulatedJSONProviderRecord.CID, NewEncapsulatedJSONProviderRecord.Addresses)
		//insert the new provider record to the slice in memory containing the provider records read
		records.EncapsulatedJSONProviderRecords = append(records.EncapsulatedJSONProviderRecords, NewEncapsulatedJSONProviderRecord)
	}
	data, err := json.MarshalIndent(&records, "", " ")
	if err != nil {
		return errors.Wrap(err, "while marshalling json data")
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.Wrap(err, "while trying to write json data to file")
	}
	return nil
}
