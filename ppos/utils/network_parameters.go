package utils

import "math/big"

type PposNetworkParameters struct {
	ChainId *big.Int
	Hrp     string
	RestrictingPlanContract string
	StakingContract string
	IncentivePoolContract string
	SlashContract string
	ProposalContract string
	RewardContract string
}


const (
	MainNetHrp = "atp"
	TestNetHrp = "atx"
)

var (
	MainNetChainId  = big.NewInt(201018)
	TestNetChainId = big.NewInt(201030)
)

var (
	MainNetParams = &PposNetworkParameters{
		MainNetChainId,
		MainNetHrp,
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzsjx8h7",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrdyjj2v",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxlcypcy",
	}

	TestNetParams = &PposNetworkParameters{
		TestNetChainId,
		TestNetHrp,
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp3yp7hw",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzlh5ge3",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrzpqayr",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyrchd9x",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq97wrcc5",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxsakwkt",
	}
)
