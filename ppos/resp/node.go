package resp

import "math/big"

type Node struct {

	NodeId string `json:"NodeId"`

	StakingAddress string `json:"StakingAddress"`

	BenefitAddress string `json:"BenefitAddress"`

	RewardPer *big.Int `json:"RewardPer"`

	NextRewardPer *big.Int `json:"NextRewardPer"`

	StakingTxIndex *big.Int `json:"StakingTxIndex"`

	ProgramVersion *big.Int `json:"ProgramVersion"`

	Status *big.Int `json:"Status"`

	StakingEpoch *big.Int `json:"StakingEpoch"`

	StakingBlockNum *big.Int `json:"StakingBlockNum"`

	Shares *big.Int `json:"Shares"`

	Released *big.Int `json:"Released"`

	ReleasedHes *big.Int `json:"ReleasedHes"`

	RestrictingPlan *big.Int `json:"RestrictingPlan"`

	RestrictingPlanHes *big.Int `json:"RestrictingPlanHes"`

	ExternalId string `json:"ExternalId"`

	NodeName string `json:"NodeName"`

	Website string `json:"Website"`

	Details string `json:"Details"`

	ValidatorTerm *big.Int `json:"ValidatorTerm"`

	DelegateEpoch *big.Int `json:"DelegateEpoch"`

	DelegateTotal *big.Int `json:"DelegateTotal"`

	DelegateTotalHes *big.Int `json:"DelegateTotalHes"`

	DelegateRewardTotal *big.Int `json:"DelegateRewardTotal"`
}
