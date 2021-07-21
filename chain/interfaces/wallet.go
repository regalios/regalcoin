package interfaces

import (
	"io"
	"regalcoin/chain/numbers"
)

type IWallet interface{
	EncryptWallet(passphrase string) bool
	IsCrypted() bool
	Lock() bool
	Unlock(passphrase string) bool
	IsLocked() bool
	ChangeWalletPassphrase(oldPass string, newPass string) bool
	AbortRescan()
	GetWalletName() string
	BackupWallet(filename string) bool
	GetKeyFromPool(internal bool, pubKey *interface{}) bool
	GetPubKey(address string, pubKey *interface{}) bool
	GetPrivKey(address string, key *interface{}) bool
	IsSpendable(dest string) bool
	HaveWatchOnly() bool
	SetAddressBook(dest string, name string, purpose string) bool
	DelAddressBook(dest string) bool
	GetAddress(dest string, name string, isMine bool, purpose string) bool
	GetAddresses() []string
	AddDestData(dest string, key string, value string) bool
	EraseDestData(dest string, key string) bool
	GetDestValue(prefix string) []string
	LockCoin(output Outpoint)
	UnlockCoin(output Outpoint)
	IsLockecCoin(output Outpoint) bool
	ListLockedCoin(outputs []*Outpoint)
	CreateTransaction(recipients []string, coinControl *interface{}, sign bool, changePos int, fee float64, failReason string) *interface{}



}

type Wallet struct {

}


type Outpoint struct {
	IOutpoint
	hash numbers.Uint256
	n uint32
}

type IOutpoint interface {
	Create(hashIn numbers.Uint256, nIn uint32) Outpoint
	Serialize(rw io.ReadWriteCloser, serAction interface{})
	SetNull()
	IsNull() bool
	ToString() string
}


