# 合约操作

Alaya sdk的`contracts`模块用于在Alaya网络上部署, 操作智能合约.

## 准备工作
要通过sdk部署和调用合约, 需要使用到`solc`, 请按照官方文档安装好solc编译器. 以及按照[Alaya-Go](https://github.com/AlayaNetwork/Alaya-Go) 的说明, 项目编译成功之后, 在`build/bin`目录下, 可以找到`abigen`工具.

为了演示如何使用, 我们准备了一个测试合约`store.sol`(关于如何编写合约, 请参考solidity官方教程):
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}

```

然后, 我们使用以下命令, 编译出我们需要的bin和abi文件:  
`solc --bin --abi store.sol -o Store`  
执行成功之后, 会在`Store`目录下生成Store.abi和Store.bin两个文件.   

接着, 我们使用`abigen`工具, 生成`Store.go`文件:  
`abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go`

至此, 我们需要的文件已经准备完毕.  

本文以下部分描述的操作, 可以在[这里](../examples/contracts-operation.go)找到完整的使用示例.  

## 部署合约
在操作合约之前, 我们需要先获得一个`Contract`对象:
```go
contract := contracts.Contract{
    Url:        PlatonEndpoint,
    PrivateKey: privateKey,
}

opts, client, err := contract.Init()
if err != nil {
    log.Fatal(err)
}
```
在Init操作完成之后, 我们获得了`opts`和`client`对象, 它们的类型分别是`*bind.TransactOpts`和`bind.ContractBackend`, 有了这两个对象, 我们就可以通过以下代码来部署合约了:  
```go
func toDeployContract(opts *bind.TransactOpts, client bind.ContractBackend) string {
	input := "1.0"
	address, tx, _, err := store.DeployStore(opts, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	// wait until tx confirmed
	time.Sleep(time.Duration(10) * time.Second)
	return address.Hex()
}
```

部署成功之后, 我们可以获得交易的hash以及合约部署的地址. 通过这引地址, 我们就可以用来调用合约的方法了.

## 调用合约
合约部署成功之后, 我们可以调用它的方法来完成工作了, 以下的代码示例说明了如果调用`Version`, `SetItems`和`Items`方法:
```go
func toCallContractMethod(hexContractAddr string, opts *bind.TransactOpts, client bind.ContractBackend) {
	addr := common.HexToAddress(hexContractAddr)
	instance, err := store.NewStore(addr, client)
	if err != nil {
		log.Fatal("new instance failed: ", err)
	}
	// to query version
	ver, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("version of contract is ", ver)

	// to set new items
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("hello"))
	copy(value[:], []byte("kitty"))

	tx, err := instance.SetItem(opts, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	time.Sleep(time.Duration(10) * time.Second)

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // kitty
}
```

## 总结
在本文中, 我们演示了怎样部署合约, 以及调用合约中的`view`方法和`非view`方法.
