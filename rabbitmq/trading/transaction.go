package trading

import (
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
)

type NotificationTransaction struct {
	CoinType            sphinxplugin.CoinType `json:"coin_type"`             // when: always
	UUID                string                `json:"uuid"`                  // when: create account; usage: for trading service to locate request when get return
	TransactionIDInsite string                `json:"transaction_id_insite"` // when: Transaction; unique
	AmountFloat64       float64               `json:"amount_float64"`        // when: Transaction
	AddressFrom         string                `json:"address_from"`          // when: Transaction
	AddressTo           string                `json:"address_to"`            // when: Transaction
	TransactionIDChain  string                `json:"transaction_id_chain"`  // when: Transaction empty when created, return when finished
	SignatureUser       string                `json:"signature_user"`        // when: Transaction preserved for 2FA verification, implement this in v2
	SignaturePlatform   string                `json:"signature_platform"`    // when: Transaction preserved for signproxy to verify host, about v3
	CreatetimeUtc       int64                 `json:"createtime_utc"`        // when: Transaction for 2FA
	UpdatetimeUtc       int64                 `json:"updatetime_utc"`        // when: Transaction for return
	IsSuccess           bool                  `json:"is_success"`            // when: Transaction return true when completed
	IsFailed            bool                  `json:"is_failed"`             // when: Transaction return true when error occurred
}
