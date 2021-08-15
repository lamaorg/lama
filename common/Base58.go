package common

import (
	"fmt"
	"math/big"
)

type Base58 struct {
}

// Decode decodes a modified base58 string to a byte slice.
func (base58 Base58) Decode(b string) []byte {
	d, err := Decode(b)
	if err != nil {
		fmt.Print(err)
		d = nil
	}
	return d
}

// Encode encodes a byte slice to a modified base58 string.
func (base58 Base58) Encode(b []byte) string {
	return Encode(b)
}

type Alphabet struct {
	decode [128]int8
	encode [58]byte
}

// NewAlphabet creates a new alphabet from the passed string.
//
// It panics if the passed string is not 58 bytes long or isn't valid ASCII.
func NewAlphabet(s string) *Alphabet {
	if len(s) != 58 {
		panic("base58 alphabets must be 58 bytes long")
	}
	ret := new(Alphabet)
	copy(ret.encode[:], s)
	for i := range ret.decode {
		ret.decode[i] = -1
	}
	for i, b := range ret.encode {
		ret.decode[b] = int8(i)
	}
	return ret
}

// BTCAlphabet is the LLX base58 alphabet.
var BAlphabet = NewAlphabet("123456789ABCDEFGHIJKMNOPQRSTUVWYZabcdefghijkmnopqrstuvwxyz")

// Decode decodes a modified base58 string to a byte slice.

// ----- Implement Function -----
var (
	bn0  = big.NewInt(0)
	bn58 = big.NewInt(58)
)

// Encode encodes the passed bytes into a base58 encoded string.
func Encode(bin []byte) string {
	return FastBase58Encoding(bin)
}

// FastBase58Encoding encodes the passed bytes into a base58 encoded string.
func FastBase58Encoding(bin []byte) string {
	return FastBase58EncodingAlphabet(bin, BAlphabet)
}

// FastBase58EncodingAlphabet encodes the passed bytes into a base58 encoded
// string with the passed alphabet.
func FastBase58EncodingAlphabet(bin []byte, alphabet *Alphabet) string {
	zero := alphabet.encode[0]

	binsz := len(bin)
	var i, j, zcount, high int
	var carry uint32

	for zcount < binsz && bin[zcount] == 0 {
		zcount++
	}

	size := (binsz-zcount)*138/100 + 1
	var buf = make([]byte, size)

	high = size - 1
	for i = zcount; i < binsz; i++ {
		j = size - 1
		for carry = uint32(bin[i]); j > high || carry != 0; j-- {
			carry = carry + 256*uint32(buf[j])
			buf[j] = byte(carry % 58)
			carry /= 58
		}
		high = j
	}

	for j = 0; j < size && buf[j] == 0; j++ {
	}

	var b58 = make([]byte, size-j+zcount)

	if zcount != 0 {
		for i = 0; i < zcount; i++ {
			b58[i] = zero
		}
	}

	for i = zcount; j < size; i++ {
		b58[i] = alphabet.encode[buf[j]]
		j++
	}

	return string(b58)
}

// Decode decodes the base58 encoded bytes.
func Decode(str string) ([]byte, error) {
	return FastBase58Decoding(str)
}

// FastBase58Decoding decodes the base58 encoded bytes.
func FastBase58Decoding(str string) ([]byte, error) {
	return FastBase58DecodingAlphabet(str, BAlphabet)
}

// FastBase58DecodingAlphabet decodes the base58 encoded bytes using the given
// b58 alphabet.
func FastBase58DecodingAlphabet(str string, alphabet *Alphabet) ([]byte, error) {
	if len(str) == 0 {
		return nil, fmt.Errorf("zero length string")
	}

	var (
		t        uint64
		zmask, c uint32
		zcount   int

		b58u  = []rune(str)
		b58sz = len(b58u)

		outisz    = (b58sz + 3) / 4 // check to see if we need to change this buffer size to optimize
		binu      = make([]byte, (b58sz+3)*3)
		bytesleft = b58sz % 4

		zero = rune(alphabet.encode[0])
	)

	if bytesleft > 0 {
		zmask = (0xffffffff << uint32(bytesleft*8))
	} else {
		bytesleft = 4
	}

	var outi = make([]uint32, outisz)

	for i := 0; i < b58sz && b58u[i] == zero; i++ {
		zcount++
	}

	for _, r := range b58u {
		if r > 127 {
			return nil, fmt.Errorf("High-bit set on invalid digit")
		}
		if alphabet.decode[r] == -1 {
			return nil, fmt.Errorf("Invalid base58 digit (%q)", r)
		}

		c = uint32(alphabet.decode[r])

		for j := (outisz - 1); j >= 0; j-- {
			t = uint64(outi[j])*58 + uint64(c)
			c = uint32(t>>32) & 0x3f
			outi[j] = uint32(t & 0xffffffff)
		}

		if c > 0 {
			return nil, fmt.Errorf("Output number too big (carry to the next int32)")
		}

		if outi[0]&zmask != 0 {
			return nil, fmt.Errorf("Output number too big (last int32 filled too far)")
		}
	}

	// the nested for-loop below is the same as the original code:
	// switch (bytesleft) {
	// 	case 3:
	// 		*(binu++) = (outi[0] & 0xff0000) >> 16;
	// 		//-fallthrough
	// 	case 2:
	// 		*(binu++) = (outi[0] & 0xff00) >>  8;
	// 		//-fallthrough
	// 	case 1:
	// 		*(binu++) = (outi[0] & 0xff);
	// 		++j;
	// 		//-fallthrough
	// 	default:
	// 		break;
	// }
	//
	// for (; j < outisz; ++j)
	// {
	// 	*(binu++) = (outi[j] >> 0x18) & 0xff;
	// 	*(binu++) = (outi[j] >> 0x10) & 0xff;
	// 	*(binu++) = (outi[j] >>    8) & 0xff;
	// 	*(binu++) = (outi[j] >>    0) & 0xff;
	// }
	var j, cnt int
	for j, cnt = 0, 0; j < outisz; j++ {
		for mask := byte(bytesleft-1) * 8; mask <= 0x18; mask, cnt = mask-8, cnt+1 {
			binu[cnt] = byte(outi[j] >> mask)
		}
		if j == 0 {
			bytesleft = 4 // because it could be less than 4 the first time through
		}
	}

	for n, v := range binu {
		if v > 0 {
			start := n - zcount
			if start < 0 {
				start = 0
			}
			return binu[start:cnt], nil
		}
	}
	return binu[:cnt], nil
}
