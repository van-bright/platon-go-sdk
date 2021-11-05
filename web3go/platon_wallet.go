package web3go

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/oldmanfan/platon-go-sdk/accounts"
	"github.com/oldmanfan/platon-go-sdk/accounts/keystore"
	"github.com/oldmanfan/platon-go-sdk/common"
	"github.com/oldmanfan/platon-go-sdk/common/hexutil"
	"github.com/oldmanfan/platon-go-sdk/core/types"
	"github.com/oldmanfan/platon-go-sdk/ethclient"
	"github.com/oldmanfan/platon-go-sdk/hdwallet"
	"github.com/oldmanfan/platon-go-sdk/network"
	"math/big"
)

type PlatonWallet struct {
	hd         *hdwallet.Wallet
	networkCfg *network.Config
}

type AccountInfo struct {
	PrivateKeyHex  string `json:"private-key"`
	MainNetAddress string `json:"mainnet-address"`
	TestNetAddress string `json:"testnet-address"`
}

type WalletExport struct {
	Mnemonic string   `json:"mnemonics"`
	Seed     string   `json:"seed"`
	Accounts []string `json:"accounts"`
}

func (w *PlatonWallet) ToString(account accounts.Account) string {
	ma, _ := account.ToMainNetAddress()
	ta, _ := account.ToTestNetAddress()
	pk, _ := w.hd.PrivateKeyHex(account)

	ai := AccountInfo{pk, ma, ta}
	jsonStr, _ := json.Marshal(ai)

	return string(jsonStr)
}

func (w *PlatonWallet) ExportHdWallet() string {
	result := WalletExport{}
	result.Mnemonic = w.hd.Mnemonic()
	result.Seed = hexutil.Encode(w.hd.Seed())
	for _, account := range w.hd.Accounts() {
		s := w.ToString(account)
		result.Accounts = append(result.Accounts, s)
	}

	val, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(val)
}

func (w *PlatonWallet) Accounts() []accounts.Account {
	accounts := w.hd.Accounts()
	return accounts
}

func (w *PlatonWallet) AccountByBech32(bech32addr string) (accounts.Account, error) {
	walletAccounts := w.Accounts()
	findAddr := common.MustBech32ToAddress(bech32addr)
	for _, a := range walletAccounts {
		if bytes.Compare(a.Address.Bytes(), findAddr.Bytes()) == 0 {
			return a, nil
		}
	}

	return accounts.Account{}, fmt.Errorf("account not found")
}

// MainNetAddress convert to main net address with prefix 'atp'
func (w *PlatonWallet) MainNetAddress(account accounts.Account) (string, error) {
	return account.ToMainNetAddress()
}

// TestNetAddress convert to test net address with prefix 'atx`
func (w *PlatonWallet) TestNetAddress(account accounts.Account) (string, error) {
	return account.ToTestNetAddress()
}

// NewAccount create a new account by hd wallet.
// `index` means the path index which from 0
func (w *PlatonWallet) NewAccount(index uint64) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := w.hd.Derive(path, true)
	if err != nil {
		return accounts.Account{}, err
	}
	return account, nil
}

// SetNetworkCfg set network config for wallet
func (w *PlatonWallet) SetNetworkCfg(cfg *network.Config) {
	w.networkCfg = cfg
}

// ExportMnemonic export the mnemonic of hd wallet, if there is no mnemonic existed, throw an error
func (w *PlatonWallet) ExportMnemonic() (string, error) {
	if len(w.hd.Mnemonic()) == 0 {
		return "", fmt.Errorf("no mnemonic exist")
	}

	return w.hd.Mnemonic(), nil
}

// ExportPrivateKey export the private key of an account
func (w *PlatonWallet) ExportPrivateKey(account accounts.Account, passphrass string) (*ecdsa.PrivateKey, error) {
	return w.hd.PrivateKey(account)
}

// ExportToKeyStore save a selected account to keystore file.
func (w *PlatonWallet) ExportToKeyStore(account accounts.Account, path string, passphrase string) error {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	privateKey, err := w.hd.PrivateKey(account)
	if err != nil {
		return err
	}

	validAccount, err := ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return err
	}

	err = ks.Update(account, passphrase, passphrase)
	if err != nil {
		return err
	}

	if validAccount.Address != account.Address {
		return fmt.Errorf("imported account is not consisted with %s", account.Address.Hex())
	}
	return nil
}

// BalanceOf query the balance of specific account
func (w *PlatonWallet) BalanceOf(owner common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(w.networkCfg.Url)
	if err != nil {
		return nil, err
	}

	common.SetAddressPrefix(network.MainNetHrp)
	addr := owner.Bech32()

	balance, err := client.BalanceAt(context.Background(), addr, "latest")
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// Transfer send value from `from` account to `to` account
// if success, return the hash of transaction.
func (w *PlatonWallet) Transfer(from common.Address, to common.Address, value *big.Int) (string, error) {
	client, err := ethclient.Dial(w.networkCfg.Url)
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	common.SetAddressPrefix(network.MainNetHrp)
	fromAddr := from.Bech32()

	nonce, err := client.NonceAt(ctx, fromAddr, "pending")
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000)

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)

	fromAccount, _ := w.AccountByBech32(fromAddr)

	signedTx, err := w.SignTx(tx, fromAccount)
	if err != nil {
		return "", err
	}

	_, err = client.SendRawTransaction(ctx, signedTx)

	return signedTx.Hash().Hex(), err
}

func (w *PlatonWallet) SignTx(tx *types.Transaction, fromAccount accounts.Account) (*types.Transaction, error) {
	var signedTx *types.Transaction
	var err error
	signedTx, err = w.hd.SignTx(fromAccount, tx, w.networkCfg.ChainId)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

// NewWallet create a new Alaya wallet with mnemonic
func NewWallet() (*PlatonWallet, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return nil, err
	}

	return NewWalletByMnemonics(mnemonic)
}

// NewWalletByMnemonics create a Alaya wallet by importing mnemonics.
func NewWalletByMnemonics(mnemonics string) (*PlatonWallet, error) {
	w, err := hdwallet.NewFromMnemonic(mnemonics)
	if err != nil {
		return nil, err
	}
	// 默认生成0地址
	aw := &PlatonWallet{w, &network.DefaultMainNetConfig}
	_, err = aw.NewAccount(0)
	if err != nil {
		return nil, err
	}
	return aw, nil
}

// NewWalletBySeed create a Alaya wallet by seed of wallet.
func NewWalletBySeed(seed []byte) (*PlatonWallet, error) {
	w, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}
	// 默认生成0地址
	aw := &PlatonWallet{w, &network.DefaultMainNetConfig}
	_, err = aw.NewAccount(0)
	if err != nil {
		return nil, err
	}
	return aw, nil
}
