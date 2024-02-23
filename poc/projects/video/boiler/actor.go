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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Actor is an object representing the database table.
type Actor struct {
	ActorID    uint16    `boil:"actor_id" json:"actor_id" toml:"actor_id" yaml:"actor_id"`
	FirstName  string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName   string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	LastUpdate time.Time `boil:"last_update" json:"last_update" toml:"last_update" yaml:"last_update"`

	R *actorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L actorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ActorColumns = struct {
	ActorID    string
	FirstName  string
	LastName   string
	LastUpdate string
}{
	ActorID:    "actor_id",
	FirstName:  "first_name",
	LastName:   "last_name",
	LastUpdate: "last_update",
}

var ActorTableColumns = struct {
	ActorID    string
	FirstName  string
	LastName   string
	LastUpdate string
}{
	ActorID:    "actor.actor_id",
	FirstName:  "actor.first_name",
	LastName:   "actor.last_name",
	LastUpdate: "actor.last_update",
}

// Generated where

type whereHelperuint16 struct{ field string }

func (w whereHelperuint16) EQ(x uint16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint16) NEQ(x uint16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint16) LT(x uint16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint16) LTE(x uint16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint16) GT(x uint16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint16) GTE(x uint16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint16) IN(slice []uint16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint16) NIN(slice []uint16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ActorWhere = struct {
	ActorID    whereHelperuint16
	FirstName  whereHelperstring
	LastName   whereHelperstring
	LastUpdate whereHelpertime_Time
}{
	ActorID:    whereHelperuint16{field: "`actor`.`actor_id`"},
	FirstName:  whereHelperstring{field: "`actor`.`first_name`"},
	LastName:   whereHelperstring{field: "`actor`.`last_name`"},
	LastUpdate: whereHelpertime_Time{field: "`actor`.`last_update`"},
}

// ActorRels is where relationship names are stored.
var ActorRels = struct {
	FilmActors string
}{
	FilmActors: "FilmActors",
}

// actorR is where relationships are stored.
type actorR struct {
	FilmActors FilmActorSlice `boil:"FilmActors" json:"FilmActors" toml:"FilmActors" yaml:"FilmActors"`
}

// NewStruct creates a new relationship struct
func (*actorR) NewStruct() *actorR {
	return &actorR{}
}

func (r *actorR) GetFilmActors() FilmActorSlice {
	if r == nil {
		return nil
	}
	return r.FilmActors
}

// actorL is where Load methods for each relationship are stored.
type actorL struct{}

var (
	actorAllColumns            = []string{"actor_id", "first_name", "last_name", "last_update"}
	actorColumnsWithoutDefault = []string{"first_name", "last_name"}
	actorColumnsWithDefault    = []string{"actor_id", "last_update"}
	actorPrimaryKeyColumns     = []string{"actor_id"}
	actorGeneratedColumns      = []string{}
)

type (
	// ActorSlice is an alias for a slice of pointers to Actor.
	// This should almost always be used instead of []Actor.
	ActorSlice []*Actor
	// ActorHook is the signature for custom Actor hook methods
	ActorHook func(context.Context, boil.ContextExecutor, *Actor) error

	actorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	actorType                 = reflect.TypeOf(&Actor{})
	actorMapping              = queries.MakeStructMapping(actorType)
	actorPrimaryKeyMapping, _ = queries.BindMapping(actorType, actorMapping, actorPrimaryKeyColumns)
	actorInsertCacheMut       sync.RWMutex
	actorInsertCache          = make(map[string]insertCache)
	actorUpdateCacheMut       sync.RWMutex
	actorUpdateCache          = make(map[string]updateCache)
	actorUpsertCacheMut       sync.RWMutex
	actorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var actorAfterSelectMu sync.Mutex
var actorAfterSelectHooks []ActorHook

var actorBeforeInsertMu sync.Mutex
var actorBeforeInsertHooks []ActorHook
var actorAfterInsertMu sync.Mutex
var actorAfterInsertHooks []ActorHook

var actorBeforeUpdateMu sync.Mutex
var actorBeforeUpdateHooks []ActorHook
var actorAfterUpdateMu sync.Mutex
var actorAfterUpdateHooks []ActorHook

var actorBeforeDeleteMu sync.Mutex
var actorBeforeDeleteHooks []ActorHook
var actorAfterDeleteMu sync.Mutex
var actorAfterDeleteHooks []ActorHook

var actorBeforeUpsertMu sync.Mutex
var actorBeforeUpsertHooks []ActorHook
var actorAfterUpsertMu sync.Mutex
var actorAfterUpsertHooks []ActorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Actor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Actor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Actor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Actor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Actor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Actor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Actor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Actor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Actor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range actorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddActorHook registers your hook function for all future operations.
func AddActorHook(hookPoint boil.HookPoint, actorHook ActorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		actorAfterSelectMu.Lock()
		actorAfterSelectHooks = append(actorAfterSelectHooks, actorHook)
		actorAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		actorBeforeInsertMu.Lock()
		actorBeforeInsertHooks = append(actorBeforeInsertHooks, actorHook)
		actorBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		actorAfterInsertMu.Lock()
		actorAfterInsertHooks = append(actorAfterInsertHooks, actorHook)
		actorAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		actorBeforeUpdateMu.Lock()
		actorBeforeUpdateHooks = append(actorBeforeUpdateHooks, actorHook)
		actorBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		actorAfterUpdateMu.Lock()
		actorAfterUpdateHooks = append(actorAfterUpdateHooks, actorHook)
		actorAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		actorBeforeDeleteMu.Lock()
		actorBeforeDeleteHooks = append(actorBeforeDeleteHooks, actorHook)
		actorBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		actorAfterDeleteMu.Lock()
		actorAfterDeleteHooks = append(actorAfterDeleteHooks, actorHook)
		actorAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		actorBeforeUpsertMu.Lock()
		actorBeforeUpsertHooks = append(actorBeforeUpsertHooks, actorHook)
		actorBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		actorAfterUpsertMu.Lock()
		actorAfterUpsertHooks = append(actorAfterUpsertHooks, actorHook)
		actorAfterUpsertMu.Unlock()
	}
}

// One returns a single actor record from the query.
func (q actorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Actor, error) {
	o := &Actor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for actor")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Actor records from the query.
func (q actorQuery) All(ctx context.Context, exec boil.ContextExecutor) (ActorSlice, error) {
	var o []*Actor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Actor slice")
	}

	if len(actorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Actor records in the query.
func (q actorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count actor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q actorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if actor exists")
	}

	return count > 0, nil
}

// FilmActors retrieves all the film_actor's FilmActors with an executor.
func (o *Actor) FilmActors(mods ...qm.QueryMod) filmActorQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`film_actor`.`actor_id`=?", o.ActorID),
	)

	return FilmActors(queryMods...)
}

// LoadFilmActors allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (actorL) LoadFilmActors(ctx context.Context, e boil.ContextExecutor, singular bool, maybeActor interface{}, mods queries.Applicator) error {
	var slice []*Actor
	var object *Actor

	if singular {
		var ok bool
		object, ok = maybeActor.(*Actor)
		if !ok {
			object = new(Actor)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeActor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeActor))
			}
		}
	} else {
		s, ok := maybeActor.(*[]*Actor)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeActor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeActor))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &actorR{}
		}
		args[object.ActorID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &actorR{}
			}
			args[obj.ActorID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`film_actor`),
		qm.WhereIn(`film_actor.actor_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load film_actor")
	}

	var resultSlice []*FilmActor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice film_actor")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on film_actor")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for film_actor")
	}

	if len(filmActorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FilmActors = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &filmActorR{}
			}
			foreign.R.Actor = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ActorID == foreign.ActorID {
				local.R.FilmActors = append(local.R.FilmActors, foreign)
				if foreign.R == nil {
					foreign.R = &filmActorR{}
				}
				foreign.R.Actor = local
				break
			}
		}
	}

	return nil
}

// AddFilmActors adds the given related objects to the existing relationships
// of the actor, optionally inserting them as new records.
// Appends related to o.R.FilmActors.
// Sets related.R.Actor appropriately.
func (o *Actor) AddFilmActors(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*FilmActor) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ActorID = o.ActorID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `film_actor` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"actor_id"}),
				strmangle.WhereClause("`", "`", 0, filmActorPrimaryKeyColumns),
			)
			values := []interface{}{o.ActorID, rel.ActorID, rel.FilmID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ActorID = o.ActorID
		}
	}

	if o.R == nil {
		o.R = &actorR{
			FilmActors: related,
		}
	} else {
		o.R.FilmActors = append(o.R.FilmActors, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &filmActorR{
				Actor: o,
			}
		} else {
			rel.R.Actor = o
		}
	}
	return nil
}

// Actors retrieves all the records using an executor.
func Actors(mods ...qm.QueryMod) actorQuery {
	mods = append(mods, qm.From("`actor`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`actor`.*"})
	}

	return actorQuery{q}
}

// FindActor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindActor(ctx context.Context, exec boil.ContextExecutor, actorID uint16, selectCols ...string) (*Actor, error) {
	actorObj := &Actor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `actor` where `actor_id`=?", sel,
	)

	q := queries.Raw(query, actorID)

	err := q.Bind(ctx, exec, actorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from actor")
	}

	if err = actorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return actorObj, err
	}

	return actorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Actor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no actor provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(actorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	actorInsertCacheMut.RLock()
	cache, cached := actorInsertCache[key]
	actorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			actorAllColumns,
			actorColumnsWithDefault,
			actorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(actorType, actorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(actorType, actorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `actor` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `actor` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `actor` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, actorPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into actor")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ActorID = uint16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == actorMapping["actor_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ActorID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for actor")
	}

CacheNoHooks:
	if !cached {
		actorInsertCacheMut.Lock()
		actorInsertCache[key] = cache
		actorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Actor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Actor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	actorUpdateCacheMut.RLock()
	cache, cached := actorUpdateCache[key]
	actorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			actorAllColumns,
			actorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update actor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `actor` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, actorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(actorType, actorMapping, append(wl, actorPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update actor row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for actor")
	}

	if !cached {
		actorUpdateCacheMut.Lock()
		actorUpdateCache[key] = cache
		actorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q actorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for actor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for actor")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ActorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `actor` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in actor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all actor")
	}
	return rowsAff, nil
}

var mySQLActorUniqueColumns = []string{
	"actor_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Actor) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no actor provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(actorColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLActorUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	actorUpsertCacheMut.RLock()
	cache, cached := actorUpsertCache[key]
	actorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			actorAllColumns,
			actorColumnsWithDefault,
			actorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			actorAllColumns,
			actorPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert actor, could not build update column list")
		}

		ret := strmangle.SetComplement(actorAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`actor`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `actor` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(actorType, actorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(actorType, actorMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for actor")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ActorID = uint16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == actorMapping["actor_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(actorType, actorMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for actor")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for actor")
	}

CacheNoHooks:
	if !cached {
		actorUpsertCacheMut.Lock()
		actorUpsertCache[key] = cache
		actorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Actor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Actor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Actor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), actorPrimaryKeyMapping)
	sql := "DELETE FROM `actor` WHERE `actor_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from actor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for actor")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q actorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no actorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from actor")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for actor")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ActorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(actorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `actor` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from actor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for actor")
	}

	if len(actorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Actor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindActor(ctx, exec, o.ActorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ActorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ActorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `actor`.* FROM `actor` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ActorSlice")
	}

	*o = slice

	return nil
}

// ActorExists checks if the Actor row exists.
func ActorExists(ctx context.Context, exec boil.ContextExecutor, actorID uint16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `actor` where `actor_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, actorID)
	}
	row := exec.QueryRowContext(ctx, sql, actorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if actor exists")
	}

	return exists, nil
}

// Exists checks if the Actor row exists.
func (o *Actor) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ActorExists(ctx, exec, o.ActorID)
}
