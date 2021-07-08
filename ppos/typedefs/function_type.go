package typedefs

type FunctionType int

const (
	/**
	 * 发起质押
	 */
	STAKING_FUNC_TYPE FunctionType = 1000
	/**
	 * 修改质押信息
	 */
	UPDATE_STAKING_INFO_FUNC_TYPE FunctionType = 1001
	/**
	 * 增持质押
	 */
	ADD_STAKING_FUNC_TYPE FunctionType = 1002
	/**
	 * 撤销质押(一次性发起全部撤销，多次到账)
	 */
	WITHDREW_STAKING_FUNC_TYPE FunctionType = 1003
	/**
	 * 发起委托
	 */
	DELEGATE_FUNC_TYPE FunctionType = 1004
	/**
	 * 减持/撤销委托(全部减持就是撤销)
	 */
	WITHDREW_DELEGATE_FUNC_TYPE FunctionType = 1005
	/**
	 * 查询当前结算周期的验证人队列
	 */
	GET_VERIFIERLIST_FUNC_TYPE FunctionType = 1100
	/**
	 * 查询当前共识周期的验证人列表
	 */
	GET_VALIDATORLIST_FUNC_TYPE FunctionType = 1101
	/**
	 * 查询所有实时的候选人列表
	 */
	GET_CANDIDATELIST_FUNC_TYPE FunctionType = 1102
	/**
	 * 查询当前账户地址所委托的节点的NodeID和质押Id
	 */
	GET_DELEGATELIST_BYADDR_FUNC_TYPE FunctionType = 1103
	/**
	 * 查询当前单个委托信息
	 */
	GET_DELEGATEINFO_FUNC_TYPE FunctionType = 1104
	/**
	 * 查询当前节点的质押信息
	 */
	GET_STAKINGINFO_FUNC_TYPE FunctionType = 1105
	/**
	 * 查询当前结算周期的区块奖励
	 */
	GET_PACKAGEREWARD_FUNC_TYPE FunctionType = 1200
	/**
	 * 查询当前结算周期的质押奖励
	 */
	GET_STAKINGREWARD_FUNC_TYPE FunctionType = 1201
	/**
	 * 查询打包区块的平均时间
	 */
	GET_AVGPACKTIME_FUNC_TYPE FunctionType = 1202
	/**
	 * 提交文本提案
	 */
	SUBMIT_TEXT_FUNC_TYPE FunctionType = 2000
	/**
	 * 提交升级提案
	 */
	SUBMIT_VERSION_FUNC_TYPE FunctionType = 2001
	/**
	 * 提交参数提案
	 */
	SUBMIT_PARAM_FUNCTION_TYPE FunctionType = 2002
	/**
	 * 给提案投票
	 */
	VOTE_FUNC_TYPE FunctionType = 2003
	/**
	 * 版本声明
	 */
	DECLARE_VERSION_FUNC_TYPE FunctionType = 2004
	/**
	 * 提交取消提案
	 */
	SUBMIT_CANCEL_FUNC_TYPE FunctionType = 2005
	/**
	 * 查询提案
	 */
	GET_PROPOSAL_FUNC_TYPE FunctionType = 2100
	/**
	 * 查询提案结果
	 */
	GET_TALLY_RESULT_FUNC_TYPE FunctionType = 2101
	/**
	 * 查询提案列表
	 */
	GET_PROPOSAL_LIST_FUNC_TYPE FunctionType = 2102
	/**
	 * 查询提案生效版本
	 */
	GET_ACTIVE_VERSION FunctionType = 2103
	/**
	 * 查询当前块高的治理参数值
	 */
	GET_GOVERN_PARAM_VALUE FunctionType = 2104
	/**
	 * 查询提案的累积可投票人数
	 */
	GET_ACCUVERIFIERS_COUNT FunctionType = 2105
	/**
	 * 查询可治理列表
	 */
	GET_PARAM_LIST FunctionType = 2106
	/**
	 * 举报双签
	 */
	REPORT_DOUBLESIGN_FUNC_TYPE FunctionType = 3000
	/**
	 * 查询节点是否已被举报过多签
	 */
	CHECK_DOUBLESIGN_FUNC_TYPE FunctionType = 3001
	/**
	 * 创建锁仓计划
	 */
	CREATE_RESTRICTINGPLAN_FUNC_TYPE FunctionType = 4000
	/**
	 * 获取锁仓信息
	 */
	GET_RESTRICTINGINFO_FUNC_TYPE FunctionType = 4100
	/**
	 * 提取账户当前所有的可提取的委托奖励
	 */
	WITHDRAW_DELEGATE_REWARD_FUNC_TYPE FunctionType = 5000
	/**
	 * 查询账户在各节点未提取委托奖励
	 */
	GET_DELEGATE_REWARD_FUNC_TYPE FunctionType = 5100
)

func (ft FunctionType) GetType() int {
	return int(ft)
}
