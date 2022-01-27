package rpc

import "github.com/HydroProtocol/hydro-sdk-backend/sdk"

type IBlockChainRPC interface {
	GetCurrentBlockNum() (uint64, error)

	GetBlockByNum(uint64) (sdk.Block, error)
	GetLiteBlockByNum(uint64) (sdk.Block, error)
	GetTransactionReceipt(txHash string) (sdk.TransactionReceipt, error)

	GetLogs(from, to uint64, address string, topics []string) ([]sdk.IReceiptLog, error)
}
