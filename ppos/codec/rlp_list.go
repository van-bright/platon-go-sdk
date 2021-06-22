package codec

type RlpList struct {
	values []interface{}
}

func (rl *RlpList) Append(args ...interface{}) {
	rl.values = append(rl.values, args...)
}

func (rl *RlpList) GetValues() []interface{} {
	return rl.values
}
