/*
Metabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metabase

import (
	"encoding/json"
)

// checks if the DatasetQueryResultsNativeForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryResultsNativeForm{}

// DatasetQueryResultsNativeForm struct for DatasetQueryResultsNativeForm
type DatasetQueryResultsNativeForm struct {
	Query *string `json:"query,omitempty"`
	// unknown type
	Params *string `json:"params,omitempty"`
}

// NewDatasetQueryResultsNativeForm instantiates a new DatasetQueryResultsNativeForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryResultsNativeForm() *DatasetQueryResultsNativeForm {
	this := DatasetQueryResultsNativeForm{}
	return &this
}

// NewDatasetQueryResultsNativeFormWithDefaults instantiates a new DatasetQueryResultsNativeForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryResultsNativeFormWithDefaults() *DatasetQueryResultsNativeForm {
	this := DatasetQueryResultsNativeForm{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DatasetQueryResultsNativeForm) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryResultsNativeForm) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DatasetQueryResultsNativeForm) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DatasetQueryResultsNativeForm) SetQuery(v string) {
	o.Query = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DatasetQueryResultsNativeForm) GetParams() string {
	if o == nil || IsNil(o.Params) {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryResultsNativeForm) GetParamsOk() (*string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *DatasetQueryResultsNativeForm) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *DatasetQueryResultsNativeForm) SetParams(v string) {
	o.Params = &v
}

func (o DatasetQueryResultsNativeForm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryResultsNativeForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableDatasetQueryResultsNativeForm struct {
	value *DatasetQueryResultsNativeForm
	isSet bool
}

func (v NullableDatasetQueryResultsNativeForm) Get() *DatasetQueryResultsNativeForm {
	return v.value
}

func (v *NullableDatasetQueryResultsNativeForm) Set(val *DatasetQueryResultsNativeForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryResultsNativeForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryResultsNativeForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryResultsNativeForm(val *DatasetQueryResultsNativeForm) *NullableDatasetQueryResultsNativeForm {
	return &NullableDatasetQueryResultsNativeForm{value: val, isSet: true}
}

func (v NullableDatasetQueryResultsNativeForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryResultsNativeForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
