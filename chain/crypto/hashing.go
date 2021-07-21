package crypto

import (
	"golang.org/x/crypto/sha3"
	"regalcoin/chain/numbers"
)

func NewHashSHA3256(data []byte) *numbers.Uint256 {

	h := sha3.New256()
	h.Reset()
	_, _ = h.Write(data)


	b := h.Sum(nil)
	num := numbers.NewUint256(b)


	return num

}
