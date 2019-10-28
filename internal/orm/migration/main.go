package migration

import (
    "fmt"

    // log "github.com/romainm/buddy-server/internal/logger"

    "github.com/romainm/buddy-server/internal/orm/migration/jobs"
    "github.com/romainm/buddy-server/internal/orm/models"
    "github.com/jinzhu/gorm"
    "gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.Transaction{},
    ).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
    // Keep a list of migrations here
    m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
    m.InitSchema(func(db *gorm.DB) error {
        // log.Info("[Migration.InitSchema] Initializing database schema")
        if err := updateMigration(db); err != nil {
            return fmt.Errorf("[Migration.InitSchema]: %v", err)
        }
        // Add more jobs, etc here
        return nil
    })
    m.Migrate()

    if err := updateMigration(db); err != nil {
        return err
    }
    m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        jobs.SeedTransactions,
    })
    return m.Migrate()
}
