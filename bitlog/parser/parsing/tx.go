package parsing

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func ParseMinerAddress(tx *wire.MsgTx) btcutil.Address {
	var minerAddress btcutil.Address
	for _, out := range tx.TxOut {
		_, minerAddresses, _, _ := txscript.ExtractPkScriptAddrs(out.PkScript, &chaincfg.TestNet3Params)
		if len(minerAddresses) > 0 {
			minerAddress = minerAddresses[0]
		}
	}
	return minerAddress
}
