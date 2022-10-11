package parsing

import (
	"fmt"
	"time"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"

	"github.com/1uvu/bitlog/pkg/utils"
)

type Block struct {
	MinerAddress btcutil.Address
	BlockHash    chainhash.Hash
	Header       wire.BlockHeader
	Txn          int64
	ConnectTime  time.Time
}

func ParseBlock(b *wire.MsgBlock, connectTime time.Time) *Block {
	genesisTX := ParseGenesisTX(b)
	minerAddress := ParseMinerAddress(genesisTX)
	return &Block{
		MinerAddress: minerAddress,
		BlockHash:    b.BlockHash(),
		Header:       b.Header,
		Txn:          int64(len(b.Transactions)),
		ConnectTime:  connectTime,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"{blockMiner=%s blockHash=%s blockHeader=%+v transactionNumber=%d timestampConnected=%s}",
		b.MinerAddress.EncodeAddress(),
		b.BlockHash.String(),
		b.Header,
		b.Txn,
		utils.TimeStrLocal(b.ConnectTime),
	)
}

func ParseGenesisTX(b *wire.MsgBlock) *wire.MsgTx {
	var genesisTX *wire.MsgTx
	for _, tx := range b.Transactions {
		if blockchain.IsCoinBaseTx(tx) {
			genesisTX = tx
			break
		}
	}
	return genesisTX
}
