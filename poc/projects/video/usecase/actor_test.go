package usecase_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	t.Parallel()

	db, err := sql.Open("txdb", "sakila")
	if err != nil {
		panic(err)
	}

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

	txdb.Register("txdb", "mysql", c.FormatDSN())

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
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

	// テスト関数毎にtxdbでトランザクションを貼るので、Loadは一回でいいはず
	if err := fixtures.Load(); err != nil {
		panic(err)
	}

	code := m.Run()
	cleanup(db)
	os.Exit(code)
}

func cleanup(db *sql.DB) {
	tables := []string{"film_actor", "actor", "film", "language"}

	for _, table := range tables {
		if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)); err != nil {
			log.Fatalf("Failed to cleanup table %s: %v", table, err)
		}
	}
}
