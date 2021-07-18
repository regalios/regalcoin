package interfaces

import (
	"net"

)

type INode interface {
	SelectParams(network string)
	GetAssumedBlockchainSize() uint64
	GetAssumedChainStateSize() uint64
	GetNetworkName() string
	InitLogging()
	InitParameterInteraction()
	GetWarnings() string
	GetLogCategories() uint32
	BaseInitialize() bool
	AppInitMain() bool
	AppShutdown()
	StartShutdown()
	ShutdownRequested() bool
	SetupServerArgs()
	MapPort(useUpnp bool)
	GetProxy(Network net.Conn, proxyInfo int)
	GetNodeCount() uint32
	GetNodeStats() bool
	GetBanned() bool
	Ban(addr net.Addr, reason string, banTimeOffset int64) bool
	Unban(addr net.Addr) bool
	Disconnect(addr net.Addr) bool
	DisconnectByNodeID(nodeID string)
	GetTotalBytesRecv() int64
	GetTotalBytesSent() int64
	GetMempoolSize() uint32
	GetMempoolDynamicUsage() uint32
	GetHeaderTip(height int, blockTime int64) bool
	GetNumBlocks() int32
	GetLastBlockTime() int64
	GetVerificationProgress() float64
	IsInitialBlockDownload() bool
	GetReindex() bool
	GetImporting() bool
	SetNetworkActive(active bool)
	GetNetworkActive() bool
	GetMaxTxFee() float64
	EstimateSmartFee(numBlocks int, conservative bool, returnedTarget int) interface{}
	GetDustRelayFee() interface{}
	ExecuteRpc(command string, params interface{}, uri string)
	ListRpcCommands() map[int]string
	GetUnspentOutput(outpoint *Output, coin *Coin) bool
	GetWalletDir() string
	ListWalletDir() map[int]string
	GetWallets() map[int]*Wallet
	LoadWallet(name string, err string, warnings string) *Wallet




}
