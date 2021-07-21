package numbers

import (
	"crypto/sha256"
	"encoding"
	"fmt"
	"bytes"
	"hash"
	"math/big"
	"encoding/hex"
)

const Uint256IdxLen = 8  // The bigger it is, the more memory is needed, but lower chance of a collision

type Uint256 struct {
	Hash [32]byte
}

func NewUint256(h []byte) (res *Uint256) {
	res = new(Uint256)
	copy(res.Hash[:], h)
	return
}

// Get from MSB hexstring
func NewUint256FromString(s string) (res *Uint256) {
	d, e := hex.DecodeString(s)
	if e != nil {
		//println("NewUint256FromString", s, e.Error())
		return
	}
	if len(d)!=32 {
		//println("NewUint256FromString", s, "not 32 bytes long")
		return
	}
	res = new(Uint256)
	for i := 0; i<32; i++ {
		res.Hash[31-i] = d[i]
	}
	return
}


func NewSha2Hash(data []byte) (res *Uint256) {
	res = new(Uint256)
	ShaHash(data, res.Hash[:])
	return
}


func (u *Uint256) Bytes() []byte {
	return u.Hash[:]
}


func (u *Uint256) String() (s string) {
	for i := 0; i<32; i++ {
		s+= fmt.Sprintf("%02x", u.Hash[31-i])
	}
	return
}

func (u *Uint256) Equal(o *Uint256) bool {
	return bytes.Equal(u.Hash[:], o.Hash[:])
}

func (u *Uint256) Calc(data []byte) {
	ShaHash(data, u.Hash[:])
}


func BIdx(hash []byte) (o [Uint256IdxLen]byte) {
	copy(o[:], hash[:Uint256IdxLen])
	return
}

func (u *Uint256) BIdx() (o [Uint256IdxLen]byte) {
	o = BIdx(u.Hash[:])
	return
}

func (u *Uint256) BigInt() *big.Int {
	var buf [32]byte
	for i := range buf {
		buf[i] = u.Hash[31-i]
	}
	return new(big.Int).SetBytes(buf[:])
}

const (
	HASHER_TAPSIGHASH = 0
	HASHER_TAPLEAF    = 1
	HASHER_TAPBRANCH  = 2
	HASHER_TAPTWEAK   = 3
)

var (
	hash_tags = []string{"TapSighash", "TapLeaf", "TapBranch", "TapTweak"}
	hashers [HASHER_TAPTWEAK + 1][]byte
)

func _TaggedHash(tag string) hash.Hash {
	sha := sha256.New()
	sha.Write([]byte(tag))
	taghash := sha.Sum(nil)
	sha.Reset()
	sha.Write(taghash)
	sha.Write(taghash)
	return sha
}

func Hasher(idx int) hash.Hash {
	s := sha256.New()
	unmarshaler, _ := s.(encoding.BinaryUnmarshaler)
	unmarshaler.UnmarshalBinary(hashers[idx])
	return s
}

func init() {
	for i, t := range hash_tags {
		sha := _TaggedHash(t)
		marshaler, _ := sha.(encoding.BinaryMarshaler)
		hashers[i], _ = marshaler.MarshalBinary()
	}
}

func ShaHash(b []byte, out []byte) {
	s := sha256.New()
	s.Write(b[:])
	tmp := s.Sum(nil)
	s.Reset()
	s.Write(tmp)
	copy(out[:], s.Sum(nil))
}

