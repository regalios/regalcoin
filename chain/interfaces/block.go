package interfaces

import (

	"github.com/libp2p/go-libp2p-core/network"
	"regalcoin/chain/numbers"
	"regalcoin/chain/numbers/uint256"
)


type IBlockHeader interface {
	SetNull()
	Serialization(stream network.Stream, op string)
	IsNull() bool
	GetHash() numbers.Uint256
	GetPowHash() uint256.Int
	GetBlockTime() int64

}

abi.U256

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

type Block struct {
	Index uint64
	Header *BlockHeader
	Hash *uint256.Int
	Height uint64
	Tx []Transaction
	Payload map[int]string
	Validators []string
	IBlock
}

type GenesisBlock struct {
	Block
	GenesisTime int64

}
