package web3go

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"platon-go-sdk/accounts"
	"platon-go-sdk/accounts/keystore"
	"platon-go-sdk/common"
	"platon-go-sdk/common/hexutil"
	"platon-go-sdk/core/types"
	"platon-go-sdk/ethclient"
	"platon-go-sdk/hdwallet"
	"strings"
)

type AlayaWallet struct {
	hd         *hdwallet.Wallet
	ks         *keystore.KeyStore
	networkCfg *NetworkCfg
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

type NetworkCfg struct {
	Url     string
	ChainId *big.Int
}

var (
	DefaultTestNetCfg = NetworkCfg{"http://47.241.91.2:6789", big.NewInt(201030)}
	DefaultMainNetCfg = NetworkCfg{"https://openapi.alaya.network/rpc", big.NewInt(201018)}
)

func (w *AlayaWallet) isHdAccount(account accounts.Account) bool {
	isKeystore := strings.HasPrefix(account.URL.String(), "keystore")
	return !isKeystore && w.hd.Contains(account)
}

func (w *AlayaWallet) isKsAccount(account accounts.Account) bool {
	if w.ks == nil {
		return false
	}

	if _, err := w.ks.Find(account); err != nil {
		return false
	}

	return true
}

func (w *AlayaWallet) ToString(account accounts.Account) string {
	ma, _ := account.ToMainNetAddress()
	ta, _ := account.ToTestNetAddress()
	pk, _ := w.hd.PrivateKeyHex(account)

	ai := AccountInfo{pk, ma, ta}
	jsonStr, _ := json.Marshal(ai)

	return string(jsonStr)
}

func (w *AlayaWallet) ExportHdWallet() string {
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

func (w *AlayaWallet) Accounts() []accounts.Account {
	accounts := w.hd.Accounts()
	if w.ks != nil {
		accounts = append(accounts, w.ks.Accounts()...)
	}

	return accounts
}

func (w *AlayaWallet) AccountByBech32(bech32addr string) (accounts.Account, error) {
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
func (w *AlayaWallet) MainNetAddress(account accounts.Account) (string, error) {
	return account.ToMainNetAddress()
}

// TestNetAddress convert to test net address with prefix 'atx`
func (w *AlayaWallet) TestNetAddress(account accounts.Account) (string, error) {
	return account.ToTestNetAddress()
}

// NewAccount create a new account by hd wallet.
// `index` means the path index which from 0
func (w *AlayaWallet) NewAccount(index uint64) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
	account, err := w.hd.Derive(path, true)
	if err != nil {
		return accounts.Account{}, err
	}
	return account, nil
}

// ImportPrivateKey import a private key to keystore file.
func (w *AlayaWallet) ImportPrivateKey(key *ecdsa.PrivateKey, ksPath string, passphrase string) (accounts.Account, error) {
	if w.ks == nil {
		dir := filepath.Dir(ksPath)
		w.ks = keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	}

	validAccount, err := w.ks.ImportECDSA(key, passphrase)
	if err != nil {
		return validAccount, err
	}

	return validAccount, nil
}

// SetNetworkCfg set network config for wallet
func (w *AlayaWallet) SetNetworkCfg(cfg *NetworkCfg) {
	w.networkCfg = cfg
}

// ExportMnemonic export the mnemonic of hd wallet, if there is no mnemonic existed, throw an error
func (w *AlayaWallet) ExportMnemonic() (string, error) {
	if len(w.hd.Mnemonic()) == 0 {
		return "", fmt.Errorf("no mnemonic exist")
	}

	return w.hd.Mnemonic(), nil
}

// ExportPrivateKey export the private key of an account
func (w *AlayaWallet) ExportPrivateKey(account accounts.Account, passphrass string) (*ecdsa.PrivateKey, error) {
	if w.isHdAccount(account) {
		return w.hd.PrivateKey(account)
	}

	if w.isKsAccount(account) {
		_, key, err := w.ks.GetDecryptedKey(account, passphrass)
		if err != nil {
			return nil, err
		}

		return key.PrivateKey, nil
	}

	return nil, fmt.Errorf("account not found")
}

// ImportFromKeyStore add a new account from a keystore file.
func (w *AlayaWallet) ImportFromKeyStore(path string, passphrase string, newpassphrase string) (accounts.Account, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return accounts.Account{}, err
	}

	dir := filepath.Dir(path)
	w.ks = keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := w.ks.Import(jsonBytes, passphrase, newpassphrase)
	if err != nil {
		return accounts.Account{}, err
	}

	if newpassphrase != passphrase {
		os.Remove(path)
	}

	return account, nil
}

// ExportToKeyStore save a selected account to keystore file.
func (w *AlayaWallet) ExportToKeyStore(account accounts.Account, path string, passphrase string) error {
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
func (w *AlayaWallet) BalanceOf(owner common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(w.networkCfg.Url)
	if err != nil {
		return nil, err
	}

	common.SetAddressPrefix(common.MainNetAddressPrefix)
	addr := owner.Bech32()

	balance, err := client.BalanceAt(context.Background(), addr, "latest")
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// Transfer send value from `from` account to `to` account
// if success, return the hash of transaction.
func (w *AlayaWallet) Transfer(from common.Address, to common.Address, value *big.Int) (string, error) {
	client, err := ethclient.Dial(w.networkCfg.Url)
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	common.SetAddressPrefix(common.MainNetAddressPrefix)
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

	err = client.SendRawTransaction(ctx, signedTx)

	return signedTx.Hash().Hex(),  err
}

func (w *AlayaWallet) SignTx(tx *types.Transaction, fromAccount accounts.Account) (*types.Transaction, error) {
	var signedTx *types.Transaction
	var err error

	switch {
	case w.isHdAccount(fromAccount):
		signedTx, err = w.hd.SignTx(fromAccount, tx, w.networkCfg.ChainId)
		if err != nil {
			return nil, err
		}
	case w.isKsAccount(fromAccount):
		signedTx, err = w.ks.SignTx(fromAccount, tx, w.networkCfg.ChainId)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown from account")
	}
	return signedTx, nil
}

// Lock lock an account
func (w *AlayaWallet) Lock(account accounts.Account) error {
	if w.isKsAccount(account) {
		return w.ks.Lock(account.Address)
	}

	return nil
}
// LockBech32 to lock an account with bech32 format.
func (w *AlayaWallet) LockBech32(bech32Address string) error {
	account, err := w.AccountByBech32(bech32Address)
	if err != nil {
		return err
	}

	return w.Lock(account)
}
// Unlock to unlock an account with wallet passphrase.
func (w *AlayaWallet) Unlock(account accounts.Account, passphrase string) error {
	if w.isKsAccount(account) {
		return w.ks.Unlock(account, passphrase)
	}
	return nil
}
// UnlockBech32 to unlock a bech32 format account
func (w *AlayaWallet) UnlockBech32(bech32Address string , passphrase string) error {
	account, err := w.AccountByBech32(bech32Address)
	if err != nil {
		return err
	}

	return w.Unlock(account, passphrase)
}
// NewWallet create a new Alaya wallet with mnemonic
func NewWallet() (*AlayaWallet, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return nil, err
	}

	return NewWalletByMnemonics(mnemonic)
}

// NewWalletByMnemonics create a Alaya wallet by importing mnemonics.
func NewWalletByMnemonics(mnemonics string) (*AlayaWallet, error) {
	w, err := hdwallet.NewFromMnemonic(mnemonics)
	if err != nil {
		return nil, err
	}
	// 默认生成0地址
	aw := &AlayaWallet{w, nil, &DefaultMainNetCfg}
	_, err = aw.NewAccount(0)
	if err != nil {
		return nil, err
	}
	return aw, nil
}

// NewWalletBySeed create a Alaya wallet by seed of wallet.
func NewWalletBySeed(seed []byte) (*AlayaWallet, error) {
	w, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}
	// 默认生成0地址
	aw := &AlayaWallet{w, nil, &DefaultMainNetCfg}
	_, err = aw.NewAccount(0)
	if err != nil {
		return nil, err
	}
	return aw, nil
}
