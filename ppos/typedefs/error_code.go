package typedefs

const (
	/**
	 * 成功
	 */
	SUCCESS = 0
	/**
	 * 系统内部错误
	 */
	SYSTEM_ERROR = 1
	/**
	 * 对象没有找到
	 */
	OBJECT_NOT_FOUND = 2
	/**
	 * 参数错误
	 */
	INVALID_PARAMETER = 3
	/**
	 * bls key 长度有误
	 */
	WRONG_BLS_KEY_LENGTH = 301000
	/**
	 * bls key 证明有误
	 */
	WRONG_BLS_KEY_PROOF = 301001
	/**
	 * 节点描述信息长度有误
	 */
	WRONG_DESCRIPTION_LENGTH = 301002
	/**
	 * 程序版本签名有误
	 */
	WRONG_PROGRAM_VERSION_SIGN = 301003
	/**
	 * 程序的版本太低
	 */
	PROGRAM_VERSION_SIGN_TOO_LOW = 301004
	/**
	 * 版本声明失败
	 */
	DELCARE_VERSION_FAILED = 301005
	/**
	 * 发起交易账户必须和发起质押账户是同一个
	 */
	ADDRESS_MUST_SAME_AS_INITIATED_STAKING = 301006

	/**
	 * 质押的金额太低
	 */
	STAKING_DEPOSIT_TOO_LOW = 301100
	/**
	 * 候选人信息已经存在
	 */
	CANDIDATE_ALREADY_EXIST = 301101
	/**
	 * 候选人信息不存在
	 */
	CANDIDATE_NOT_EXIST = 301102
	/**
	 * 候选人状态已失效
	 */
	CANDIDATE_STATUS_INVALIDED = 301103
	/**
	 * 增持质押金额太低
	 */
	INCREASE_STAKE_TOO_LOW = 301104
	/**
	 * 委托金额太低
	 */
	DELEGATE_DEPOSIT_TOO_LOW = 301105
	/**
	 * 该账户不允许发起委托
	 */
	ACCOUNT_NOT_ALLOWED_DELEGATING = 301106
	/**
	 * 该候选人不接受委托
	 */
	CANDIDATE_NOT_ALLOWED_DELEGATE = 301107
	/**
	 * 撤销委托的金额太低
	 */
	WITHDRAW_DELEGATE_NOT_EXIST = 301108
	/**
	 * 委托详情不存在
	 */
	DELEGATE_NOT_EXIST = 301109
	/**
	 * von操作类型有误 (非自由金额或非锁仓金额)
	 */
	WRONG_VON_OPERATION_TYPE = 301110
	/**
	 * 账户的余额不足
	 */
	ACCOUNT_BALANCE_NOT_ENOUGH = 301111
	/**
	 * 区块高度和预期不匹配
	 */
	BLOCKNUMBER_DISORDERED = 301112
	/**
	 * 委托信息中余额不足
	 */
	DELEGATE_VON_NOT_ENOUGH = 301113
	/**
	 * 撤销委托时金额计算有误
	 */
	WRONG_WITHDRAW_DELEGATE_VON = 301114
	/**
	 * 验证人信息不存在
	 */
	VALIDATOR_NOT_EXIST = 301115
	/**
	 * 参数有误
	 */
	WRONG_FUNCTION_PARAM = 301116
	/**
	 * 惩罚类型有误
	 */
	WRONG_SLASH_TYPE = 301117
	/**
	 * 惩罚扣除的金额溢出
	 */
	SLASH_AMOUNT_TOO_LARGE = 301118
	/**
	 * 惩罚削减质押信息时金额计算有误
	 */
	WRONG_SLASH_CANDIDATE_VON = 301119
	/**
	 * 拉取结算周期验证人列表失败
	 */
	GETTING_VERIFIERLIST_FAILED = 301200
	/**
	 * 拉取共识周期验证人列表失败
	 */
	pGETTING_VALIDATORLIST_FAILED = 301201
	/**
	 * 拉取候选人列表失败
	 */
	GETTING_CANDIDATELIST_FAILED = 301202
	/**
	 * 拉取委托关联映射关系失败
	 */
	GETTING_DELEGATE_FAILED = 301203
	/**
	 * 查询候选人详情失败
	 */
	QUERY_CANDIDATE_INFO_FAILED = 301204
	/**
	 * 查询委托详情失败
	 */
	QUERY_DELEGATE_INFO_FAILED = 301205

	/**
	 * 链上生效版本没有找到
	 */
	ACTIVE_VERSION_NOT_FOUND = 302001
	/**
	 * 投票选项错误
	 */
	VOTE_OPTION_ERROR = 302002
	/**
	 * 提案类型错误
	 */
	PROPOSAL_TYPE_ERROR = 302003
	/**
	 * 提案ID为空
	 */
	PROPOSAL_ID_EMPTY = 302004
	/**
	 * 提案ID已经存在
	 */
	PROPOSAL_ID_ALREADY_EXISTS = 302005
	/**
	 * 提案没有找到
	 */
	PROPOSAL_NOT_FOUND = 302006
	/**
	 * PIPID为空
	 */
	PIPID_EMPTY = 302007
	/**
	 * PIPID已经存在
	 */
	PIPID_ALREADY_EXISTS = 302008
	/**
	 * 投票持续的共识轮数量太小
	 */
	ENDVOTINGROUNDS_TOO_SMALL = 302009
	/**
	 * 投票持续的共识轮数量太大
	 */
	ENDVOTINGROUNDS_TOO_LARGE = 302010
	/**
	 * 新版本的大版本应该大于当前生效版本的大版本
	 */
	NEWVERSION_SHOULD_LARGE_CURRENT_ACTIVE_VERSION = 302011
	/**
	 * 有另一个在投票期的升级提案
	 */
	ANOTHER_VERSION_PROPOSAL_AT_VOTING_STAGE = 302012
	/**
	 * 有另一个预生效的升级提案
	 */
	ANOTHER_VERSION_PROPOSAL_AT_PRE_ACTIVE_STAGE = 302013
	/**
	 * 有另一个在投票期的取消提案
	 */
	ANOTHER_CANCEL_PROPOSAL_AT_VOTING_STAGE = 302014
	/**
	 * 待取消的(升级)提案没有找到
	 */
	CANCELED_PROPOSAL_NOT_FOUND = 302015
	/**
	 * 待取消的提案不是升级提案
	 */
	CANCELED_PROPOSAL_NOT_VERSION_TYPE = 302016
	/**
	 * 待取消的(升级)提案不在投票期
	 */
	CANCELED_PROPOSAL_NOT_AT_VOTING_STAGE = 302017
	/**
	 * 提案人NodeID为空
	 */
	PROPOSER_EMPTY = 302018
	/**
	 * 验证人详情没有找到
	 */
	VERIFIER_DETAIL_INFO_NOT_FOUND = 302019
	/**
	 * 验证人状态为无效状态
	 */
	VERIFIER_STATUS_INVALID = 302020
	/**
	 * 发起交易账户和发起质押账户不是同一个
	 */
	TX_CALLER_DIFFER_FROM_STAKING = 302021
	/**
	 * 发起交易的节点不是验证人
	 */
	TX_CALLER_NOT_VERIFIER = 302022
	/**
	 * 发起交易的节点不是候选人
	 */
	TX_CALLER_NOT_CANDIDATE = 302023
	/**
	 * 版本签名错误
	 */
	VERSION_SIGN_ERROR = 302024
	/**
	 * 验证人没有升级到新版本
	 */
	VERIFIER_NOT_UPGRADED = 302025
	/**
	 * 提案不在投票期
	 */
	PROPOSAL_NOT_AT_VOTING_STAGE = 302026
	/**
	 * 投票重复
	 */
	VOTE_DUPLICATED = 302027
	/**
	 * 声明的版本错误
	 */
	DECLARE_VERSION_ERROR = 302028
	/**
	 * 把节点声明的版本通知Staking时出错
	 */
	NOTIFY_STAKING_DECLARED_VERSION_ERROR = 302029
	/**
	 * 提案结果没有找到
	 */
	TALLY_RESULT_NOT_FOUND = 302030
	/**
	 * 不支持的治理参数
	 */
	UNSUPPORTED_GOVERN_PARAMETER = 302031
	/**
	 * 有另一个在投票期的参数提案
	 */
	ANOTHER_PARAM_PROPOSAL_AT_VOTING_STAGE = 302032
	/**
	 * 参数提案的的参数值错误
	 */
	GOVERN_PARAMETER_VALUE_ERROR = 302033
	/**
	 * 参数提案的值必须和旧值不同
	 */
	PARAMETER_PROPOSAL_NEW_VALUE_SAME_AS_OLD_VALUE = 302034
	/**
	 * 双签证据校验失败
	 */
	DUPLICATE_SIGNATURE_VERIFICATION_FAILED = 303000
	/**
	 * 已根据该证据执行过惩罚
	 */
	PUNISHMENT_HAS_BEEN_IMPLEMENTED = 303001
	/**
	 * 举报的双签块高比当前区块高
	 */
	BLOCKNUMBER_TOO_HIGH = 303002
	/**
	 * 举报的证据超过有效期
	 */
	EVIDENCE_INTERVAL_TOO_LONG = 303003
	/**
	 * 获取举报的验证人信息失败
	 */
	GET_CERTIFIER_INFOMATION_FAILED = 303004
	/**
	 * 证据的地址和验证人的地址不匹配
	 */
	ADDRESS_NOT_MATCH = 303005
	/**
	 * 证据的节点ID和验证人的节点ID不匹配
	 */
	NODEID_NOT_MATCH = 303006
	/**
	 * 证据的blsPubKey和验证人的blsPubKey不匹配
	 */
	BLS_PUBKEY_NOT_MATCH = 303007
	/**
	 * 惩罚节点失败
	 */
	SLASH_NODE_FAILED = 303008
	/**
	 * 创建锁仓计划数不能为0或者大于36
	 */
	PARAM_EPOCH_CANNOT_BE_ZERO = 304001
	/**
	 * 创建锁仓计划数不能为0或者大于36
	 */
	RESTRICTING_PLAN_NUMBER_CANNOT_BE_0_OR_MORE_THAN_36 = 304002
	/**
	 * 锁仓创建总金额不能小于1E18
	 */
	TOTAL_RESTRICTING_AMOUNT_SHOULD_MORE_THAN_ONE = 304003
	/**
	 * 账户余额不够支付锁仓
	 */
	BALANCE_NOT_ENOUGH_FOR_RESTRICT = 304004
	/**
	 * 没有在锁仓合约中找到该账户
	 */
	RESTRICTING_CONTRACT_AMOUNT_NOT_FOUND = 304005
	/**
	 * 惩罚金额大于质押金额
	 */
	SLASH_AMOUNT_LARGER_THAN_STAKING_AMOUNT = 304006
	/**
	 * 惩罚锁仓账户的质押金额不能为0
	 */
	STAKING_AMOUNT_ZERO = 304007
	/**
	 * 锁仓转质押后回退的金额不能小于0
	 */
	AMOUNT_CANNOT_LESS_THAN_ZERO = 304008
	/**
	 * 锁仓信息中的质押金额小于回退的金额
	 */
	WRONG_STAKING_RETURN_AMOUNT = 304009
)
