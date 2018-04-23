package entity

type Authorization struct {
	Account    string `json:"account"`
	Permission string `json:"permission"`
}

type Action struct {
	Code          string          `json:"code"`
	Type          string          `json:"type"`
	Recipients    []string        `json:"recipients"`
	Authorization []Authorization `json:"authorization"`
	Data          string          `json:"data"`
}
type TransferDat struct {
	RefBlockNum    string   `json:"ref_block_num"`
	RefBlockPrefix string   `json:"ref_block_prefix"`
	Expiration     string   `json:"expiration"`
	Scope          []string `json:"scope"`
	Actions        []Action `json:"actions"`
	Signatures     []string `json:"signatures"`
	Authorizations []string `json:"authorizations"`
}

type AuthorizationV3 struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type ActionV3 struct {
	Account       string            `json:"account"`
	Name          string            `json:"name"`
	Authorization []AuthorizationV3 `json:"authorization"`
	Data          string            `json:"data"`
}
type TransferV3 struct {
	Region             int      `json:"region"`
	RefBlockNum        string   `json:"ref_block_num"`
	RefBlockPrefix     string   `json:"ref_block_prefix"`
	Expiration         string   `json:"expiration"`
	MaxNetUsageWords   int      `json:"max_net_usage_words"`
	MaxKcpuUsage       int      `json:"max_kcpu_usage"`
	DelaySec           int      `json:"delay_sec"`
	ContextFreeActions []string `json:"context_free_actions"`
	Actions            []ActionV3 `json:"actions"`
}
