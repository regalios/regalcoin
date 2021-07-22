package interfaces

import (
	"encoding/json"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/regalios/regalcoin/crypto"
	"regalcoin/chain/numbers"
	"time"
)


type IBlockHeader interface {
	SetNull()
	Serialization(stream network.Stream, op string)
	IsNull() bool
	GetHash() *numbers.Uint256
	GetPowHash() *numbers.Uint256
	GetBlockTime() int64

}



type BlockHeader struct {
	ChainID string
	Version uint16
	HashPrevBlock *numbers.Uint256
	HashMerkleRoot *numbers.Uint256
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

	Index int `storm:"id,increment"`
	Header *BlockHeader
	Hash string
	Height uint64 `storm:"index,increment"`
	Tx []Transaction
	Payload map[int]string
	Validators []string
	Size int
	IBlock `json:"-"`
}

type GenesisBlock struct {
	b *Block
	GenesisTime int64
}

func (b Block) NewBlock(chain *RegalChain) *RegalChain {

	block := new(Block)
	block.Header = new(BlockHeader)
	block.Header.Version = chain.Blocks[0].Header.Version
	block.Header.ChainID = chain.ChainID
	block.Header.HashMerkleRoot = new(numbers.Uint256)
	block.Header.HashPrevBlock = new(numbers.Uint256)
	block.Header.Bits = 0
	block.Header.Nonce = 0
	block.Header.Timestamp = time.Now().UnixNano()
	block.Index = chain.NumBlocks
	block.Height = uint64(chain.NumBlocks+1)

	ser, _ := json.Marshal(block)
	hash := crypto.NewHashSHA3256(ser)
	block.Hash = hash.String()
	block.Header.HashPrevBlock = crypto.NewHashSHA3256([]byte(chain.Blocks[block.Index-1].Hash))

	chain.Blocks = append(chain.Blocks, block)
	chain.NumBlocks++

	StoreBlock(chain.NetworkType, *block)

	return chain

}

func (g GenesisBlock) Create(chain *RegalChain) *GenesisBlock {

	genesis := new(GenesisBlock)
	genesis.b = new(Block)
	genesis.b.Header = new(BlockHeader)
	genesis.b.Header.Version = 0
	genesis.b.Header.HashMerkleRoot = new(numbers.Uint256)
	genesis.b.Header.HashPrevBlock = new(numbers.Uint256)
	genesis.b.Header.Bits = 0
	genesis.b.Header.Nonce = 0
	genesis.b.Header.Timestamp = time.Now().UnixNano()
	genesis.b.Index = 0
	genesis.b.Hash = ""
	genesis.b.Height = 0
	genesis.b.Tx = make([]Transaction,0)
	genesis.b.Payload = make(map[int]string, 0)
	genesis.b.Validators = make([]string, 0)
	genesis.GenesisTime = 	genesis.b.Header.Timestamp
	genesis.b.Header.ChainID = chain.ChainID


	ser := genesis.Serialize()

	hash := crypto.NewHashSHA3256(ser)
	genesis.b.Hash = hash.String()

	ser = genesis.Serialize()

	genesis.b.Size = len(ser[:])



	return genesis



}

func calculateHash(data []byte) *numbers.Uint256 {
	return numbers.NewUint256(data)
}


func (g *GenesisBlock) Serialize() []byte {


	s, _ := json.Marshal(g.b)
	return s

}
