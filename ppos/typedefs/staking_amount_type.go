package typedefs

type StakingAmountType uint16

const (
	FREE_AMOUNT_TYPE        StakingAmountType = 0
	RESTRICTING_AMOUNT_TYPE StakingAmountType = 1
)

func (sat StakingAmountType) GetValue() uint16 {
	return uint16(sat)
}
