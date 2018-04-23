package enum

/*URL 访问地址*/
var URL = ""

const (
	V1          = "/v1"
	V1ChainUri  = V1 + "/chain/"
	V1WalletUri = V1 + "/wallet/"

	Get_info_v1          = V1ChainUri + "get_info"
	Get_block_v1         = V1ChainUri + "get_block"
	Get_account_v1       = V1ChainUri + "get_account"
	Get_code_v1          = V1ChainUri + "get_code"
	Get_table_rows_v1    = V1ChainUri + "get_table_rows"
	Abi_json_to_bin_v1   = V1ChainUri + "abi_json_to_bin"
	Abi_bin_to_json_v1   = V1ChainUri + "abi_bin_to_json"
	Push_transaction_v1  = V1ChainUri + "push_transaction"
	Push_transactions_v1 = V1ChainUri + "push_transactions"
	Get_required_keys_v1 = V1ChainUri + "get_required_keys"

	Wallet_create_v1          = V1WalletUri + "create"
	Wallet_open_v1            = V1WalletUri + "open"
	Wallet_lock_v1            = V1WalletUri + "lock"
	Wallet_lock_all_v1        = V1WalletUri + "lock_all"
	Wallet_unlock_v1          = V1WalletUri + "unlock"
	Wallet_import_key_v1      = V1WalletUri + "import_key"
	Wallet_list_v1            = V1WalletUri + "list_wallets"
	Wallet_list_keys_v1       = V1WalletUri + "list_keys"
	Wallet_get_public_keys_v1 = V1WalletUri + "get_public_keys"
	Wallet_set_timeout_v1     = V1WalletUri + "set_timeout"
	Wallet_sign_trx_v1        = V1WalletUri + "sign_transaction"
)
