package common

type DuplicateSignType int
const (
	PREPARE_BLOCK DuplicateSignType = 1
	PREPARE_VOTE  DuplicateSignType = 2
	VIEW_CHANGE   DuplicateSignType = 3
)

func (dst DuplicateSignType) GetValue() int {
	return int(dst)
}
