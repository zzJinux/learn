package conditional

import (
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"

	"github.com/stretchr/testify/assert"
)

func TestUpdateConditions(t *testing.T) {
	type Human struct {
		ID   string
		Name string
		Job  string
		Loc  string
	}
	sess := Connect(database.EnvPort)
	GormMust(sess.Exec("DROP TABLE IF EXISTS human"))
	Must(sess.Migrator().CreateTable(Human{}))
	sess = sess.Debug()

	GormMust(sess.Create([]Human{
		{ID: "123", Name: "John", Job: "Engineer", Loc: "Seoul"},
		{ID: "234", Name: "Baba", Job: "Engineer", Loc: "Seoul"},
		{ID: "345", Name: "Chop", Job: "Doctor", Loc: "Seoul"},
		{ID: "456", Name: "Daniel", Job: "Doctor", Loc: "San Francisco"},
	}))

	// Model object contains primary key
	GormMust(sess.Model(Human{ID: "123"}).Update("loc", "Incheon"))
	{
		var h Human
		h.ID = "123"
		GormMust(sess.First(&h))
		assert.Equal(t, "Incheon", h.Loc)
	}

	// WHERE is must
	GormMust(sess.Where(Human{Job: "Engineer"}).Updates(Human{Loc: "Busan"}))
	{
		var h Human
		h.Name = "John"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "Busan", h.Loc)
	}
	{
		var h Human
		h.Name = "Baba"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "Busan", h.Loc)
	}
	{
		var h Human
		h.Name = "Chop"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "Seoul", h.Loc)
	}

	// Engineer is ignored
	GormMust(sess.Model(Human{Job: "Engineer"}).Where(Human{Job: "Doctor"}).Update("loc", "New York"))
	{
		var h Human
		h.Name = "Baba"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "Busan", h.Loc)
	}
	{
		var h Human
		h.Name = "Chop"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "New York", h.Loc)
	}
	{
		var h Human
		h.Name = "Daniel"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "New York", h.Loc)
	}
}

func TestUpdateCompositeKey(t *testing.T) {
	type Foo struct {
		Group string `gorm:"primaryKey"`
		Name  string `gorm:"primaryKey"`
		Value string
	}
	sess := Connect(database.EnvPort)
	GormMust(sess.Exec("DROP TABLE IF EXISTS foo"))
	Must(sess.Migrator().CreateTable(Foo{}))
	sess = sess.Debug()

	GormMust(sess.Create([]Foo{
		{Group: "A", Name: "John", Value: "aaa"},
		{Group: "A", Name: "Ferry", Value: "bbb"},
		{Group: "B", Name: "John", Value: "ccc"},
		{Group: "B", Name: "Ferry", Value: "ddd"},
	}))

	GormMust(sess.Model(Foo{Group: "A", Name: "John"}).Updates(Foo{Value: "yay"}))
	{
		var f Foo
		f.Group = "A"
		f.Name = "John"
		GormMust(sess.First(&f))
		assert.Equal(t, "yay", f.Value)

		f.Group = "A"
		f.Name = "Ferry"
		GormMust(sess.First(&f))
		assert.Equal(t, "bbb", f.Value)
	}
}
