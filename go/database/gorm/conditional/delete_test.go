package conditional

import (
	"testing"

	"learn/database"
	. "learn/database/gorm"
	. "learn/database/u"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDeleteConditions(t *testing.T) {
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

	// GormMust(sess.Delete(Human{}, Human{Job: "Engineer"}))

	// GormMust(sess.Where(Human{Job: "Engineer"}).Delete(Human{}))

	h := Human{Job: "Engineer"}
	GormMust(sess.Delete(h, h))
	{
		var h Human
		h.Name = "John"
		tx := sess.First(&h, h)
		assert.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)
	}
	{
		var h Human
		h.Name = "Baba"
		tx := sess.First(&h, h)
		assert.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)
	}
	{
		var h Human
		h.Name = "Chop"
		GormMust(sess.First(&h, h))
		assert.Equal(t, "Seoul", h.Loc)
	}

	GormMust(sess.Delete(Human{ID: "456"}))
	{
		var h Human
		h.ID = "456"
		tx := sess.First(&h)
		assert.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)
	}
}
