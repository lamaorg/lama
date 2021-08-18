package common

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

/*

	As this is super cool Llama no drama chain
	all general hex must start with Ll 0x108 (lowercase) 0x76 (capitalized)
	if the hex is having a suffix then it must be the final of the word drama 0x77 0x65 (capitalized)

	the context is a 8 byte context to inform about the purpose of the hexlified content

	8 BYTE CONTEXTS POSSIBLE (some may be added later)

	IMAGES = 0x494d4753
	BINARY DATA = 0x42444154
	ADDRESS = 0x41444452
	INTERNAL = 0x494e5452 (equivalent of system purpose only)
	LAMA FUN / MEME = 0x4c414d41 (LAMA)
	LAMA CASH = 0x4c414d24 (LAM$)
	SCRIPT = 0x53435258 (SCRX)
	Transaction = 0x4c415458 (LATX)
    NONE = 0x4e554c4c (NULL)
	WALLET = 0x4c57414c (LWAL)


*/

var HexContextsMap map[string]uint = make(map[string]uint)

func setupContexts() {
	HexContextsMap["IMAGES"] = 0x494d4753
	HexContextsMap["BINARY"] = 0x42444154
	HexContextsMap["ADDRESS"] = 0x41444452
	HexContextsMap["INTERNAL"] = 0x494e5452
	HexContextsMap["FUNMEME"] = 0x4c414d41
	HexContextsMap["COIN"] = 0x4c414d24
	HexContextsMap["SCRIPT"] = 0x53435258
	HexContextsMap["TX"] = 0x4c415458
	HexContextsMap["NULL"] = 0x4e554c4c
	HexContextsMap["WALLET"] = 0x4c57414c
}

type DecodingError struct {
	msg string
}

func (err DecodingError) Error() string {
	return err.msg
}

type Hexlified struct {
	hexEncodeFromByteNum int
	hexEncodeToByteNum   int
	hexLength            int
	hexContext           [8]byte
}

type HexWithPrefix struct {
	Hexlified
	hexPrefix [2]byte
}

type HexWithSuffix struct {
	Hexlified
	hexSuffix [2]byte
}

type HexTools interface {
	Encode()
	EncodeWithPrefix()
	EncodeWithSuffix()
	DecodeToString()
	DecodeToBytes()
	ValidateHex()
}

type HexOperator struct {
	HexTools
}

func (h *HexOperator) Encode() {}

func (h *HexOperator) EncodeWithPrefix() {}

func (h *HexOperator) EncodeWithSuffix() {}

func (h *HexOperator) DecodeToString() {}

func (h *HexOperator) DecodeToBytes() {}

func (h *HexOperator) ValidateHex() {}

func (h HexOperator) EncodeHex(toHash []byte, context string) string {

	ctx, err := h.GetContext(context)
	if err != nil {
		panic(err)
	}

	enc := make([]byte, len(toHash)*2+6)

	//ctxs := strconv.Itoa(int(ctx))
	//log.Println(ctxs)
	c := EncodeUint64(uint64(ctx))
	copy(enc, "Llx")
	copy(enc[3:], c)
	hex.Encode(enc[3:], toHash[:])
	return string(enc[:len(enc)-3]) + c

}

func (h HexOperator) GetContext(context string) (uint, error) {

	if val, ok := HexContextsMap[context]; ok {

		return val, nil

	}
	return uint(0), errors.New("invalid context")

}

func GetEncoderDecoder() *HexOperator {

	op := new(HexOperator)
	return op

}

func EncodeNewHex(toEncode []byte, ctx string) string {

	var h HexOperator
	hext := h.EncodeHex(toEncode, ctx)

	return hext

}

var bigWordNibbles int

func init() {
	setupContexts()
	b, _ := new(big.Int).SetString("FFFFFFFFFF", 16)

	switch len(b.Bits()) {
	case 1:
		bigWordNibbles = 16
	case 2:
		bigWordNibbles = 8
	default:
		panic("weird big.Word size")
	}
	h := EncodeNewHex([]byte("this is a shitty sentence to test"), "ADDRESS")
	DecodeHex(h)
}

func EncodeBig(bigint *big.Int) string {
	nbits := bigint.BitLen()
	if nbits == 0 {
		return "Llx0"
	}
	return fmt.Sprintf("%#x", bigint)
}

// DecodeUint64 decodes a hex string with 0x prefix as a quantity.
/*
func DecodeUint64(input string) (uint64, error) {
	raw, err := checkNumber(input)
	if err != nil {
		return 0, nil
	}
	dec, err := strconv.ParseUint(raw, 16, 64)
	if err != nil {
		err = mapError(err)
	}
	return dec, nil
}

func mapError(err interface{}) interface{} {
	return nil
}

func checkNumber(input string) (interface{}, interface{}) {
	return nil, nil
}
*/
// MustDecodeUint64 decodes a hex string with LLx prefix as a quantity.
// It panics for invalid input.
/*func MustDecodeUint64(input string) uint64 {
	dec, err := DecodeUint64(input)
	if err != nil {
		panic(err)
	}
	return dec
}*/

// EncodeUint64 encodes i as a hex string with LLx prefix.
func EncodeUint64(i uint64) string {
	enc := make([]byte, 2, 10)
	copy(enc, "Lx")
	return string(strconv.AppendUint(enc, i, 16))
}

func (h HexOperator) DecodeHex(hexStr string) *DecodedHex {

	hexWithoutLLx := hexStr[3:]
	realHex := hexWithoutLLx[:len(hexWithoutLLx)-10]
	hexType := hexWithoutLLx[len(hexWithoutLLx)-8:]
	//log.Println("hex", realHex)
	//log.Println("type", hexType)

	var found bool

	var realType string
	found = false
	for found != true {
		for key, val := range HexContextsMap {

			uintval := EncodeUint64(uint64(val))[2:]
			//log.Println(uintval)

			//log.Println(hexType)
			if uintval == hexType {
				realType = key
				found = true
			}

		}
	}
	//log.Println(realType)
	dh := new(DecodedHex)
	dh.HexType = realType
	dh.HexVal = realHex
	dh.ByteVal = []byte(realHex)
	byteVal := make([]byte, len(dh.ByteVal)/2)
	_, err := hex.Decode(byteVal, []byte(realHex))
	if err != nil {
		return nil
	}
	dh.ToBytes = byteVal
	dh.StrVal = string(byteVal)
	return dh

}

func DecodeHex(h string) *DecodedHex {
	var h1 HexOperator
	return h1.DecodeHex(h)
}

type DecodedHex struct {
	HexType string
	HexVal  string
	ByteVal []byte
	StrVal  string
	ToBytes []byte
}
