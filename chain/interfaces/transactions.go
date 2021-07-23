package interfaces

import "encoding/json"

func NewAccount(value string) string {
	return MasterKey([]byte(value)).Address()
}

type Transaction struct {
	From string	`json:"from"`
	To string	`json:"to"`
	Value uint	`json:"value"`
	Nonce uint	`json:"nonce"`
	Data TxData	`json:"data"`
	Time uint64	`json:"time"`
}

type TxData struct {
	txType string
	txFee uint
	txTypeImpl func()
}

type SignedTx struct {
	Transaction
	Sig []byte `json:"signature"`
}

type NFTTx struct {
	Transaction
	IpfsAssetLink string `json:"nft_enc_url"`
	Sig []byte `json:"signature"`
}


func NewTransaction(from, to string, value, nonce uint, data string) Transaction {
	return Transaction{}
}

func NewSignedTx(tx Transaction, sig []byte) SignedTx {
	return SignedTx{}
}

func (t Transaction) IsReward() bool {
	return t.Data.txType == "reward"
}

func (t Transaction) IsCoinbase() bool {
	return t.Data.txType == "coinbase"
}

func (t Transaction) IsNFT() bool {
	return t.Data.txType == "nft"
}

func (t Transaction) IsPrivate() bool {
	return t.Data.txType == "private"
}

func (t Transaction) Cost() uint {
	return t.Value + t.Data.txFee
}

func (t Transaction) Hash() (string, error) {
	txJson := t.Encode()

	return calculateHash(txJson), nil
}

func (t Transaction) Encode() []byte {
	te, _ :=  json.Marshal(t)
	return te
}

func (t SignedTx) Hash() (string, error) {
	txJson := t.Encode()

	return calculateHash(txJson), nil
}

func (t SignedTx) IsAuthentic() (bool, error) {
	return false, nil
}