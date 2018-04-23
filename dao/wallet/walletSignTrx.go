package wallet

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

func WalletSignTrx(para []interface{}) ([]byte, error) {

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Wallet_sign_trx_v1, string(jsonStr))
	return result, err
}
