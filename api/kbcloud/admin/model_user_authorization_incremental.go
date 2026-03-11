// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// UserAuthorizationIncremental Incremental user authorization update.
type UserAuthorizationIncremental struct {
	// Add or update user roles in the specified organizations.
	Grant []UserOrgRole `json:"grant,omitempty"`
	// Remove user roles from the specified organizations (only orgName is strictly required).
	Remove []UserOrgRole `json:"remove,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserAuthorizationIncremental instantiates a new UserAuthorizationIncremental object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserAuthorizationIncremental() *UserAuthorizationIncremental {
	this := UserAuthorizationIncremental{}
	return &this
}

// NewUserAuthorizationIncrementalWithDefaults instantiates a new UserAuthorizationIncremental object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserAuthorizationIncrementalWithDefaults() *UserAuthorizationIncremental {
	this := UserAuthorizationIncremental{}
	return &this
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *UserAuthorizationIncremental) GetGrant() []UserOrgRole {
	if o == nil || o.Grant == nil {
		var ret []UserOrgRole
		return ret
	}
	return o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthorizationIncremental) GetGrantOk() (*[]UserOrgRole, bool) {
	if o == nil || o.Grant == nil {
		return nil, false
	}
	return &o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *UserAuthorizationIncremental) HasGrant() bool {
	return o != nil && o.Grant != nil
}

// SetGrant gets a reference to the given []UserOrgRole and assigns it to the Grant field.
func (o *UserAuthorizationIncremental) SetGrant(v []UserOrgRole) {
	o.Grant = v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *UserAuthorizationIncremental) GetRemove() []UserOrgRole {
	if o == nil || o.Remove == nil {
		var ret []UserOrgRole
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthorizationIncremental) GetRemoveOk() (*[]UserOrgRole, bool) {
	if o == nil || o.Remove == nil {
		return nil, false
	}
	return &o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *UserAuthorizationIncremental) HasRemove() bool {
	return o != nil && o.Remove != nil
}

// SetRemove gets a reference to the given []UserOrgRole and assigns it to the Remove field.
func (o *UserAuthorizationIncremental) SetRemove(v []UserOrgRole) {
	o.Remove = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserAuthorizationIncremental) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Grant != nil {
		toSerialize["grant"] = o.Grant
	}
	if o.Remove != nil {
		toSerialize["remove"] = o.Remove
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserAuthorizationIncremental) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Grant  []UserOrgRole `json:"grant,omitempty"`
		Remove []UserOrgRole `json:"remove,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"grant", "remove"})
	} else {
		return err
	}
	o.Grant = all.Grant
	o.Remove = all.Remove

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
