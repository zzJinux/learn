package learngorm

import (
	"testing"

	"gorm.io/gorm"

	"learn/database"

	rq "github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type Model struct {
		ID   string
		Name string
		Foo  string
		Bar  string
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	tx := db.Create([]Model{
		{ID: "id1", Name: "name1", Foo: "foo1", Bar: "bar1"},
		{ID: "id2", Name: "name2", Foo: "foo2", Bar: "bar2"},
		{ID: "id3", Name: "name3", Foo: "foo3", Bar: "bar3"},
	})
	rq.NoError(t, tx.Error)

	// Update single column
	tx = db.Model(Model{}).Where("id = ?", "id1").Update("foo", "foo1-updated")
	rq.NoError(t, tx.Error)
	r1 := Model{ID: "id1"}
	db.First(&r1)
	rq.Equal(t, "foo1-updated", r1.Foo)

	// Update multiple columns
	tx = db.Model(Model{}).Where("id = ?", "id2").Updates(
		Model{Foo: "foo2-updated", Bar: "bar2-updated"},
	)
	rq.NoError(t, tx.Error)
	r2 := Model{ID: "id2"}
	db.First(&r2)
	rq.Equal(t, "foo2-updated", r2.Foo)
	rq.Equal(t, "bar2-updated", r2.Bar)

	// Does 'Updates' modify args? -> No
	f3 := &Model{Foo: "foo2-updated", Bar: "bar2-updated"}
	tx = db.Model(Model{}).Where("id = ?", "id2").Updates(f3)
	rq.NoError(t, tx.Error)
	rq.Empty(t, f3.ID)
	rq.Empty(t, f3.Name)
}

func TestUpdateExpr(t *testing.T) {
	type Model struct {
		ID   string
		Name string
		V    int
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	GormMust(db.Create([]Model{
		{ID: "id1", Name: "name1", V: 10},
	}))

	GormMust(db.Model(Model{ID: "id1"}).Updates(map[string]any{
		"v": gorm.Expr("v + ?", 1),
	}))

	{
		var m Model
		m.ID = "id1"
		GormMust(db.Find(&m))
		rq.Equal(t, 11, m.V)
	}
}
