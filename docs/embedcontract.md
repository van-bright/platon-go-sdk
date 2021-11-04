# 内置合约

`ppos`模块用来操作PlatON网络的内置合约, 包括`delegate_contract`, `node_contract`, `proposal_contract`, `restricting_plan_contract`, `reward_contract`, `slash_contract`, `staking_contract`.   
下面我们分别描述如何调用合约中的方法. 详细的使用方式, 请参考各内置合约的test文件, 如`delegate_contract_test.go`, 以获得完整的使用示例.

在调用内置合约之前, 我们需要提供`network.PposNetworkParameters`配置以及`typedefs.Credentials`配置, 这两个配置对于所有的内置合约都是适用的. 在以下的描述中, 我们假设都使用以下的`PposNetworkParameters`和`Credentials`:  
```go
const PrivateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"
var credentials, _ = typedefs.NewCredential(PrivateKey, network.MainNetHrp)

config := network.PposMainNetParams
```

## delegate_contract合约
delegate合约用于向节点发起/取消委托.

* 发起委托  
`func Delegate(nodeId string, stakingAmountType typedefs.StakingAmountType, amount *big.Int) (typedefs.TransactionHash, error)`  
Delegate方法接收以下参数:
    * nodeId: 委托节点的id
    * stakingAmountType: 委托类型, 可以是Free模式或锁仓模式
    * amount: 委托数量

返回交易的hash或error.

```go
config := network.PposMainNetParams
dc := NewDelegateContract(config, credentials)

nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"

amount := new(big.Int)
amount.SetString("200000000000000000000", 10)
list, err := dc.Delegate(nodeId, typedefs.FREE_AMOUNT_TYPE, amount)
if err != nil {
    t.Errorf("DelegateContract.Delegate failed: %s", err)
}

result, err := json.Marshal(list)
if err != nil {
    t.Errorf("Marshal of list failed: %s", err)
}

fmt.Println(string(result))
```
* 取消委托  
`func UnDelegate(nodeId string, stakingBlockNum *big.Int, amount *big.Int) (typedefs.TransactionHash, error)`  
参数说明:
    * nodeId: 委托节点的Id
    * stakingBlockNum: 质押节点的某个时刻块高
    * amount: 取消的数量

示例:
```go
config := network.PposMainNetParams
dc := NewDelegateContract(config, credentials)

nodeId := "0x77fffc999d9f9403b65009f1eb27bae65774e2d8ea36f7b20a89f82642a5067557430e6edfe5320bb81c3666a19cf4a5172d6533117d7ebcd0f2c82055499050"
stakingBlockNumber := big.NewInt(129518)
amount := new(big.Int)
amount.SetString("100000000000000000000", 10)
list, err := dc.UnDelegate(nodeId, stakingBlockNumber, amount)
if err != nil {
    t.Errorf("DelegateContract.UnDelegate failed: %s", err)
}

result, err := json.Marshal(list)
if err != nil {
    t.Errorf("Marshal of list failed: %s", err)
}

fmt.Println(string(result))
```

## node_contract合约
Node合约用于查询当前网络中各节点的状态. 

* GetVerifierList() 查询当前周期Verifier列表
* GetValidatorList() 查询当前周期Validator列表
* GetCandidateList() 查询当前周期Candidate列表

## proposal_contract合约
proposal合约用于管理提案.

* `func GetProposal(proposalId string) (resp.Proposal, error)`   
用于查询提案的详情.  
参数说明:
    * proposalId: 提案的id

* `func GetTallyResult(proposalId string) (resp.TallyResult, error)`  
用于查询提案的投票结果.  
参数说明:
    * proposalId: 提案的id  
    
* `func GetProposalList() ([]resp.Proposal, error)`  
用于查询提案列表.  

* `func GetActiveVersion() (uint64, error)`  
用于查询已经生效的版本.  

* `func GetGovernParamValue(module string, name string) (string, error)`  
用于查询模块的参数.
参数说明:
    * module 模块名
    * name   参数名
    
* `func GetAccuVerifiersCount(proposalId string, blockHash string) (*big.Int, error)`  
查询指定时间, 提案的投票人数.
参数说明:
    * proposalId: 提案id
    * blockHash: 截至的块hash
    
* `func GetParamList(module string) ([]resp.GovernParam, error)`  
查询模块所有可治理的参数列表.  
参数说明: 
    * module: 模块名

* `func Vote(programVersion typedefs.ProgramVersion, voteOption typedefs.VoteOption, proposalID string, nodeId string) (typedefs.TransactionHash, error)`  
用于给提案投票.
参数说明:
    * programVersion: 版本号
    * voteOption: 投票内容
    * proposalId: 提案id
    * nodeId: 节点id
    
* `func DeclareVersion(programVersion typedefs.ProgramVersion, verifier string) (typedefs.TransactionHash, error)`  
声明版本号, 只能由验证人/候选人发起.

* `func SubmitProposal(proposal *resp.Proposal) (typedefs.TransactionHash, error)`  
提交提案.
参数说明: 
    * proposal: 提案的详细内容
    
## restricting_plan_contract合约
restricting plan合约用于管理锁仓计划.

* `func CreateRestrictingPlan(account common2.Address, restrictingPlanList []resp.RestrictingPlan) (typedefs.TransactionHash, error)`  
用于创建锁仓计划.
参数说明:
    * account: 受益人账号
    * restrictingPlanList: 锁仓计划
    
* `func GetRestrictingInfo(account common2.Address) (resp.RestrictingItem, error)`  
查询锁仓计划.
参数说明:
    * account: 受益人账号
    
## reward_contract合约
reward合约用于管理委托收益.

* `func WithdrawDelegateReward() (typedefs.TransactionHash, error)`  
提取委托收益. 受益人为交易发起人.

* `func GetDelegateReward(address common2.Address, nodeIdList []string) ([]resp.Reward, error)`  
查询委托人的收益.
参数说明:
    * address: 委托人
    * nodeIdList: 委托的节点id列表.

## slash_contract合约
slash合约用于管理节点作弊.

* `func ReportDoubleSign(duplicateSignType typedefs.DuplicateSignType, data string) (typedefs.TransactionHash, error)`  
举报双签.
参数说明:
    * duplicateSignType: 双签类型
    * data: 证据文件, 一般为JSON格式.
    
* `func CheckDoubleSign(doubleSignType typedefs.DuplicateSignType, nodeId string, blockNumber *big.Int) (string, error) `  
查询节点是否有指定违规类型的双签行为.

## staking_contract合约
staking合约用于管理节点质押操作.

* `func GetStakingInfo(nodeId string) (resp.Node, error)`  
查询节点的质押信息.

* `func GetPackageReward() (*big.Int, error)`  
查询当前结算周期的出块奖励.

* `func GetStakingReward() (*big.Int, error)`  
查询当前结算周期的质押奖励.

* `func GetAvgPackTime() (*big.Int, error)`  
查询区块打包的平均时间.

* `func Staking(stakingParam req.StakingParam) (typedefs.TransactionHash, error)`  
发起质押

* `func UnStaking(nodeId string) (typedefs.TransactionHash, error)`  
取消质押

* `func UpdateStakingInfo(updateStakingParam req.UpdateStakingParam) (typedefs.TransactionHash, error)`  
更新质押信息.

* `func AddStaking(nodeId string, stakingAmountType typedefs.StakingAmountType, amount *big.Int) (typedefs.TransactionHash, error) `  
增持质押.
