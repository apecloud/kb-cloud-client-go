// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// CkAccount A ClickHouse account.
type CkAccount struct {
	// The account (user) name.
	Name *string `json:"name,omitempty"`
	// Host IPs the account is allowed to log in from.
	HostIp []string `json:"hostIp,omitempty"`
	// Host names the account is allowed to log in from.
	HostNames []string `json:"hostNames,omitempty"`
	// Roles granted to the account, separated by commas.
	Roles *string `json:"roles,omitempty"`
	// The account role type.
	Type *CkAccountType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCkAccount instantiates a new CkAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCkAccount() *CkAccount {
	this := CkAccount{}
	return &this
}

// NewCkAccountWithDefaults instantiates a new CkAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCkAccountWithDefaults() *CkAccount {
	this := CkAccount{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CkAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CkAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CkAccount) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CkAccount) SetName(v string) {
	o.Name = &v
}

// GetHostIp returns the HostIp field value if set, zero value otherwise.
func (o *CkAccount) GetHostIp() []string {
	if o == nil || o.HostIp == nil {
		var ret []string
		return ret
	}
	return o.HostIp
}

// GetHostIpOk returns a tuple with the HostIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CkAccount) GetHostIpOk() (*[]string, bool) {
	if o == nil || o.HostIp == nil {
		return nil, false
	}
	return &o.HostIp, true
}

// HasHostIp returns a boolean if a field has been set.
func (o *CkAccount) HasHostIp() bool {
	return o != nil && o.HostIp != nil
}

// SetHostIp gets a reference to the given []string and assigns it to the HostIp field.
func (o *CkAccount) SetHostIp(v []string) {
	o.HostIp = v
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *CkAccount) GetHostNames() []string {
	if o == nil || o.HostNames == nil {
		var ret []string
		return ret
	}
	return o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CkAccount) GetHostNamesOk() (*[]string, bool) {
	if o == nil || o.HostNames == nil {
		return nil, false
	}
	return &o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *CkAccount) HasHostNames() bool {
	return o != nil && o.HostNames != nil
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *CkAccount) SetHostNames(v []string) {
	o.HostNames = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CkAccount) GetRoles() string {
	if o == nil || o.Roles == nil {
		var ret string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CkAccount) GetRolesOk() (*string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CkAccount) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given string and assigns it to the Roles field.
func (o *CkAccount) SetRoles(v string) {
	o.Roles = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CkAccount) GetType() CkAccountType {
	if o == nil || o.Type == nil {
		var ret CkAccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CkAccount) GetTypeOk() (*CkAccountType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CkAccount) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CkAccountType and assigns it to the Type field.
func (o *CkAccount) SetType(v CkAccountType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CkAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.HostIp != nil {
		toSerialize["hostIp"] = o.HostIp
	}
	if o.HostNames != nil {
		toSerialize["hostNames"] = o.HostNames
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CkAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string        `json:"name,omitempty"`
		HostIp    []string       `json:"hostIp,omitempty"`
		HostNames []string       `json:"hostNames,omitempty"`
		Roles     *string        `json:"roles,omitempty"`
		Type      *CkAccountType `json:"type,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "hostIp", "hostNames", "roles", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.HostIp = all.HostIp
	o.HostNames = all.HostNames
	o.Roles = all.Roles
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
