// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// FilmList is an object representing the database table.
type FilmList struct {
	FID         int                `boil:"FID" json:"FID" toml:"FID" yaml:"FID"`
	Title       string             `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description null.String        `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Category    null.String        `boil:"category" json:"category,omitempty" toml:"category" yaml:"category,omitempty"`
	Price       types.Decimal      `boil:"price" json:"price" toml:"price" yaml:"price"`
	Length      null.Int           `boil:"length" json:"length,omitempty" toml:"length" yaml:"length,omitempty"`
	Rating      FilmListNullRating `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`
	Actors      null.String        `boil:"actors" json:"actors,omitempty" toml:"actors" yaml:"actors,omitempty"`
}

var FilmListColumns = struct {
	FID         string
	Title       string
	Description string
	Category    string
	Price       string
	Length      string
	Rating      string
	Actors      string
}{
	FID:         "FID",
	Title:       "title",
	Description: "description",
	Category:    "category",
	Price:       "price",
	Length:      "length",
	Rating:      "rating",
	Actors:      "actors",
}

var FilmListTableColumns = struct {
	FID         string
	Title       string
	Description string
	Category    string
	Price       string
	Length      string
	Rating      string
	Actors      string
}{
	FID:         "film_list.FID",
	Title:       "film_list.title",
	Description: "film_list.description",
	Category:    "film_list.category",
	Price:       "film_list.price",
	Length:      "film_list.length",
	Rating:      "film_list.rating",
	Actors:      "film_list.actors",
}

// Generated where

type whereHelperFilmListNullRating struct{ field string }

func (w whereHelperFilmListNullRating) EQ(x FilmListNullRating) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelperFilmListNullRating) NEQ(x FilmListNullRating) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelperFilmListNullRating) LT(x FilmListNullRating) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperFilmListNullRating) LTE(x FilmListNullRating) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperFilmListNullRating) GT(x FilmListNullRating) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperFilmListNullRating) GTE(x FilmListNullRating) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperFilmListNullRating) IN(slice []FilmListNullRating) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperFilmListNullRating) NIN(slice []FilmListNullRating) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelperFilmListNullRating) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperFilmListNullRating) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

var FilmListWhere = struct {
	FID         whereHelperint
	Title       whereHelperstring
	Description whereHelpernull_String
	Category    whereHelpernull_String
	Price       whereHelpertypes_Decimal
	Length      whereHelpernull_Int
	Rating      whereHelperFilmListNullRating
	Actors      whereHelpernull_String
}{
	FID:         whereHelperint{field: "`film_list`.`FID`"},
	Title:       whereHelperstring{field: "`film_list`.`title`"},
	Description: whereHelpernull_String{field: "`film_list`.`description`"},
	Category:    whereHelpernull_String{field: "`film_list`.`category`"},
	Price:       whereHelpertypes_Decimal{field: "`film_list`.`price`"},
	Length:      whereHelpernull_Int{field: "`film_list`.`length`"},
	Rating:      whereHelperFilmListNullRating{field: "`film_list`.`rating`"},
	Actors:      whereHelpernull_String{field: "`film_list`.`actors`"},
}

var (
	filmListAllColumns            = []string{"FID", "title", "description", "category", "price", "length", "rating", "actors"}
	filmListColumnsWithoutDefault = []string{"title", "description", "category", "length", "actors"}
	filmListColumnsWithDefault    = []string{"FID", "price", "rating"}
	filmListPrimaryKeyColumns     = []string{}
	filmListGeneratedColumns      = []string{}
)

type (
	// FilmListSlice is an alias for a slice of pointers to FilmList.
	// This should almost always be used instead of []FilmList.
	FilmListSlice []*FilmList
	// FilmListHook is the signature for custom FilmList hook methods
	FilmListHook func(context.Context, boil.ContextExecutor, *FilmList) error

	filmListQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	filmListType           = reflect.TypeOf(&FilmList{})
	filmListMapping        = queries.MakeStructMapping(filmListType)
	filmListInsertCacheMut sync.RWMutex
	filmListInsertCache    = make(map[string]insertCache)
	filmListUpdateCacheMut sync.RWMutex
	filmListUpdateCache    = make(map[string]updateCache)
	filmListUpsertCacheMut sync.RWMutex
	filmListUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
	// These are used in some views
	_ = fmt.Sprintln("")
	_ = reflect.Int
	_ = strings.Builder{}
	_ = sync.Mutex{}
	_ = strmangle.Plural("")
	_ = strconv.IntSize
)

var filmListAfterSelectMu sync.Mutex
var filmListAfterSelectHooks []FilmListHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FilmList) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filmListAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFilmListHook registers your hook function for all future operations.
func AddFilmListHook(hookPoint boil.HookPoint, filmListHook FilmListHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		filmListAfterSelectMu.Lock()
		filmListAfterSelectHooks = append(filmListAfterSelectHooks, filmListHook)
		filmListAfterSelectMu.Unlock()
	}
}

// One returns a single filmList record from the query.
func (q filmListQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FilmList, error) {
	o := &FilmList{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for film_list")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FilmList records from the query.
func (q filmListQuery) All(ctx context.Context, exec boil.ContextExecutor) (FilmListSlice, error) {
	var o []*FilmList

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FilmList slice")
	}

	if len(filmListAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FilmList records in the query.
func (q filmListQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count film_list rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q filmListQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if film_list exists")
	}

	return count > 0, nil
}

// FilmLists retrieves all the records using an executor.
func FilmLists(mods ...qm.QueryMod) filmListQuery {
	mods = append(mods, qm.From("`film_list`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`film_list`.*"})
	}

	return filmListQuery{q}
}
