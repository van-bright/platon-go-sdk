package codec

import (
	"math/big"
)

type RlpString struct {
	value []byte
}

func removeLeadingZeros(arr []byte) []byte {
	for i, b := range arr {
		if b != 0 {
			return arr[i:]
		}
	}
	return []byte{}
}

func RlpStringFromBig(v *big.Int) RlpString {
	str := &RlpString{}
	str.FromBig(v)
	return *str
}

func (rs *RlpString) FromBig(v *big.Int) {
	bytes := v.Bytes()
	rs.value = bytes //removeLeadingZeros(bytes)
}

func (rs *RlpString) GetBytes() []byte {
	return rs.value
}
