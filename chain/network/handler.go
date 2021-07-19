package network

import (
	"github.com/libp2p/go-libp2p-core/network"
	"sync"
)

func (n *Node) StreamHandler(protocolID string, f ...func(s network.Stream)) int {



	if len(f) > 1 {

		var wg sync.WaitGroup
		wg.Add(len(f))
		wg.Wait()

		for _, fu := range f {

			go fu(n.Stream)

		}

		wg.Done()
		return 0
	}

	go f[0](n.Stream)

	return 0


}
