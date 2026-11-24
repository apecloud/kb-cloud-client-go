// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeDefinition struct {
	Name             string                          `json:"name"`
	SupportsColumns  *bool                           `json:"supportsColumns,omitempty"`
	GrantOptionScope ObjectPrivilegeGrantOptionScope `json:"grantOptionScope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeDefinition instantiates a new ObjectPrivilegeDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeDefinition(name string, grantOptionScope ObjectPrivilegeGrantOptionScope) *ObjectPrivilegeDefinition {
	this := ObjectPrivilegeDefinition{}
	this.Name = name
	this.GrantOptionScope = grantOptionScope
	return &this
}

// NewObjectPrivilegeDefinitionWithDefaults instantiates a new ObjectPrivilegeDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeDefinitionWithDefaults() *ObjectPrivilegeDefinition {
	this := ObjectPrivilegeDefinition{}
	return &this
}

// GetName returns the Name field value.
func (o *ObjectPrivilegeDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObjectPrivilegeDefinition) SetName(v string) {
	o.Name = v
}

// GetSupportsColumns returns the SupportsColumns field value if set, zero value otherwise.
func (o *ObjectPrivilegeDefinition) GetSupportsColumns() bool {
	if o == nil || o.SupportsColumns == nil {
		var ret bool
		return ret
	}
	return *o.SupportsColumns
}

// GetSupportsColumnsOk returns a tuple with the SupportsColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeDefinition) GetSupportsColumnsOk() (*bool, bool) {
	if o == nil || o.SupportsColumns == nil {
		return nil, false
	}
	return o.SupportsColumns, true
}

// HasSupportsColumns returns a boolean if a field has been set.
func (o *ObjectPrivilegeDefinition) HasSupportsColumns() bool {
	return o != nil && o.SupportsColumns != nil
}

// SetSupportsColumns gets a reference to the given bool and assigns it to the SupportsColumns field.
func (o *ObjectPrivilegeDefinition) SetSupportsColumns(v bool) {
	o.SupportsColumns = &v
}

// GetGrantOptionScope returns the GrantOptionScope field value.
func (o *ObjectPrivilegeDefinition) GetGrantOptionScope() ObjectPrivilegeGrantOptionScope {
	if o == nil {
		var ret ObjectPrivilegeGrantOptionScope
		return ret
	}
	return o.GrantOptionScope
}

// GetGrantOptionScopeOk returns a tuple with the GrantOptionScope field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeDefinition) GetGrantOptionScopeOk() (*ObjectPrivilegeGrantOptionScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantOptionScope, true
}

// SetGrantOptionScope sets field value.
func (o *ObjectPrivilegeDefinition) SetGrantOptionScope(v ObjectPrivilegeGrantOptionScope) {
	o.GrantOptionScope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.SupportsColumns != nil {
		toSerialize["supportsColumns"] = o.SupportsColumns
	}
	toSerialize["grantOptionScope"] = o.GrantOptionScope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                          `json:"name"`
		SupportsColumns  *bool                            `json:"supportsColumns,omitempty"`
		GrantOptionScope *ObjectPrivilegeGrantOptionScope `json:"grantOptionScope"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.GrantOptionScope == nil {
		return fmt.Errorf("required field grantOptionScope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "supportsColumns", "grantOptionScope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.SupportsColumns = all.SupportsColumns
	if !all.GrantOptionScope.IsValid() {
		hasInvalidField = true
	} else {
		o.GrantOptionScope = *all.GrantOptionScope
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
