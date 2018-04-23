package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
	"Gcoin/entity"
)

func GetRequiredKeys(transferDat entity.TransferDat, availableKeys []string) ([]byte, error) {

	para := map[string]interface{}{
		"transaction":    transferDat,
		"available_keys": availableKeys,
	}

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Get_required_keys_v1, string(jsonStr))
	return result, err
}
