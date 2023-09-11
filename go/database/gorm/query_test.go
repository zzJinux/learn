package learngorm

import (
	"testing"

	"learn/database"

	rq "github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestQueryLimit(t *testing.T) {
	type Model struct {
		ID   string
		Name string
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	tx := db.Create([]Model{
		{ID: "id1", Name: "name1"},
		{ID: "id2", Name: "name2"},
		{ID: "id3", Name: "name3"},
	})
	rq.NoError(t, tx.Error)

	var ents []Model
	tx = db.Limit(-1).Find(&ents)
	rq.NoError(t, tx.Error)
	rq.Len(t, ents, 3)

	// There is a bug that Limit(0) returns all rows. See https://github.com/go-gorm/gorm/pull/6191
	tx = db.Limit(0).Find(&ents)
	rq.NoError(t, tx.Error)
	rq.Len(t, ents, 0)

	tx = db.Limit(1).Find(&ents)
	rq.NoError(t, tx.Error)
	rq.Len(t, ents, 1)
}

func TestQuerySelectColumn(t *testing.T) {
	t.Run("column w/o default value", func(t *testing.T) {
		var tx *gorm.DB
		type Model struct {
			ID    string
			Name  string
			Alias string
			Count int
		}
		db := Connect(database.EnvPort)
		ResetTable(db, Model{})
		db = db.Debug()

		ent := Model{
			ID:    "idid1",
			Name:  "namename",
			Alias: "aliasalias",
			Count: 999,
		}
		tx = db.Select("name", "count").Create(ent)
		rq.Error(t, tx.Error)
	})

	// PK has default value
	t.Run("column w/ default value", func(t *testing.T) {
		var tx *gorm.DB
		type Model struct {
			ID    string `gorm:"primaryKey"`
			Name  string
			Alias string
			Count int
		}
		db := Connect(database.EnvPort)
		ResetTable(db, Model{})
		db = db.Debug()

		ent := Model{
			ID:    "idid1",
			Name:  "namename",
			Alias: "aliasalias",
			Count: 999,
		}
		tx = db.Select("name", "count").Create(ent)
		rq.Error(t, tx.Error)

		ent.ID = "idid2"
		ent.Name = "gobbbb"
		tx = db.Create(ent)
		rq.NoError(t, tx.Error)
	})
}
