package service

import (
	"Gcoin/dao/chain"
	"github.com/tidwall/gjson"
	"fmt"
	"Gcoin/dao/wallet"
	"encoding/json"
	"Gcoin/entity"
	"Gcoin/utils"
)

type CurrencyCmdParams struct {
	Action     string
	Code       string
	Permission string
	Context    string
}

func CurrencyCmd(para CurrencyCmdParams) ([]byte, error) {
	paraStr := "cleos push action " +
		para.Code + " " +
		para.Action + " " +
		"'" + para.Context + "' " + //context
		"--permission " + para.Permission + "@active" + //permission
		" -j" //json out put
	print(paraStr)
	out, err := utils.ExecShell(paraStr)
	return out, err
}

func Transfer(from string, to string, act chain.JsonToBin) {
	//TODO open wallet
	//get all available public keys
	publicKeysDat, _ := wallet.GetPublickKeys()
	//
	binDat, _ := chain.AbiJsonToBin(&act)
	bin := gjson.GetBytes(binDat, "binargs").String()
	//get last irreversible block num
	infoDat, _ := chain.GetInfo()

	blockNum := gjson.GetBytes(infoDat, "last_irreversible_block_num").String()
	expiration := gjson.GetBytes(infoDat, "head_block_time").String()

	//get last irreversible block info
	blockDat, _ := chain.GetBlock(blockNum)

	refBlockPrefix := gjson.GetBytes(blockDat, "ref_block_prefix").String()
	//expiration := gjson.GetBytes(blockDat, "timestamp").String()

	//build data
	scope := []string{from, to}
	auth := entity.Authorization{
		Account:    from,
		Permission: "active",
	}
	action := entity.Action{
		Code:          act.Code,
		Type:          "transfer",
		Recipients:    scope,
		Authorization: []entity.Authorization{auth},
		Data:          bin,
	}
	trans := entity.TransferDat{
		RefBlockNum:    blockNum,
		RefBlockPrefix: refBlockPrefix,
		Expiration:     expiration,
		Scope:          scope,
		Actions:        []entity.Action{action},
		Signatures:     []string{},
		Authorizations: []string{},
	}

	var publicKeys []string
	json.Unmarshal(publicKeysDat, &publicKeys)

	//requiredKeysDat, _ := chain.GetRequiredKeys(trans, publicKeys)
	//requiredKeys := gjson.GetBytes(requiredKeysDat, "required_keys").Array()
	//fmt.Println(publicKeys)
	sigPara := []interface{}{
		trans, []string{"EOS7KPzd3KwhR4jKUFyMXUhyvxUP6NV9Sf8f9v5CC6cs9PQ6AzgWb"}, "",
	}
	sigDat, _ := wallet.WalletSignTrx(sigPara)
	signatures := gjson.GetBytes(sigDat, "signatures").Array()
	for _, s := range signatures {
		trans.Signatures = append(trans.Signatures, s.String())
	}
	//trans.Signatures=signatures

	responce, _ := chain.PushTransaction(trans)
	json, _ := json.Marshal(trans)
	fmt.Println(string(json))
	fmt.Print(string(responce))
}
func TransferV3(from string, to string, act chain.JsonToBin) {
	//TODO open wallet
	//get all available public keys
	publicKeysDat, _ := wallet.GetPublickKeys()
	//
	binDat, _ := chain.AbiJsonToBin(&act)
	bin := gjson.GetBytes(binDat, "binargs").String()
	//get last irreversible block num
	infoDat, _ := chain.GetInfo()

	blockNum := gjson.GetBytes(infoDat, "last_irreversible_block_num").String()
	expiration := gjson.GetBytes(infoDat, "head_block_time").String()

	//get last irreversible block info
	blockDat, _ := chain.GetBlock(blockNum)

	refBlockPrefix := gjson.GetBytes(blockDat, "ref_block_prefix").String()
	//expiration := gjson.GetBytes(blockDat, "timestamp").String()

	//build data
	//scope := []string{from, to}
	auth := entity.AuthorizationV3{
		Actor:      from,
		Permission: "active",
	}
	action := entity.ActionV3{
		Account:       act.Code,
		Name:          act.Action,
		Authorization: []entity.AuthorizationV3{auth},
		Data:          bin,
	}
	trans := entity.TransferV3{
		RefBlockNum:        blockNum,
		RefBlockPrefix:     refBlockPrefix,
		Expiration:         expiration,
		ContextFreeActions: []string{},
		Actions:            []entity.ActionV3{action},
	}

	var publicKeys []string
	json.Unmarshal(publicKeysDat, &publicKeys)

	//requiredKeysDat, _ := chain.GetRequiredKeys(trans, publicKeys)
	//requiredKeys := gjson.GetBytes(requiredKeysDat, "required_keys").Array()
	//fmt.Println(publicKeys)
	sigPara := []interface{}{
		trans, []string{"EOS7KPzd3KwhR4jKUFyMXUhyvxUP6NV9Sf8f9v5CC6cs9PQ6AzgWb"}, "",
	}
	sigDat, _ := wallet.WalletSignTrx(sigPara)
	signatures := gjson.GetBytes(sigDat, "signatures").Array()
	var sigs []string
	for _, s := range signatures {
		sigs = append(sigs, s.String())
	}
	var transMap = map[string]interface{}{
		"signatures":        sigs,
		"compression":       "none",
		"context_free_data": []string{},
		"transaction":       trans,
	}
	//trans.Signatures=signatures

	responce, _ := chain.PushTransaction(transMap)
	json, _ := json.Marshal(trans)
	fmt.Println(string(json))
	fmt.Print(string(responce))
}
