package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"

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

func new_host(ctx context.Context, priv_key crypto.PrivKey, ip string, port string) (*Host, error) {
	log.Infof("Creating new host")
	multiaddress, err := ma.NewMultiaddr(fmt.Sprintf("/ipv4/%s/tcp/%s", ip, port))
	if err != nil {
		return nil, err
	}

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.ListenAddrs(multiaddress),
		libp2p.Identity(priv_key),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			//TODO missing kaddht.MessageSenderImpl(msgSender.init) from options
			//TODO missing kaddht.NetworkSizeHook(newHost.SaveNetworkSizeEstimate) from options
			dht, err := kaddht.New(ctx, h)
			return dht, err
		}))

	if err != nil {
		panic(err)
	}

	if dht == nil {
		return nil, errors.New("error - no IPFS dht server has been initialized")
	}

	new_host := &Host{
		ctx:  ctx,
		Host: h,
		DHT:  dht,
	}

	log.Infof("New local node with ID: %s", h.ID().String())
	return new_host, nil
}

func (h *Host) bootstrap(ctx context.Context) error {
	log.Infof("Trying to initiliaze nodes with bootstraps")
	successful_connections := int64(0)
	var wg sync.WaitGroup
	//IPFS bootstrap peers
	for _, p := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		log.Infof("Connecting to bootstrap peer %s", p.ID.String())
		wg.Add(1)
		go func(bootstrap_node peer.AddrInfo) {
			defer wg.Done()
			if err := h.Connect(ctx, p); err != nil {
				log.Errorf("unable to connect to: %s", p.String())
			} else {
				atomic.AddInt64(&successful_connections, 1)
			}
		}(p)
	}
	wg.Wait()
	if successful_connections > 0 {
		log.Infof("%d of successful connections", successful_connections)
	} else {
		return errors.New("Error trying to connect to bootstrap peers")
	}
	return nil
}

func (h *Host) peer_id() string {
	return h.ID().String()
}

func (h *Host) close() error {
	return h.Close()
}
