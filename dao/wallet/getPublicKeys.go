package wallet

import (
	"Gcoin/utils"
	"Gcoin/enum"
)

func GetPublickKeys() ([]byte, error) {

	result, err := utils.Get(enum.URL + enum.Wallet_get_public_keys_v1)
	return result, err
}
