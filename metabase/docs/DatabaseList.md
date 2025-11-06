# DatabaseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Database**](Database.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewDatabaseList

`func NewDatabaseList() *DatabaseList`

NewDatabaseList instantiates a new DatabaseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseListWithDefaults

`func NewDatabaseListWithDefaults() *DatabaseList`

NewDatabaseListWithDefaults instantiates a new DatabaseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatabaseList) GetData() []Database`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatabaseList) GetDataOk() (*[]Database, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatabaseList) SetData(v []Database)`

SetData sets Data field to given value.

### HasData

`func (o *DatabaseList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *DatabaseList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DatabaseList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DatabaseList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DatabaseList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


