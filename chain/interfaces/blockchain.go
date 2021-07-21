package interfaces

import (
	"regalcoin/chain/numbers"
)

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
	GetBlocks() []*Block
	GetGenesis()
	GetSuperValidator(address string) (error, *SuperValidator)
	GetValidators() (error, []*Node)
	GenerateNewBlock(Validator *Node)
}


type RegalChain struct {

	IBlockchain
	Version uint32
	ChainID string
	Genesis string
	LastHeight uint64
	SuperValidators []string
	Validators []*Validator
	Blocks map[string]*Block
	BlockCandidates []*Block

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