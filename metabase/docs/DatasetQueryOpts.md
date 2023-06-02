# DatasetQueryOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **interface{}** | unknown type | [optional] 

## Methods

### NewDatasetQueryOpts

`func NewDatasetQueryOpts() *DatasetQueryOpts`

NewDatasetQueryOpts instantiates a new DatasetQueryOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryOptsWithDefaults

`func NewDatasetQueryOptsWithDefaults() *DatasetQueryOpts`

NewDatasetQueryOptsWithDefaults instantiates a new DatasetQueryOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasetQueryOpts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetQueryOpts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetQueryOpts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetQueryOpts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *DatasetQueryOpts) GetEntityType() interface{}`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DatasetQueryOpts) GetEntityTypeOk() (*interface{}, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DatasetQueryOpts) SetEntityType(v interface{})`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DatasetQueryOpts) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *DatasetQueryOpts) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *DatasetQueryOpts) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


