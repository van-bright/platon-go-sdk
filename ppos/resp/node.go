package resp

import "math/big"

type Node struct {
	NodeId string `json:"NodeId"`

	BlsPubKey string `json:"BlsPubKey"`

	StakingAddress string `json:"StakingAddress"`

	BenefitAddress string `json:"BenefitAddress"`

	RewardPer *big.Int `json:"RewardPer"`

	NextRewardPer *big.Int `json:"NextRewardPer"`

	StakingTxIndex *big.Int `json:"StakingTxIndex"`

	ProgramVersion *big.Int `json:"ProgramVersion"`

	Status *big.Int `json:"Status"`

	StakingEpoch *big.Int `json:"StakingEpoch"`

	StakingBlockNum *big.Int `json:"StakingBlockNum"`

	Shares string `json:"Shares"`

	Released string `json:"Released"`

	ReleasedHes string `json:"ReleasedHes"`

	RestrictingPlan string `json:"RestrictingPlan"`

	RestrictingPlanHes string `json:"RestrictingPlanHes"`

	DelegateEpoch *big.Int `json:"DelegateEpoch"`

	DelegateTotal string `json:"DelegateTotal"`

	DelegateTotalHes string `json:"DelegateTotalHes"`

	DelegateRewardTotal string `json:"DelegateRewardTotal"`

	ExternalId string `json:"ExternalId"`

	NodeName string `json:"NodeName"`

	Website string `json:"Website"`

	Details string `json:"Details"`

	//ValidatorTerm *big.Int `json:"ValidatorTerm"`
}
