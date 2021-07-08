package typedefs

import "math/big"

type ProgramVersion struct {
	Version *big.Int `json:"Version"`
	Sign    string   `json:"Sign"`
}
