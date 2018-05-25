package ethornge

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
)

func NewLedgerProvider() (err error) {
	var hub *usbwallet.Hub
	if hub, err = usbwallet.NewLedgerHub(); err != nil {
		return
	}

	var sink = make(chan accounts.WalletEvent)
	var subs = hub.Subscribe(sink)
	defer subs.Unsubscribe()

	for {
		select {
		case evt := <-sink:
			switch evt.Kind {
			case accounts.WalletOpened:
				break
			case accounts.WalletArrived:
				break
			case accounts.WalletDropped:
				break
			}
		}
	}

	return
}
