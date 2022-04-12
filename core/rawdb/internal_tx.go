package rawdb

import (
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

func ReadInternalTxs(db ethdb.Reader, hash common.Hash, number uint64) (result []*types.InternalTx) {
	result = make([]*types.InternalTx, 0)

	data, err := db.Get(blockInternalTxsKey(number, hash))
	if err != nil {
		if err != leveldb.ErrNotFound {
			log.Error("Read DB error", "err", err)
		}
		return
	}

	err = rlp.DecodeBytes(data, &result)
	if err != nil {
		log.Error("Decode data error", "err", err)
		return
	}

	return
}

// WriteInternalTxs stores all the internal transactions belonging to a block.
func WriteInternalTxs(db ethdb.KeyValueWriter, hash common.Hash, number uint64, itxs types.InternalTxs) {
	// Convert the receipts into their storage form and serialize them
	storageITxs := make([]*types.InternalTxForStorage, len(itxs))
	for i, tx := range itxs {
		storageITxs[i] = (*types.InternalTxForStorage)(tx)
	}
	bytes, err := rlp.EncodeToBytes(storageITxs)
	if err != nil {
		log.Crit("Failed to encode block internal txs", "err", err)
	}
	log.Debug("internal txs", "hash", hash.String(), "number", number, "lens", len(itxs))
	// Store the flattened receipt slice
	if err := db.Put(blockInternalTxsKey(number, hash), bytes); err != nil {
		log.Crit("Failed to encode block internal txs", "err", err)
	}
}

// DeleteInternalTxs removes all internal transactions associated with a block hash.
func DeleteInternalTxs(db ethdb.KeyValueWriter, hash common.Hash, number uint64) {
	if err := db.Delete(blockInternalTxsKey(number, hash)); err != nil {
		log.Crit("Failed to delete block internal txs", "err", err)
	}
}
