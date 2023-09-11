package customdatatype

import (
	"fmt"
	"testing"

	"learn/database"
	. "learn/database/gorm"
)

func TestByteSliceColumn(t *testing.T) {
	type Model struct {
		ID  string
		Foo []byte `gorm:"type:json;serializer:json"`
	}

	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	{
		tx := db.Create(&Model{
			ID:  "hello",
			Foo: []byte(`{"a": 1, "b": 2}`),
		})
		GormMust(tx)
	}
	{
		tx := db.Model(Model{}).Create(map[string]any{
			"id":  "world",
			"foo": []byte(`{"c": 3, "d": 4}`),
		})
		GormMust(tx)
	}
	{
		tx := db.Where("id", "hello").Updates(&Model{
			Foo: []byte(`{"a": 1, "b": 2}`),
		})
		GormMust(tx)
	}
	{
		tx := db.Where("id", "world").Model(Model{}).Updates(map[string]any{
			"foo": []byte(`{"c": 3, "d": 4}`),
		})
		GormMust(tx)
	}
	{
		tx := db.Where("id", "world").Model(Model{}).Update("foo", []byte(`{"c": 3, "d": 4}`))
		GormMust(tx)
	}
	{
		var obj Model
		obj.ID = "hello"
		tx := db.First(&obj)
		GormMust(tx)
		fmt.Println(obj)
	}
	{
		var obj Model
		obj.ID = "world"
		tx := db.First(&obj)
		GormMust(tx)
		fmt.Println(obj)
	}
}

func TestAnyColumn(t *testing.T) {
	type Foo struct {
		A int
		B struct {
			Nested1 string
			Nested2 string
		}
	}
	type Model struct {
		ID  string
		Foo any `gorm:"type:json;serializer:json"`
	}

	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	{
		var foo Foo
		foo.A = 1
		foo.B.Nested1 = "nested1"
		foo.B.Nested2 = "nested2"
		tx := db.Create(&Model{
			ID:  "hello",
			Foo: foo,
		})
		GormMust(tx)
	}

	{
		var m Model
		var foo Foo
		m.ID = "hello"
		m.Foo = &foo

		tx := db.First(&m)
		GormMust(tx)
		fmt.Println(m)
	}
}

func TestPseudoSumType(t *testing.T) {
	type Foo1 struct {
		Foo1Field string
	}
	type Foo2 struct {
		Foo2Field string
	}
	type Foo struct {
		*Foo1
		*Foo2
	}
	type Model struct {
		ID  string
		Foo any `gorm:"type:json;serializer:json"`
	}

	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	{
		var foo Foo
		foo.A = 1
		foo.B.Nested1 = "nested1"
		foo.B.Nested2 = "nested2"
		tx := db.Create(&Model{
			ID:  "hello",
			Foo: foo,
		})
		GormMust(tx)
	}

	{
		var m Model
		var foo Foo
		m.ID = "hello"
		m.Foo = &foo

		tx := db.First(&m)
		GormMust(tx)
		fmt.Println(m)
	}
}
