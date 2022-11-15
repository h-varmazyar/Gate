package buffer

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"sync"
)

type walletBuffer struct {
	wallets    map[string]*chipmunkApi.Wallet
	references map[string]*chipmunkApi.Reference
}

var (
	walletLock *sync.Mutex
	Wallets    *walletBuffer
)

func init() {
	walletLock = new(sync.Mutex)
	Wallets = new(walletBuffer)
	Wallets.wallets = make(map[string]*chipmunkApi.Wallet)
	Wallets.references = make(map[string]*chipmunkApi.Reference)
}

func (buffer *walletBuffer) AddOrUpdate(input *chipmunkApi.Wallet) {
	walletLock.Lock()
	buffer.wallets[input.AssetName] = input
	walletLock.Unlock()
}

func (buffer *walletBuffer) AddOrUpdateList(input []*chipmunkApi.Wallet) {
	walletLock.Lock()
	buffer.wallets = make(map[string]*chipmunkApi.Wallet)
	for _, wallet := range input {
		buffer.wallets[wallet.AssetName] = wallet
	}
	walletLock.Unlock()
}

func (buffer *walletBuffer) FetchWallet(assetName string) *chipmunkApi.Wallet {
	if w, ok := buffer.wallets[assetName]; ok {
		return w
	}
	return nil
}

func (buffer *walletBuffer) FetchReference(refName string) *chipmunkApi.Reference {
	if w, ok := buffer.references[refName]; ok {
		return w
	}
	return nil
}

func (buffer *walletBuffer) FetchAll() *chipmunkApi.Wallets {
	wallets := new(chipmunkApi.Wallets)
	wallets.Elements = make([]*chipmunkApi.Wallet, 0)
	for key, value := range buffer.wallets {
		wallet := &chipmunkApi.Wallet{
			BlockedBalance: value.BlockedBalance,
			ActiveBalance:  value.ActiveBalance,
			TotalBalance:   value.TotalBalance,
			AssetName:      key,
		}
		wallets.Elements = append(wallets.Elements, wallet)
	}
	return wallets
}

func (buffer *walletBuffer) UpdateReferences(newRefs map[string]*chipmunkApi.Reference) {
	walletLock.Lock()
	buffer.references = newRefs
	walletLock.Unlock()
}

func (buffer *walletBuffer) Flush() {
	Wallets.wallets = make(map[string]*chipmunkApi.Wallet)
	Wallets.references = make(map[string]*chipmunkApi.Reference)
}
