package chain

import (
	"Gcoin/utils"
	"Gcoin/enum"
	"encoding/json"
)

type TableRowsParams struct {
	Scope      string `json:"scope"`
	Code       string `json:"code"`
	Table      string `json:"table"`
	Json       bool   `json:"json"`
	LowerBound int    `json:"lower_bound"`
	UpperBound int    `json:"upper_bound"`
	Limit      int    `json:"limit"`
}

func Gettablerows(s *TableRowsParams) ([]byte, error) {
	jsonStr, _ := json.Marshal(s)

	result, err := utils.Post(enum.URL+enum.Get_table_rows_v1, string(jsonStr))
	return result, err
}
func GetTableRows(scope string, code string, table string) ([]byte, error) {
	para := TableRowsParams{
		Scope:      scope,
		Code:       code,
		Table:      table,
		Json:       true,
		LowerBound: 0,
		UpperBound: -1,
		Limit:      20,
	}
	return Gettablerows(&para)
}
