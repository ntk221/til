package usecase_test

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"time"

	"example.com/m/models"
	"example.com/m/usecase"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/require"
)

func TestActorUsecase_GetActorInfo(t *testing.T) {
	db, err := sql.Open("txdb", "sakila")
	require.NoError(t, err)

	a := &models.Actor{
		ActorID: 1,
	}

	actorUsecase, err := usecase.NewActorUsecase(a, db)
	require.NoError(t, err)

	films, err := actorUsecase.GetAllFilmsForActor(context.Background())
	require.NoError(t, err)
	require.True(t, len(films) > 0)
}

func TestMain(m *testing.M) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		DBName:    "sakila",
		User:      "root",
		Passwd:    "root",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	prepare(c)

	txdb.Register("txdb", "mysql", c.FormatDSN())

	code := m.Run()
	os.Exit(code)
}

func prepare(c mysql.Config) {
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 現在の作業ディレクトリの絶対パスを取得
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// testdata/fixtures への正確なパスを構築
	fixturesPath := filepath.Join(currentDir, "..", "testdata", "fixtures")

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(fixturesPath),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
