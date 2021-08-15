package Transaction

var TxType map[string]uint = make(map[string]uint, 0)

func init() {
	TxType["COINBASE"] = 0x0000
	TxType["WALLET"] = 0x1010
	TxType["CONTENT"] = 0x0201
	TxType["MINT"] = 0x0301
	TxType["BURN"] = 0x0302
	TxType["TRANSFER"] = 0x0303
	TxType["ALLOW"] = 0x0401
	TxType["SECURITY"] = 0x0555
	TxType["SYSTEM"] = 0x0999
	TxType["EVENT"] = 0x0606
	TxType["REVERT"] = 0x0990
	TxType["TOKEN"] = 0x0001
	TxType["VM"] = 0x7777
	TxType["CONTRACT"] = 0x7771
	TxType["SWAP"] = 0x7772
	TxType["EXCHANGE"] = 0x7773
	TxType["BUY"] = 0x7774
	TxType["REWARD"] = 0x7775

}
