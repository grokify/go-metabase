# DatasetQueryResultsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]DatasetQueryResultsMetadataColumn**](DatasetQueryResultsMetadataColumn.md) |  | [optional] 

## Methods

### NewDatasetQueryResultsMetadata

`func NewDatasetQueryResultsMetadata() *DatasetQueryResultsMetadata`

NewDatasetQueryResultsMetadata instantiates a new DatasetQueryResultsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryResultsMetadataWithDefaults

`func NewDatasetQueryResultsMetadataWithDefaults() *DatasetQueryResultsMetadata`

NewDatasetQueryResultsMetadataWithDefaults instantiates a new DatasetQueryResultsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *DatasetQueryResultsMetadata) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *DatasetQueryResultsMetadata) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *DatasetQueryResultsMetadata) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *DatasetQueryResultsMetadata) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetColumns

`func (o *DatasetQueryResultsMetadata) GetColumns() []DatasetQueryResultsMetadataColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DatasetQueryResultsMetadata) GetColumnsOk() (*[]DatasetQueryResultsMetadataColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DatasetQueryResultsMetadata) SetColumns(v []DatasetQueryResultsMetadataColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DatasetQueryResultsMetadata) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


