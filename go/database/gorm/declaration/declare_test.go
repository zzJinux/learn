package declaration

import (
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"
)

// Go Struct embedding, Gorm embedding
func TestEmbedding1(t *testing.T) {
	type Human struct {
		ID   string
		Name string
	}
	type Superhuman struct {

		// Pure embedding
		Human

		Power string
	}
	sess := Connect(database.EnvPort)
	tb := Superhuman{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS superhuman"))

	sess = sess.Debug()
	Must(sess.Migrator().CreateTable(tb))
}

func TestEmbedding2(t *testing.T) {
	type Human struct {
		ID   string
		Name string
	}
	// No struct tags
	type Superhuman struct {

		// Ordinary field
		Human Human

		Power string
	}
	sess := Connect(database.EnvPort)
	tb := Superhuman{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS superhuman"))

	sess = sess.Debug()
	Must(sess.Migrator().CreateTable(tb))
	// Panic!
}

func TestEmbedding3(t *testing.T) {
	type Human struct {
		ID   string
		Name string
	}
	// No struct tags
	type Superhuman struct {

		// Ordinary field + gorm embedding
		Human Human `gorm:"embedded"`

		Power string
	}
	sess := Connect(database.EnvPort)
	tb := Superhuman{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS superhuman"))

	sess = sess.Debug()
	Must(sess.Migrator().CreateTable(tb))
}

func TestColumnName(t *testing.T) {
	type Human struct {
		ID   string
		Name string `gorm:"column:full_name"`
	}
	sess := Connect(database.EnvPort)
	tb := Human{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS human"))

	sess = sess.Debug()
	Must(sess.Migrator().CreateTable(tb))
}
