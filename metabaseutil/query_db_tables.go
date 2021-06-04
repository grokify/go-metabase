package metabaseutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/grokify/go-metabase/metabase"
)

func GetTableBySQLName(apiClient *metabase.APIClient, tableName string) (*metabase.Database, *metabase.DatabaseTable, error) {
	if apiClient == nil {
		return nil, nil, errors.New("`apiClient *metabase.APIClient` cannot be nil")
	}

	opts := metabase.ListDatabasesOpts{
		IncludeTables: optional.NewBool(true)}

	info, resp, err := apiClient.DatabaseApi.ListDatabases(
		context.Background(), &opts)
	if err != nil {
		return nil, nil, err
	} else if resp.StatusCode >= 300 {
		return nil, nil, fmt.Errorf("bad API ssatus code [%v]", resp.StatusCode)
	}

	for _, db := range info {
		for _, tb := range db.Tables {
			if tb.Name == tableName {
				return &db, &tb, nil
			}
		}
	}
	return nil, nil, fmt.Errorf("table [%s] not found", tableName)
}
