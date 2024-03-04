package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"example.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrDbNil         = errors.New("db is nil")
	ErrActorNotFound = errors.New("actor not found")
)

type actorUsecase struct {
	db *sql.DB
}

func NewActorUsecase(db *sql.DB) (*actorUsecase, error) {
	if db == nil {
		log.Print("db is nil")
		return nil, ErrDbNil
	}

	return &actorUsecase{
		db,
	}, nil
}

func (a *actorUsecase) GetAllFilmsForActor(ctx context.Context, actorID int) ([]*models.Film, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Actorに紐づくFilmActorsを読み込み、その中のFilmも読み込む
	films, err := models.Films(
		qm.Load(models.FilmRels.FilmActors, models.FilmActorWhere.ActorID.EQ(actorID)),
		qm.Load(models.ActorRels.FilmActors+"."+models.FilmActorRels.Film),
	).All(ctx, tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return films, nil
}
