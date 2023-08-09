package models

type Contract struct {
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	ChainId         int    `json:"chainId"`
	Verified        bool   `json:"verified"`
	StorefrontId    string `json:"storefrontId"`
	ContractId      string `json:"contractId"`
}
