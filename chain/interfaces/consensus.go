package interfaces

import "errors"

type Consensus struct {
	Blockchain RegalChain
}

func (c Consensus) ValidateBlockchain() error {

	blocks := c.Blockchain.Blocks
	if len(blocks) <= 1 {
		return nil
	}

	currBlockIdx := len(blocks)-1
	prevBlockIdx := len(blocks)-2

	for prevBlockIdx >= 0 {
		currBlock := blocks[currBlockIdx]
		prevBlock := blocks[prevBlockIdx]
		if currBlock.Header.HashPrevBlock != prevBlock.Hash {
			return errors.New("blockchain has inconsistent hashes")
		}

		if currBlock.Header.Timestamp <= prevBlock.Header.Timestamp {
			return errors.New("blockchain has inconsistent timestamps")
		}
		currBlockIdx--
		prevBlockIdx--

	}

	return nil

}

func (c Consensus) ValidateGenesisBlock() error {

	return nil
}