package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type AccountParams struct {
	AccountName string `json:"account_name"`
}

func GetAccount(accountName string) ([]byte, error) {
	para := AccountParams{
		AccountName: accountName,
	}

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Get_account_v1, string(jsonStr))
	return result, err
}
