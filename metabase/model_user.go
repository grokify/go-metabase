package metabase

import (
	"time"
	"encoding/json"
)


type User struct {
	Id  						*int32		`json:"id,omitempty"`
	Email  						*string		`json:"email,omitempty"`
	CommonName  				*string		`json:"common_name,omitempty"`
	FirstName  					*string		`json:"first_name,omitempty"`
	LastName  					*string		`json:"last_name,omitempty"`
	Locale  					*string		`json:"locale,omitempty"`
	LastLogin  					*time.Time	`json:"last_login,omitempty"`
	IsActive  					*bool		`json:"is_active,omitempty"`
	UpdatedAt  					*time.Time	`json:"updated_at,omitempty"`
	DateJoined  				*time.Time	`json:"date_joined,omitempty"`
	GroupIDs  					[]int32		`json:"group_ids,omitempty"`
	IsSuperuser  				*bool		`json:"is_superuser,omitempty"`
	PersonalCollectionID  		*int32		`json:"personal_collection_id,omitempty"`
	IsQbnewb  					*bool		`json:"is_qbnewb,omitempty"`
	JwtAttributes  				*string		`json:"jwt_attributes,omitempty"`
	// LoginAttributes  			*struct		`json:"login_attributes,omitempty"`
	Tenant_id  					*string		`json:"tenant_id,omitempty"`
	SSOSource  					*string		`json:"sso_source,omitempty"`
}


func NewUser() *User {
	this := User{}
	return &this
}

func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.CommonName) {
		toSerialize["common_name"] = o.CommonName
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DateJoined) {
		toSerialize["date_joined"] = o.DateJoined
	}
	if !IsNil(o.GroupIDs) {
		toSerialize["group_ids"] = o.GroupIDs
	}
	if !IsNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if !IsNil(o.PersonalCollectionID) {
		toSerialize["personal_collection_id"] = o.PersonalCollectionID
	}
	if !IsNil(o.IsQbnewb) {
		toSerialize["is_qbnewb"] = o.IsQbnewb
	}
	if !IsNil(o.JwtAttributes) {
		toSerialize["jwt_attributes"] = o.JwtAttributes
	}
	// if !IsNil(o.LoginAttributes) {
	// 	toSerialize["login_attributes"] = o.LoginAttributes
	// }
	if !IsNil(o.Tenant_id) {
		toSerialize["tenant_id"] = o.Tenant_id
	}
	if !IsNil(o.SSOSource) {
		toSerialize["sso_source"] = o.SSOSource
	}
	return toSerialize, nil
}



type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



// Wrapper for a slice of Users
type Users struct {
	Data 	[]User 	`json:"data,omitempty"`
	Total 	*int32 	`json:"total,omitempty"`
	Limit 	*int32 	`json:"limit,omitempty"`
	Offset 	*int32 	`json:"offset,omitempty"`
}

func NewUsers(data []User) *Users {
    return &Users{Data: data}
}

func NewUsersWithDefaults() *Users {
    return &Users{}
}

func (u Users) MarshalJSON() ([]byte, error) {
    toSerialize := map[string]interface{}{}
    if u.Data != nil {
        toSerialize["data"] = u.Data
    }
    if u.Total != nil {
        toSerialize["total"] = u.Total
    }
    if u.Limit != nil {
        toSerialize["limit"] = u.Limit
    }
    if u.Offset != nil {
        toSerialize["offset"] = u.Offset
    }
    return json.Marshal(toSerialize)
}

func (u *Users) UnmarshalJSON(b []byte) error {
    var aux struct {
        Data   []User `json:"data"`
        Total  *int32 `json:"total"`
        Limit  *int32 `json:"limit"`
        Offset *int32 `json:"offset"`
    }
    if err := json.Unmarshal(b, &aux); err != nil {
        return err
    }
    *u = Users(aux)
    return nil
}

func (u Users) ToMap() (map[string]interface{}, error) {
    data, err := json.Marshal(u)
    if err != nil {
        return nil, err
    }
    var m map[string]interface{}
    err = json.Unmarshal(data, &m)
    return m, err
}


type NullableUsers struct {
    value *Users
    isSet bool
}

func (v NullableUsers) Get() *Users {
    return v.value
}

func (v *NullableUsers) Set(val *Users) {
    v.value = val
    v.isSet = true
}

func (v NullableUsers) IsSet() bool {
    return v.isSet
}

func (v *NullableUsers) Unset() {
    v.value = nil
    v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
    return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
    return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
    v.isSet = true
    return json.Unmarshal(src, &v.value)
}
