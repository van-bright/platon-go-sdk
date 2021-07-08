package typedefs

import (
	"reflect"
)

const (
	/**
	 * 默认gasLimit固定值
	 */
	BASE_DEFAULT_GAS_LIMIT = uint64(21000)
	/**
	 * 非0值gasLimit基数
	 */
	BASE_NON_ZERO_GAS_LIMIT = uint64(68)
	/**
	 * 0值gasLimit基数
	 */
	BASE_ZERO_GAS_LIMIT = uint64(4)
)

func IsLocalSupportFunction(ftype FunctionType) bool {
	switch ftype {
	case
		STAKING_FUNC_TYPE,
		UPDATE_STAKING_INFO_FUNC_TYPE,
		ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE,
		SUBMIT_TEXT_FUNC_TYPE,
		SUBMIT_VERSION_FUNC_TYPE,
		SUBMIT_PARAM_FUNCTION_TYPE,
		SUBMIT_CANCEL_FUNC_TYPE,
		VOTE_FUNC_TYPE,
		DECLARE_VERSION_FUNC_TYPE,
		REPORT_DOUBLESIGN_FUNC_TYPE,
		CREATE_RESTRICTINGPLAN_FUNC_TYPE:
		return true
	default:
		return false
	}
}

func GetGasLimit(f *Function) uint64 {
	if IsLocalSupportFunction(f.Type) {
		bytes := f.ToBytes()
		return BASE_DEFAULT_GAS_LIMIT +
			getContractGasLimit(f.Type) +
			getFunctionGasLimit(f.Type) +
			getInterfaceDynamicGasLimit(f.Type, f.InputParams) +
			getDataGasLimit(bytes)
	}
	return 0
}

func getContractGasLimit(ftype FunctionType) uint64 {
	switch ftype {
	case
		STAKING_FUNC_TYPE,
		UPDATE_STAKING_INFO_FUNC_TYPE,
		ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE,
		DELEGATE_FUNC_TYPE,
		WITHDREW_DELEGATE_FUNC_TYPE:
		return uint64(6000)
	case
		SUBMIT_TEXT_FUNC_TYPE,
		SUBMIT_VERSION_FUNC_TYPE,
		SUBMIT_PARAM_FUNCTION_TYPE,
		SUBMIT_CANCEL_FUNC_TYPE,
		VOTE_FUNC_TYPE,
		DECLARE_VERSION_FUNC_TYPE:
		return uint64(9000)
	case REPORT_DOUBLESIGN_FUNC_TYPE:
		return uint64(21000)
	case CREATE_RESTRICTINGPLAN_FUNC_TYPE:
		return uint64(18000)
	default:
		return uint64(0)
	}
}

/**
 * 根据接口类型，获取接口固定的gas消耗
 *
 * @param type
 * @return
 */
func getFunctionGasLimit(ftype FunctionType) uint64 {
	switch ftype {
	case SUBMIT_PARAM_FUNCTION_TYPE,
		SUBMIT_CANCEL_FUNC_TYPE:
		return uint64(500000)
	case SUBMIT_VERSION_FUNC_TYPE:
		return uint64(450000)
	case SUBMIT_TEXT_FUNC_TYPE:
		return uint64(320000)
	case STAKING_FUNC_TYPE:
		return uint64(32000)
	case UPDATE_STAKING_INFO_FUNC_TYPE:
		return uint64(12000)
	case ADD_STAKING_FUNC_TYPE,
		WITHDREW_STAKING_FUNC_TYPE:
		return uint64(20000)
	case DELEGATE_FUNC_TYPE:
		return uint64(16000)
	case WITHDREW_DELEGATE_FUNC_TYPE,
		CREATE_RESTRICTINGPLAN_FUNC_TYPE:
		return uint64(8000)
	case DECLARE_VERSION_FUNC_TYPE:
		return uint64(3000)
	case VOTE_FUNC_TYPE:
		return uint64(2000)
	case REPORT_DOUBLESIGN_FUNC_TYPE:
		return uint64(42000)
	default:
		return uint64(0)
	}
}

func getInterfaceDynamicGasLimit(ftype FunctionType, inputParameters []interface{}) uint64 {
	if ftype == CREATE_RESTRICTINGPLAN_FUNC_TYPE {
		if len(inputParameters) > 1 && reflect.Array == reflect.ValueOf(inputParameters[1]).Kind() {
			itemSize := reflect.ValueOf(inputParameters[1]).Len()
			return BASE_DEFAULT_GAS_LIMIT * uint64(itemSize)
		}
	}
	return uint64(0)
}

func getDataGasLimit(rlpData []byte) uint64 {
	var nonZeroSize uint64 = 0
	var zeroSize uint64 = 0

	for _, b := range rlpData {
		if b != 0 {
			nonZeroSize++
		} else {
			zeroSize++
		}
	}
	return nonZeroSize*BASE_NON_ZERO_GAS_LIMIT + zeroSize*BASE_ZERO_GAS_LIMIT
}
