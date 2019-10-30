package jobs

import (
    "github.com/romainm/buddy-server/internal/orm/models"
    "github.com/jinzhu/gorm"
    "gopkg.in/gormigrate.v1"
)

var (
    firstTransaction   *models.Transaction = &models.Transaction{
        Name:        "Optus",
		Amount:		 69.99,
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
