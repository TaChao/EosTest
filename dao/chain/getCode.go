package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type CodeParams struct {
	AccountName string `json:"account_name"`
}

func GetCode(accountName string) ([]byte, error) {
	para := CodeParams{
		AccountName: accountName,
	}

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Get_code_v1, string(jsonStr))
	return result, err
}