# DatasetQueryResultsCol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**TableId** | Pointer to **int64** |  | [optional] 
**SchemaName** | Pointer to **string** |  | [optional] 
**SpecialType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**RemappedFrom** | Pointer to **string** | unknown type | [optional] 
**ExtraInfo** | Pointer to **map[string]interface{}** | can be &#39;{\&quot;target_table_id\&quot;:517}&#39; | [optional] 
**FkFieldId** | Pointer to **string** | unknown type | [optional] 
**RemappedTo** | Pointer to **string** | unknown type | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**VisibilityType** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**DatasetQueryResultsColTarget**](DatasetQueryResultsColTarget.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to [**DatasetQueryResultsColFingerprint**](DatasetQueryResultsColFingerprint.md) |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasetQueryResultsCol

`func NewDatasetQueryResultsCol() *DatasetQueryResultsCol`

NewDatasetQueryResultsCol instantiates a new DatasetQueryResultsCol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetQueryResultsColWithDefaults

`func NewDatasetQueryResultsColWithDefaults() *DatasetQueryResultsCol`

NewDatasetQueryResultsColWithDefaults instantiates a new DatasetQueryResultsCol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatasetQueryResultsCol) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatasetQueryResultsCol) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatasetQueryResultsCol) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatasetQueryResultsCol) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTableId

`func (o *DatasetQueryResultsCol) GetTableId() int64`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *DatasetQueryResultsCol) GetTableIdOk() (*int64, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *DatasetQueryResultsCol) SetTableId(v int64)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *DatasetQueryResultsCol) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetSchemaName

`func (o *DatasetQueryResultsCol) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *DatasetQueryResultsCol) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *DatasetQueryResultsCol) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *DatasetQueryResultsCol) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetSpecialType

`func (o *DatasetQueryResultsCol) GetSpecialType() string`

GetSpecialType returns the SpecialType field if non-nil, zero value otherwise.

### GetSpecialTypeOk

`func (o *DatasetQueryResultsCol) GetSpecialTypeOk() (*string, bool)`

GetSpecialTypeOk returns a tuple with the SpecialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialType

`func (o *DatasetQueryResultsCol) SetSpecialType(v string)`

SetSpecialType sets SpecialType field to given value.

### HasSpecialType

`func (o *DatasetQueryResultsCol) HasSpecialType() bool`

HasSpecialType returns a boolean if a field has been set.

### GetName

`func (o *DatasetQueryResultsCol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetQueryResultsCol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetQueryResultsCol) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasetQueryResultsCol) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *DatasetQueryResultsCol) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatasetQueryResultsCol) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatasetQueryResultsCol) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DatasetQueryResultsCol) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRemappedFrom

`func (o *DatasetQueryResultsCol) GetRemappedFrom() string`

GetRemappedFrom returns the RemappedFrom field if non-nil, zero value otherwise.

### GetRemappedFromOk

`func (o *DatasetQueryResultsCol) GetRemappedFromOk() (*string, bool)`

GetRemappedFromOk returns a tuple with the RemappedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemappedFrom

`func (o *DatasetQueryResultsCol) SetRemappedFrom(v string)`

SetRemappedFrom sets RemappedFrom field to given value.

### HasRemappedFrom

`func (o *DatasetQueryResultsCol) HasRemappedFrom() bool`

HasRemappedFrom returns a boolean if a field has been set.

### GetExtraInfo

`func (o *DatasetQueryResultsCol) GetExtraInfo() map[string]interface{}`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *DatasetQueryResultsCol) GetExtraInfoOk() (*map[string]interface{}, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *DatasetQueryResultsCol) SetExtraInfo(v map[string]interface{})`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *DatasetQueryResultsCol) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetFkFieldId

`func (o *DatasetQueryResultsCol) GetFkFieldId() string`

GetFkFieldId returns the FkFieldId field if non-nil, zero value otherwise.

### GetFkFieldIdOk

`func (o *DatasetQueryResultsCol) GetFkFieldIdOk() (*string, bool)`

GetFkFieldIdOk returns a tuple with the FkFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkFieldId

`func (o *DatasetQueryResultsCol) SetFkFieldId(v string)`

SetFkFieldId sets FkFieldId field to given value.

### HasFkFieldId

`func (o *DatasetQueryResultsCol) HasFkFieldId() bool`

HasFkFieldId returns a boolean if a field has been set.

### GetRemappedTo

`func (o *DatasetQueryResultsCol) GetRemappedTo() string`

GetRemappedTo returns the RemappedTo field if non-nil, zero value otherwise.

### GetRemappedToOk

`func (o *DatasetQueryResultsCol) GetRemappedToOk() (*string, bool)`

GetRemappedToOk returns a tuple with the RemappedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemappedTo

`func (o *DatasetQueryResultsCol) SetRemappedTo(v string)`

SetRemappedTo sets RemappedTo field to given value.

### HasRemappedTo

`func (o *DatasetQueryResultsCol) HasRemappedTo() bool`

HasRemappedTo returns a boolean if a field has been set.

### GetId

`func (o *DatasetQueryResultsCol) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetQueryResultsCol) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetQueryResultsCol) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DatasetQueryResultsCol) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVisibilityType

`func (o *DatasetQueryResultsCol) GetVisibilityType() string`

GetVisibilityType returns the VisibilityType field if non-nil, zero value otherwise.

### GetVisibilityTypeOk

`func (o *DatasetQueryResultsCol) GetVisibilityTypeOk() (*string, bool)`

GetVisibilityTypeOk returns a tuple with the VisibilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityType

`func (o *DatasetQueryResultsCol) SetVisibilityType(v string)`

SetVisibilityType sets VisibilityType field to given value.

### HasVisibilityType

`func (o *DatasetQueryResultsCol) HasVisibilityType() bool`

HasVisibilityType returns a boolean if a field has been set.

### GetTarget

`func (o *DatasetQueryResultsCol) GetTarget() DatasetQueryResultsColTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DatasetQueryResultsCol) GetTargetOk() (*DatasetQueryResultsColTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DatasetQueryResultsCol) SetTarget(v DatasetQueryResultsColTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DatasetQueryResultsCol) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDisplayName

`func (o *DatasetQueryResultsCol) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasetQueryResultsCol) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasetQueryResultsCol) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasetQueryResultsCol) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFingerprint

`func (o *DatasetQueryResultsCol) GetFingerprint() DatasetQueryResultsColFingerprint`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *DatasetQueryResultsCol) GetFingerprintOk() (*DatasetQueryResultsColFingerprint, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *DatasetQueryResultsCol) SetFingerprint(v DatasetQueryResultsColFingerprint)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *DatasetQueryResultsCol) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetBaseType

`func (o *DatasetQueryResultsCol) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *DatasetQueryResultsCol) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *DatasetQueryResultsCol) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *DatasetQueryResultsCol) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


