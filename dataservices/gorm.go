package dataservices

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMSSQLGormDB creates a GORM DB session towards a Microsoft SQL Server
// it uses the supplied Logger to create the GORM Logger.
func NewMSSQLGormDB(
	conf *gorm.Config,
	connectionString string,
	maxOpenConn,
	maxIdleConn int,
	maxConnLifeTime time.Duration,
	gormMigrations []*gormigrate.Migration) (*gorm.DB, error) {

	if conf == nil {
		conf = &gorm.Config{}
	}
	gormDB, err := gorm.Open(sqlserver.Open(connectionString), conf)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create the GORM session towards the Microsoft SQL server")
	}
	// Execute migrations on if they are present
	if gormMigrations != nil {
		migrator := gormigrate.New(gormDB, gormigrate.DefaultOptions, gormMigrations)
		if err = migrator.Migrate(); err != nil {
			return nil, errors.Wrap(err, "Failed to migrate the table")
		}
	}
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	if maxConnLifeTime > 0 {
		db.SetConnMaxLifetime(maxConnLifeTime)
	}
	return gormDB, nil
}
