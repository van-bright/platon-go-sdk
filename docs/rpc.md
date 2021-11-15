# web3go rpc 接口使用说明

**web3go rpc** 接口定义了一系列与**PlatON** 网络进行交互的RPC调用, 方便利用go语言开发基于PlatON链的应用.

在使用以下接口之前, 我们需要先创建一个`web3go` 对象:
```go
const platonEndpoint = "http://localhost:6789"
web3g, err := web3go.New(platonEndpoint)
```
所有的接口, 都通过`web3g`对象进行操作.

### Accounts
查询当前client的账号
```go
accounts, err := web3g.Accounts()
fmt.Println("accounts: ", accounts)
```
### NetworkID() (*big.Int, error)
查询当前网络的id.
### NetListening() (bool, error)
查询节点是否已经接入网络.
### NetPeerCount() (uint64, error)
查询节点数量
### ClientVersion() (string, error)
查询客户端版本号
### Sha3(str string) (common.Hash, error)
计算一个utf-8格式的string的hash值
### ProtocolVersion() (uint64, error)
查询platon的版本号
### Syncing() (*platon.SyncProgress, error)
查询当前的是否在同步块信息.
如果没在进行同步, 返回nil.
如果正在进行同步, 返回一个同步块信息的对象.
### GasPrice() (*big.Int, error)
查询当前的Gas Price
### EstimateGasLimit(msg platon.CallMsg) (uint64, error)
预估调用msg的Gas费.
### BlockNumber() (uint64, error)
返回当前块高.
### BalanceAt(bech32Account string, pos interface{}) (*big.Int, error)
查询账号`bech32Account`在`pos`位置的余额, `pos`可以是`latest`, `pending`, 或者`earliest`
### StorageAt(bech32Account string, key common.Hash, pos interface{}) ([]byte, error)
查询账号`bech32Account`在`pos`位置的一个`key`对应的存储数据, `pos`可以是`latest`, `pending`, 或者`earliest`
### CodeAt(bech32Account string, pos interface{}) ([]byte, error)
查询账号`bech32Account`在`pos`位置存储的合约代码, `pos`可以是`latest`, `pending`, 或者`earliest`.
如果是EOA账号, 则无代码返回.
### HeaderByNumber(number *big.Int) (*types.Header, error)
通过块高查询header块的信息. 如果查询`latest`的头信息, 使用参数`nil`
### HeaderByHash(hash common.Hash) (*types.Header, error)
通过hash查询header的信息.
### NonceAt(bech32Account string, pos interface{}) (uint64, error)
查询账号`bech32Account`在`pos`位置的`nonce`值, `pos`可以是`latest`, `pending`, 或者`earliest`
### TransactionCountByHash(blockHash common.Hash) (uint, error)
查询hash值为`blockHash`的块内交易的数量.
### TransactionCountByNumber(option interface{}) (uint, error)
查询块高为`option`的块内交易的数量, `option`可是一个具体的int值, 也可以是`latest`, `pending`, 或者`earliest`.
### BlockByHash(hash common.Hash) (*types.Block, error)
查询hash为`hash`的块内的数据.
### BlockByNumber(option interface{}) (*types.Block, error)
查询块高为`option`的块的数据, `option`可是一个具体的int值, 也可以是`latest`, `pending`, 或者`earliest`.
### TransactionByHash(hash common.Hash) (tx *types.Transaction, isPending bool, err error)
查询hash值为`hash`的交易的信息.
### TransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (*types.Transaction, error)
查询hash值为`blockHash`的块内第`index`个交易的信息.
### TransactionByBlockNumberAndIndex(option interface{}, index uint) (*types.Transaction, error)
查询块高为`option`的块内第`index`个交易的信息. `option`可是一个具体的int值, 也可以是`latest`, `pending`, 或者`earliest`.
### TransactionReceipt(txHash common.Hash) (*types.Receipt, error)
查询交易的receipt.
### CallContract(msg platon.CallMsg, option interface{}) ([]byte, error)
在指定`option`的块上调用合约交易.
### SendRawTransaction(tx *types.Transaction) error
广播已经签名的tx到网络中.
### SendTransaction(tx *types.Transaction) error
广播未签名的tx.
### AdminGetProgramVersion() (*params.ProgramVersion, error)
查询程序的版本号
### AdminDataDir() (string, error)
查询data的存储路径.
### AdminPeers() ([]string, error)
查询节点连接的节点信息.
### AdminAddPeer(peer string) (bool, error)
添加节点.
### GetSchnorrNIZKProve() (string, error)
查询SchnorrNIZKProve.
### Evidences() (string, error)
查询多签证明.
### Sign(req *platon.SignReq) (string, error)
使用已经解锁的账号, 对交易hash进行签名.
### AdminNodeInfo
查询所连接的Node信息
```go
info, err := web3g.AdminNodeInfo()
fmt.Println("admin node info: ", info, err)
```
如果没有错误发生, 将产生的如下的输出:
```json
{
    "id":"3537fb8faebef1838ddf3f9cc828be000cda442862b6ca1dc2804e6e42e41345ae0d8377b526732042dc04f7b53ef1304540c0bdfb718fd2294b5172dd019653",
    "name":"PlatONnetwork/platon-node/v0.15.0-unstable-9867ee68/linux-amd64/go1.13.4",
    "blsPubKey":"bbcf9c796585a58d44e3e1f0c899e8dcae3704d9dae78cec9a8a5ae3c4551517449580dc6198205c74186b7223622d060101210926ff603c29686810c8dc63ad595e49e2b12cb116471c3691f943a06b91a3547e034d45c1bf373b19e648e395",
    "enode":"enode://3537fb8faebef1838ddf3f9cc828be000cda442862b6ca1dc2804e6e42e41345ae0d8377b526732042dc04f7b53ef1304540c0bdfb718fd2294b5172dd019653@[::]:16789",
    "ip":"::",
    "ports":{
        "discovery":16789,
        "listener":16789
    },
    "listenAddr":"[::]:16789",
    "protocols":{
        "cbft":{
            "config":{
                "sys":{
                    "period":20000,
                    "amount":10,
                    "initialNodes":[
                        {
                            "node":"enode://ba0f7995ee0cf98e18e82daf9560f2007eb46a510be0b3bcc6ec7ab6f4e7f611ce783036916aa0a0af95ddc1ad691df1fe84cfab926f358aec1a786e79d9eb7b@161.117.253.224:16789",
                            "blsPubKey":"ccb421751a3ae082df3c134cbcc0f9e36c0ef2ddb641b1cfa4f647abe522bc5db1c432eb94368c7624a340a9021980084c0247240830f7da28d5a290c6b1e018d2f7f1b4c4a72b0dca4503f7fc6d43b63bee3534d2fb592d20f97be3e3a6048e"
                        },
                        {
                            "node":"enode://b273bd13e4d82793ff5a874ef514934db333bcb3b12330ecf91e831ef6e536db61b4179bea8886095e931c1c0e749367eaf84266e083fa55c48c9fe1d50e26df@47.254.174.63:16789",
                            "blsPubKey":"c7e05f20d53715719bd22b5e4cec453157ae5e81b3edf737e33c37b0839c630a77129e12550ae1571d1bec399b12a619ed987307435304d4b169e460929516be8d298a7493afc7449359ce76731938bb57b550578843c3ae6ba41e3d63ffd88a"
                        },
                        {
                            "node":"enode://45cc53b5279503b6476bb2e5bcc8f0d3f1e7f6526ade9f27f9fea528f9bda9f5c6b8a5379235363b158188add776d14a5baebe55e8e8d57cdaefcea05307aa95@54.255.41.125:16789",
                            "blsPubKey":"2d317f94fa7583de8cd0d72069f3b104bcd6d8c1bc24265d5e66218d31eb4db19a2ddef75b74090fc34b951bfe4dac1535f91c5819569997131d1de688c1d3a2b7ae4ec3bfccf5a7d6af0412e01e1c6937311232ecc0a1c5bce951e84d425900"
                        },
                        {
                            "node":"enode://695055cfccec7536b3f1e12208af7742a8aa30f0d7ed916e557a7cd50f3bedc9044e3916ce6d4d43aff4ae09e1df0e06842f7eb9dfd226fe49f4e0a6c1ad7a0a@15.207.25.137:16789",
                            "blsPubKey":"8f8d77b2e4d2e2e36000f56a26562fa5a3780b3263ca0e1f0533b9a14a3530d55da4dd84fe0a8713639f54cf0e7564141f1ce2779feee022e25a29bc20f2eed9c2367bba535e92f1057a1601b368a3cd2366a929e2c5f74d5e5bf45096063d18"
                        },
                        {
                            "node":"enode://bc93ef4c138c4e010a57ba7c4565faaf07f5cdbe4bff58f92f7aa3e9a10e830dcb4a6f63fe04e61872c5ea1f98a302cfb6cb0004538db0cf3c7c8d697dc5bfee@35.183.248.205:16789",
                            "blsPubKey":"dd2be8a4a50cda5c8cbe6cc8d38aff0e18991749fff68018f86c4c46fe7d48f301a216648136fae5d27c6db128761910488f2ae886390384fffed9068fc18132c60e2874d210a1ad890eca1f99051e3aa9d7ef748df57ed9a2b93bb370e9fd80"
                        },
                        {
                            "node":"enode://a860a56adaa9d3bd66494973d44b54bdfdadaeabdf0377b40e9dd99235a1d764dfc6791a19e744bef1d5e1962cc8de8afbf6202b922e2a411783a3ed9d622c2b@54.73.237.127:16789",
                            "blsPubKey":"ce809302eb36897b937874ad84bc1e506c543252510f1ee20ab6578457b321667f8faf18fb1143b7e1af418e56033603062dc809705e220a39032e86f3b49ada23224e28798d4b50b82b80f03d609c555acce12ddb7b3f7259d23c769b708506"
                        },
                        {
                            "node":"enode://25af23c768bb57bbb5b72e349cf23bbb371e7359a3c0436cc3c22f28edbfa3429a511cd1f05783f4b385f84cd0649884fd36b8d3018b0a108ed7e7b189f41566@35.157.79.228:16789",
                            "blsPubKey":"05fee124fdb890a4c795142228c8812308138640fb3a173324af4e8e2df13a6b5a93af19edeb730ec884104e5dd4f70e60e578470a730fe5b3a99dd52590304603df7bc189bdac4d556e736d5659ad2ffd20a14cd5fbf567952fb4b9a678118d"
                        }
                    ],
                    "validatorMode":"ppos"
                },
                "option":{
                    "nodeID":"3537fb8faebef1838ddf3f9cc828be000cda442862b6ca1dc2804e6e42e41345ae0d8377b526732042dc04f7b53ef1304540c0bdfb718fd2294b5172dd019653",
                    "walMode":true,
                    "peerMsgQueueSize":1024,
                    "evidenceDir":"evidence",
                    "maxPingLatency":5000,
                    "maxQueuesLimit":4096,
                    "blacklistDeadline":60,
                    "period":20000,
                    "amount":10
                }
            }
        },
        "platon":{
            "network":1,
            "genesis":"0xfb787fede6752e1a5ad85d2c6fc140454759be5e69d86d2425ceac22c23bd419",
            "config":{
                "chainId":201018,
                "emptyBlock":"on",
                "eip155Block":1,
                "cbft":{
                    "period":20000,
                    "amount":10,
                    "initialNodes":[
                        {
                            "node":"enode://ba0f7995ee0cf98e18e82daf9560f2007eb46a510be0b3bcc6ec7ab6f4e7f611ce783036916aa0a0af95ddc1ad691df1fe84cfab926f358aec1a786e79d9eb7b@161.117.253.224:16789",
                            "blsPubKey":"ccb421751a3ae082df3c134cbcc0f9e36c0ef2ddb641b1cfa4f647abe522bc5db1c432eb94368c7624a340a9021980084c0247240830f7da28d5a290c6b1e018d2f7f1b4c4a72b0dca4503f7fc6d43b63bee3534d2fb592d20f97be3e3a6048e"
                        },
                        {
                            "node":"enode://b273bd13e4d82793ff5a874ef514934db333bcb3b12330ecf91e831ef6e536db61b4179bea8886095e931c1c0e749367eaf84266e083fa55c48c9fe1d50e26df@47.254.174.63:16789",
                            "blsPubKey":"c7e05f20d53715719bd22b5e4cec453157ae5e81b3edf737e33c37b0839c630a77129e12550ae1571d1bec399b12a619ed987307435304d4b169e460929516be8d298a7493afc7449359ce76731938bb57b550578843c3ae6ba41e3d63ffd88a"
                        },
                        {
                            "node":"enode://45cc53b5279503b6476bb2e5bcc8f0d3f1e7f6526ade9f27f9fea528f9bda9f5c6b8a5379235363b158188add776d14a5baebe55e8e8d57cdaefcea05307aa95@54.255.41.125:16789",
                            "blsPubKey":"2d317f94fa7583de8cd0d72069f3b104bcd6d8c1bc24265d5e66218d31eb4db19a2ddef75b74090fc34b951bfe4dac1535f91c5819569997131d1de688c1d3a2b7ae4ec3bfccf5a7d6af0412e01e1c6937311232ecc0a1c5bce951e84d425900"
                        },
                        {
                            "node":"enode://695055cfccec7536b3f1e12208af7742a8aa30f0d7ed916e557a7cd50f3bedc9044e3916ce6d4d43aff4ae09e1df0e06842f7eb9dfd226fe49f4e0a6c1ad7a0a@15.207.25.137:16789",
                            "blsPubKey":"8f8d77b2e4d2e2e36000f56a26562fa5a3780b3263ca0e1f0533b9a14a3530d55da4dd84fe0a8713639f54cf0e7564141f1ce2779feee022e25a29bc20f2eed9c2367bba535e92f1057a1601b368a3cd2366a929e2c5f74d5e5bf45096063d18"
                        },
                        {
                            "node":"enode://bc93ef4c138c4e010a57ba7c4565faaf07f5cdbe4bff58f92f7aa3e9a10e830dcb4a6f63fe04e61872c5ea1f98a302cfb6cb0004538db0cf3c7c8d697dc5bfee@35.183.248.205:16789",
                            "blsPubKey":"dd2be8a4a50cda5c8cbe6cc8d38aff0e18991749fff68018f86c4c46fe7d48f301a216648136fae5d27c6db128761910488f2ae886390384fffed9068fc18132c60e2874d210a1ad890eca1f99051e3aa9d7ef748df57ed9a2b93bb370e9fd80"
                        },
                        {
                            "node":"enode://a860a56adaa9d3bd66494973d44b54bdfdadaeabdf0377b40e9dd99235a1d764dfc6791a19e744bef1d5e1962cc8de8afbf6202b922e2a411783a3ed9d622c2b@54.73.237.127:16789",
                            "blsPubKey":"ce809302eb36897b937874ad84bc1e506c543252510f1ee20ab6578457b321667f8faf18fb1143b7e1af418e56033603062dc809705e220a39032e86f3b49ada23224e28798d4b50b82b80f03d609c555acce12ddb7b3f7259d23c769b708506"
                        },
                        {
                            "node":"enode://25af23c768bb57bbb5b72e349cf23bbb371e7359a3c0436cc3c22f28edbfa3429a511cd1f05783f4b385f84cd0649884fd36b8d3018b0a108ed7e7b189f41566@35.157.79.228:16789",
                            "blsPubKey":"05fee124fdb890a4c795142228c8812308138640fb3a173324af4e8e2df13a6b5a93af19edeb730ec884104e5dd4f70e60e578470a730fe5b3a99dd52590304603df7bc189bdac4d556e736d5659ad2ffd20a14cd5fbf567952fb4b9a678118d"
                        }
                    ],
                    "validatorMode":"ppos"
                },
                "genesisVersion":3330
            },
            "head":"0xfb787fede6752e1a5ad85d2c6fc140454759be5e69d86d2425ceac22c23bd419"
        }
    }
}
```

