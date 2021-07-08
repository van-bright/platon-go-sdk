package typedefs

import "math/big"

type DuplicateSignType int

const (
	PREPARE_BLOCK DuplicateSignType = 1
	PREPARE_VOTE  DuplicateSignType = 2
	VIEW_CHANGE   DuplicateSignType = 3
)

func (dst DuplicateSignType) GetValue() *big.Int {
	return big.NewInt(int64(dst))
}
