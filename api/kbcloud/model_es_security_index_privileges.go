// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityIndexPrivileges struct {
	Names                  []string                 `json:"names"`
	Privileges             []string                 `json:"privileges"`
	FieldSecurity          *ESSecurityFieldSecurity `json:"field_security,omitempty"`
	Query                  *string                  `json:"query,omitempty"`
	AllowRestrictedIndices common.NullableBool      `json:"allow_restricted_indices,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityIndexPrivileges instantiates a new ESSecurityIndexPrivileges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityIndexPrivileges(names []string, privileges []string) *ESSecurityIndexPrivileges {
	this := ESSecurityIndexPrivileges{}
	this.Names = names
	this.Privileges = privileges
	return &this
}

// NewESSecurityIndexPrivilegesWithDefaults instantiates a new ESSecurityIndexPrivileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityIndexPrivilegesWithDefaults() *ESSecurityIndexPrivileges {
	this := ESSecurityIndexPrivileges{}
	return &this
}

// GetNames returns the Names field value.
func (o *ESSecurityIndexPrivileges) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *ESSecurityIndexPrivileges) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *ESSecurityIndexPrivileges) SetNames(v []string) {
	o.Names = v
}

// GetPrivileges returns the Privileges field value.
func (o *ESSecurityIndexPrivileges) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *ESSecurityIndexPrivileges) GetPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *ESSecurityIndexPrivileges) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetFieldSecurity returns the FieldSecurity field value if set, zero value otherwise.
func (o *ESSecurityIndexPrivileges) GetFieldSecurity() ESSecurityFieldSecurity {
	if o == nil || o.FieldSecurity == nil {
		var ret ESSecurityFieldSecurity
		return ret
	}
	return *o.FieldSecurity
}

// GetFieldSecurityOk returns a tuple with the FieldSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityIndexPrivileges) GetFieldSecurityOk() (*ESSecurityFieldSecurity, bool) {
	if o == nil || o.FieldSecurity == nil {
		return nil, false
	}
	return o.FieldSecurity, true
}

// HasFieldSecurity returns a boolean if a field has been set.
func (o *ESSecurityIndexPrivileges) HasFieldSecurity() bool {
	return o != nil && o.FieldSecurity != nil
}

// SetFieldSecurity gets a reference to the given ESSecurityFieldSecurity and assigns it to the FieldSecurity field.
func (o *ESSecurityIndexPrivileges) SetFieldSecurity(v ESSecurityFieldSecurity) {
	o.FieldSecurity = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ESSecurityIndexPrivileges) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityIndexPrivileges) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ESSecurityIndexPrivileges) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ESSecurityIndexPrivileges) SetQuery(v string) {
	o.Query = &v
}

// GetAllowRestrictedIndices returns the AllowRestrictedIndices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ESSecurityIndexPrivileges) GetAllowRestrictedIndices() bool {
	if o == nil || o.AllowRestrictedIndices.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowRestrictedIndices.Get()
}

// GetAllowRestrictedIndicesOk returns a tuple with the AllowRestrictedIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ESSecurityIndexPrivileges) GetAllowRestrictedIndicesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowRestrictedIndices.Get(), o.AllowRestrictedIndices.IsSet()
}

// HasAllowRestrictedIndices returns a boolean if a field has been set.
func (o *ESSecurityIndexPrivileges) HasAllowRestrictedIndices() bool {
	return o != nil && o.AllowRestrictedIndices.IsSet()
}

// SetAllowRestrictedIndices gets a reference to the given common.NullableBool and assigns it to the AllowRestrictedIndices field.
func (o *ESSecurityIndexPrivileges) SetAllowRestrictedIndices(v bool) {
	o.AllowRestrictedIndices.Set(&v)
}

// SetAllowRestrictedIndicesNil sets the value for AllowRestrictedIndices to be an explicit nil.
func (o *ESSecurityIndexPrivileges) SetAllowRestrictedIndicesNil() {
	o.AllowRestrictedIndices.Set(nil)
}

// UnsetAllowRestrictedIndices ensures that no value is present for AllowRestrictedIndices, not even an explicit nil.
func (o *ESSecurityIndexPrivileges) UnsetAllowRestrictedIndices() {
	o.AllowRestrictedIndices.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityIndexPrivileges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["names"] = o.Names
	toSerialize["privileges"] = o.Privileges
	if o.FieldSecurity != nil {
		toSerialize["field_security"] = o.FieldSecurity
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.AllowRestrictedIndices.IsSet() {
		toSerialize["allow_restricted_indices"] = o.AllowRestrictedIndices.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityIndexPrivileges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Names                  *[]string                `json:"names"`
		Privileges             *[]string                `json:"privileges"`
		FieldSecurity          *ESSecurityFieldSecurity `json:"field_security,omitempty"`
		Query                  *string                  `json:"query,omitempty"`
		AllowRestrictedIndices common.NullableBool      `json:"allow_restricted_indices,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"names", "privileges", "field_security", "query", "allow_restricted_indices"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Names = *all.Names
	o.Privileges = *all.Privileges
	if all.FieldSecurity != nil && all.FieldSecurity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FieldSecurity = all.FieldSecurity
	o.Query = all.Query
	o.AllowRestrictedIndices = all.AllowRestrictedIndices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
