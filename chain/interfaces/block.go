package interfaces

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v3"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/regalios/regalcoin/crypto"
	log "github.com/sirupsen/logrus"
	"regalcoin/chain/numbers"
	"time"
	"github.com/davecgh/go-spew/spew"
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
	Index uint64
	Header *BlockHeader
	Hash *numbers.Uint256
	Height uint64
	Tx []Transaction
	Payload map[int]string
	Validators []string
	Size int
	IBlock
}

type GenesisBlock struct {
	b *Block
	GenesisTime int64
}

func (g GenesisBlock) Create() *GenesisBlock {

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

	/*err := genesis.AddToQueue()
	if err != nil {
		panic(err)
	}*/
	spew.Dump(genesis)
	spew.Dump(ser)

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
	bq := new(BlockQueue)
	bq.GetInstance()
	 g := new(GenesisBlock)
	 g.Create()

}