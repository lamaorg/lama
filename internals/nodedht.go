package internals

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	disc "github.com/libp2p/go-libp2p-discovery"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"io"
	mrand "math/rand"
)

type Node struct {
	privKey ci.PrivKey
	pubKey  ci.PubKey
	Addr    multiaddr.Multiaddr
	Host    host.Host
}

func (p *Node) NewHost(ctx context.Context, seed int64, port int) host.Host {
	var r io.Reader
	if seed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(seed))
	}
	priv, pub, err := ci.GenerateKeyPairWithReader(ci.RSA, 2048, r)
	if err != nil {
		panic(err)
	}
	p.privKey = priv
	p.pubKey = pub

	addr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	p.Addr = addr
	l, _ := libp2p.New(ctx, libp2p.ListenAddrs(addr), libp2p.Identity(priv))
	p.Host = l
	return l
}

func NewKDHT(ctx context.Context, host host.Host, bootstrapPeers []multiaddr.Multiaddr) (*disc.RoutingDiscovery, error) {
	var options []dht.Option
	var wg sync.WaitGroup

	if len(bootstrapPeers) == 0 {
		options = append(options, dht.Mode(dht.ModeServer))
	}

	kdht, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err = kdht.Bootstrap(ctx); err != nil {
		return nil, err
	}

	for _, peerAddr := range bootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				log.Printf("Error while connecting to node %q: %-v", peerinfo, err)
			} else {
				log.Printf("Connection established with bootstrap node: %q", *peerinfo)
			}
		}()
	}
	wg.Wait()

	return disc.NewRoutingDiscovery(kdht), nil
}
