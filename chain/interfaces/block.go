package interfaces

import (
	"encoding/json"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/regalios/regalcoin/crypto"
	"io/ioutil"
	"regalcoin/chain/numbers"
	"time"
	mh "github.com/multiformats/go-multihash"
	cbor "github.com/ipfs/go-ipld-cbor"

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
	HashPrevBlock string
	HashMerkleRoot string
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

	CID *cid.Cid
	PrevCID *cid.Cid
	Index int `storm:"id,increment"`
	Header *BlockHeader
	Hash string
	Height uint64 `storm:"index,increment"`
	Tx []Transaction
	Payload map[string]string
	Size int
	IBlock `json:"-"`
}

type GenesisBlock struct {
	B *Block
	GenesisTime int64 `json:"genesis_time"`
	ChainID string	`json:"chain_id"`
	Symbol string `json:"symbol"`
	Balances map[string]uint `json:"balances"`
	Validators map[string]uint `json:"validators"`
}

func (b Block) NewBlock(chain *RegalChain) *RegalChain {

	block := new(Block)
	block.Header = new(BlockHeader)
	block.Header.Version = 0
	block.Header.ChainID = chain.ChainID
	block.Header.HashMerkleRoot = new(numbers.Uint256).String()
	block.Header.HashPrevBlock = new(numbers.Uint256).String()
	block.Header.Bits = 0
	block.Header.Nonce = 0
	block.Header.Timestamp = time.Now().UnixNano()
	block.Index = chain.NumBlocks
	block.Height = uint64(chain.NumBlocks+1)

	ser, _ := json.Marshal(block)
	hash := crypto.NewHashSHA3256(ser)
	block.Hash = hash.String()
	block.Header.HashPrevBlock = chain.GetHead().CID.String()

	chain.Blocks = append(chain.Blocks, block)
	chain.NumBlocks++

	StoreBlock(chain.NetworkType, *block)

	return chain

}

func init() {
	cbor.RegisterCborType(Block{})
	cbor.RegisterCborType(Transaction{})
	cbor.RegisterCborType(BlockHeader{})
}

func (b *Block) GetCid() cid.Cid {
	nd, err := cbor.WrapObject(b, mh.SHA3_256, -1)
	if err != nil {
		panic(err)
	}

	return nd.Cid()
}

func (g GenesisBlock) Create(chain *RegalChain) *GenesisBlock {

	genesis := new(GenesisBlock)
	genesis.B = new(Block)
	genesis.B.Header = new(BlockHeader)
	genesis.B.Header.Version = 0
	genesis.B.Header.HashMerkleRoot = ""
	genesis.B.Header.HashPrevBlock = ""
	genesis.B.Header.Bits = 0
	genesis.B.Header.Nonce = 0
	genesis.B.Header.Timestamp = time.Now().UnixNano()
	genesis.B.Index = 0
	genesis.B.Hash = ""
	genesis.B.Height = 0
	genesis.B.Tx = make([]Transaction,0)
	genesis.B.Payload = make(map[string]string, 0)
	genesis.Validators = make(map[string]uint, 0)
	genesis.GenesisTime = 	genesis.B.Header.Timestamp
	genesis.B.Header.ChainID = chain.ChainID
	genesis.ChainID = "regalcoin"


	ser := genesis.Serialize()

	hash := crypto.NewHashSHA3256(ser)
	genesis.B.Hash = hash.String()

	ser = genesis.Serialize()

	genesis.B.Size = len(ser[:])



	return genesis



}

func calculateHash(data []byte) string {
	return crypto.NewHashSHA3256(data).String()
}


func (g *GenesisBlock) Serialize() []byte {


	s, _ := json.Marshal(g.B)
	return s

}

func loadGenesis(path string) (GenesisBlock, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return GenesisBlock{}, err
		}

	var loadedGenesis GenesisBlock
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return GenesisBlock{}, err
	}
	return loadedGenesis, nil

}

func writeGenesisToDisk(path string, genesis []byte) error {
	return ioutil.WriteFile(path, genesis, 0644)
}