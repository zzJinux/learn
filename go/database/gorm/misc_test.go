package learngorm

import (
	"testing"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"learn/database"
	. "learn/database/u"
)

func TestCreateIndex1(t *testing.T) {
	type User1 struct {
		ID   string
		Name string `gorm:"size:255;index:idx_name,unique"`
		Foo  string
		Bar  string
	}
	db := Connect(database.EnvPort)

	tb := User1{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))
}

func TestCreateIndex2(t *testing.T) {
	type User2 struct {
		ID   string
		Name string `gorm:"size:255;index:idx_name,unique"`
		Foo  string
		Bar  string
	}
	db := Connect(database.EnvPort)

	tb := User2{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))

	Must(db.Migrator().CreateIndex(tb, "Name"))
	// Duplicate key name 'idx_name'
}

func TestCreateIndex3(t *testing.T) {
	type User3 struct {
		ID   string
		Name string `gorm:"size:255;index:idx_name,unique"`
		Foo  string
		Bar  string
	}
	db := Connect(database.EnvPort)

	tb := User3{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))

	Must(db.Migrator().CreateIndex(tb, "idx_name"))
	// Duplicate key name 'idx_name'
}

func TestCreateIndex4(t *testing.T) {
	type User4 struct {
		ID   string
		Name string `gorm:"size:255;index:idx_name,unique"`
		Foo  string
		Bar  string
	}
	db := Connect(database.EnvPort)

	tb := User4{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))

	db = db.Debug()
	Must(db.Migrator().CreateIndex(tb, "idx_name"))
	// Table user4 doesn't exist
}

func TestCreateIndex5(t *testing.T) {
	type User5 struct {
		ID   string
		Name string
		Foo  string `gorm:"size:255;uniqueIndex:idx_aaa"`
		Bar  string `gorm:"size:255;uniqueIndex:idx_aaa"`
	}
	db := Connect(database.EnvPort)

	tb := User5{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))
	db = db.Debug()
}

func TestCreateIndex6(t *testing.T) {
	type User6 struct {
		ID   string
		Name string
		Foo  string `gorm:"size:255"`
		Bar  string `gorm:"size:255"`
	}
	db := Connect(database.EnvPort)

	tb := User6{}
	if db.Migrator().HasTable(tb) {
		Must(db.Migrator().DropTable(tb))
	}
	db = db.Debug()
	Must(db.Migrator().CreateTable(tb))

	Must(CreateUniqueIndexIfNotExists(db, tb, "idx_hoo", "foo", "bar"))
}

func getTableName(db *gorm.DB, table any) (string, error) {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(table)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get table name")
	}
	return stmt.Schema.Table, nil
}

func CreateUniqueIndexIfNotExists(db *gorm.DB, table any, indexName string, columns ...string) error {
	if db.Migrator().HasIndex(table, indexName) {
		return nil
	}
	tableName, err := getTableName(db, table)
	if err != nil {
		return nil
	}

	sqlText := "CREATE UNIQUE INDEX ? ON ? ?"
	values := []any{clause.Column{Name: indexName}, clause.Table{Name: tableName}, nil}
	columns2 := make([]clause.Column, 0, len(columns))
	for _, column := range columns {
		columns2 = append(columns2, clause.Column{Name: column})
	}
	values[2] = columns2

	tx := db.Exec(sqlText, values...)
	if tx.Error != nil {
		return errors.Wrapf(tx.Error, "failed to create unique index")
	}
	return nil
}
