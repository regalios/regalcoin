package interfaces

import "github.com/wealdtech/go-merkletree"

func NewMerkleTree(blocks []*Block) *merkletree.MerkleTree {
	data := make([][]byte,0)
	for _, block := range blocks {
		data = append(data, []byte(block.Hash))
	}

	tree, err := merkletree.New(data)
	if err != nil {
		panic(err)
	}

	return tree
}