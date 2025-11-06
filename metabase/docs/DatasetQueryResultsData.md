# DatasetQueryResultsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**Rows** | Pointer to **[][]interface{}** |  | [optional] 
**NativeForm** | Pointer to [**DatasetQueryResultsNativeForm**](DatasetQueryResultsNativeForm.md) |  | [optional] 
**Cols** | Pointer to [**[]DatasetQueryResultsCol**](DatasetQueryResultsCol.md) |  | [optional] 
**ResultsMetadata** | Pointer to [**DatasetQueryResultsMetadata**](DatasetQueryResultsMetadata.md) |  | [optional] 
**RowsTruncated** | Pointer to **int64** |  | [optional] 

## Methods

### NewDatasetQueryResultsData

`func NewDatasetQueryResultsData() *DatasetQueryResultsData`

NewDatasetQueryResultsData instantiates a new DatasetQueryResultsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryResultsDataWithDefaults

`func NewDatasetQueryResultsDataWithDefaults() *DatasetQueryResultsData`

NewDatasetQueryResultsDataWithDefaults instantiates a new DatasetQueryResultsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *DatasetQueryResultsData) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DatasetQueryResultsData) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DatasetQueryResultsData) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DatasetQueryResultsData) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *DatasetQueryResultsData) GetRows() [][]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatasetQueryResultsData) GetRowsOk() (*[][]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatasetQueryResultsData) SetRows(v [][]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatasetQueryResultsData) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetNativeForm

`func (o *DatasetQueryResultsData) GetNativeForm() DatasetQueryResultsNativeForm`

GetNativeForm returns the NativeForm field if non-nil, zero value otherwise.

### GetNativeFormOk

`func (o *DatasetQueryResultsData) GetNativeFormOk() (*DatasetQueryResultsNativeForm, bool)`

GetNativeFormOk returns a tuple with the NativeForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeForm

`func (o *DatasetQueryResultsData) SetNativeForm(v DatasetQueryResultsNativeForm)`

SetNativeForm sets NativeForm field to given value.

### HasNativeForm

`func (o *DatasetQueryResultsData) HasNativeForm() bool`

HasNativeForm returns a boolean if a field has been set.

### GetCols

`func (o *DatasetQueryResultsData) GetCols() []DatasetQueryResultsCol`

GetCols returns the Cols field if non-nil, zero value otherwise.

### GetColsOk

`func (o *DatasetQueryResultsData) GetColsOk() (*[]DatasetQueryResultsCol, bool)`

GetColsOk returns a tuple with the Cols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCols

`func (o *DatasetQueryResultsData) SetCols(v []DatasetQueryResultsCol)`

SetCols sets Cols field to given value.

### HasCols

`func (o *DatasetQueryResultsData) HasCols() bool`

HasCols returns a boolean if a field has been set.

### GetResultsMetadata

`func (o *DatasetQueryResultsData) GetResultsMetadata() DatasetQueryResultsMetadata`

GetResultsMetadata returns the ResultsMetadata field if non-nil, zero value otherwise.

### GetResultsMetadataOk

`func (o *DatasetQueryResultsData) GetResultsMetadataOk() (*DatasetQueryResultsMetadata, bool)`

GetResultsMetadataOk returns a tuple with the ResultsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsMetadata

`func (o *DatasetQueryResultsData) SetResultsMetadata(v DatasetQueryResultsMetadata)`

SetResultsMetadata sets ResultsMetadata field to given value.

### HasResultsMetadata

`func (o *DatasetQueryResultsData) HasResultsMetadata() bool`

HasResultsMetadata returns a boolean if a field has been set.

### GetRowsTruncated

`func (o *DatasetQueryResultsData) GetRowsTruncated() int64`

GetRowsTruncated returns the RowsTruncated field if non-nil, zero value otherwise.

### GetRowsTruncatedOk

`func (o *DatasetQueryResultsData) GetRowsTruncatedOk() (*int64, bool)`

GetRowsTruncatedOk returns a tuple with the RowsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsTruncated

`func (o *DatasetQueryResultsData) SetRowsTruncated(v int64)`

SetRowsTruncated sets RowsTruncated field to given value.

### HasRowsTruncated

`func (o *DatasetQueryResultsData) HasRowsTruncated() bool`

HasRowsTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


