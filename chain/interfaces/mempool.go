package interfaces

type LockPoints struct {
	height int
	timestamp int64
	maxInputBlock interface{}
}

type ILockPoints interface {
	create(height int, timestamp int64, maxInputBlock interface{})
}

type MempoolEntry struct {
	Tx *interface{}
	Fee float64
	TxWeight uint64
	UsageSize uint64
	Timestamp int64
	EntryHeight uint
	SpendsCoinbase bool
	SigOpCost int64
	FeeDelta int64
	Lockpoints LockPoints
	CountWithDescendants uint64
	SizeWithDescendants uint64

}