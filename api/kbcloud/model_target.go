// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// Target Identifies which component and role the scheduling rule applies to.
type Target struct {
	// Component name to match. Use "*" to match all components.
	//
	Component *string `json:"component,omitempty"`
	// Role name to match. Use "*" to match all roles (or no role if the component has none).
	//
	Role *Target_role `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTarget instantiates a new Target object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTarget() *Target {
	this := Target{}
	var component string = "*"
	this.Component = &component
	var role Target_role = Target_role
	this.Role = &role
	return &this
}

// NewTargetWithDefaults instantiates a new Target object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTargetWithDefaults() *Target {
	this := Target{}
	var component string = "*"
	this.Component = &component
	var role Target_role = Target_role
	this.Role = &role
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Target) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Target) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *Target) SetComponent(v string) {
	o.Component = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Target) GetRole() Target_role {
	if o == nil || o.Role == nil {
		var ret Target_role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetRoleOk() (*Target_role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Target) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given Target_role and assigns it to the Role field.
func (o *Target) SetRole(v Target_role) {
	o.Role = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Target) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Target) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string      `json:"component,omitempty"`
		Role      *Target_role `json:"role,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "role"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = all.Component
	if all.Role != nil && !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
