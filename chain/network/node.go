package network

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/helpers"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	"regalcoin/chain/interfaces"
)



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
