// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
    // log "github.com/romainm/buddy-server/internal/logger"

    "github.com/romainm/buddy-server/internal/orm/migration"

    "github.com/romainm/buddy-server/pkg/utils"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/jinzhu/gorm"
)

var autoMigrate, logMode bool
var dbPath, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
    DB *gorm.DB
}

func init() {
    dialect = utils.MustGet("GORM_DIALECT")
    dbPath = utils.MustGet("GORM_DB_PATH")
    logMode = utils.MustGetBool("GORM_LOGMODE")
    autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
    db, err := gorm.Open(dialect, dbPath)
    if err != nil {
        // log.Panic("[ORM] err: ", err)
        // log.Panic("[ORM] err: ", err)
    }
    orm := &ORM{
        DB: db,
    }
    // Log every SQL command on dev, @prod: this should be disabled?
    db.LogMode(logMode)
    // Automigrate tables
    if autoMigrate {
        err = migration.ServiceAutoMigration(orm.DB)
    }
    // log.Info("[ORM] Database connection initialized.")
    return orm, err
}
