package jobs

import (
    "github.com/romainm/buddy-server/internal/orm/models"
    "github.com/jinzhu/gorm"
    "gopkg.in/gormigrate.v1"
)

var (
	name = "Optus"
	amount = 123.45
    firstTransaction   *models.Transaction = &models.Transaction{
        Name:        &name,
		Amount:		 &amount,
    }
)

// SeedTransactions inserts the first transactions
var SeedTransactions *gormigrate.Migration = &gormigrate.Migration{
    ID: "SEED_TRANSACTIONS",
    Migrate: func(db *gorm.DB) error {
        return db.Create(&firstTransaction).Error
    },
    Rollback: func(db *gorm.DB) error {
        return db.Delete(&firstTransaction).Error
    },
}
