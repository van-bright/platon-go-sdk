# Alaya Wallet使用说明

**Alaya Wallet**用来管理Alaya网络的账号.

本钱包同时支持**HDWallet**和**KeyStore Wallet**. 
默认情况下使用HDWallet管理账号, 如果仅通过导入私钥或导入KeyStore文件的方式导入账户,
则使用KeyStore存储账户信息.

## 创建钱包
创建钱包接口, 默认创建HDWallet钱包. 可以通过以下方式创建:
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

* `NewWalletBySeed`方法:  
调用该方法用来通过钱包的seed导入钱包.

  
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
下面的代码示例, 演示了如果在Alaya测试网上列出账户余额和转账的操作.

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
包括Alaya主网地址, 测试网地址, 私钥.

## 导出KeyStore中账户的私钥
使用`ExportPrivateKey(account accounts.Account, passphrass string) (*ecdsa.PrivateKey, error)`方法
导出KeyStore中账户的私钥. `passphrase`是KeyStore文件的密码.

## 导出KeyStore文件
使用`ExportToKeyStore(account accounts.Account, path string, passphrase string) error`将账号`account`导出到`path`指定的路径下, 
并设置KeyStore文件的密码为`passphrase`.

## 导入私钥
仅仅只有私钥的情况下, 只能将账号的私钥导入到KeyStore文件中. 通过`ImportPrivateKey(key *ecdsa.PrivateKey, ksPath string, passphrase string) (accounts.Account, error)`
方法, 将私钥`key`导入到`ksPath`指定的KeyStore文件中, 并设置密码为`passphrase`.

## 导入KeyStore文件
通过`ImportFromKeyStore(path string, passphrase string, newpassphrase string) (accounts.Account, error)`方法
导入账号. 其中`path`是KeyStore文件的路径, `passphrase`是KeyStore文件的当前密码, `newpassphrase`是更新之后的KeyStore文件的密码.

## 签名交易但是不广播
在`Transfer`方法的说明中, 当我们进行转账操作的时候, 默认会通过`rpc.SendRawTransaction`方法, 将交易广播到网络中.
如果我们只需要签名但是无需广播到网络中, 则可以使用`SignTx(tx *types.Transaction, fromAccount accounts.Account) (*types.Transaction, error)`
方法对交易进行签名, 但是不广播到网络中, 同时获取签名之后的`V`, `R`, `S`数据.
