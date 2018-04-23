package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type BlockParams struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

func GetBlock(blockNumOrId string) ([]byte, error) {
	para := BlockParams{
		BlockNumOrId: blockNumOrId,
	}

	jsonStr, _ := json.Marshal(para)

	result, err := utils.Post(enum.URL+enum.Get_block_v1, string(jsonStr))
	return result, err
}
