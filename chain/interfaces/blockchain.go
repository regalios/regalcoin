package interfaces

import (
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
	GetHead() *Block
	FindBlock(hash uint256.Int)
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
	TotalStaked uint256.Int
	Children []*Validator
	RewardSettings *ValidatorRewardSettings
}

type ValidatorRewardSettings struct {

}