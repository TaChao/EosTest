package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type JsonToBin struct {
	Code   string                 `json:"code"`
	Action string                 `json:"action"`
	Args   map[string]interface{} `json:"args"`
}

func AbiJsonToBin(s *JsonToBin) ([]byte, error) {

	jsonStr, _ := json.Marshal(s)

	result, err := utils.Post(enum.URL+enum.Abi_json_to_bin_v1, string(jsonStr))
	return result, err
}
