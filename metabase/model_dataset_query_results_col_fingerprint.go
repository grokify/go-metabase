/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

type DatasetQueryResultsColFingerprint struct {
	Global DatasetQueryResultsColFingerprintGlobal `json:"global,omitempty"`
	// map[string]DatasetQueryResultsColFingerprintType results in map[string]interface{}
	Type map[string]interface{} `json:"type,omitempty"`
}