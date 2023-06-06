# DatasetQueryResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**JsonQuery** | Pointer to [**DatasetQueryJsonQuery**](DatasetQueryJsonQuery.md) |  | [optional] 
**AverageExecutionTime** | Pointer to **string** | type unknown | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**RowCount** | Pointer to **int64** |  | [optional] 
**RunningTime** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**DatasetQueryResultsData**](DatasetQueryResultsData.md) |  | [optional] 

## Methods

### NewDatasetQueryResults

`func NewDatasetQueryResults() *DatasetQueryResults`

NewDatasetQueryResults instantiates a new DatasetQueryResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryResultsWithDefaults

`func NewDatasetQueryResultsWithDefaults() *DatasetQueryResults`

NewDatasetQueryResultsWithDefaults instantiates a new DatasetQueryResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *DatasetQueryResults) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DatasetQueryResults) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DatasetQueryResults) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DatasetQueryResults) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetJsonQuery

`func (o *DatasetQueryResults) GetJsonQuery() DatasetQueryJsonQuery`

GetJsonQuery returns the JsonQuery field if non-nil, zero value otherwise.

### GetJsonQueryOk

`func (o *DatasetQueryResults) GetJsonQueryOk() (*DatasetQueryJsonQuery, bool)`

GetJsonQueryOk returns a tuple with the JsonQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonQuery

`func (o *DatasetQueryResults) SetJsonQuery(v DatasetQueryJsonQuery)`

SetJsonQuery sets JsonQuery field to given value.

### HasJsonQuery

`func (o *DatasetQueryResults) HasJsonQuery() bool`

HasJsonQuery returns a boolean if a field has been set.

### GetAverageExecutionTime

`func (o *DatasetQueryResults) GetAverageExecutionTime() string`

GetAverageExecutionTime returns the AverageExecutionTime field if non-nil, zero value otherwise.

### GetAverageExecutionTimeOk

`func (o *DatasetQueryResults) GetAverageExecutionTimeOk() (*string, bool)`

GetAverageExecutionTimeOk returns a tuple with the AverageExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageExecutionTime

`func (o *DatasetQueryResults) SetAverageExecutionTime(v string)`

SetAverageExecutionTime sets AverageExecutionTime field to given value.

### HasAverageExecutionTime

`func (o *DatasetQueryResults) HasAverageExecutionTime() bool`

HasAverageExecutionTime returns a boolean if a field has been set.

### GetStatus

`func (o *DatasetQueryResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasetQueryResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasetQueryResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatasetQueryResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContext

`func (o *DatasetQueryResults) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DatasetQueryResults) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DatasetQueryResults) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DatasetQueryResults) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRowCount

`func (o *DatasetQueryResults) GetRowCount() int64`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatasetQueryResults) GetRowCountOk() (*int64, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatasetQueryResults) SetRowCount(v int64)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatasetQueryResults) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetRunningTime

`func (o *DatasetQueryResults) GetRunningTime() int64`

GetRunningTime returns the RunningTime field if non-nil, zero value otherwise.

### GetRunningTimeOk

`func (o *DatasetQueryResults) GetRunningTimeOk() (*int64, bool)`

GetRunningTimeOk returns a tuple with the RunningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTime

`func (o *DatasetQueryResults) SetRunningTime(v int64)`

SetRunningTime sets RunningTime field to given value.

### HasRunningTime

`func (o *DatasetQueryResults) HasRunningTime() bool`

HasRunningTime returns a boolean if a field has been set.

### GetData

`func (o *DatasetQueryResults) GetData() DatasetQueryResultsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetQueryResults) GetDataOk() (*DatasetQueryResultsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetQueryResults) SetData(v DatasetQueryResultsData)`

SetData sets Data field to given value.

### HasData

`func (o *DatasetQueryResults) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


