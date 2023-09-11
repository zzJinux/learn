package learngorm

import (
	"testing"

	"learn/database"

	rq "github.com/stretchr/testify/require"
	"gorm.io/gorm/clause"
)

func TestCompositeKey(t *testing.T) {
	type Model struct {
		PrimaryFoo string `gorm:"primaryKey"`
		PrimaryBar string `gorm:"primaryKey"`
		Val        string
	}
	db := Connect(database.EnvPort)
	ResetTable(db, Model{})
	db = db.Debug()

	// OnConflict: https://gorm.io/gen/create.html#Upsert-x2F-On-Conflict

	// Prepare initial data
	{
		tx := db.Create([]Model{
			{
				PrimaryFoo: "a1",
				PrimaryBar: "b1",
				Val:        "val1",
			},
			{
				PrimaryFoo: "a2",
				PrimaryBar: "b2",
				Val:        "val2",
			},
		})
		rq.NoError(t, tx.Error)

		found1 := Model{
			PrimaryFoo: "a1",
			PrimaryBar: "b1",
		}
		tx = db.Find(&found1)
		rq.NoError(t, tx.Error)
		rq.Equal(t, "val1", found1.Val)

		found2 := Model{
			PrimaryFoo: "a2",
			PrimaryBar: "b2",
		}
		tx = db.Find(&found2)
		rq.NoError(t, tx.Error)
		rq.Equal(t, "val2", found2.Val)
	}

	// Upsert using OnConflict
	{
		tx := db.Clauses(clause.OnConflict{UpdateAll: true}).Create([]Model{
			{
				PrimaryFoo: "a1",
				PrimaryBar: "b1",
				Val:        "qqq1",
			},
			{
				PrimaryFoo: "a2",
				PrimaryBar: "b2",
				Val:        "qqq2",
			},
		})
		rq.NoError(t, tx.Error)

		found1 := Model{
			PrimaryFoo: "a1",
			PrimaryBar: "b1",
		}
		tx = db.Debug().Find(&found1)
		rq.NoError(t, tx.Error)
		rq.Equal(t, "qqq1", found1.Val)

		found2 := Model{
			PrimaryFoo: "a2",
			PrimaryBar: "b2",
		}
		tx = db.Debug().Find(&found2)
		rq.NoError(t, tx.Error)
		rq.Equal(t, "qqq2", found2.Val)
	}
}
