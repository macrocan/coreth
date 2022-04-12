package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Action struct {
	From         common.Address `gencodec:"required" json:"from"`
	To           common.Address `gencodec:"required" json:"to"`
	Value        *big.Int       `gencodec:"required" json:"value"`
	Success      bool           `gencodec:"required" json:"success"`
	OpCode       string         `gencodec:"required" json:"opcode"`
	Depth        uint64         `gencodec:"required" json:"depth"`
	Gas          uint64         `gencodec:"required" json:"gas"`
	GasUsed      uint64         `gencodec:"required" json:"gas_used"`
	TraceAddress []uint64       `gencodec:"required" json:"trace_address"`
	Error        string         `gencodec:"required" json:"error,omitempty"`
}

type ActionFrame struct {
	Action
	Calls []ActionFrame
}

type InternalTx struct {
	TxHash      common.Hash `json:"transactionHash" gencodec:"required"`
	BlockHash   common.Hash `json:"blockHash,omitempty"`
	BlockNumber *big.Int    `json:"blockNumber,omitempty"`
	Actions     []*Action   `json:"logs" gencodec:"required"`
}

type InternalTxForStorage InternalTx

type InternalTxs []*InternalTx
