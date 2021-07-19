package crypto

import (
	"encoding/hex"
	"regalcoin/chain/numbers/uint256"
	"golang.org/x/crypto/sha3"
)

func NewHashSHA3256(data []byte) *uint256.Int {

	h := sha3.New256()
	h.Reset()
	_, _ = h.Write(data)

	hx :=  hex.EncodeToString(h.Sum(nil))
	num, err := uint256.FromHex(hx)
	if err != nil {
		panic(err)
	}

	return num

}
