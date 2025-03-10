package webapp

import "github.com/MyriadFlow/storefront-gateway/models"

type WebappPayload struct {
	SubgraphId string `json:"subgraphId,omitempty"`
}

type WebappResponse struct {
	Storefront          models.Storefront `json:"Storefront,omitempty"`
	TradehubAddress     string            `json:"TradehubAddress,omitempty"`
	AccessMasterAddress string            `json:"accessMasterAddress,omitempty"`
	BaseUrlGateway      string            `json:"baseUrlGateway,omitempty"`
	IpfsGateway         string            `json:"ipfsGateway,omitempty"`
	Profile             models.User       `json:"profile,omitempty"`
	RkProjectId         string            `json:"rkProjectId,omitempty"`
	AlchemyId           string            `json:"alchemyId,omitempty"`
}

type Contract struct {
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	CollectionName  string `json:"collectionName"`
	Thumbnail       string `json:"thumbnail"`
	CoverImage      string `json:"coverImage"`
	GraphUrl        string `json:"graphUrl"`
	Drops           bool   `json:"drops"`
}
