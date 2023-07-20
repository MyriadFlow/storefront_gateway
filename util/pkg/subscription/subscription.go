package subscription

import (
	"time"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/google/uuid"
)

func CreateSubscription(name string, owner string, plan string, cost int, currency string, createdBy string, updatedBy string, image string) error {
	db := dbconfig.GetDb()
	subscription := models.Subscription{
		Id:        uuid.New(),
		Name:      name,
		Owner:     owner,
		Plan:      plan,
		Cost:      cost,
		Currency:  currency,
		Status:    "active",
		Validity:  time.Now().AddDate(1, 0, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
		Image:     image,
	}
	result := db.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
