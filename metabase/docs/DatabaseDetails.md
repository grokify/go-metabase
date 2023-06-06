# DatabaseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Dbname** | Pointer to **string** |  | [optional] 
**AuthMech** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ConnProperties** | Pointer to **string** |  | [optional] 
**LetUserControlScheduling** | Pointer to **bool** |  | [optional] 

## Methods

### NewDatabaseDetails

`func NewDatabaseDetails() *DatabaseDetails`

NewDatabaseDetails instantiates a new DatabaseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseDetailsWithDefaults

`func NewDatabaseDetailsWithDefaults() *DatabaseDetails`

NewDatabaseDetailsWithDefaults instantiates a new DatabaseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DatabaseDetails) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseDetails) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseDetails) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DatabaseDetails) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *DatabaseDetails) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseDetails) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseDetails) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatabaseDetails) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDbname

`func (o *DatabaseDetails) GetDbname() string`

GetDbname returns the Dbname field if non-nil, zero value otherwise.

### GetDbnameOk

`func (o *DatabaseDetails) GetDbnameOk() (*string, bool)`

GetDbnameOk returns a tuple with the Dbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbname

`func (o *DatabaseDetails) SetDbname(v string)`

SetDbname sets Dbname field to given value.

### HasDbname

`func (o *DatabaseDetails) HasDbname() bool`

HasDbname returns a boolean if a field has been set.

### GetAuthMech

`func (o *DatabaseDetails) GetAuthMech() int32`

GetAuthMech returns the AuthMech field if non-nil, zero value otherwise.

### GetAuthMechOk

`func (o *DatabaseDetails) GetAuthMechOk() (*int32, bool)`

GetAuthMechOk returns a tuple with the AuthMech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMech

`func (o *DatabaseDetails) SetAuthMech(v int32)`

SetAuthMech sets AuthMech field to given value.

### HasAuthMech

`func (o *DatabaseDetails) HasAuthMech() bool`

HasAuthMech returns a boolean if a field has been set.

### GetUser

`func (o *DatabaseDetails) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DatabaseDetails) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DatabaseDetails) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DatabaseDetails) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *DatabaseDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DatabaseDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConnProperties

`func (o *DatabaseDetails) GetConnProperties() string`

GetConnProperties returns the ConnProperties field if non-nil, zero value otherwise.

### GetConnPropertiesOk

`func (o *DatabaseDetails) GetConnPropertiesOk() (*string, bool)`

GetConnPropertiesOk returns a tuple with the ConnProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnProperties

`func (o *DatabaseDetails) SetConnProperties(v string)`

SetConnProperties sets ConnProperties field to given value.

### HasConnProperties

`func (o *DatabaseDetails) HasConnProperties() bool`

HasConnProperties returns a boolean if a field has been set.

### GetLetUserControlScheduling

`func (o *DatabaseDetails) GetLetUserControlScheduling() bool`

GetLetUserControlScheduling returns the LetUserControlScheduling field if non-nil, zero value otherwise.

### GetLetUserControlSchedulingOk

`func (o *DatabaseDetails) GetLetUserControlSchedulingOk() (*bool, bool)`

GetLetUserControlSchedulingOk returns a tuple with the LetUserControlScheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetUserControlScheduling

`func (o *DatabaseDetails) SetLetUserControlScheduling(v bool)`

SetLetUserControlScheduling sets LetUserControlScheduling field to given value.

### HasLetUserControlScheduling

`func (o *DatabaseDetails) HasLetUserControlScheduling() bool`

HasLetUserControlScheduling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


