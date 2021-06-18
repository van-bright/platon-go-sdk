package resp

import (
	"platon-go-sdk/common/hexutil"
)

type Reward struct {
	NodeId     string       `json:"nodeId"`
	StakingNum int          `json:"stakingNum"`
	Reward     *hexutil.Big `json:"reward"`
}
