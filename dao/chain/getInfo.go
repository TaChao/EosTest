package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
)

func GetInfo() ([]byte, error) {

	result, err := utils.Get(enum.URL + enum.Get_info_v1)
	return result, err
}
