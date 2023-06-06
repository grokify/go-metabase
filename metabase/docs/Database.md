# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**IsFullSync** | Pointer to **bool** |  | [optional] 
**IsSample** | Pointer to **bool** |  | [optional] 
**CacheFieldValuesSchedule** | Pointer to **string** |  | [optional] 
**MetadataSyncSchedule** | Pointer to **string** |  | [optional] 
**Caveats** | Pointer to **string** | type unknown | [optional] 
**Engine** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**NativePermissions** | Pointer to **string** |  | [optional] 
**PointsOfInterest** | Pointer to **string** | type unknown | [optional] 
**Details** | Pointer to [**DatabaseDetails**](DatabaseDetails.md) |  | [optional] 
**Tables** | Pointer to [**[]DatabaseTable**](DatabaseTable.md) |  | [optional] 

## Methods

### NewDatabase

`func NewDatabase(id int64, ) *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Database) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Database) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Database) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Database) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Database) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Database) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Database) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *Database) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Database) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Database) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Database) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIsFullSync

`func (o *Database) GetIsFullSync() bool`

GetIsFullSync returns the IsFullSync field if non-nil, zero value otherwise.

### GetIsFullSyncOk

`func (o *Database) GetIsFullSyncOk() (*bool, bool)`

GetIsFullSyncOk returns a tuple with the IsFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullSync

`func (o *Database) SetIsFullSync(v bool)`

SetIsFullSync sets IsFullSync field to given value.

### HasIsFullSync

`func (o *Database) HasIsFullSync() bool`

HasIsFullSync returns a boolean if a field has been set.

### GetIsSample

`func (o *Database) GetIsSample() bool`

GetIsSample returns the IsSample field if non-nil, zero value otherwise.

### GetIsSampleOk

`func (o *Database) GetIsSampleOk() (*bool, bool)`

GetIsSampleOk returns a tuple with the IsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSample

`func (o *Database) SetIsSample(v bool)`

SetIsSample sets IsSample field to given value.

### HasIsSample

`func (o *Database) HasIsSample() bool`

HasIsSample returns a boolean if a field has been set.

### GetCacheFieldValuesSchedule

`func (o *Database) GetCacheFieldValuesSchedule() string`

GetCacheFieldValuesSchedule returns the CacheFieldValuesSchedule field if non-nil, zero value otherwise.

### GetCacheFieldValuesScheduleOk

`func (o *Database) GetCacheFieldValuesScheduleOk() (*string, bool)`

GetCacheFieldValuesScheduleOk returns a tuple with the CacheFieldValuesSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFieldValuesSchedule

`func (o *Database) SetCacheFieldValuesSchedule(v string)`

SetCacheFieldValuesSchedule sets CacheFieldValuesSchedule field to given value.

### HasCacheFieldValuesSchedule

`func (o *Database) HasCacheFieldValuesSchedule() bool`

HasCacheFieldValuesSchedule returns a boolean if a field has been set.

### GetMetadataSyncSchedule

`func (o *Database) GetMetadataSyncSchedule() string`

GetMetadataSyncSchedule returns the MetadataSyncSchedule field if non-nil, zero value otherwise.

### GetMetadataSyncScheduleOk

`func (o *Database) GetMetadataSyncScheduleOk() (*string, bool)`

GetMetadataSyncScheduleOk returns a tuple with the MetadataSyncSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSyncSchedule

`func (o *Database) SetMetadataSyncSchedule(v string)`

SetMetadataSyncSchedule sets MetadataSyncSchedule field to given value.

### HasMetadataSyncSchedule

`func (o *Database) HasMetadataSyncSchedule() bool`

HasMetadataSyncSchedule returns a boolean if a field has been set.

### GetCaveats

`func (o *Database) GetCaveats() string`

GetCaveats returns the Caveats field if non-nil, zero value otherwise.

### GetCaveatsOk

`func (o *Database) GetCaveatsOk() (*string, bool)`

GetCaveatsOk returns a tuple with the Caveats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaveats

`func (o *Database) SetCaveats(v string)`

SetCaveats sets Caveats field to given value.

### HasCaveats

`func (o *Database) HasCaveats() bool`

HasCaveats returns a boolean if a field has been set.

### GetEngine

`func (o *Database) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *Database) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *Database) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *Database) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Database) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Database) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Database) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Database) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Database) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Database) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Database) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Database) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNativePermissions

`func (o *Database) GetNativePermissions() string`

GetNativePermissions returns the NativePermissions field if non-nil, zero value otherwise.

### GetNativePermissionsOk

`func (o *Database) GetNativePermissionsOk() (*string, bool)`

GetNativePermissionsOk returns a tuple with the NativePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePermissions

`func (o *Database) SetNativePermissions(v string)`

SetNativePermissions sets NativePermissions field to given value.

### HasNativePermissions

`func (o *Database) HasNativePermissions() bool`

HasNativePermissions returns a boolean if a field has been set.

### GetPointsOfInterest

`func (o *Database) GetPointsOfInterest() string`

GetPointsOfInterest returns the PointsOfInterest field if non-nil, zero value otherwise.

### GetPointsOfInterestOk

`func (o *Database) GetPointsOfInterestOk() (*string, bool)`

GetPointsOfInterestOk returns a tuple with the PointsOfInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsOfInterest

`func (o *Database) SetPointsOfInterest(v string)`

SetPointsOfInterest sets PointsOfInterest field to given value.

### HasPointsOfInterest

`func (o *Database) HasPointsOfInterest() bool`

HasPointsOfInterest returns a boolean if a field has been set.

### GetDetails

`func (o *Database) GetDetails() DatabaseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Database) GetDetailsOk() (*DatabaseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Database) SetDetails(v DatabaseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Database) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTables

`func (o *Database) GetTables() []DatabaseTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *Database) GetTablesOk() (*[]DatabaseTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *Database) SetTables(v []DatabaseTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *Database) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


