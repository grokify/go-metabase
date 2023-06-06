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

// checks if the DatasetQueryNative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryNative{}

// DatasetQueryNative struct for DatasetQueryNative
type DatasetQueryNative struct {
	Query *string `json:"query,omitempty"`
}

// NewDatasetQueryNative instantiates a new DatasetQueryNative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryNative() *DatasetQueryNative {
	this := DatasetQueryNative{}
	return &this
}

// NewDatasetQueryNativeWithDefaults instantiates a new DatasetQueryNative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryNativeWithDefaults() *DatasetQueryNative {
	this := DatasetQueryNative{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DatasetQueryNative) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryNative) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DatasetQueryNative) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DatasetQueryNative) SetQuery(v string) {
	o.Query = &v
}

func (o DatasetQueryNative) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryNative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableDatasetQueryNative struct {
	value *DatasetQueryNative
	isSet bool
}

func (v NullableDatasetQueryNative) Get() *DatasetQueryNative {
	return v.value
}

func (v *NullableDatasetQueryNative) Set(val *DatasetQueryNative) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryNative) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryNative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryNative(val *DatasetQueryNative) *NullableDatasetQueryNative {
	return &NullableDatasetQueryNative{value: val, isSet: true}
}

func (v NullableDatasetQueryNative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryNative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
