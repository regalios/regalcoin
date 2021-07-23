package interfaces

import "github.com/asdine/storm/v3"

const TxFee = float32(0.001)

type State struct {
	Balances map[string]uint
	Account2Nonce map[string]uint
	db *storm.DB
	latestBlock Block
	latestBlockHash string
	hasGenesisBlock bool
	validators []*Validator
	difficulty uint
	feePerBytes uint

}

func NewState(dataDir string, difficulty uint, validators []*Validator) (*State, error) {
	return &State{}, nil
}

func (s *State) AddBlocks(blocks []Block) error {
	for _, b := range blocks {
		err := s.AddBlock(b)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *State) AddBlock(block Block) error {
	return nil
}

func (s *State) Copy() State {
	c := State{}
	c.hasGenesisBlock = s.hasGenesisBlock
	c.latestBlock = s.latestBlock
	c.latestBlockHash = s.latestBlockHash
	c.Balances = make(map[string]uint)
	c.Account2Nonce = make(map[string]uint)
	c.difficulty = s.difficulty
	c.validators = s.validators

	for acc, balance := range s.Balances {
		c.Balances[acc] = balance
	}

	for acc, nonce := range s.Account2Nonce {
		c.Account2Nonce[acc] = nonce
	}

	return c
}

func (s *State) Close() error {
	return s.db.Close()
}

func (s *State) NextBlockNumber() int {
	if !s.hasGenesisBlock {
		return int(0)
	}

	return s.LatestBlock().Index + 1
}

func (s *State) LatestBlock() Block {
	return s.latestBlock
}

func (s *State) LatestBlockHash() string {
	return s.latestBlockHash
}

func (s *State) GetNextAccountNonce(account string) uint {
	return s.Account2Nonce[account] + 1
}

func (c *State) ChangeDifficulty(newDifficulty uint) {
	c.difficulty = newDifficulty
}