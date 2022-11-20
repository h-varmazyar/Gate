package buffer

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"sync"
)

type WalletBuffer struct {
	wallets    map[string]*chipmunkApi.Wallet
	references map[string]*chipmunkApi.Reference
	lock       *sync.Mutex
}

var walletInstance *WalletBuffer

func NewWalletInstance(configs *Configs) *WalletBuffer {
	if walletInstance == nil {
		walletInstance = &WalletBuffer{
			lock:       new(sync.Mutex),
			wallets:    make(map[string]*chipmunkApi.Wallet),
			references: make(map[string]*chipmunkApi.Reference),
		}
	}
	return walletInstance
}

func (buffer *WalletBuffer) AddOrUpdate(input *chipmunkApi.Wallet) {
	buffer.lock.Lock()
	buffer.wallets[input.AssetName] = input
	buffer.lock.Unlock()
}

func (buffer *WalletBuffer) AddOrUpdateList(input []*chipmunkApi.Wallet) {
	buffer.lock.Lock()
	buffer.wallets = make(map[string]*chipmunkApi.Wallet)
	for _, wallet := range input {
		buffer.wallets[wallet.AssetName] = wallet
	}
	buffer.lock.Unlock()
}

func (buffer *WalletBuffer) FetchWallet(assetName string) *chipmunkApi.Wallet {
	if w, ok := buffer.wallets[assetName]; ok {
		return w
	}
	return nil
}

func (buffer *WalletBuffer) FetchReference(refName string) *chipmunkApi.Reference {
	if w, ok := buffer.references[refName]; ok {
		return w
	}
	return nil
}

func (buffer *WalletBuffer) FetchAll() *chipmunkApi.Wallets {
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

func (buffer *WalletBuffer) UpdateReferences(newRefs map[string]*chipmunkApi.Reference) {
	buffer.lock.Lock()
	buffer.references = newRefs
	buffer.lock.Unlock()
}

func (buffer *WalletBuffer) Flush() {
	buffer.lock.Lock()
	buffer.wallets = make(map[string]*chipmunkApi.Wallet)
	buffer.references = make(map[string]*chipmunkApi.Reference)
	buffer.lock.Unlock()
}
