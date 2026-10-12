// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TablePrivilegeChange struct {
	Target TablePrivilegeTarget `json:"target"`
	// Operation, either grant or revoke.
	Action string `json:"action"`
	// Required unless revoking the MySQL table-wide GRANT OPTION.
	Privileges []TablePrivilegeSpec `json:"privileges,omitempty"`
	// PostgreSQL grant only. Explicitly grant database CONNECT and schema USAGE in the same transaction; never grants sequence privileges. Table/column revocation does not revoke these prerequisites.
	GrantPrerequisites *bool `json:"grantPrerequisites,omitempty"`
	// Grant with delegation rights. MySQL applies GRANT OPTION to the whole table, including existing column/table privileges.
	GrantOption *bool `json:"grantOption,omitempty"`
	// Revoke delegation only. PostgreSQL applies to listed privileges; MySQL requires an empty privileges list and revokes the table-wide GRANT OPTION.
	GrantOptionOnly *bool `json:"grantOptionOnly,omitempty"`
	// PostgreSQL revoke only; defaults to RESTRICT. Explicitly allows revoking dependent grants.
	Cascade *bool `json:"cascade,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTablePrivilegeChange instantiates a new TablePrivilegeChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTablePrivilegeChange(target TablePrivilegeTarget, action string) *TablePrivilegeChange {
	this := TablePrivilegeChange{}
	this.Target = target
	this.Action = action
	return &this
}

// NewTablePrivilegeChangeWithDefaults instantiates a new TablePrivilegeChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTablePrivilegeChangeWithDefaults() *TablePrivilegeChange {
	this := TablePrivilegeChange{}
	return &this
}

// GetTarget returns the Target field value.
func (o *TablePrivilegeChange) GetTarget() TablePrivilegeTarget {
	if o == nil {
		var ret TablePrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetTargetOk() (*TablePrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *TablePrivilegeChange) SetTarget(v TablePrivilegeTarget) {
	o.Target = v
}

// GetAction returns the Action field value.
func (o *TablePrivilegeChange) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *TablePrivilegeChange) SetAction(v string) {
	o.Action = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *TablePrivilegeChange) GetPrivileges() []TablePrivilegeSpec {
	if o == nil || o.Privileges == nil {
		var ret []TablePrivilegeSpec
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetPrivilegesOk() (*[]TablePrivilegeSpec, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *TablePrivilegeChange) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given []TablePrivilegeSpec and assigns it to the Privileges field.
func (o *TablePrivilegeChange) SetPrivileges(v []TablePrivilegeSpec) {
	o.Privileges = v
}

// GetGrantPrerequisites returns the GrantPrerequisites field value if set, zero value otherwise.
func (o *TablePrivilegeChange) GetGrantPrerequisites() bool {
	if o == nil || o.GrantPrerequisites == nil {
		var ret bool
		return ret
	}
	return *o.GrantPrerequisites
}

// GetGrantPrerequisitesOk returns a tuple with the GrantPrerequisites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetGrantPrerequisitesOk() (*bool, bool) {
	if o == nil || o.GrantPrerequisites == nil {
		return nil, false
	}
	return o.GrantPrerequisites, true
}

// HasGrantPrerequisites returns a boolean if a field has been set.
func (o *TablePrivilegeChange) HasGrantPrerequisites() bool {
	return o != nil && o.GrantPrerequisites != nil
}

// SetGrantPrerequisites gets a reference to the given bool and assigns it to the GrantPrerequisites field.
func (o *TablePrivilegeChange) SetGrantPrerequisites(v bool) {
	o.GrantPrerequisites = &v
}

// GetGrantOption returns the GrantOption field value if set, zero value otherwise.
func (o *TablePrivilegeChange) GetGrantOption() bool {
	if o == nil || o.GrantOption == nil {
		var ret bool
		return ret
	}
	return *o.GrantOption
}

// GetGrantOptionOk returns a tuple with the GrantOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetGrantOptionOk() (*bool, bool) {
	if o == nil || o.GrantOption == nil {
		return nil, false
	}
	return o.GrantOption, true
}

// HasGrantOption returns a boolean if a field has been set.
func (o *TablePrivilegeChange) HasGrantOption() bool {
	return o != nil && o.GrantOption != nil
}

// SetGrantOption gets a reference to the given bool and assigns it to the GrantOption field.
func (o *TablePrivilegeChange) SetGrantOption(v bool) {
	o.GrantOption = &v
}

// GetGrantOptionOnly returns the GrantOptionOnly field value if set, zero value otherwise.
func (o *TablePrivilegeChange) GetGrantOptionOnly() bool {
	if o == nil || o.GrantOptionOnly == nil {
		var ret bool
		return ret
	}
	return *o.GrantOptionOnly
}

// GetGrantOptionOnlyOk returns a tuple with the GrantOptionOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetGrantOptionOnlyOk() (*bool, bool) {
	if o == nil || o.GrantOptionOnly == nil {
		return nil, false
	}
	return o.GrantOptionOnly, true
}

// HasGrantOptionOnly returns a boolean if a field has been set.
func (o *TablePrivilegeChange) HasGrantOptionOnly() bool {
	return o != nil && o.GrantOptionOnly != nil
}

// SetGrantOptionOnly gets a reference to the given bool and assigns it to the GrantOptionOnly field.
func (o *TablePrivilegeChange) SetGrantOptionOnly(v bool) {
	o.GrantOptionOnly = &v
}

// GetCascade returns the Cascade field value if set, zero value otherwise.
func (o *TablePrivilegeChange) GetCascade() bool {
	if o == nil || o.Cascade == nil {
		var ret bool
		return ret
	}
	return *o.Cascade
}

// GetCascadeOk returns a tuple with the Cascade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeChange) GetCascadeOk() (*bool, bool) {
	if o == nil || o.Cascade == nil {
		return nil, false
	}
	return o.Cascade, true
}

// HasCascade returns a boolean if a field has been set.
func (o *TablePrivilegeChange) HasCascade() bool {
	return o != nil && o.Cascade != nil
}

// SetCascade gets a reference to the given bool and assigns it to the Cascade field.
func (o *TablePrivilegeChange) SetCascade(v bool) {
	o.Cascade = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TablePrivilegeChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["action"] = o.Action
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.GrantPrerequisites != nil {
		toSerialize["grantPrerequisites"] = o.GrantPrerequisites
	}
	if o.GrantOption != nil {
		toSerialize["grantOption"] = o.GrantOption
	}
	if o.GrantOptionOnly != nil {
		toSerialize["grantOptionOnly"] = o.GrantOptionOnly
	}
	if o.Cascade != nil {
		toSerialize["cascade"] = o.Cascade
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TablePrivilegeChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target             *TablePrivilegeTarget `json:"target"`
		Action             *string               `json:"action"`
		Privileges         []TablePrivilegeSpec  `json:"privileges,omitempty"`
		GrantPrerequisites *bool                 `json:"grantPrerequisites,omitempty"`
		GrantOption        *bool                 `json:"grantOption,omitempty"`
		GrantOptionOnly    *bool                 `json:"grantOptionOnly,omitempty"`
		Cascade            *bool                 `json:"cascade,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"target", "action", "privileges", "grantPrerequisites", "grantOption", "grantOptionOnly", "cascade"})
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
	o.GrantPrerequisites = all.GrantPrerequisites
	o.GrantOption = all.GrantOption
	o.GrantOptionOnly = all.GrantOptionOnly
	o.Cascade = all.Cascade

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
