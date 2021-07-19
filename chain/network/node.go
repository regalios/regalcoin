package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/helpers"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"os"
	"os/signal"
	"regalcoin/chain/interfaces"
	"syscall"
	"time"
)

const discoveryNamespace = "regalcoin"

type Node struct {
	ctx context.Context
	Host host.Host
	config *libp2p.Config
	options *libp2p.Option
	Addr *peer.AddrInfo
	ID peer.ID
	ma []multiaddr.Multiaddr
	Stream network.Stream
	interfaces.INode

}





func CreateNetwork(peerAddr string) {
	h, err := libp2p.New(context.Background(), libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		panic(err)
	}

	// Print this node's addresses and ID
	fmt.Println("Addresses:", h.Addrs())
	fmt.Println("ID:", h.ID())


	discoveryService, err := discovery.NewMdnsService(
		context.Background(),
		h,
		time.Second,
		discoveryNamespace,
	)

	if err != nil {
		panic(err)
	}


	defer discoveryService.Close()
	discoveryService.RegisterNotifee(&discoveryNotifee{h: h})

	if  peerAddr != "" {

		peerma, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/80")
		if err != nil {
			panic(err)
		}
		peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerma)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", peerAddrInfo.String())

		if err := h.Connect(context.Background(), *peerAddrInfo); err != nil {
			panic(err)
		}
		fmt.Println("Connected to", peerAddrInfo.String())

		/*s, err := h.NewStream(context.Background(), peerAddrInfo.ID, "/regalcoin/1.0.0")
		if err != nil {
			panic(err)
		}

		*/

	}


	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh

}

func (n *Node) SetHandlers(s network.Stream, protocolID protocol.ID) {

	mfunc, err := helpers.MultistreamSemverMatcher(protocolID)
	if err != nil {
		panic(err)
	}

	connectedOn := make(chan protocol.ID)

	handler := func(s network.Stream) {
		connectedOn <- s.Protocol()
		_ = s.Close()
	}

	n.Host.SetStreamHandlerMatch(protocolID, mfunc, handler)


}


type discoveryNotifee struct{
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(peerInfo peer.AddrInfo) {
	fmt.Println("found peer", peerInfo.String())

}