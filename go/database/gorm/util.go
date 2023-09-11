package learngorm

import (
	"fmt"

	. "learn/database/u"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Connect creates a connection to the MySQL database.
func Connect(port uint16) *gorm.DB {
	// DSN format: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string.
	// The date or datetime like 0000-00-00 00:00:00 is converted into zero value of time.Time.
	// https://github.com/go-sql-driver/mysql#parsetime
	dbname := "learn_sandbox"
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:%d)/?parseTime=true&multiStatements=true", port)
	precision := 6
	db := Must1(gorm.Open(mysql.New(mysql.Config{DSN: dsn, DefaultDatetimePrecision: &precision}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}))

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if _, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + "; USE " + dbname + ";"); err != nil {
		panic(err)
	}
	return db
}

func GormMust(tx *gorm.DB) *gorm.DB {
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx
}

func ResetTable(db *gorm.DB, m any) {
	if db.Migrator().HasTable(m) {
		Must(db.Migrator().DropTable(m))
	}
	Must(db.Migrator().CreateTable(m))
}
