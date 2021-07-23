package interfaces


type Coin struct {
	Address string
	Signature []byte
	Coinbase int
	History coinHistory

}

type coinHistory struct {

	CreatedAt int64
	BlockTime int64
	BlockIndex uint32

}

type CoinSpecs struct {
	Creator Accounts
	Denom string
	DenomCode [8]byte
	TotalSupply Amount
	MaxMoney Amount
	Decimals int
}

type ICoin interface {

}