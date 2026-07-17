// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityRoleMapping struct {
	Enabled       common.NullableBool      `json:"enabled,omitempty"`
	Roles         []string                 `json:"roles,omitempty"`
	RoleTemplates []ESSecurityRoleTemplate `json:"role_templates,omitempty"`
	Rules         map[string]interface{}   `json:"rules,omitempty"`
	Metadata      map[string]interface{}   `json:"metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityRoleMapping instantiates a new ESSecurityRoleMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityRoleMapping() *ESSecurityRoleMapping {
	this := ESSecurityRoleMapping{}
	return &this
}

// NewESSecurityRoleMappingWithDefaults instantiates a new ESSecurityRoleMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityRoleMappingWithDefaults() *ESSecurityRoleMapping {
	this := ESSecurityRoleMapping{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ESSecurityRoleMapping) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ESSecurityRoleMapping) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *ESSecurityRoleMapping) HasEnabled() bool {
	return o != nil && o.Enabled.IsSet()
}

// SetEnabled gets a reference to the given common.NullableBool and assigns it to the Enabled field.
func (o *ESSecurityRoleMapping) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil.
func (o *ESSecurityRoleMapping) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil.
func (o *ESSecurityRoleMapping) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ESSecurityRoleMapping) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleMapping) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ESSecurityRoleMapping) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ESSecurityRoleMapping) SetRoles(v []string) {
	o.Roles = v
}

// GetRoleTemplates returns the RoleTemplates field value if set, zero value otherwise.
func (o *ESSecurityRoleMapping) GetRoleTemplates() []ESSecurityRoleTemplate {
	if o == nil || o.RoleTemplates == nil {
		var ret []ESSecurityRoleTemplate
		return ret
	}
	return o.RoleTemplates
}

// GetRoleTemplatesOk returns a tuple with the RoleTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleMapping) GetRoleTemplatesOk() (*[]ESSecurityRoleTemplate, bool) {
	if o == nil || o.RoleTemplates == nil {
		return nil, false
	}
	return &o.RoleTemplates, true
}

// HasRoleTemplates returns a boolean if a field has been set.
func (o *ESSecurityRoleMapping) HasRoleTemplates() bool {
	return o != nil && o.RoleTemplates != nil
}

// SetRoleTemplates gets a reference to the given []ESSecurityRoleTemplate and assigns it to the RoleTemplates field.
func (o *ESSecurityRoleMapping) SetRoleTemplates(v []ESSecurityRoleTemplate) {
	o.RoleTemplates = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ESSecurityRoleMapping) GetRules() map[string]interface{} {
	if o == nil || o.Rules == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleMapping) GetRulesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ESSecurityRoleMapping) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given map[string]interface{} and assigns it to the Rules field.
func (o *ESSecurityRoleMapping) SetRules(v map[string]interface{}) {
	o.Rules = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ESSecurityRoleMapping) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleMapping) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ESSecurityRoleMapping) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ESSecurityRoleMapping) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityRoleMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.RoleTemplates != nil {
		toSerialize["role_templates"] = o.RoleTemplates
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityRoleMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       common.NullableBool      `json:"enabled,omitempty"`
		Roles         []string                 `json:"roles,omitempty"`
		RoleTemplates []ESSecurityRoleTemplate `json:"role_templates,omitempty"`
		Rules         map[string]interface{}   `json:"rules,omitempty"`
		Metadata      map[string]interface{}   `json:"metadata,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "roles", "role_templates", "rules", "metadata"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Roles = all.Roles
	o.RoleTemplates = all.RoleTemplates
	o.Rules = all.Rules
	o.Metadata = all.Metadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
