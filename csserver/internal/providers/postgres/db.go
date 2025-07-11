package postgres

import (
	"context"
	"fmt"
	"time"

	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TableName string

// GetDB returns a connection pool from a connection string
func GetDB(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	//urlExample := "postgres://username:password@localhost:5432/database_name"
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(ctx, config)
}

// GetDBFromConfig returns a connection pool from a database config object
func GetDBFromConfig(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)

	return GetDB(ctx, url)
}

// UpdateObject create or update an object in the given table
func UpdateObject[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	obj T,
	table TableName,
	id string) (*common.BaseModel[T], error) {

	return UpdateObjectWithParent(ctx, db, obj, table, id, nil)
}

// UpdateObjectWithParent create or update an object with a parent context
func UpdateObjectWithParent[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	obj T,
	table TableName,
	id string,
	parentID *string) (*common.BaseModel[T], error) {

	usr := config.GetUserEmailFromContext(ctx)

	if len(id) == 0 {
		id = utils.GenerateBase64UUID()
	}

	dbObj := common.BaseModel[T]{
		ID:        id,
		ParentID:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: usr,
		UpdatedBy: usr,
		Data:      obj,
	}

	sql := fmt.Sprintf(`
		INSERT INTO %s 
			(id, parent_id, created_at, created_by, updated_at, updated_by, data) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		ON CONFLICT(id)
		DO UPDATE SET
			updated_at = EXCLUDED.updated_at,
			updated_by = EXCLUDED.updated_by,
			data = EXCLUDED.data
		RETURNING id;
		`, table)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(ctx, sql,
		dbObj.ID,
		dbObj.ParentID,
		dbObj.CreatedAt,
		dbObj.CreatedBy,
		dbObj.UpdatedAt,
		dbObj.UpdatedBy,
		dbObj.Data)

	if err != nil {
		tx.Rollback(ctx)
	} else {
		tx.Commit(ctx)
	}

	return GetObjectByID[T](ctx, db, table, id)

}

// GetObjectByID return a single object from the database based on its ID
func GetObjectByID[T any](ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	id string) (*common.BaseModel[T], error) {

	sql := fmt.Sprintf(`
		SELECT id, parent_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, data
		FROM %s
		WHERE true 
			AND deleted_at is null
			AND id = $1;
		`, table)

	return GetObject[T](ctx, db, table, sql, id)
}

// GetObject return a single object from the database
func GetObject[T any](ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	sql string,
	params ...any) (*common.BaseModel[T], error) {

	rows, err := db.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output, err := pgx.CollectRows(rows, pgx.RowToStructByName[common.BaseModel[T]])
	if err != nil {
		return nil, err
	}

	if len(output) > 0 {
		return &output[0], nil
	}

	return nil, nil
}

// FindAllObjects return all objects in a given database table that are not soft-deleted
func FindAllObjects[T any](ctx context.Context,
	db *pgxpool.Pool,
	table TableName) (common.PagedResults[common.BaseModel[T]], error) {

	pf := common.NewPagedResultsForAllRecords[T]()

	sql := fmt.Sprintf(`
		SELECT id, parent_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, data
		FROM %s
		WHERE true 
			AND deleted_at is null
		`, table)

	return FindPagedObjects[T](ctx, db, sql, pf.Pagination, pf.Filters, []any{})
}

// FindAllChildren return all children of a parent object
func FindAllChildren[T any](ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	parentID string) (common.PagedResults[common.BaseModel[T]], error) {

	pf := common.NewPagedResultsForAllRecords[T]()

	pf.Filters.AddFilter(common.Filter{
		Key:       "parent_id",
		Value:     parentID,
		Operation: common.FilterOperationEqual,
	})

	sql := fmt.Sprintf(`
		SELECT id, parent_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, data
		FROM %s
		WHERE true 
			AND deleted_at is null
			and parent_id = $1
		`, table)

	return FindPagedObjects[T](ctx, db, sql, pf.Pagination, pf.Filters, pf.Filters.GetFiltersOrderedValues())
}

// SoftDelete soft-delete an object from the DB
func SoftDelete(
	ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	id string) error {

	usr := config.GetUserEmailFromContext(ctx)

	sql := fmt.Sprintf(`
			UPDATE %s 
			SET deleted_at = $1,
				deleted_by = $2
			WHERE true
				AND id = $3;
			`, table)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql,
		time.Now(),
		usr,
		id)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	return nil
}

// UndoSoftDelete remove the delete attributes of a record
func UndoSoftDelete(
	ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	id string) error {

	usr := config.GetUserEmailFromContext(ctx)

	sql := fmt.Sprintf(`
			UPDATE %s 
			SET updated_at = $1,
				updated_by = $2,
				deleted_at = null,
				deleted_by = null
			WHERE true
				AND id = $3;
			`, table)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql,
		time.Now(),
		usr,
		id)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	return nil
}

// Delete hard-delete an object from the database
func Delete(
	ctx context.Context,
	db *pgxpool.Pool,
	table TableName,
	id string) error {

	sql := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE true
				AND id = $1;
			`, table)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql, id)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	return nil
}

// Exec execute a sql statement in the DB
func ExecTx(
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	params ...any) error {

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sql, params...)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	return nil
}

// Exec execute a sql statement in the DB
func Exec(
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	params ...any) error {

	_, err := db.Exec(ctx, sql, params...)

	return err
}

// FindPagedObjects find objects based on a given sql statement
func FindPagedObjects[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	paging common.Pagination,
	filters common.Filters,
	params []any) (common.PagedResults[common.BaseModel[T]], error) {

	pageSql := GetPageSql(sql, len(params)+1)
	countSql := GetCountSql(sql)
	out := common.NewPagedResults[common.BaseModel[T]](paging, filters)

	count, err := GetScalar[int](ctx, db, countSql, params...)
	if err != nil {
		return out, err
	}

	params = append(params, *paging.ResultsPerPage)
	params = append(params, paging.GetOffset())

	rows, err := db.Query(ctx, pageSql, params...)
	if err != nil {
		return out, err
	}
	defer rows.Close()

	output, err := pgx.CollectRows(rows, pgx.RowToStructByName[common.BaseModel[T]])
	if err != nil {
		return out, err
	}

	out.Results = output
	out.Pagination.TotalResults = count

	return out, nil
}

// GetScalar return a single bit of data from the database
func GetScalar[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	params ...any) (*T, error) {

	var output T
	db.QueryRow(ctx, sql, params...).Scan(&output)

	return &output, nil
}
