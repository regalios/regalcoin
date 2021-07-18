package interfaces

import (
	"fmt"
	"io"
	"regalcoin/chain/numbers/uint256"
	"sync"
)

// SEQUENCE_FINAL is the default for TxIn Sequence
const SEQUENCE_FINAL uint32 = 0xffffffff
const SEQUENCE_LOCKTIME_DISABLE_FLAG uint32 = 1 << 31
const SEQUENCE_LOCKTIME_TYPE_FLAG uint32 = 1 << 22
const SEQUENCE_LOCKTIME_MASK uint32 = 0x0000ffff
const SEQUENCE_LOCKTIME_GRANULARITY int = 9


type TxOut struct {
	value Amount
	scriptPubKey string
	ITxOut
}

type ITxOut interface {
	Create(value Amount, scriptPubKey string)
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
	ITxIn
}

type ITxIn interface {
	Create(prevoutIn Outpoint, scriptSigIn string, SequenceIn uint32)
	Serialize(rw io.ReadWriter, op string)
	ToString() string
}


type Transaction struct {
	mu sync.RWMutex
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
	Unserialize(tx string, rw io.ReadWriter)
	SerializeTx(tx string, rw io.ReadWriter)

}

type MutableTransaction struct {
	Vin []TxIn
	Vout []TxOut
	Version int32
	LockTime uint32
}

func (o Outpoint) ToString() string {
	return fmt.Sprintf("Outpoint(%s, %u)", o.hash.String()[0:10], o.n)
}

func (txin TxIn) Create(prevoutIn Outpoint, scriptSigIn string, SequenceIn uint32) {


	txin.PrevOut = prevoutIn
	txin.ScriptSig = scriptSigIn
	txin.Sequence = SequenceIn

}

func (txin TxIn) ToString() string {
	str := "TxIn("

	str += txin.PrevOut.ToString()
	if txin.PrevOut.IsNull() {
		str += fmt.Sprintf(", coinbase %s", txin.ScriptSig)
	}
	if !txin.PrevOut.IsNull() {
		str += fmt.Sprintf(", scriptSig = %s", txin.ScriptSig[0:24])
	}
	if txin.Sequence != SEQUENCE_FINAL {
		str += fmt.Sprintf(", Sequence=%u", txin.Sequence)
	}
	str += ")"
	return str

}

func (txout TxOut) 	Create(value Amount, scriptPubKey string) {
	txout.value = value
	txout.scriptPubKey = scriptPubKey
}

func (txout TxOut) ToString() string {
	return fmt.Sprintf("TxOut(Value=%d.%08d, scriptPubKey=%s", txout.value % 1e9, txout.scriptPubKey[0:30])

}
