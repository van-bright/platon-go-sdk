# Platon Wallet使用说明

**Platon Wallet**用来管理Platon网络的账号.

本钱包支持**HDWallet**.

## 创建钱包
创建钱包接口, 可以通过以下方式创建:
* `NewWallet`方法:  
调用该方法将创建一个HDWallet钱包, 同时会生成新的一组助记词和`DerivationPath`的index为0的账号.  
```go
w, err := web3go.NewWallet()
if err != nil {
    fmt.Println("new wall error: ", err)
    return
}
account0, _ := w.Accounts()[0]  // go the default 0 account
```
* `NewWalletByMnemonic`方法:  
调用该方法用来通过助记词导入钱包.
```go
mnemonic := "always brick access science decade nasty marriage attack fame topple pen add"
w, err := NewWalletByMnemonic(mnemonic)
```

* `NewWalletBySeed`方法:  
调用该方法用来通过钱包的seed导入钱包.
```go
seed := "0x9dfc7e3f52c4438d04db5488e13672faa37920ec62bacdc333a83974cb07bfdd893bfd46940dedfeb7ef30a142c4d07d552dd6589b40d3a58b941b7e9d6dae7e"
seedBytes, _ := hexutil.Decode(seed)
w, _ := NewWalletBySeed(seedBytes)
```
  
  下面的接口中, 我们默认已经通过上面的方法, 成功创建了一个钱包`w`.
## 创建新的账号
在钱包创建成功之后, 通过方法`NewAccount(index uint64)`方法, 创建一个序号为`index`的账号.
```go
account, err := w.NewAccount(1)
```
上面的示例代码将生成序号为`1`的账户.  

## 列出账号信息
使用`Accounts() ([]accounts.Account, error)`方法, 将列出钱包中所有的账号信息.

## 查询账号余额
使用`BalanceOf(owner common.Address) (*big.Int, error)`方法, 查询指定地址的余额.  

## 转账
使用`Transfer(from common.Address, to common.Address, value *big.Int) (string, error)`方法用来转账.
下面的代码示例, 演示了如果在Platon测试网上列出账户余额和转账的操作.

```go
const mnemonic = "always brick access science decade nasty marriage attack fame topple pen add"
	w, err := web3go.NewWalletByMnemonics(mnemonic)
	if err != nil {
		fmt.Println("import wall error: ", err)
		return
	}

	w.SetNetworkCfg(&web3go.DefaultTestNetCfg)

	accounts := w.Accounts()
	for _, account := range accounts {
		b, _ := w.BalanceOf(account.Address)
		addr, _ := account.ToMainNetAddress()
		fmt.Printf("balance of %s is %s\n", addr, b.String())
	}

	digest, err := w.Transfer(accounts[0].Address, common.MustBech32ToAddress("atp1ydaqepg8s48gqhz29kk0wrf3lqdtj38d8mkcz3"), big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("transfer failed: ", err)
		return
	}

	fmt.Println("tx send: ", digest)
```
## 导出HDWallet的助记词
使用`ExportMnemonic() (string, error)`方法导出钱包的助记词.

## 导出HDWallet的账户详情
使用`ToString(account accounts.Account) string`方法导出一个HDWallet账号的详情,
包括Platon主网地址, 测试网地址, 私钥.

## 导出KeyStore文件
使用`ExportToKeyStore(account accounts.Account, path string, passphrase string) error`将账号`account`导出到`path`指定的路径下, 
并设置KeyStore文件的密码为`passphrase`.

## 导入私钥
仅仅只有私钥的情况下, 只能将账号的私钥导入到KeyStore文件中. 通过`ImportPrivateKey(key *ecdsa.PrivateKey, ksPath string, passphrase string) (accounts.Account, error)`
方法, 将私钥`key`导入到`ksPath`指定的KeyStore文件中, 并设置密码为`passphrase`.

## 签名交易但是不广播
在`Transfer`方法的说明中, 当我们进行转账操作的时候, 默认会通过`rpc.SendRawTransaction`方法, 将交易广播到网络中.
如果我们只需要签名但是无需广播到网络中, 则可以使用`SignTx(tx *types.Transaction, fromAccount accounts.Account) (*types.Transaction, error)`
方法对交易进行签名, 但是不广播到网络中, 同时获取签名之后的`V`, `R`, `S`数据.
```go
w, _ := NewWalletByMnemonics(mnemonic)
	w.NewAccount(1)
	fromAccount := w.Accounts()[0]
	toAccount := w.Accounts()[1]

	nonce := uint64(1)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(5000000000)
	tx := types.NewTransaction(nonce, toAccount.Address, big.NewInt(100000), gasLimit, gasPrice, nil)

	signedTx, _ := w.SignTx(tx, fromAccount)

	s, _ := json.Marshal(signedTx)

	fmt.Println("signed Tx: ", string(s))
```
上面的代码将产生类似的输出:
```json
{
    "nonce":"0x1",
    "gasPrice":"0x12a05f200",
    "gas":"0x5208",
    "to":"atx1u6vtwsz2fqw5ufnm3tm070k43scxhhc8r3nnts",
    "value":"0x186a0",
    "input":"0x",
    "v":"0x62297",
    "r":"0xd6904b0251615d525f6b3c699047977676baf2ba385f3ec8737e1a530c88796d",
    "s":"0x796efcf6b6bede3a49ec6c7969136d50a8fcbb5f43a8d8595e77c6b2b813732b",
    "hash":"0x3b8c1ef129e4b1e65ed527cbae7718c32a0a80c8a432d3c8b5c083e613485139"
}
```
