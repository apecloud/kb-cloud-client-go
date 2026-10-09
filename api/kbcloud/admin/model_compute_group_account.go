// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComputeGroupAccount Grants reported by the engine can include role inheritance. Revoking a direct grant does not remove inherited permissions.
type ComputeGroupAccount struct {
	Identity      *string `json:"identity,omitempty"`
	AccountName   *string `json:"accountName,omitempty"`
	Manageable    *bool   `json:"manageable,omitempty"`
	GlobalGrants  *string `json:"globalGrants,omitempty"`
	ComputeGrants *string `json:"computeGrants,omitempty"`
	Roles         *string `json:"roles,omitempty"`
	IsDefault     *bool   `json:"isDefault,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComputeGroupAccount instantiates a new ComputeGroupAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupAccount() *ComputeGroupAccount {
	this := ComputeGroupAccount{}
	return &this
}

// NewComputeGroupAccountWithDefaults instantiates a new ComputeGroupAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupAccountWithDefaults() *ComputeGroupAccount {
	this := ComputeGroupAccount{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasIdentity() bool {
	return o != nil && o.Identity != nil
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *ComputeGroupAccount) SetIdentity(v string) {
	o.Identity = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasAccountName() bool {
	return o != nil && o.AccountName != nil
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ComputeGroupAccount) SetAccountName(v string) {
	o.AccountName = &v
}

// GetManageable returns the Manageable field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetManageable() bool {
	if o == nil || o.Manageable == nil {
		var ret bool
		return ret
	}
	return *o.Manageable
}

// GetManageableOk returns a tuple with the Manageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetManageableOk() (*bool, bool) {
	if o == nil || o.Manageable == nil {
		return nil, false
	}
	return o.Manageable, true
}

// HasManageable returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasManageable() bool {
	return o != nil && o.Manageable != nil
}

// SetManageable gets a reference to the given bool and assigns it to the Manageable field.
func (o *ComputeGroupAccount) SetManageable(v bool) {
	o.Manageable = &v
}

// GetGlobalGrants returns the GlobalGrants field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetGlobalGrants() string {
	if o == nil || o.GlobalGrants == nil {
		var ret string
		return ret
	}
	return *o.GlobalGrants
}

// GetGlobalGrantsOk returns a tuple with the GlobalGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetGlobalGrantsOk() (*string, bool) {
	if o == nil || o.GlobalGrants == nil {
		return nil, false
	}
	return o.GlobalGrants, true
}

// HasGlobalGrants returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasGlobalGrants() bool {
	return o != nil && o.GlobalGrants != nil
}

// SetGlobalGrants gets a reference to the given string and assigns it to the GlobalGrants field.
func (o *ComputeGroupAccount) SetGlobalGrants(v string) {
	o.GlobalGrants = &v
}

// GetComputeGrants returns the ComputeGrants field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetComputeGrants() string {
	if o == nil || o.ComputeGrants == nil {
		var ret string
		return ret
	}
	return *o.ComputeGrants
}

// GetComputeGrantsOk returns a tuple with the ComputeGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetComputeGrantsOk() (*string, bool) {
	if o == nil || o.ComputeGrants == nil {
		return nil, false
	}
	return o.ComputeGrants, true
}

// HasComputeGrants returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasComputeGrants() bool {
	return o != nil && o.ComputeGrants != nil
}

// SetComputeGrants gets a reference to the given string and assigns it to the ComputeGrants field.
func (o *ComputeGroupAccount) SetComputeGrants(v string) {
	o.ComputeGrants = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetRoles() string {
	if o == nil || o.Roles == nil {
		var ret string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetRolesOk() (*string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given string and assigns it to the Roles field.
func (o *ComputeGroupAccount) SetRoles(v string) {
	o.Roles = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ComputeGroupAccount) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupAccount) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ComputeGroupAccount) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ComputeGroupAccount) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.Manageable != nil {
		toSerialize["manageable"] = o.Manageable
	}
	if o.GlobalGrants != nil {
		toSerialize["globalGrants"] = o.GlobalGrants
	}
	if o.ComputeGrants != nil {
		toSerialize["computeGrants"] = o.ComputeGrants
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Identity      *string `json:"identity,omitempty"`
		AccountName   *string `json:"accountName,omitempty"`
		Manageable    *bool   `json:"manageable,omitempty"`
		GlobalGrants  *string `json:"globalGrants,omitempty"`
		ComputeGrants *string `json:"computeGrants,omitempty"`
		Roles         *string `json:"roles,omitempty"`
		IsDefault     *bool   `json:"isDefault,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"identity", "accountName", "manageable", "globalGrants", "computeGrants", "roles", "isDefault"})
	} else {
		return err
	}
	o.Identity = all.Identity
	o.AccountName = all.AccountName
	o.Manageable = all.Manageable
	o.GlobalGrants = all.GlobalGrants
	o.ComputeGrants = all.ComputeGrants
	o.Roles = all.Roles
	o.IsDefault = all.IsDefault

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
