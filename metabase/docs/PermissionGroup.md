# PermissionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPermissionGroup

`func NewPermissionGroup() *PermissionGroup`

NewPermissionGroup instantiates a new PermissionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGroupWithDefaults

`func NewPermissionGroupWithDefaults() *PermissionGroup`

NewPermissionGroupWithDefaults instantiates a new PermissionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PermissionGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMemberCount

`func (o *PermissionGroup) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *PermissionGroup) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *PermissionGroup) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *PermissionGroup) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


