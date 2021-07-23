package interfaces

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	CreateTransaction(recipients []string, coinControl *interface{}, sign bool, changePos int, fee float64, failReason string) *interface{}



}

const DefaultWalletPath = "data/wallets/wallet.dat"

type Wallet struct {
	IWallet
	mWallet *HDWallet
	mAddress string
	addrBook *AddressBook
}

type AddressBook struct {
	items []*Address
}

type Address struct {
	id uint32
	address string
	wallet *HDWallet
	txs []Transaction
}

func NewWallet() *Wallet {
	seed, _ := GenSeed(1024)
	mWallet := MasterKey(seed)
	w := new(Wallet)
	w.mWallet = mWallet
	w.mAddress = mWallet.Address()
	w.addrBook = new(AddressBook)
	return w

}

func (w *Wallet) NewAddress() {

	if len(w.addrBook.items) >= 1 {
		id := w.addrBook.items[len(w.addrBook.items)-1].id+1
		newAddr, _ := w.mWallet.Child(id)
		Addr := new(Address)
		Addr.id = id
		Addr.address = newAddr.Address()
		Addr.wallet = newAddr
		w.addrBook.items = append(w.addrBook.items, Addr)
		w.WriteToDisk(DefaultWalletPath)
	} else {
		id := 1
		newAddr, _ := w.mWallet.Child(uint32(id))
		Addr := new(Address)
		Addr.id = uint32(id)
		Addr.address = newAddr.Address()
		Addr.wallet = newAddr
		w.addrBook.items = append(w.addrBook.items, Addr)
		w.WriteToDisk(DefaultWalletPath)
	}

}

func (w *Wallet) WriteToDisk(path string) {

	ser, _ := json.Marshal(w)
	err := ioutil.WriteFile(path, ser, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadWalletFromDisk(path string) *Wallet {


	var w Wallet
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			data, _ = json.Marshal(NewWallet())
		}

	}
	_ = json.Unmarshal(data, &w)
	return &w

}

func (w *Wallet) GetBalance() float32 {

	return 0.00

}