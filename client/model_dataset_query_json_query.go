/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

type DatasetQueryJsonQuery struct {
	Database int32           `json:"database,omitempty"`
	Type     string          `json:"type,omitempty"`
	Query    DatasetQueryDsl `json:"query,omitempty"`
}