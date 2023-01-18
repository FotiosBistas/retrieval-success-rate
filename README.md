# retrieval-success-rate
This repository aims to track the retrieval success rate of the provider records using the Optimistic Provide algorithm proposed by [Dennis Trautwein](https://github.com/dennis-tra "dennis-tra"), using [Mikel Corte's](https://github.com/cortze "cortze") repository [IPFS CID HOARDER](https://github.com/cortze/ipfs-cid-hoarder "hoarder"). The ``go-libp2p-kad-dht`` submodule uses the ``ProvidersFile`` branch for all the following workflow. 

# How to run 
Example run: 
``go run retrieval_success_rate.go run_optimistic_provide --cid-number 4000 --log-level debug`` 

# Process 

## Provider Record format 
```golang
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
    ID              string   `json:"PeerID"`
    CID             string   `json:"ContentID"`
    Creator         string   `json:"Creator"`
    PublicationTime string   `json:"PublicationTime"`
    ProvideTime     string   `json:"ProvideTime"`
    UserAgent       string   `json:"UserAgent"`
    Addresses       []string `json:"PeerMultiaddresses"`
}
```

## Saving to JSON files 

If the sample size is small ( <= 1000 CIDs) the provider records are writen into a JSON file and inserted into the hoarder. The hoarder is then responsible for extracting the metadata from the JSON file, inserting the CIDs into the database and proceeding the ping them, gathering the necessary data to determine whether the Provider Records are retrievable. 

## Publishing to HTTP server 

If the cid-number is large (> 1000 CIDs) the HTTP server, inside the hoarder, is configured to listen to ``port:8080`` and the hoarder's ``localhost`` address. The publisher, using for example the address ``http://localhost:8080/`` connects to the server, and must use an HTTP ``POST`` method to send the CIDs to the server. The CIDs must follow the following format: 
```golang
The HTTP server has an internal queue that contains an array of ``[]EncapsulatedJSONProviderRecord`` and using a ``GET`` request the hoarder can receive an element from the array. 
```golang
type HttpCidSource struct {
    port            int
    hostname        string
    lock            sync.Mutex
    server          *http.Server
    providerRecords []ProviderRecords
    isStarted       bool
}
```
This means that the publisher must aggregate the ``EncapsulatedJSONProviderRecord`` for each CID into an array and then send it to the server. Then the same process as in the JSON file section is followed by the hoarder. 
