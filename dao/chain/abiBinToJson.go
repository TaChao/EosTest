package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type BinToJson struct {
	Code    string `json:"code"`
	Action  string `json:"action"`
	Binargs string `json:"binargs"`
}

func abiBinToJson(s *BinToJson) ([]byte, error) {
	jsonStr, _ := json.Marshal(s)

	result, err := utils.Post(enum.URL+enum.Abi_bin_to_json_v1, string(jsonStr))
	return result, err
}
