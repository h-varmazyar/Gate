package buffer

import (
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"sync"
)

type walletBuffer struct {
	wallets    map[string]*brokerageApi.Wallet
	references map[string]*Reference
}

type Reference struct {
	AssetName string
	Total     float64
	Active    float64
	Blocked   float64
}

var (
	walletLock *sync.Mutex
	Wallets    *walletBuffer
)

func init() {
	walletLock = new(sync.Mutex)
	Wallets = new(walletBuffer)
	Wallets.wallets = make(map[string]*brokerageApi.Wallet)
	Wallets.references = make(map[string]*Reference)
}

func (buffer *walletBuffer) AddOrUpdate(input *brokerageApi.Wallet) {
	walletLock.Lock()
	buffer.wallets[input.AssetName] = input
	walletLock.Unlock()
}

func (buffer *walletBuffer) AddOrUpdateList(input []*brokerageApi.Wallet) {
	walletLock.Lock()
	buffer.wallets = make(map[string]*brokerageApi.Wallet)
	for _, wallet := range input {
		buffer.wallets[wallet.AssetName] = wallet
	}
	walletLock.Unlock()
}

func (buffer *walletBuffer) Fetch(assetName string) *chipmunkApi.Wallet {
	if w, ok := buffer.wallets[assetName]; ok {
		return w
	}
	return nil
}

func (buffer *walletBuffer) FetchAll() *brokerageApi.Wallets {
	wallets := new(brokerageApi.Wallets)
	wallets.Wallets = make([]*brokerageApi.Wallet, 0)
	for key, value := range buffer.wallets {
		wallet := &brokerageApi.Wallet{
			BlockedBalance: value.BlockedBalance,
			ActiveBalance:  value.ActiveBalance,
			TotalBalance:   value.TotalBalance,
			AssetName:      key,
		}
		wallets.Wallets = append(wallets.Wallets, wallet)
	}
	return wallets
}

func (buffer *walletBuffer) UpdateReferences(newRefs map[string]*Reference) {
	walletLock.Lock()
	buffer.references = newRefs
	walletLock.Unlock()
}

func (buffer *walletBuffer) Flush() {
	Wallets.wallets = make(map[string]*brokerageApi.Wallet)
	Wallets.references = make(map[string]*Reference)
}
