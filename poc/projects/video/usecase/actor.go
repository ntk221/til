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

type ActorUsecase struct {
	*models.Actor

	db *sql.DB
}

func NewActorUsecase(a *models.Actor, db *sql.DB) (*ActorUsecase, error) {
	if db == nil {
		log.Print("db is nil")
		return nil, ErrDbNil
	}
	if a == nil {
		log.Print("actor is nil")
		return nil, ErrActorNotFound
	}

	return &ActorUsecase{
		a,
		db,
	}, nil
}

func (a *ActorUsecase) GetAllFilmsForActor(ctx context.Context) ([]*models.Film, error) {
	// Actorに紐づくFilmActorsを読み込み、その中のFilmも読み込む
	films, err := models.Films(
		qm.Load(models.FilmRels.FilmActors, models.FilmActorWhere.ActorID.EQ(a.ActorID)),
		qm.Load(models.ActorRels.FilmActors+"."+models.FilmActorRels.Film),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	return films, nil
}
