package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

func PushTransaction(para interface{}) ([]byte, error) {

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Push_transaction_v1, string(jsonStr))
	return result, err
}

func PushTransactions(para interface{}) ([]byte, error) {

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Push_transaction_v1, string(jsonStr))
	return result, err
}
