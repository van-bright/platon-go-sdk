package web3g

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

const uintBits = 32 << (uint64(^uint(0)) >> 63)

var (
	ErrEmptyString   = &decError{"empty hex string"}
	ErrSyntax        = &decError{"invalid hex string"}
	ErrMissingPrefix = &decError{"hex string without 0x prefix"}
	ErrOddLength     = &decError{"hex string of odd length"}
	ErrEmptyNumber   = &decError{"hex string \"0x\""}
	ErrLeadingZero   = &decError{"hex number with leading zero digits"}
	ErrUint64Range   = &decError{"hex number > 64 bits"}
	ErrUintRange     = &decError{fmt.Sprintf("hex number > %d bits", uintBits)}
	ErrBig256Range   = &decError{"hex number > 256 bits"}
)

var (
	bytesT = reflect.TypeOf(Bytes(nil))
	bigT    = reflect.TypeOf((*BigInt)(nil))
	uintT   = reflect.TypeOf(Uint(0))
	uint64T = reflect.TypeOf(Uint64(0))
)

type decError struct{ msg string }

func (err decError) Error() string { return err.msg }

type BigInt big.Int
type Bytes []byte
type Uint uint
type Uint64 uint64

var bigWordNibbles int

func init() {
	// This is a weird way to compute the number of nibbles required for big.Word.
	// The usual way would be to use constant arithmetic but go vet can't handle that.
	b, _ := new(big.Int).SetString("FFFFFFFFFF", 16)
	switch len(b.Bits()) {
	case 1:
		bigWordNibbles = 16
	case 2:
		bigWordNibbles = 8
	default:
		panic("weird big.Word size")
	}
}

func isString(input []byte) bool {
	return len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"'
}

func bytesHave0xPrefix(input []byte) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func checkNumberText(input []byte) (raw []byte, err error) {
	if len(input) == 0 {
		return nil, nil // empty strings are allowed
	}
	if !bytesHave0xPrefix(input) {
		return nil, ErrMissingPrefix
	}
	input = input[2:]
	if len(input) == 0 {
		return nil, ErrEmptyNumber
	}
	if len(input) > 1 && input[0] == '0' {
		return nil, ErrLeadingZero
	}
	return input, nil
}

func wrapTypeError(err error, typ reflect.Type) error {
	if _, ok := err.(*decError); ok {
		return &json.UnmarshalTypeError{Value: err.Error(), Type: typ}
	}
	return err
}

func errNonString(typ reflect.Type) error {
	return &json.UnmarshalTypeError{Value: "non-string", Type: typ}
}

// EncodeBig encodes bigint as a hex string with 0x prefix.
// The sign of the integer is ignored.
func EncodeBig(bigint *big.Int) string {
	nbits := bigint.BitLen()
	if nbits == 0 {
		return "0x0"
	}
	return fmt.Sprintf("%#x", bigint)
}

// ToInt converts b to a big.Int.
func (b *BigInt) ToInt() *big.Int {
	return (*big.Int)(b)
}

//
//// String returns the hex encoding of b.
func (b *BigInt) String() string {
	return EncodeBig(b.ToInt())
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *BigInt) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return errNonString(bigT)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), bigT)
}

//func (b *BigInt) MarshalJSON() ([]byte, error) {
//	fmt.Println("marshalJson: ", b.String())
//	return []byte(b.String()), nil
//}

func (b *BigInt) MarshalText() ([]byte, error) {
	return []byte(EncodeBig((*big.Int)(b))), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (b *BigInt) UnmarshalText(input []byte) error {
	raw, err := checkNumberText(input)
	if err != nil {
		return err
	}
	if len(raw) > 64 {
		return ErrBig256Range
	}
	words := make([]big.Word, len(raw)/bigWordNibbles+1)
	end := len(raw)
	for i := range words {
		start := end - bigWordNibbles
		if start < 0 {
			start = 0
		}
		for ri := start; ri < end; ri++ {
			nib := decodeNibble(raw[ri])
			if nib == badNibble {
				return ErrSyntax
			}
			words[i] *= 16
			words[i] += big.Word(nib)
		}
		end = start
	}
	var dec big.Int
	dec.SetBits(words)
	*b = (BigInt)(dec)
	return nil
}

const badNibble = ^uint64(0)

func decodeNibble(in byte) uint64 {
	switch {
	case in >= '0' && in <= '9':
		return uint64(in - '0')
	case in >= 'A' && in <= 'F':
		return uint64(in - 'A' + 10)
	case in >= 'a' && in <= 'f':
		return uint64(in - 'a' + 10)
	default:
		return badNibble
	}
}

//func (trl *PlatonTransactionLog) UnmarshalJSON(input []byte) error {
//	var logs []string
//	err := json.Unmarshal(input, &logs);
//	if err == nil {
//		trl.Hashes = logs
//	} else {
//		var pt PlatonTransactionLog
//		err = json.Unmarshal(input, &pt)
//	}
//
//	return err
//}
