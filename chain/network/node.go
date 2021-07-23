package network

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
	"github.com/regalios/regalcoin/interfaces"
)

const endpointBlocks = "regalcoin-blocks"
const endpointTxs = "regalcoin-txs"


type Node struct {
	Host host.Host
	config *libp2p.Config
	options *libp2p.Option
	Addr *peer.AddrInfo
	ID peer.ID
	ma []multiaddr.Multiaddr
	Stream network.Stream
	wallet *interfaces.Wallet
	chain *interfaces.RegalChain
	pubsub *pubsub.PubSub


}

func CreateNode(ctx context.Context) *Node {

	var node Node
	newNode, err := libp2p.New(ctx, libp2p.Defaults)
	if err != nil {
		panic(err)
	}
	ps, err := pubsub.NewFloodSub(ctx, newNode)
	if err != nil {
		panic(err)
	}

	for i, addr := range newNode.Addrs() {
		fmt.Printf("%d: %s/ipfs/%s\n", i, addr, newNode.ID().Pretty())
	}

	node.Host = newNode
	node.pubsub = ps
	node.wallet = new(interfaces.Wallet)
	node.chain = interfaces.NewChain(newNode, "local", 0)


	go interfaces.AddBlocksAtInterval(node.chain, 2)
	select {}

	return &node

}

