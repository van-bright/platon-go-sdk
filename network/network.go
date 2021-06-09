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
	MainNetParams = &PposNetworkParameters{
		DefaultMainNetConfig,
		MainNetHrp,
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp7pn3ep",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzsjx8h7",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrdyjj2v",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyva9ztf",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq93t3hkm",
		"lat1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxlcypcy",
	}

	TestNetParams = &PposNetworkParameters{
		DefaultTestNetConfig,
		TestNetHrp,
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp3yp7hw",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzlh5ge3",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqrzpqayr",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqyrchd9x",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq97wrcc5",
		"lax1zqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxsakwkt",
	}
)
