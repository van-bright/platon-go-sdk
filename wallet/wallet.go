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

func createWalletBySeed(seed []byte) (*AlayaWallet, error) {
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return &AlayaWallet{}, err
	}

	return &AlayaWallet{Wallet: wallet}, nil
}

// Create an Alaya account, and store it to KeyStore file
func CreateAlayaWalletBySeed(seed []byte) (*AlayaWallet, error) {
	return createWalletBySeed(seed)
}

func CreateAlayaWalletByMnemonic(mnemonic string) (*AlayaWallet, error) {
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := createWalletBySeed(seed)
	if err != nil {
		return &AlayaWallet{}, err
	}

	wallet.Mnemonics = mnemonic
	return wallet, nil
}
// To create an Alaya wallet with mnemonic.
func CreateAlayaWalletWithMnemonic() (*AlayaWallet, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return &AlayaWallet{}, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return &AlayaWallet{}, err
	}

	return CreateAlayaWalletByMnemonic(mnemonic)
}

func (w *AlayaWallet) CreateAccount(pathIndex int32) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", pathIndex)) //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := w.Derive(path, true)
	if err != nil {
		return accounts.Account{}, err
	}

	return account, nil
}

func (w *AlayaWallet) EthAccount(account accounts.Account) (string, error) {
	address := account.Address.Hex()
	return address, nil
}

func (w *AlayaWallet) MainNetAccount(account accounts.Account) (string, error) {
	ethAccount, err := w.EthAccount(account)
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
	ethAccount, err := w.EthAccount(account)
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
	PrivateKey string `json:"private-key"`
	EthAccounts string `json:"eth-address"`
	MainNetAccounts string `json:"mainnet-addresses"`
	TestNetAccounts string `json:"testnet-addresses"`
	Path            string `json:"path"`
}

type AlayaAccounts struct {
	Mnemonics string `json:"mnemonics"`
	Accounts []ExportedAccount `json:"accounts"`
}

func (w *AlayaWallet) Export() *AlayaAccounts {
	tAccounts := &AlayaAccounts{}
	tAccounts.Mnemonics = w.Mnemonics
	//tAccounts.seed =

	accounts := w.Accounts()
	expAccounts := make([]ExportedAccount, len(accounts))

	for i, acc := range accounts {
		expAccounts[i].PrivateKey, _ = w.PrivateKeyHex(acc)
		expAccounts[i].EthAccounts , _ = w.EthAccount(acc)
		expAccounts[i].MainNetAccounts, _ = w.MainNetAccount(acc)
		expAccounts[i].TestNetAccounts, _ = w.TestNetAccount(acc)
		expAccounts[i].Path, _ = w.Path(acc)
	}

	tAccounts.Accounts = expAccounts
	return tAccounts
}
