package pkg

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/routing"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
)

type Host struct {
	ctx context.Context
	host.Host
	DHT *kaddht.IpfsDHT
}

func NewHost(ctx context.Context, ip string, port string) (*Host, error) {
	log.Infof("Creating new host")
	multiaddress, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", ip, port))
	if err != nil {
		return nil, err
	}

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.ListenAddrs(multiaddress),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			//TODO missing kaddht.MessageSenderImpl(msgSender.init) from options
			//TODO missing kaddht.NetworkSizeHook(newHost.SaveNetworkSizeEstimate) from options
			dht, err := kaddht.New(ctx, h)
			return dht, err
		}))

	if err != nil {
		return nil, errors.Wrap(err, " while creating libp2p node or dht")
	}

	new_host := &Host{
		ctx:  ctx,
		Host: h,
		DHT:  dht,
	}

	log.Infof("New local node with ID: %s", h.ID().String())
	return new_host, nil
}

func (h *Host) Bootstrap(ctx context.Context) error {
	log.Infof("Trying to initiliaze nodes with bootstraps")
	successful_connections := int64(0)
	var wg sync.WaitGroup
	//IPFS bootstrap peers
	for _, p := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		log.Infof("Connecting to bootstrap peer %s", p.ID.String())
		wg.Add(1)
		go func(bootstrap_node *peer.AddrInfo) {
			defer wg.Done()
			if err := h.Connect(ctx, p); err != nil {
				log.Errorf("unable to connect to: %s", p.String())
			} else {
				atomic.AddInt64(&successful_connections, 1)
			}
		}(&p)
	}
	wg.Wait()
	if successful_connections > 0 {
		log.Infof("%d of successful connections", successful_connections)
	} else {
		return errors.New("Error trying to connect to bootstrap peers")
	}
	return nil
}

func (h *Host) PeerId() string {
	return h.ID().String()
}

func (h *Host) Close() error {
	return h.Close()
}
