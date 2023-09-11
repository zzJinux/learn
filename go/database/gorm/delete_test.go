package learngorm

import (
	"testing"
	"time"

	"learn/database"
	. "learn/database/u"

	rq "github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func TestSoftDelete1(t *testing.T) {
	type SoftDelete1 struct {
		ID  string `gorm:"primaryKey"`
		Val string

		// gorm treats this as a normal field, not a soft-delete field.
		DeletedAt *time.Time
	}
	db := Connect(database.EnvPort)
	ResetTable(db, SoftDelete1{})
	db = db.Debug()

	// Explicitly set DeletedAt to a non-nil value.
	tx := db.Create(SoftDelete1{
		ID:        "aaa",
		Val:       "val",
		DeletedAt: Ptr(time.Now()), // This won't soft-delete the record.
	})
	rq.NoError(t, tx.Error)

	found := SoftDelete1{ID: "aaa"}
	tx = db.First(&found)
	rq.NoError(t, tx.Error)
	rq.Equal(t, "val", found.Val)
}

func TestSoftDelete2(t *testing.T) {
	type Model struct {
		ID  string `gorm:"primaryKey"`
		Val string

		// gorm's way
		// https://gorm.io/docs/delete.html#Soft-Delete
		DeletedAt gorm.DeletedAt // works
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	{
		// Explicitly set DeletedAt to a non-nil value.
		tx := db.Create(Model{
			ID:  "aaa",
			Val: "val",
			// Since Valid is false, it won't soft-delete the record.
			DeletedAt: gorm.DeletedAt{Time: time.Now()},
		})
		rq.NoError(t, tx.Error)

		found := Model{ID: "aaa"}
		tx = db.First(&found)
		rq.NoError(t, tx.Error)
		rq.Equal(t, "val", found.Val)
	}

	{
		// Explicitly set DeletedAt to a non-nil value.
		tx := db.Create(Model{
			ID:        "bbb",
			Val:       "val",
			DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true},
		})
		rq.NoError(t, tx.Error)

		found := Model{ID: "bbb"}
		tx = db.First(&found)
		rq.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)
	}
}

// Check what SQLs are generated.
func TestSoftDeleteVarious(t *testing.T) {
	type ExampleSoftDel1 struct {
		Name string `gorm:"primaryKey"`
		// unix seconds
		soft_delete.DeletedAt
	}
	type ExampleSoftDel2 struct {
		Name string `gorm:"primaryKey"`
		// unix milliseconds
		soft_delete.DeletedAt `gorm:"softDelete:milli"`
	}
	type ExampleSoftDel3 struct {
		Name string `gorm:"primaryKey"`
		gorm.DeletedAt
	}
	type ExampleSoftDel4 struct {
		Name           string `gorm:"primaryKey"`
		gorm.DeletedAt `gorm:"precision:3"`
	}
	db := Connect(database.EnvPort)
	ResetTable(db, ExampleSoftDel1{})
	ResetTable(db, ExampleSoftDel2{})
	ResetTable(db, ExampleSoftDel3{})
	ResetTable(db, ExampleSoftDel4{})
	db = db.Debug()

	var tx *gorm.DB

	tx = db.Create(&ExampleSoftDel1{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Create(&ExampleSoftDel2{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Create(&ExampleSoftDel3{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Create(&ExampleSoftDel4{Name: "a"})
	rq.NoError(t, tx.Error)

	tx = db.Delete(&ExampleSoftDel1{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Delete(&ExampleSoftDel2{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Delete(&ExampleSoftDel3{Name: "a"})
	rq.NoError(t, tx.Error)
	tx = db.Delete(&ExampleSoftDel4{Name: "a"})
	rq.NoError(t, tx.Error)
}

func TestDoubleDelete(t *testing.T) {
	type Model struct {
		ID  string
		Val string
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	tx := db.Create(Model{
		ID:  "aaa",
		Val: "val",
	})
	rq.NoError(t, tx.Error)

	// first delete
	tx = db.Delete(Model{ID: "aaa"})
	rq.NoError(t, tx.Error)
	rq.Equal(t, int64(1), tx.RowsAffected)

	// second delete
	tx = db.Delete(Model{ID: "aaa"})
	rq.NoError(t, tx.Error)
	rq.Equal(t, int64(0), tx.RowsAffected)
}
