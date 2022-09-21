package pkg

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
)

type Host struct {
	ctx context.Context
	host.Host
	DHT *kaddht.IpfsDHT
}
