package interfaces

import (
	"regalcoin/chain/numbers/uint256"
	"io"
)

type TxOut struct {
	value Amount
	scriptPubKey interface{}
	ITxOut
}

type ITxOut interface {
	Create(value Amount, scriptPubKey interface{})
	Serialize(rw io.ReadWriter, op string)
	SetNull()
	IsNull() bool
	ToString() string
}


type TxIn struct {
	PrevOut Outpoint
	ScriptSig string
	Sequence uint32
	ScriptWitness []byte
}


type Transaction struct {
	CurrentVersion int32
	MaxStdVersion int32
	Vin []TxIn
	Vout []TxOut
	Version int32
	LockTime uint32
	hash uint256.Int
	witnessHash uint256.Int
	ITransaction
}

type ITransaction interface {
	CreateNullTx() *Transaction
	ConvertMutableToTx() *Transaction
	Serialize(rw io.ReadWriter)
	IsNull() bool
	GetHash() *uint256.Int
	GetWitnessHash() *uint256.Int
	ComputeHash() *uint256.Int
	ComputeWitnessHash() *uint256.Int
	GetValueOut() Amount
	GetTotalSize() int
	IsCoinbase() bool
	ToString() string
	HasWitness() bool

}

type MutableTransaction struct {
	Vin []TxIn
	Vout []TxOut
	
}