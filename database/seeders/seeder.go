package seeders

import (
	"github.com/Anka-Abdullah/Go-toko-1/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegistrationSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegistrationSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
