package serialize

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"learn/database"

	. "learn/database/gorm"
)

func TestInsertUpdateSelect(t *testing.T) {
	type SomeStruct struct {
		A, B, C string
		D       struct {
			E, F string
		}
	}
	type Model struct {
		ID   string
		Name string
		S    SomeStruct `gorm:"type:json;serializer:json"`
	}

	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	{
		var s SomeStruct
		s.A = "a"
		s.B = "b"
		s.C = "c"
		s.D.E = "e"
		s.D.F = "f"
		tx := db.Create(&Model{
			ID:   "hello",
			Name: "world",
			S:    s,
		})
		GormMust(tx)
	}

	{
		var m Model
		tx := db.First(&m, Model{ID: "hello"})
		GormMust(tx)
		assert.Equal(t, "a", m.S.A)
		assert.Equal(t, "b", m.S.B)
		assert.Equal(t, "c", m.S.C)
		assert.Equal(t, "e", m.S.D.E)
		assert.Equal(t, "f", m.S.D.F)
	}

	{
		var s SomeStruct
		s.A = "a1"
		s.B = "b1"
		s.C = "c1"
		s.D.E = "e1"
		s.D.F = "f1"
		{
			tx := db.Model(Model{ID: "hello"}).Update("s", s)
			GormMust(tx)
		}

		var m Model
		{
			tx := db.First(&m, Model{ID: "hello"})
			GormMust(tx)
		}
		assert.Equal(t, "a1", m.S.A)
		assert.Equal(t, "b1", m.S.B)
		assert.Equal(t, "c1", m.S.C)
		assert.Equal(t, "e1", m.S.D.E)
		assert.Equal(t, "f1", m.S.D.F)
	}
}
