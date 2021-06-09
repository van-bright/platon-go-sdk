package common

import "platon-go-sdk/rlp"

type Function struct {
	Type FunctionType
	InputParams []interface{}
}

func NewFunction(typ FunctionType, params []interface{}) *Function {
	return &Function{typ, params}
}

func (f *Function) ToBytes() ([]byte, error) {
	data, err := rlp.EncodeToBytes(f.Type.GetType())
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(f.InputParams); i++ {
		d, err := rlp.EncodeToBytes(f.InputParams[i])
		if err != nil {
			return nil, err
		}
		data = append(data, d...)
	}

	return data, nil
 }
