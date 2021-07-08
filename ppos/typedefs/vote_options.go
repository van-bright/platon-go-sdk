package typedefs

type VoteOption uint16

const (
	YEAS        VoteOption = 1
	NAYS        VoteOption = 2
	ABSTENTIONS VoteOption = 3
)

func (vo VoteOption) GetValue() uint16 {
	return uint16(vo)
}
