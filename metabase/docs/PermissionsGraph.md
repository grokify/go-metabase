# PermissionsGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **int32** |  | [optional] 
**Groups** | Pointer to [**map[string]map[string]PermissionGraphData**](map.md) |  | [optional] 

## Methods

### NewPermissionsGraph

`func NewPermissionsGraph() *PermissionsGraph`

NewPermissionsGraph instantiates a new PermissionsGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsGraphWithDefaults

`func NewPermissionsGraphWithDefaults() *PermissionsGraph`

NewPermissionsGraphWithDefaults instantiates a new PermissionsGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *PermissionsGraph) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PermissionsGraph) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PermissionsGraph) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PermissionsGraph) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetGroups

`func (o *PermissionsGraph) GetGroups() map[string]map[string]PermissionGraphData`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PermissionsGraph) GetGroupsOk() (*map[string]map[string]PermissionGraphData, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PermissionsGraph) SetGroups(v map[string]map[string]PermissionGraphData)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PermissionsGraph) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


