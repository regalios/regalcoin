package interfaces

type Consensus struct {
	Blockchain RegalChain
}

func (c Consensus) ValidateBlockchain() error {



	return nil

}

func (c Consensus) ValidateGenesisBlock() error {

	return nil
}