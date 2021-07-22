package interfaces

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regalcoin/chain/config"
	"regalcoin/chain/numbers"
	"time"
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
	Name string
	NetworkType string
	Version uint32
	ChainID string
	Genesis string
	LastHeight uint64
	SuperValidators []string
	Validators []*Validator
	Blocks []*Block
	BlockCandidates map[int]*Block
	BlockMemStorage map[int]*Block
	NumBlocks int
	//our priority queue
	blockQueue Queue
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

func NewChain(networkType string, version uint32)  {


	r := new(RegalChain)
	r.ChainID = uuid.New().String()
	r.Name = "RegalChain"
	r.NetworkType = networkType
	r.Version = version
	r.BlockCandidates = make(map[int]*Block, 0)
	r.config = (*config.Config)(config.ChainConfig)
	r.NewGenesis()


	_ = r.StoreValidBlock(r.Blocks[0])

	r.GetTotalBlocks()
	// start node here



	go AddBlocksAtInterval(r, 2)
	select {}



}

func AddBlocksAtInterval(r *RegalChain ,n time.Duration) {
	var B Block

	for now := range time.Tick(time.Second * n) {

		log.Infoln(now)
		r1 := B.NewBlock(r)


		log.Infoln(fmt.Sprintf("new block added at height: %v with hash: %s ", r1.NumBlocks, r1.Blocks[r1.NumBlocks-1].Hash))

	}
}

func (r *RegalChain) StoreValidBlock(b *Block) error {

	StoreBlock(r.NetworkType, *b)
	return nil

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

func (r *RegalChain) NewGenesis() {
	genesis := new(GenesisBlock)
	g := genesis.Create(r)
	r.Blocks = make([]*Block, 0)

	r.Blocks = append(r.Blocks, g.b)
	r.Genesis = g.b.Hash

	IOF := new(IO)
	genBytes, _ := json.Marshal(g.b)

	_ = IOF.WriteToDisk(genBytes, "data/chain/"+r.NetworkType+"/genesis.dat")

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