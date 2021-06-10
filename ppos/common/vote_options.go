package common

type VoteOption int

const (
	YEAS        VoteOption = 1
	NAYS        VoteOption = 2
	ABSTENTIONS VoteOption = 3
)

func (vo VoteOption) GetValue() int {
	return int(vo)
}
