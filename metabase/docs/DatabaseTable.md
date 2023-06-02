# DatabaseTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** | unknown type | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**RawTableId** | Pointer to **string** | unknown type | [optional] 
**ShowInGettingStarted** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Caveats** | Pointer to **string** | unknown type | [optional] 
**Rows** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**EntityName** | Pointer to **string** | unknown type | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**DbId** | Pointer to **int64** |  | [optional] 
**VisibilityType** | Pointer to **string** | unknown type | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**PointsOfInterest** | Pointer to **string** | unknown type | [optional] 

## Methods

### NewDatabaseTable

`func NewDatabaseTable() *DatabaseTable`

NewDatabaseTable instantiates a new DatabaseTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTableWithDefaults

`func NewDatabaseTableWithDefaults() *DatabaseTable`

NewDatabaseTableWithDefaults instantiates a new DatabaseTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatabaseTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *DatabaseTable) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DatabaseTable) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DatabaseTable) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DatabaseTable) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetSchema

`func (o *DatabaseTable) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DatabaseTable) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DatabaseTable) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DatabaseTable) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetRawTableId

`func (o *DatabaseTable) GetRawTableId() string`

GetRawTableId returns the RawTableId field if non-nil, zero value otherwise.

### GetRawTableIdOk

`func (o *DatabaseTable) GetRawTableIdOk() (*string, bool)`

GetRawTableIdOk returns a tuple with the RawTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTableId

`func (o *DatabaseTable) SetRawTableId(v string)`

SetRawTableId sets RawTableId field to given value.

### HasRawTableId

`func (o *DatabaseTable) HasRawTableId() bool`

HasRawTableId returns a boolean if a field has been set.

### GetShowInGettingStarted

`func (o *DatabaseTable) GetShowInGettingStarted() bool`

GetShowInGettingStarted returns the ShowInGettingStarted field if non-nil, zero value otherwise.

### GetShowInGettingStartedOk

`func (o *DatabaseTable) GetShowInGettingStartedOk() (*bool, bool)`

GetShowInGettingStartedOk returns a tuple with the ShowInGettingStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInGettingStarted

`func (o *DatabaseTable) SetShowInGettingStarted(v bool)`

SetShowInGettingStarted sets ShowInGettingStarted field to given value.

### HasShowInGettingStarted

`func (o *DatabaseTable) HasShowInGettingStarted() bool`

HasShowInGettingStarted returns a boolean if a field has been set.

### GetName

`func (o *DatabaseTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatabaseTable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCaveats

`func (o *DatabaseTable) GetCaveats() string`

GetCaveats returns the Caveats field if non-nil, zero value otherwise.

### GetCaveatsOk

`func (o *DatabaseTable) GetCaveatsOk() (*string, bool)`

GetCaveatsOk returns a tuple with the Caveats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaveats

`func (o *DatabaseTable) SetCaveats(v string)`

SetCaveats sets Caveats field to given value.

### HasCaveats

`func (o *DatabaseTable) HasCaveats() bool`

HasCaveats returns a boolean if a field has been set.

### GetRows

`func (o *DatabaseTable) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatabaseTable) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatabaseTable) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatabaseTable) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DatabaseTable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabaseTable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabaseTable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DatabaseTable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEntityName

`func (o *DatabaseTable) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *DatabaseTable) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *DatabaseTable) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *DatabaseTable) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetActive

`func (o *DatabaseTable) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DatabaseTable) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DatabaseTable) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DatabaseTable) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *DatabaseTable) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseTable) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseTable) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DatabaseTable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDbId

`func (o *DatabaseTable) GetDbId() int64`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *DatabaseTable) GetDbIdOk() (*int64, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *DatabaseTable) SetDbId(v int64)`

SetDbId sets DbId field to given value.

### HasDbId

`func (o *DatabaseTable) HasDbId() bool`

HasDbId returns a boolean if a field has been set.

### GetVisibilityType

`func (o *DatabaseTable) GetVisibilityType() string`

GetVisibilityType returns the VisibilityType field if non-nil, zero value otherwise.

### GetVisibilityTypeOk

`func (o *DatabaseTable) GetVisibilityTypeOk() (*string, bool)`

GetVisibilityTypeOk returns a tuple with the VisibilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityType

`func (o *DatabaseTable) SetVisibilityType(v string)`

SetVisibilityType sets VisibilityType field to given value.

### HasVisibilityType

`func (o *DatabaseTable) HasVisibilityType() bool`

HasVisibilityType returns a boolean if a field has been set.

### GetDisplayName

`func (o *DatabaseTable) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatabaseTable) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatabaseTable) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatabaseTable) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DatabaseTable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseTable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseTable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DatabaseTable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPointsOfInterest

`func (o *DatabaseTable) GetPointsOfInterest() string`

GetPointsOfInterest returns the PointsOfInterest field if non-nil, zero value otherwise.

### GetPointsOfInterestOk

`func (o *DatabaseTable) GetPointsOfInterestOk() (*string, bool)`

GetPointsOfInterestOk returns a tuple with the PointsOfInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsOfInterest

`func (o *DatabaseTable) SetPointsOfInterest(v string)`

SetPointsOfInterest sets PointsOfInterest field to given value.

### HasPointsOfInterest

`func (o *DatabaseTable) HasPointsOfInterest() bool`

HasPointsOfInterest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


