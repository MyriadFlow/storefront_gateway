package flowid

import (
	"github.com/TheLazarusNetwork/marketplace-engine/db"
	"github.com/TheLazarusNetwork/marketplace-engine/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, update bool, flowIdType models.FlowIdType, relatedRoleId string) (string, error) {
	flowId := uuid.NewString()
	if update {
		// User exist so update
		association := db.Db.Model(&models.User{
			WalletAddress: walletAddress,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		association.Append(&models.FlowId{FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId})

	} else {
		// User doesn't exist so create

		newUser := &models.User{
			WalletAddress: walletAddress,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId,
			}},
		}
		if err := db.Db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
