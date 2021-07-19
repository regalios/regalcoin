package interfaces

import (
	"github.com/libp2p/go-libp2p-core/network"
	"regalcoin/chain/numbers/uint256"
)

type IBlockchain interface {
	GetHeight() uint64
	GetBlockHeight(hash uint256.Int) uint64
	GetBlockDepth(hash uint256.Int) uint64
	GetBlockHash(height int) uint256.Int
	GetBlockTime(height int) int64
	GetBlockMedianTimePast(height int) int64
	HaveBlockOnDisk(height int) bool
	FindFirstBlockWithTime(timestamp int64, hash uint256.Int) int
	FindFirstBlockWithTimeAndHeight(timestamp int64, height int) int
	FindPruned(start_height, stop_height int) int
	FindFork(hash uint256.Int, height int) int
	IsPotentialTip(hash uint256.Int) bool
	GetTipLocator() interface{}
	FindBlock(hash uint256.Int)
	GetGenesis()

}

type IBlockHeader interface {
	SetNull()
	Serialization(stream network.Stream, op string)
	IsNull() bool
	GetHash() uint256.Int
	GetPowHash() uint256.Int
	GetBlockTime() int64

}

type BlockHeader struct {
	Version uint16
	HashPrevBlock uint256.Int
	HashMerkleRoot uint256.Int
	Timestamp int64
	Bits uint32
	Nonce uint32
}

type IBlock interface {
	IBlockHeader
	GetBlockHeader() BlockHeader
	ToString() string
}

type IBlockLocator interface {
	SetNull()
	Serialization(stream network.Stream, op string)
	IsNull() bool
}

type RegalChain struct {

	IBlockchain
	Version uint32
	ChainID string
	Genesis string
	LastHeight uint64



}