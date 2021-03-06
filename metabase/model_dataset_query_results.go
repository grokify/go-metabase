/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

import (
	"time"
)

type DatasetQueryResults struct {
	StartedAt time.Time             `json:"started_at,omitempty"`
	JsonQuery DatasetQueryJsonQuery `json:"json_query,omitempty"`
	// type unknown
	AverageExecutionTime string                  `json:"average_execution_time,omitempty"`
	Status               string                  `json:"status,omitempty"`
	Context              string                  `json:"context,omitempty"`
	RowCount             int64                   `json:"row_count,omitempty"`
	RunningTime          int64                   `json:"running_time,omitempty"`
	Data                 DatasetQueryResultsData `json:"data,omitempty"`
}
