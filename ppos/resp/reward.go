package resp

import "math/big"

type Reward struct {
	NodeId     string
	StakingNum *big.Int
	Reward     *big.Int
}
