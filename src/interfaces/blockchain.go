package interfaces

import (
	"github.com/regalios/regalcoin/uint256"
)

type Blockchain interface {
	GetHeight() uint64
	GetBlockHeight() uint64
	GetBlockDepth() uint64
	GetBlockHash() uint256

}