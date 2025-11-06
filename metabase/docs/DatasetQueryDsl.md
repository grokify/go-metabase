# DatasetQueryDsl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceTable** | Pointer to **int64** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to [**DatasetQueryDslPage**](DatasetQueryDslPage.md) |  | [optional] 

## Methods

### NewDatasetQueryDsl

`func NewDatasetQueryDsl() *DatasetQueryDsl`

NewDatasetQueryDsl instantiates a new DatasetQueryDsl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryDslWithDefaults

`func NewDatasetQueryDslWithDefaults() *DatasetQueryDsl`

NewDatasetQueryDslWithDefaults instantiates a new DatasetQueryDsl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceTable

`func (o *DatasetQueryDsl) GetSourceTable() int64`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *DatasetQueryDsl) GetSourceTableOk() (*int64, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *DatasetQueryDsl) SetSourceTable(v int64)`

SetSourceTable sets SourceTable field to given value.

### HasSourceTable

`func (o *DatasetQueryDsl) HasSourceTable() bool`

HasSourceTable returns a boolean if a field has been set.

### GetLimit

`func (o *DatasetQueryDsl) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DatasetQueryDsl) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DatasetQueryDsl) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DatasetQueryDsl) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *DatasetQueryDsl) GetPage() DatasetQueryDslPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DatasetQueryDsl) GetPageOk() (*DatasetQueryDslPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DatasetQueryDsl) SetPage(v DatasetQueryDslPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *DatasetQueryDsl) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


