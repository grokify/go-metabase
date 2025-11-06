# DatasetQueryJsonQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Native** | Pointer to [**DatasetQueryNative**](DatasetQueryNative.md) |  | [optional] 
**Query** | Pointer to [**DatasetQueryDsl**](DatasetQueryDsl.md) |  | [optional] 
**Constraints** | Pointer to [**DatasetQueryConstraints**](DatasetQueryConstraints.md) |  | [optional] 

## Methods

### NewDatasetQueryJsonQuery

`func NewDatasetQueryJsonQuery() *DatasetQueryJsonQuery`

NewDatasetQueryJsonQuery instantiates a new DatasetQueryJsonQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryJsonQueryWithDefaults

`func NewDatasetQueryJsonQueryWithDefaults() *DatasetQueryJsonQuery`

NewDatasetQueryJsonQueryWithDefaults instantiates a new DatasetQueryJsonQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *DatasetQueryJsonQuery) GetDatabase() int64`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DatasetQueryJsonQuery) GetDatabaseOk() (*int64, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DatasetQueryJsonQuery) SetDatabase(v int64)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DatasetQueryJsonQuery) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetType

`func (o *DatasetQueryJsonQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetQueryJsonQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetQueryJsonQuery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatasetQueryJsonQuery) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNative

`func (o *DatasetQueryJsonQuery) GetNative() DatasetQueryNative`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *DatasetQueryJsonQuery) GetNativeOk() (*DatasetQueryNative, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *DatasetQueryJsonQuery) SetNative(v DatasetQueryNative)`

SetNative sets Native field to given value.

### HasNative

`func (o *DatasetQueryJsonQuery) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetQuery

`func (o *DatasetQueryJsonQuery) GetQuery() DatasetQueryDsl`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DatasetQueryJsonQuery) GetQueryOk() (*DatasetQueryDsl, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DatasetQueryJsonQuery) SetQuery(v DatasetQueryDsl)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DatasetQueryJsonQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetConstraints

`func (o *DatasetQueryJsonQuery) GetConstraints() DatasetQueryConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *DatasetQueryJsonQuery) GetConstraintsOk() (*DatasetQueryConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *DatasetQueryJsonQuery) SetConstraints(v DatasetQueryConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *DatasetQueryJsonQuery) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


