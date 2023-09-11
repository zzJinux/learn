package transaction

import (
	"log"
	"testing"

	"learn/database"
	"learn/database/gorm"
	. "learn/database/u"

	"github.com/cockroachdb/errors"
	rq "github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCallbackTransaction(t *testing.T) {
	db := learngorm.Connect(database.EnvPort)

	t.Run("success", func(t *testing.T) {
		prepare(db)
		db = db.Debug()

		err := db.Transaction(databaseWork)
		rq.NoError(t, err)

		a := Alpha{ID: "idid"}
		rq.NoError(t, db.Find(&a).Error)
		rq.Equal(t, "walwal", a.Val)

		b := Beta{UID: "zzzz"}
		rq.NoError(t, db.Find(&b).Error)
		rq.Equal(t, "kokoko", b.Data)
	})

	t.Run("fail (rollback)", func(t *testing.T) {
		prepare(db)
		db = db.Debug()

		err := db.Transaction(databaseWorkWithError)
		rq.Error(t, err)

		a := Alpha{ID: "idid"}
		rq.NoError(t, db.Find(&a).Error)
		rq.Equal(t, "valval", a.Val)

		b := Beta{UID: "zzzz"}
		rq.NoError(t, db.Find(&b).Error)
		rq.Equal(t, "datata", b.Data)
	})
}

func TestManualTransaction(t *testing.T) {
	db := learngorm.Connect(database.EnvPort)

	t.Run("success", func(t *testing.T) {
		prepare(db)
		db = db.Debug()

		err := manualCommitRollback(db, databaseWork)
		rq.NoError(t, err)

		a := Alpha{ID: "idid"}
		rq.NoError(t, db.Find(&a).Error)
		rq.Equal(t, "walwal", a.Val)

		b := Beta{UID: "zzzz"}
		rq.NoError(t, db.Find(&b).Error)
		rq.Equal(t, "kokoko", b.Data)
	})

	t.Run("fail (rollback)", func(t *testing.T) {
		prepare(db)
		db = db.Debug()

		err := manualCommitRollback(db, databaseWorkWithError)
		rq.Error(t, err)

		a := Alpha{ID: "idid"}
		rq.NoError(t, db.Find(&a).Error)
		rq.Equal(t, "valval", a.Val)

		b := Beta{UID: "zzzz"}
		rq.NoError(t, db.Find(&b).Error)
		rq.Equal(t, "datata", b.Data)
	})
}

func prepare(db *gorm.DB) {
	learngorm.ResetTable(db, Alpha{})
	learngorm.ResetTable(db, Beta{})

	Must(db.Create(Alpha{
		ID:  "idid",
		Val: "valval",
	}).Error)
	Must(db.Create(Beta{
		UID:  "zzzz",
		Data: "datata",
	}).Error)
}

func databaseWork(db *gorm.DB) error {
	db.Save(Alpha{
		ID:  "idid",
		Val: "walwal",
	})
	db.Save(Beta{
		UID:  "zzzz",
		Data: "kokoko",
	})
	return nil
}

func databaseWorkWithError(db *gorm.DB) error {
	db.Save(Alpha{
		ID:  "idid",
		Val: "walwal",
	})
	db.Save(Beta{
		UID:  "zzzz",
		Data: "kokoko",
	})
	return errors.New("boo")
}

func manualCommitRollback(db *gorm.DB, fn func(*gorm.DB) error) (err error) {
	panicked := true

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if panicked {
			rollbackErr := tx.Rollback().Error
			// You have to handle the rollback error here.
			// Because manualCommitRollback hasn't returned yet, its caller can't see the error it returns.
			log.Print(rollbackErr)
		} else if err != nil {
			rollbackErr := tx.Rollback().Error
			err = errors.WithSecondaryError(err, rollbackErr)
		}
	}()

	if err = fn(tx); err == nil {
		panicked = false
		return tx.Commit().Error
	}

	panicked = false
	return
}
