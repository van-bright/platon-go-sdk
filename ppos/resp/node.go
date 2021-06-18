package resp

import (
	"platon-go-sdk/common/hexutil"
)

type Node struct {
	NodeId string `json:"NodeId"`

	BlsPubKey string `json:"BlsPubKey"`

	StakingAddress string `json:"StakingAddress"`

	BenefitAddress string `json:"BenefitAddress"`

	RewardPer int `json:"RewardPer"`

	NextRewardPer int `json:"NextRewardPer"`

	StakingTxIndex int `json:"StakingTxIndex"`

	ProgramVersion int `json:"ProgramVersion"`

	Status int `json:"Status"`

	StakingEpoch int `json:"StakingEpoch"`

	StakingBlockNum uint64 `json:"StakingBlockNum"`

	Shares *hexutil.Big `json:"Shares"`

	Released *hexutil.Big `json:"Released"`

	ReleasedHes *hexutil.Big `json:"ReleasedHes"`

	RestrictingPlan *hexutil.Big `json:"RestrictingPlan"`

	RestrictingPlanHes *hexutil.Big `json:"RestrictingPlanHes"`

	DelegateEpoch int `json:"DelegateEpoch"`

	DelegateTotal *hexutil.Big `json:"DelegateTotal"`

	DelegateTotalHes *hexutil.Big `json:"DelegateTotalHes"`

	DelegateRewardTotal *hexutil.Big `json:"DelegateRewardTotal"`

	ExternalId string `json:"ExternalId"`

	NodeName string `json:"NodeName"`

	Website string `json:"Website"`

	Details string `json:"Details"`

	//ValidatorTerm *big.Int `json:"ValidatorTerm"`
}
