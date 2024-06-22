package surreal

import (
	"csserver/internal/common"
	"csserver/internal/interfaces"

	"github.com/surrealdb/surrealdb.go"
)

type DBClient struct {
	Client *surrealdb.DB
	Logger interfaces.Logger
}

func NewDBClient(client surrealdb.DB, logger interfaces.Logger) *DBClient {
	return &DBClient{
		Client: &client,
		Logger: logger,
	}
}

// GetObjectById returns an object from the database
func (db *DBClient) GetObjectById(id string) (interface{}, error) {
	objectData, err := db.Client.Select(id)
	if err != nil {
		return common.HandleReturnWithValue(db.Logger, &objectData, err)
	}

	err = surrealdb.Unmarshal(objectData, &objectData)
	if err != nil {
		return common.HandleReturnWithValue(db.Logger, &objectData, err)
	}

	return common.HandleReturnWithValue(db.Logger, &objectData, err)
}

// GetObjectById returns an object from the database
func (db *DBClient) GetObject(
	sql string,
	key string,
	criteria map[string]interface{}) (interface{}, error) {

	//outlist := make([]interface{}, 1)

	objectData, err := db.Client.Query(sql, criteria)
	return common.HandleReturnWithValue(db.Logger, &objectData, err)
}

// DeleteObject delete an object from the DB based on the ID
func (db *DBClient) DeleteObject(
	userID string,
	id string) error {

	db.Logger.Infof("Deleting database object: %v", id)

	if _, err := db.Client.Delete(id); err != nil {
		db.Logger.Error(err)
		return err
	}

	return nil
}

// SoftDeleteObject mark an obect as deleted but leave its record in the database
func (db *DBClient) SoftDeleteObject(
	userID string,
	input interfaces.DBObject) error {

	db.Logger.Infof("Soft deleting database object: %v", input)
	input.SetDeleteInfo(userID)

	_, err := db.Client.Update(input.GetID(), input)
	return common.HandleReturn(db.Logger, err)
}

// CreateObject create an object of the passed in type
func (db *DBClient) CreateObject(
	userID string,
	objectName string,
	input interfaces.DBObject) (interface{}, error) {

	input.SetCreateInfo(userID)
	data, err := db.Client.Create(objectName, input)
	if err != nil {
		return common.HandleReturnWithValue(db.Logger, &data, err)
	}

	return data, nil
}

// CreateObject create an object of the passed in type
func (db *DBClient) UpdateObject(
	userID string,
	objectID string,
	input interfaces.DBObject) (interface{}, error) {

	input.SetUpdateInfo(userID)
	data, err := db.Client.Update(objectID, input)
	if err != nil {
		return common.HandleReturnWithValue(db.Logger, &data, err)
	}

	return data, nil
}

// CreateOrUpdateObject if the passed in object contains an ID, it will update.  otherwise it will create.
func (db *DBClient) CreateOrUpdateObject(
	userID string,
	objectID string,
	input interfaces.DBObject) (interface{}, error) {

	if input.HasID() {
		return db.UpdateObject(userID, objectID, input)
	} else {
		return db.CreateObject(userID, objectID, input)
	}
}

// FindPagedObjects return paged objects from the database
func (db *DBClient) FindPagedObjects(sql string, paging common.Pagination, filters map[string]interface{}) (interface{}, int, error) {
	pageSql := getPageSql(sql)

	filters["start"] = paging.GetOffset()
	filters["limit"] = *paging.ResultsPerPage

	resultsData, err := db.Client.Query(pageSql, filters)
	if err != nil {
		db.Logger.Error(err)
		return nil, -1, err
	}

	count, err := db.GetCount(sql, filters)
	if err != nil {
		db.Logger.Error(err)
		return nil, -1, err
	}

	return resultsData, *count, nil
}

// GetCount returns an int representing the number of records in a database
func (db *DBClient) GetCount(sql string, filters map[string]interface{}) (*int, error) {
	countSql := getCountSql(sql)
	countData, err := db.Client.Query(countSql, filters)
	if err != nil {
		return common.HandleReturnWithValue[int](db.Logger, nil, err)
	}

	count, err := parseCountFromSurrealResult(db.Logger, countData)
	return common.HandleReturnWithValue[int](db.Logger, count, err)
}

// GetScalar returns a single value from the database
// func  (db *DBClient) GetScalar(sql string, keyName string, filters map[string]interface{}) (interface{}, error) {
// 	raw, err := db.Client.Query(sql, filters)
// 	if err != nil {
// 		return common.HandleReturnWithValue[int](db.Logger, nil, err)
// 	}

// 	var results []map[string]interface{}

// 	surrealdb.UnmarshalRaw(raw, &results)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	if len(results) == 0 {
// 		return "", nil
// 	}

// 	rm := results[0]
// 	result := rm[keyName]

// 	return result, nil
// }

// Execute runs a statement against the database
func (db *DBClient) Execute(sql string, vars map[string]interface{}) error {
	_, err := db.Client.Query(sql, vars)
	return common.HandleReturn(db.Logger, err)
}
