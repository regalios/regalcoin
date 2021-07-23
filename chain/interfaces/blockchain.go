package interfaces

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/multiformats/go-multihash"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regalcoin/chain/config"
	"regalcoin/chain/numbers"
	"time"
	"github.com/davecgh/go-spew/spew"
	bserv "github.com/ipfs/go-blockservice"
	bitswap "github.com/ipfs/go-bitswap"
	bnetwork "github.com/ipfs/go-bitswap/network"
	norouting "github.com/ipfs/go-ipfs-routing/none"
	bstore "github.com/ipfs/go-ipfs-blockstore"
)

var (
	Public []byte
	Private []byte
	PublicTest []byte
	PrivateTest []byte
)

func init() {
	Public, _ = hex.DecodeString("08b1d2be")
	Private, _ = hex.DecodeString("08b1d5bf")
	PublicTest, _ = hex.DecodeString("08b10808")
	PrivateTest, _ = hex.DecodeString("08b1b1b1")
}

type IBlockchain interface {
	GetHeight() uint64
	GetBlockHeight(hash numbers.Uint256) uint64
	GetBlockDepth(hash numbers.Uint256) uint64
	GetBlockHash(height int) numbers.Uint256
	GetBlockTime(height int) int64
	GetBlockMedianTimePast(height int) int64
	HaveBlockOnDisk(height int) bool
	FindFirstBlockWithTime(timestamp int64, hash numbers.Uint256) int
	FindFirstBlockWithTimeAndHeight(timestamp int64, height int) int
	FindPruned(start_height, stop_height int) int
	FindFork(hash numbers.Uint256, height int) int
	IsPotentialTip(hash numbers.Uint256) bool
	GetTipLocator() interface{}
	GetHead() *Block
	FindBlock(hash numbers.Uint256)
	GetBlocks() map[string]*Block
	GetGenesis()
	GetSuperValidator(address string) (error, *SuperValidator)
	GetValidators() (error, []*Node)
	GenerateNewBlock(Validator *Node)
	StoreValidBlock(b *Block) error
}


type RegalChain struct {

	IBlockchain `json:"-"`
	Head Block
	Name string
	NetworkType string
	Version uint32
	ChainID string
	Genesis string
	LastHeight uint64
	SuperValidators []string
	Validators []*Validator
	Blocks []*Block
	BlockStore bstore.Blockstore
	BlockCandidates map[int]*Block
	BlockMemStorage map[int]*Block
	NumBlocks int
	//our priority queue
	blockQueue Queue
	ChainFetcher bserv.BlockService
	config *config.Config

}


type Validator struct {
	Address string
	Staked Amount
	Parent *SuperValidator
}

type SuperValidator struct {
	Name string
	Url string
	Rank int
	Address string
	TotalStaked numbers.Uint256
	Children []*Validator
	RewardSettings *ValidatorRewardSettings
}

type ValidatorRewardSettings struct {

}

func NewChain(h host.Host, networkType string, version uint32) *RegalChain  {

	nr, _ := norouting.ConstructNilRouting(nil,nil,nil, nil)
	bsnet := bnetwork.NewFromIpfsHost(h, nr)
	dstore := datastore.NewMapDatastore()
	blocks := bstore.NewBlockstore(dstore)
	bswap := bitswap.New(context.Background(), bsnet,blocks)
	bservice := bserv.New(blocks, bswap)


	r := new(RegalChain)
	r.ChainID = uuid.New().String()
	r.Name = "RegalChain"
	r.NetworkType = networkType
	r.Version = version
	r.config = (*config.Config)(config.ChainConfig)
	r.ChainFetcher = bservice
	r.BlockStore = blocks
	block := r.NewGenesis()
	blkcopy := *block
	r.Head = blkcopy


	validBlockCID , err := r.StoreValidBlock(bservice, block)
	r.Head.CID = &validBlockCID
	if err != nil {
		panic(err)
	}


	r.GetTotalBlocks()
	// start node here




	return r

}

func (r *RegalChain) GetHead() Block {
	return r.Head
}

func AddBlocksAtInterval(r *RegalChain ,n time.Duration) {
	var B Block

	for now := range time.Tick(time.Second * n) {

		log.Infoln(now)
		r1 := B.NewBlock(r)
	//	B.Header.ChainID = r.ChainID
		r.Blocks = append(r.Blocks, &B)

		r.Head = B


		spew.Dump(r.Head)

		log.Infoln(fmt.Sprintf("new block added at height: %v with hash: %s ", r1.NumBlocks, r1.Blocks[r1.NumBlocks-1].Hash))

	}
}

func (r *RegalChain) StoreValidBlock(bs bserv.BlockService, b *Block) (cid.Cid, error)  {


	nd, err := cbor.WrapObject(b, multihash.BLAKE2B_MIN+31, 32)
	if err != nil {
		return cid.Cid{}, err
	}

	if err := bs.AddBlock(nd); err != nil {
		return cid.Cid{}, err
	}

	r.Head = *b

	return nd.Cid(), nil

}

func (r *RegalChain) GetBlocks() {

	allBlocks := GetAllBlocks(r.NetworkType)
	r.Blocks = allBlocks

}

func (r *RegalChain) resetBlockMemStorage() {
	r.BlockMemStorage = nil
}



func (r *RegalChain) GetTotalBlocks() {

	r.NumBlocks = len(r.Blocks)

}

func (r *RegalChain) NewGenesis() *Block {
	genesis := new(GenesisBlock)
	g := genesis.Create(r)
	r.Blocks = make([]*Block, 0)

	_, _ = r.StoreValidBlock(r.ChainFetcher, g.B)
	r.Genesis = g.B.Hash

	IOF := new(IO)
	genBytes, _ := json.Marshal(g.B)

	_ = IOF.WriteToDisk(genBytes, "data/chain/"+r.NetworkType+"/genesis.dat")

	return g.B

}

type IO struct {

}

func (iob *IO) WriteToDisk(item []byte, filename string) error {
	err := ioutil.WriteFile(filename, item, 0o777)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func (iob *IO) ReadFromDisk(filename string) ([]byte, error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return b, nil

}