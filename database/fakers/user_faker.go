package fakers

import (
	"time"

	"github.com/Anka-Abdullah/Go-toko-1/app/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		Addresses:     nil,
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.cGFzc3dvcmQ.gGOQG7Yl79jqan3QSjNJQPMrw81gbtpUW4fnKVM4WwA",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
