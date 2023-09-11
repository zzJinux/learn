package conditional

import (
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"

	"github.com/stretchr/testify/assert"
)

func TestInlinedCondition(t *testing.T) {
	type Human struct {
		ID   string
		Name string
	}
	sess := Connect(database.EnvPort)
	tb := Human{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS human"))
	Must(sess.Migrator().CreateTable(tb))
	sess = sess.Debug()

	orig := Human{ID: "123", Name: "John"}
	another := Human{ID: "234", Name: "Baba"}
	GormMust(sess.Create(orig))
	GormMust(sess.Create(another))

	// SELECT (First)
	{
		var got Human
		got.ID = "123"
		// NOTE: First finds records by primary key
		GormMust(sess.First(&got))
		assert.Equal(t, orig, got)
	}
	// SELECT (Find)
	{
		var got Human
		got.ID = "123"
		// NOTE: First finds records by primary key
		GormMust(sess.Find(&got))
		assert.Equal(t, orig, got)
	}

	// UPDATE
	GormMust(sess.Save(Human{ID: "123", Name: "new John"}))
	var got3 Human
	got3.ID = "123"
	GormMust(sess.First(&got3))
	assert.Equal(t, "new John", got3.Name)
}

func TestConditionByStruct(t *testing.T) {
	type Human struct {
		ID    string
		Name  string
		Value string
	}
	sess := Connect(database.EnvPort)
	tb := Human{}
	GormMust(sess.Exec("DROP TABLE IF EXISTS human"))
	Must(sess.Migrator().CreateTable(tb))
	sess = sess.Debug()

	GormMust(sess.Create([]Human{
		{ID: "123", Name: "John", Value: "aaa"},
		{ID: "234", Name: "John", Value: "bbb"},
		{ID: "345", Name: "Ferry", Value: "ccc"},
	}))

	var got []Human
	GormMust(sess.Find(&got, Human{Name: "John"}))
	assert.Len(t, got, 2)
	assert.ElementsMatch(t, []string{"123", "234"}, []string{got[0].ID, got[1].ID})

	// UPDATE
	{
		GormMust(sess.Where(Human{Name: "John"}).Updates(Human{Value: "yay"}))
		var got []Human
		GormMust(sess.Find(&got, Human{Name: "John"}))
		assert.Len(t, got, 2)
		assert.ElementsMatch(t, []string{"yay", "yay"}, []string{got[0].Value, got[1].Value})
	}
}

func TestCompositeKey(t *testing.T) {
	type Param struct {
		Family string `gorm:"primaryKey"`
		Key    string `gorm:"primaryKey"`
		Value  string
	}
	sess := Connect(database.EnvPort)
	GormMust(sess.Exec("DROP TABLE IF EXISTS param"))
	Must(sess.Migrator().CreateTable(Param{}))
	sess = sess.Debug()

	GormMust(sess.Create([]Param{
		{Family: "moo", Key: "John", Value: "aaa"},
		{Family: "moo", Key: "Ferry", Value: "bbb"},
		{Family: "zoo", Key: "John", Value: "asdf"},
	}))

	var got []Param

	GormMust(sess.Find(&got, Param{Family: "moo"}))
	assert.Len(t, got, 2)
	assert.ElementsMatch(t, []string{"aaa", "bbb"}, []string{got[0].Value, got[1].Value})

	GormMust(sess.Find(&got, Param{Family: "moo", Key: "John"}))
	assert.Len(t, got, 1)
	assert.ElementsMatch(t, []string{"aaa"}, []string{got[0].Value})

	{
		got := Param{Family: "moo", Key: "John"}
		GormMust(sess.First(&got))
		assert.Equal(t, "aaa", got.Value)
	}
}
