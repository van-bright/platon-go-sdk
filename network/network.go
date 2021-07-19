package network

import "math/big"

const (
	MainNetHrp = "atp"
	TestNetHrp = "atx"
)

var (
	MainNetChainId = big.NewInt(201018)
	TestNetChainId = big.NewInt(201030)
)

type Config struct {
	Url     string
	ChainId *big.Int
}

var (
	DefaultTestNetConfig = Config{"http://47.241.91.2:6789", TestNetChainId}
	DefaultMainNetConfig = Config{"https://openapi.alaya.network/rpc", MainNetChainId}
	//DefaultMainNetConfig = Config{"http://172.16.64.132:6789", MainNetChainId}
)

type PposNetworkParameters struct {
	Config
	Hrp                     string
	RestrictingPlanContract string
	StakingContract         string
	IncentivePoolContract   string
	SlashContract           string
	ProposalContract        string
	RewardContract          string
}

var (
	PposMainNetParams = &PposNetworkParameters{
		DefaultMainNetConfig,
		MainNetHrp,
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp8h9fxw",
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzfyslg3",
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqr5jy24r",
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqy4tn65x",
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq9ga80f5",
		"atp1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxxwje8t",
	}

	PposTestNetParams = &PposNetworkParameters{
		DefaultTestNetConfig,
		TestNetHrp,
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp3yp7hw",
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzlh5ge3",
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrzpqayr",
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyrchd9x",
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq97wrcc5",
		"atx1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxsakwkt",
	}
)
