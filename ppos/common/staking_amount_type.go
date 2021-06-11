package common

type StakingAmountType int

const (
	FREE_AMOUNT_TYPE        StakingAmountType = 0
	RESTRICTING_AMOUNT_TYPE StakingAmountType = 1
)

func (sat StakingAmountType) GetValue() int {
	return int(sat)
}
