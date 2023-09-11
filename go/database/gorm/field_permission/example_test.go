package field_permission

import (
	"fmt"
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"
)

// https://gorm.io/docs/models.html#Field-Level-Permission

// Caveat:
// The SQL statement generation is not affected by field permission tags. (huh?)
// Even if a field is marked as IGNORED, the corresponding column is not excluded from the SQL statement.

func TestExample(t *testing.T) {
	type Foo struct {
		ID string

		ColIgnored  string `gorm:"-"`  // ignored
		ColReadOnly string `gorm:"->"` // read-only
	}

	type FooEx struct {
		ID string

		ColIgnored  string
		ColReadOnly string
	}

	sess := Connect(database.EnvPort)
	tb := FooEx{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS foo"))

	sess = sess.Debug()
	Must(sess.Table("foo").Migrator().CreateTable(tb))
	GormMust(sess.Table("foo").Create(&Foo{ID: "id1", ColIgnored: "ignored", ColReadOnly: "read-only"}))
	{
		foo := Foo{
			ID: "id1",
		}
		sess.First(&foo)
		fmt.Printf("%+v\n", foo)
	}

	GormMust(sess.Table("foo").Create(&FooEx{ID: "id2", ColIgnored: "ignored", ColReadOnly: "read-only"}))
	{
		foo := Foo{
			ID: "id2",
		}
		sess.First(&foo)
		fmt.Printf("%+v\n", foo)
	}
}
