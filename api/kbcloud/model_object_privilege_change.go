// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeChange struct {
	Target ObjectPrivilegeTarget `json:"target"`
	// Operation, either grant or revoke.
	Action string `json:"action"`
	// Required except when revoking MySQL scope-wide GRANT OPTION. Dynamic global privileges require explicit names.
	Privileges []ObjectPrivilegeSpec `json:"privileges,omitempty"`
	// Grant delegation. MySQL static GRANT OPTION is scope-wide; dynamic global privileges each have their own grant option.
	GrantOption *bool `json:"grantOption,omitempty"`
	// Revoke delegation only. PostgreSQL applies to listed privileges. MySQL requires an empty privileges list and revokes GRANT OPTION at the entire scope; global dynamic grant options can also be removed.
	GrantOptionOnly *bool `json:"grantOptionOnly,omitempty"`
	// Versioned namespaced JSON extensions. Response keys unknown to a client may be ignored. Change keys must be registered and strictly decoded by the current engine; PG/MySQL currently accept no change extensions.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// Explicitly grant advertised access prerequisites. PG schema-contained grants add CONNECT and schema USAGE in the same transaction; never sequence privileges. Defaults false; later revoke does not undo prerequisites.
	GrantPrerequisites *bool `json:"grantPrerequisites,omitempty"`
	// Explicit native revoke behavior from capabilities; valid only for revoke. Omitted uses the advertised default.
	RevokeBehavior *string `json:"revokeBehavior,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeChange instantiates a new ObjectPrivilegeChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeChange(target ObjectPrivilegeTarget, action string) *ObjectPrivilegeChange {
	this := ObjectPrivilegeChange{}
	this.Target = target
	this.Action = action
	return &this
}

// NewObjectPrivilegeChangeWithDefaults instantiates a new ObjectPrivilegeChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeChangeWithDefaults() *ObjectPrivilegeChange {
	this := ObjectPrivilegeChange{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ObjectPrivilegeChange) GetTarget() ObjectPrivilegeTarget {
	if o == nil {
		var ret ObjectPrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetTargetOk() (*ObjectPrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObjectPrivilegeChange) SetTarget(v ObjectPrivilegeTarget) {
	o.Target = v
}

// GetAction returns the Action field value.
func (o *ObjectPrivilegeChange) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObjectPrivilegeChange) SetAction(v string) {
	o.Action = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetPrivileges() []ObjectPrivilegeSpec {
	if o == nil || o.Privileges == nil {
		var ret []ObjectPrivilegeSpec
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetPrivilegesOk() (*[]ObjectPrivilegeSpec, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given []ObjectPrivilegeSpec and assigns it to the Privileges field.
func (o *ObjectPrivilegeChange) SetPrivileges(v []ObjectPrivilegeSpec) {
	o.Privileges = v
}

// GetGrantOption returns the GrantOption field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetGrantOption() bool {
	if o == nil || o.GrantOption == nil {
		var ret bool
		return ret
	}
	return *o.GrantOption
}

// GetGrantOptionOk returns a tuple with the GrantOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetGrantOptionOk() (*bool, bool) {
	if o == nil || o.GrantOption == nil {
		return nil, false
	}
	return o.GrantOption, true
}

// HasGrantOption returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasGrantOption() bool {
	return o != nil && o.GrantOption != nil
}

// SetGrantOption gets a reference to the given bool and assigns it to the GrantOption field.
func (o *ObjectPrivilegeChange) SetGrantOption(v bool) {
	o.GrantOption = &v
}

// GetGrantOptionOnly returns the GrantOptionOnly field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetGrantOptionOnly() bool {
	if o == nil || o.GrantOptionOnly == nil {
		var ret bool
		return ret
	}
	return *o.GrantOptionOnly
}

// GetGrantOptionOnlyOk returns a tuple with the GrantOptionOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetGrantOptionOnlyOk() (*bool, bool) {
	if o == nil || o.GrantOptionOnly == nil {
		return nil, false
	}
	return o.GrantOptionOnly, true
}

// HasGrantOptionOnly returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasGrantOptionOnly() bool {
	return o != nil && o.GrantOptionOnly != nil
}

// SetGrantOptionOnly gets a reference to the given bool and assigns it to the GrantOptionOnly field.
func (o *ObjectPrivilegeChange) SetGrantOptionOnly(v bool) {
	o.GrantOptionOnly = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *ObjectPrivilegeChange) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetGrantPrerequisites returns the GrantPrerequisites field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetGrantPrerequisites() bool {
	if o == nil || o.GrantPrerequisites == nil {
		var ret bool
		return ret
	}
	return *o.GrantPrerequisites
}

// GetGrantPrerequisitesOk returns a tuple with the GrantPrerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetGrantPrerequisitesOk() (*bool, bool) {
	if o == nil || o.GrantPrerequisites == nil {
		return nil, false
	}
	return o.GrantPrerequisites, true
}

// HasGrantPrerequisites returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasGrantPrerequisites() bool {
	return o != nil && o.GrantPrerequisites != nil
}

// SetGrantPrerequisites gets a reference to the given bool and assigns it to the GrantPrerequisites field.
func (o *ObjectPrivilegeChange) SetGrantPrerequisites(v bool) {
	o.GrantPrerequisites = &v
}

// GetRevokeBehavior returns the RevokeBehavior field value if set, zero value otherwise.
func (o *ObjectPrivilegeChange) GetRevokeBehavior() string {
	if o == nil || o.RevokeBehavior == nil {
		var ret string
		return ret
	}
	return *o.RevokeBehavior
}

// GetRevokeBehaviorOk returns a tuple with the RevokeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeChange) GetRevokeBehaviorOk() (*string, bool) {
	if o == nil || o.RevokeBehavior == nil {
		return nil, false
	}
	return o.RevokeBehavior, true
}

// HasRevokeBehavior returns a boolean if a field has been set.
func (o *ObjectPrivilegeChange) HasRevokeBehavior() bool {
	return o != nil && o.RevokeBehavior != nil
}

// SetRevokeBehavior gets a reference to the given string and assigns it to the RevokeBehavior field.
func (o *ObjectPrivilegeChange) SetRevokeBehavior(v string) {
	o.RevokeBehavior = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["action"] = o.Action
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.GrantOption != nil {
		toSerialize["grantOption"] = o.GrantOption
	}
	if o.GrantOptionOnly != nil {
		toSerialize["grantOptionOnly"] = o.GrantOptionOnly
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.GrantPrerequisites != nil {
		toSerialize["grantPrerequisites"] = o.GrantPrerequisites
	}
	if o.RevokeBehavior != nil {
		toSerialize["revokeBehavior"] = o.RevokeBehavior
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target             *ObjectPrivilegeTarget `json:"target"`
		Action             *string                `json:"action"`
		Privileges         []ObjectPrivilegeSpec  `json:"privileges,omitempty"`
		GrantOption        *bool                  `json:"grantOption,omitempty"`
		GrantOptionOnly    *bool                  `json:"grantOptionOnly,omitempty"`
		Extensions         map[string]interface{} `json:"extensions,omitempty"`
		GrantPrerequisites *bool                  `json:"grantPrerequisites,omitempty"`
		RevokeBehavior     *string                `json:"revokeBehavior,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"target", "action", "privileges", "grantOption", "grantOptionOnly", "extensions", "grantPrerequisites", "revokeBehavior"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Action = *all.Action
	o.Privileges = all.Privileges
	o.GrantOption = all.GrantOption
	o.GrantOptionOnly = all.GrantOptionOnly
	o.Extensions = all.Extensions
	o.GrantPrerequisites = all.GrantPrerequisites
	o.RevokeBehavior = all.RevokeBehavior

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
