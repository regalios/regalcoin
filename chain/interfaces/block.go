package interfaces

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/regalios/regalcoin/crypto"
	log "github.com/sirupsen/logrus"
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
	Index uint64	`badgerhold:"key"`
	Header *BlockHeader
	Hash *numbers.Uint256 `badgerholdIndex:"IdxBlockHash"`
	Height uint64 `badgerholdIndex:"IdxBlockHeight"`
	Tx []Transaction `badgerholdIndex:"IdxBlockTx"`
	Payload map[int]string
	Validators []string `badgerholdIndex:"IdxBlockValidators"`
	Size int
	IBlock
}

type GenesisBlock struct {
	b *Block
	GenesisTime int64
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
	genesis.b.Hash = nil
	genesis.b.Height = 0
	genesis.b.Tx = make([]Transaction,0)
	genesis.b.Payload = make(map[int]string, 0)
	genesis.b.Validators = make([]string, 0)
	genesis.GenesisTime = time.Now().UnixNano()

	ser := genesis.Serialize()

	hash := crypto.NewHashSHA3256(ser)
	genesis.b.Hash = hash

	ser = genesis.Serialize()

	genesis.b.Size = len(ser[:])

	chain.BlockCandidates[0] = genesis.b


	return genesis



}

func calculateHash(data []byte) *numbers.Uint256 {
	return numbers.NewUint256(data)
}

func (g *GenesisBlock) AddToQueue() error {

	txn := blockQueueDB.NewTransaction(true)
	defer txn.Discard()
	ser := g.Serialize()
	size := len(ser)
	g.b.Size = size
	g.b.Hash = calculateHash(ser)
	err := txn.Set([]byte("block0"), ser)
	if err != nil {
		log.Errorln("cannot add genesis block to queue", err)
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	log.Infoln("added genesis block to queue...")

	return nil

}

func (g *GenesisBlock) Serialize() []byte {


	s, _ := json.Marshal(g.b)
	return s

}

type BlockQueue struct {}

var blockQueueDB *badger.DB

func (q BlockQueue) GetInstance() {


	blockQueueDB, err := badger.Open(badger.DefaultOptions("data/chain/blockqueue"))
	if err != nil {
		log.Fatalln(err)
	}
	defer  blockQueueDB.Close()

}

func (q BlockQueue) ProcessBlocks() error {



	var block Block
	err := blockQueueDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				block := json.Unmarshal(v, &block)
				log.Infoln(k, block)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil

	})

	if err != nil {
		return err
	}


	return nil

}

func init() {

	chain := new(RegalChain)
	chain.ChainID = uuid.New().String()
	chain.BlockCandidates = make(map[int]*Block, 0)



	 g := new(GenesisBlock)
	g.Create(chain)
	 blockHead := chain.BlockCandidates[0]
	 if blockHead.Index == 0 {
	 	chain.Blocks = make(map[int]map[string]*Block, 0)
	 	blockHash := chain.BlockCandidates[0].Hash.String()
	 	chain.Blocks[0][blockHash] =  chain.BlockCandidates[0]
	 	chain.BlockCandidates[0] = nil
	 	chain.Genesis = blockHash
	 }

	 spew.Dump(chain)

}