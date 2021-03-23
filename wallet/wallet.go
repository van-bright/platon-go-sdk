package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"platon-go-sdk/bech32"
	"platon-go-sdk/common"
)

type AlayaWallet struct {
	*hdwallet.Wallet
	Mnemonics string
}

func importPrivateKey(seed []byte) (*AlayaWallet, error) {
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return &AlayaWallet{}, err
	}

	return &AlayaWallet{Wallet: wallet}, nil
}

// 打开指定keystore文件的钱包
func Open(ksPath string, passphrase string) (*AlayaWallet, error) {
	return nil, nil
}

// 保存keystore文件到指定地址
func Save(ksPath string, passphrase string) {

}

// 创建一个钱包, 并生成一组助记词
func Create() (*AlayaWallet, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return &AlayaWallet{}, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return &AlayaWallet{}, err
	}

	return ImportMnemonic(mnemonic)
}

// 通过助记词导入钱包
func ImportMnemonic(mnemonic string) (*AlayaWallet, error) {
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := importPrivateKey(seed)
	if err != nil {
		return &AlayaWallet{}, err
	}

	wallet.Mnemonics = mnemonic
	return wallet, nil
}

//
func ImportPrivateKey(seed []byte) (*AlayaWallet, error) {
	return importPrivateKey(seed)
}

func (w *AlayaWallet) CreateAccount(pathIndex int32) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", pathIndex)) //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := w.Derive(path, true)
	if err != nil {
		return accounts.Account{}, err
	}

	return account, nil
}

func (w *AlayaWallet) RawAccount(account accounts.Account) (string, error) {
	address := account.Address.Hex()
	return address, nil
}

func (w *AlayaWallet) MainNetAccount(account accounts.Account) (string, error) {
	ethAccount, err := w.RawAccount(account)
	if err != nil {
		return "", err
	}

	mainAccount, err := bech32.EncodeAddress(common.MainNetAddressPrefix, ethAccount)
	if err != nil {
		return "", err
	}

	return mainAccount, nil
}

func (w *AlayaWallet) TestNetAccount(account accounts.Account) (string, error) {
	ethAccount, err := w.RawAccount(account)
	if err != nil {
		return "", err
	}

	testAccount, err := bech32.EncodeAddress(common.TestNetAddressPrefix, ethAccount)
	if err != nil {
		return "", err
	}

	return testAccount, nil
}

type ExportedAccount struct {
	PrivateKey      string `json:"private-key"`
	EthAccounts     string `json:"eth-address"`
	MainNetAccounts string `json:"mainnet-addresses"`
	TestNetAccounts string `json:"testnet-addresses"`
	Path            string `json:"path"`
}

type ExportedWallet struct {
	Mnemonics string            `json:"mnemonics"`
	Accounts  []ExportedAccount `json:"accounts"`
}

func (w *AlayaWallet) Export() *ExportedWallet {
	tAccounts := &ExportedWallet{}
	tAccounts.Mnemonics = w.Mnemonics
	//tAccounts.seed =

	accounts := w.Accounts()
	expAccounts := make([]ExportedAccount, len(accounts))

	for i, acc := range accounts {
		expAccounts[i].PrivateKey, _ = w.PrivateKeyHex(acc)
		expAccounts[i].EthAccounts, _ = w.RawAccount(acc)
		expAccounts[i].MainNetAccounts, _ = w.MainNetAccount(acc)
		expAccounts[i].TestNetAccounts, _ = w.TestNetAccount(acc)
		expAccounts[i].Path, _ = w.Path(acc)
	}

	tAccounts.Accounts = expAccounts
	return tAccounts
}
