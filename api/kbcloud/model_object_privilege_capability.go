// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeCapability struct {
	ObjectType           string                      `json:"objectType"`
	Privileges           []ObjectPrivilegeDefinition `json:"privileges"`
	RequiredTargetFields []string                    `json:"requiredTargetFields"`
	Warnings             []string                    `json:"warnings,omitempty"`
	// Fields required to enumerate objects of this kind; distinct from fields needed to identify an individual target.
	RequiredDiscoveryFields []string                         `json:"requiredDiscoveryFields"`
	RevokeGrantOptionScope  *ObjectPrivilegeGrantOptionScope `json:"revokeGrantOptionScope,omitempty"`
	// Versioned namespaced JSON extensions. Response keys unknown to a client may be ignored. Change keys must be registered and strictly decoded by the current engine; PG/MySQL currently accept no change extensions.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// Allows explicitly granting the documented access prerequisites together with an object grant.
	SupportsGrantPrerequisites *bool `json:"supportsGrantPrerequisites,omitempty"`
	// When prerequisite grants are supported, whether they and the requested grant are atomic.
	GrantPrerequisitesAtomic *bool `json:"grantPrerequisitesAtomic,omitempty"`
	// Explicit revoke behaviors supported by this object kind. Omitted change behavior uses defaultRevokeBehavior.
	RevokeBehaviors []string `json:"revokeBehaviors"`
	// Default native revoke behavior. PostgreSQL restrict; MySQL native.
	DefaultRevokeBehavior string `json:"defaultRevokeBehavior"`
	// Versioned namespaces accepted for change extensions by this engine/version; empty means none.
	ChangeExtensions []string `json:"changeExtensions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeCapability instantiates a new ObjectPrivilegeCapability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeCapability(objectType string, privileges []ObjectPrivilegeDefinition, requiredTargetFields []string, requiredDiscoveryFields []string, revokeBehaviors []string, defaultRevokeBehavior string) *ObjectPrivilegeCapability {
	this := ObjectPrivilegeCapability{}
	this.ObjectType = objectType
	this.Privileges = privileges
	this.RequiredTargetFields = requiredTargetFields
	this.RequiredDiscoveryFields = requiredDiscoveryFields
	this.RevokeBehaviors = revokeBehaviors
	this.DefaultRevokeBehavior = defaultRevokeBehavior
	return &this
}

// NewObjectPrivilegeCapabilityWithDefaults instantiates a new ObjectPrivilegeCapability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeCapabilityWithDefaults() *ObjectPrivilegeCapability {
	this := ObjectPrivilegeCapability{}
	return &this
}

// GetObjectType returns the ObjectType field value.
func (o *ObjectPrivilegeCapability) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value.
func (o *ObjectPrivilegeCapability) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPrivileges returns the Privileges field value.
func (o *ObjectPrivilegeCapability) GetPrivileges() []ObjectPrivilegeDefinition {
	if o == nil {
		var ret []ObjectPrivilegeDefinition
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetPrivilegesOk() (*[]ObjectPrivilegeDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *ObjectPrivilegeCapability) SetPrivileges(v []ObjectPrivilegeDefinition) {
	o.Privileges = v
}

// GetRequiredTargetFields returns the RequiredTargetFields field value.
func (o *ObjectPrivilegeCapability) GetRequiredTargetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RequiredTargetFields
}

// GetRequiredTargetFieldsOk returns a tuple with the RequiredTargetFields field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetRequiredTargetFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredTargetFields, true
}

// SetRequiredTargetFields sets field value.
func (o *ObjectPrivilegeCapability) SetRequiredTargetFields(v []string) {
	o.RequiredTargetFields = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ObjectPrivilegeCapability) SetWarnings(v []string) {
	o.Warnings = v
}

// GetRequiredDiscoveryFields returns the RequiredDiscoveryFields field value.
func (o *ObjectPrivilegeCapability) GetRequiredDiscoveryFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RequiredDiscoveryFields
}

// GetRequiredDiscoveryFieldsOk returns a tuple with the RequiredDiscoveryFields field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetRequiredDiscoveryFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredDiscoveryFields, true
}

// SetRequiredDiscoveryFields sets field value.
func (o *ObjectPrivilegeCapability) SetRequiredDiscoveryFields(v []string) {
	o.RequiredDiscoveryFields = v
}

// GetRevokeGrantOptionScope returns the RevokeGrantOptionScope field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetRevokeGrantOptionScope() ObjectPrivilegeGrantOptionScope {
	if o == nil || o.RevokeGrantOptionScope == nil {
		var ret ObjectPrivilegeGrantOptionScope
		return ret
	}
	return *o.RevokeGrantOptionScope
}

// GetRevokeGrantOptionScopeOk returns a tuple with the RevokeGrantOptionScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetRevokeGrantOptionScopeOk() (*ObjectPrivilegeGrantOptionScope, bool) {
	if o == nil || o.RevokeGrantOptionScope == nil {
		return nil, false
	}
	return o.RevokeGrantOptionScope, true
}

// HasRevokeGrantOptionScope returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasRevokeGrantOptionScope() bool {
	return o != nil && o.RevokeGrantOptionScope != nil
}

// SetRevokeGrantOptionScope gets a reference to the given ObjectPrivilegeGrantOptionScope and assigns it to the RevokeGrantOptionScope field.
func (o *ObjectPrivilegeCapability) SetRevokeGrantOptionScope(v ObjectPrivilegeGrantOptionScope) {
	o.RevokeGrantOptionScope = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *ObjectPrivilegeCapability) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetSupportsGrantPrerequisites returns the SupportsGrantPrerequisites field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetSupportsGrantPrerequisites() bool {
	if o == nil || o.SupportsGrantPrerequisites == nil {
		var ret bool
		return ret
	}
	return *o.SupportsGrantPrerequisites
}

// GetSupportsGrantPrerequisitesOk returns a tuple with the SupportsGrantPrerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetSupportsGrantPrerequisitesOk() (*bool, bool) {
	if o == nil || o.SupportsGrantPrerequisites == nil {
		return nil, false
	}
	return o.SupportsGrantPrerequisites, true
}

// HasSupportsGrantPrerequisites returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasSupportsGrantPrerequisites() bool {
	return o != nil && o.SupportsGrantPrerequisites != nil
}

// SetSupportsGrantPrerequisites gets a reference to the given bool and assigns it to the SupportsGrantPrerequisites field.
func (o *ObjectPrivilegeCapability) SetSupportsGrantPrerequisites(v bool) {
	o.SupportsGrantPrerequisites = &v
}

// GetGrantPrerequisitesAtomic returns the GrantPrerequisitesAtomic field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetGrantPrerequisitesAtomic() bool {
	if o == nil || o.GrantPrerequisitesAtomic == nil {
		var ret bool
		return ret
	}
	return *o.GrantPrerequisitesAtomic
}

// GetGrantPrerequisitesAtomicOk returns a tuple with the GrantPrerequisitesAtomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetGrantPrerequisitesAtomicOk() (*bool, bool) {
	if o == nil || o.GrantPrerequisitesAtomic == nil {
		return nil, false
	}
	return o.GrantPrerequisitesAtomic, true
}

// HasGrantPrerequisitesAtomic returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasGrantPrerequisitesAtomic() bool {
	return o != nil && o.GrantPrerequisitesAtomic != nil
}

// SetGrantPrerequisitesAtomic gets a reference to the given bool and assigns it to the GrantPrerequisitesAtomic field.
func (o *ObjectPrivilegeCapability) SetGrantPrerequisitesAtomic(v bool) {
	o.GrantPrerequisitesAtomic = &v
}

// GetRevokeBehaviors returns the RevokeBehaviors field value.
func (o *ObjectPrivilegeCapability) GetRevokeBehaviors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RevokeBehaviors
}

// GetRevokeBehaviorsOk returns a tuple with the RevokeBehaviors field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetRevokeBehaviorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevokeBehaviors, true
}

// SetRevokeBehaviors sets field value.
func (o *ObjectPrivilegeCapability) SetRevokeBehaviors(v []string) {
	o.RevokeBehaviors = v
}

// GetDefaultRevokeBehavior returns the DefaultRevokeBehavior field value.
func (o *ObjectPrivilegeCapability) GetDefaultRevokeBehavior() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultRevokeBehavior
}

// GetDefaultRevokeBehaviorOk returns a tuple with the DefaultRevokeBehavior field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetDefaultRevokeBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRevokeBehavior, true
}

// SetDefaultRevokeBehavior sets field value.
func (o *ObjectPrivilegeCapability) SetDefaultRevokeBehavior(v string) {
	o.DefaultRevokeBehavior = v
}

// GetChangeExtensions returns the ChangeExtensions field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapability) GetChangeExtensions() []string {
	if o == nil || o.ChangeExtensions == nil {
		var ret []string
		return ret
	}
	return o.ChangeExtensions
}

// GetChangeExtensionsOk returns a tuple with the ChangeExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapability) GetChangeExtensionsOk() (*[]string, bool) {
	if o == nil || o.ChangeExtensions == nil {
		return nil, false
	}
	return &o.ChangeExtensions, true
}

// HasChangeExtensions returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapability) HasChangeExtensions() bool {
	return o != nil && o.ChangeExtensions != nil
}

// SetChangeExtensions gets a reference to the given []string and assigns it to the ChangeExtensions field.
func (o *ObjectPrivilegeCapability) SetChangeExtensions(v []string) {
	o.ChangeExtensions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["objectType"] = o.ObjectType
	toSerialize["privileges"] = o.Privileges
	toSerialize["requiredTargetFields"] = o.RequiredTargetFields
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["requiredDiscoveryFields"] = o.RequiredDiscoveryFields
	if o.RevokeGrantOptionScope != nil {
		toSerialize["revokeGrantOptionScope"] = o.RevokeGrantOptionScope
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.SupportsGrantPrerequisites != nil {
		toSerialize["supportsGrantPrerequisites"] = o.SupportsGrantPrerequisites
	}
	if o.GrantPrerequisitesAtomic != nil {
		toSerialize["grantPrerequisitesAtomic"] = o.GrantPrerequisitesAtomic
	}
	toSerialize["revokeBehaviors"] = o.RevokeBehaviors
	toSerialize["defaultRevokeBehavior"] = o.DefaultRevokeBehavior
	if o.ChangeExtensions != nil {
		toSerialize["changeExtensions"] = o.ChangeExtensions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeCapability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObjectType                 *string                          `json:"objectType"`
		Privileges                 *[]ObjectPrivilegeDefinition     `json:"privileges"`
		RequiredTargetFields       *[]string                        `json:"requiredTargetFields"`
		Warnings                   []string                         `json:"warnings,omitempty"`
		RequiredDiscoveryFields    *[]string                        `json:"requiredDiscoveryFields"`
		RevokeGrantOptionScope     *ObjectPrivilegeGrantOptionScope `json:"revokeGrantOptionScope,omitempty"`
		Extensions                 map[string]interface{}           `json:"extensions,omitempty"`
		SupportsGrantPrerequisites *bool                            `json:"supportsGrantPrerequisites,omitempty"`
		GrantPrerequisitesAtomic   *bool                            `json:"grantPrerequisitesAtomic,omitempty"`
		RevokeBehaviors            *[]string                        `json:"revokeBehaviors"`
		DefaultRevokeBehavior      *string                          `json:"defaultRevokeBehavior"`
		ChangeExtensions           []string                         `json:"changeExtensions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ObjectType == nil {
		return fmt.Errorf("required field objectType missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	if all.RequiredTargetFields == nil {
		return fmt.Errorf("required field requiredTargetFields missing")
	}
	if all.RequiredDiscoveryFields == nil {
		return fmt.Errorf("required field requiredDiscoveryFields missing")
	}
	if all.RevokeBehaviors == nil {
		return fmt.Errorf("required field revokeBehaviors missing")
	}
	if all.DefaultRevokeBehavior == nil {
		return fmt.Errorf("required field defaultRevokeBehavior missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"objectType", "privileges", "requiredTargetFields", "warnings", "requiredDiscoveryFields", "revokeGrantOptionScope", "extensions", "supportsGrantPrerequisites", "grantPrerequisitesAtomic", "revokeBehaviors", "defaultRevokeBehavior", "changeExtensions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ObjectType = *all.ObjectType
	o.Privileges = *all.Privileges
	o.RequiredTargetFields = *all.RequiredTargetFields
	o.Warnings = all.Warnings
	o.RequiredDiscoveryFields = *all.RequiredDiscoveryFields
	if all.RevokeGrantOptionScope != nil && !all.RevokeGrantOptionScope.IsValid() {
		hasInvalidField = true
	} else {
		o.RevokeGrantOptionScope = all.RevokeGrantOptionScope
	}
	o.Extensions = all.Extensions
	o.SupportsGrantPrerequisites = all.SupportsGrantPrerequisites
	o.GrantPrerequisitesAtomic = all.GrantPrerequisitesAtomic
	o.RevokeBehaviors = *all.RevokeBehaviors
	o.DefaultRevokeBehavior = *all.DefaultRevokeBehavior
	o.ChangeExtensions = all.ChangeExtensions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
