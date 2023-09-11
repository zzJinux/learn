package learnsqlmock

import (
	"log/slog"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	rq "github.com/stretchr/testify/require"
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	AddSource: true,
}))

func TestXXX(t *testing.T) {
	db, mock, err := sqlmock.New()
	rq.NoError(t, err)
	// Don't call it
	// defer Must(db.Close())

	_, err = db.Query("SELECT id FROM aaa")
	// expect_queue = [|]
	// -> all expectations are fulfilled
	rq.Error(t, err)
	logger.Info(err.Error())

	// statement A
	mock.ExpectQuery("SELECT id FROM aaa").WillReturnRows(sqlmock.NewRows([]string{}))
	// expect_queue = [|A]

	_, err = db.Query("SELECT id FROM aaa")
	// expect_queue = [A|] *done*
	// -> Match. The queue advances to next expectation. It turns out that all exps are now fulfilled.
	rq.NoError(t, err)
	logger.Info("<nil>")

	_, err = db.Query("SELECT id FROM aaa")
	// expect_queue = [A|] *done*
	// -> You are executing a query on already fulfilled exps.
	rq.Error(t, err)
	logger.Info(err.Error())

	_, err = db.Query("SELECT name FROM aaa")
	// expect_queue = [A|] *done*
	// -> You are executing a query on already fulfilled exps.
	rq.Error(t, err)
	logger.Info(err.Error())

	err = mock.ExpectationsWereMet()
	// expect_queue = [A|] *done*
	// -> Already fulfilled
	rq.NoError(t, err)

	// expect_queue remains the same
	// expect_queue = [A]

	// B
	mock.ExpectQuery("SELECT id FROM aaa").WillReturnRows(sqlmock.NewRows([]string{}))
	// C
	mock.ExpectQuery("SELECT name FROM aaa").WillReturnRows(sqlmock.NewRows([]string{}))
	// expect_queue = [A| B C]

	_, err = db.Query("SELECT id FROM aaa")
	// expect_queue = [A B| C]
	// -> Match. The queue advances to next expectation.
	rq.NoError(t, err)

	_, err = db.Query("SELECT id FROM aaa")
	// expect_queue = [A B| C]
	// -> Mismatch.
	rq.Error(t, err)
	logger.Info(err.Error())

	_, err = db.Query("SELECT name FROM aaa")
	// expect_queue = [A B C|] *done* (???)
	// -> Match. The queue advances to next expectation. It turns out that all exps are now fulfilled.
	rq.NoError(t, err)

	err = mock.ExpectationsWereMet()
	rq.NoError(t, err)
}
