package network

const (
	FULLSYNC_ADDR = "/regalcoin/chain/fullsync/1.0.0"
	SYNCFROM_ADDR = "/regalcoin/chain/syncfrom/1.0.0"
	SENDTX_ADDR = "/regalcoin/tx/send/1.0.0"
	GETTX_ADDR = "/regalcoin/tx/get/1.0.0"
	GETBLOCK_ADDR = "/regalcoin/chain/block/1.0.0"
	GETINFO_ADDR = "/regalcoin/chain/info/1.0.0"

	)

var ServiceDefinitions = []string{FULLSYNC_ADDR, SYNCFROM_ADDR, SENDTX_ADDR, GETTX_ADDR, GETBLOCK_ADDR, GETINFO_ADDR}