/*
 * Metabase
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metabase

type DatasetQueryResultsColFingerprintType struct {
	PercentJson   float64 `json:"percent-json,omitempty"`
	PercentUrl    float64 `json:"percent-url,omitempty"`
	PercentEmail  float64 `json:"percent-email,omitempty"`
	AverageLength float64 `json:"average-length,omitempty"`
}
