package pkg

import (
	"context"
	"crypto"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/routing"
)

type Host struct {
	ctx context.Context
	host.Host
	DHT *kaddht.IpfsDHT
}

func new_host(ctx context.Context, priv_key crypto.PrivateKey) (*Host, error) {

	log.Debug("Creating new host")
	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			dht, err := kaddht.New(ctx, h)
			return dht, err
		}))

	if err != nil {
		panic(err)
	}

	for _, p := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		if err = h.Connect(ctx, p); err != nil {
			panic(err)
		}
	}

	if dht == nil {
		return nil, errors.New("error - no IPFS dht server has been initialized")
	}

	new_host := &Host{
		ctx:  ctx,
		Host: h,
		DHT:  dht,
	}

	log.Debugf("New peer with ID: %s", h.ID().String())
	return new_host, nil
}
