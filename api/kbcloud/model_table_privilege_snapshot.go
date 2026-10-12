// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TablePrivilegeSnapshot struct {
	Target                    TablePrivilegeTarget `json:"target"`
	Username                  string               `json:"username"`
	SupportedTablePrivileges  []string             `json:"supportedTablePrivileges"`
	SupportedColumnPrivileges []string             `json:"supportedColumnPrivileges"`
	// Exact table/column grants to this account. Does not represent all effective access.
	DirectPrivileges []TablePrivilegeGrant `json:"directPrivileges"`
	// True for PostgreSQL native privilege checks. False for MySQL because active roles depend on the login session; do not interpret an omitted effectivePrivileges list as no access.
	EffectivePrivilegesKnown bool `json:"effectivePrivilegesKnown"`
	// PostgreSQL has_table_privilege/has_column_privilege including inherited and PUBLIC grants; excludes RLS, session SET ROLE and other execution prerequisites.
	EffectivePrivileges []TablePrivilegeGrant `json:"effectivePrivileges,omitempty"`
	// PostgreSQL CONNECT prerequisite; omitted for MySQL.
	DatabaseConnect *bool `json:"databaseConnect,omitempty"`
	// PostgreSQL USAGE prerequisite; omitted for MySQL.
	SchemaUsage *bool `json:"schemaUsage,omitempty"`
	// MySQL SHOW GRANTS output including database/global grants and role memberships; roles are not expanded or assumed active.
	AccountGrants []string `json:"accountGrants,omitempty"`
	Warnings      []string `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTablePrivilegeSnapshot instantiates a new TablePrivilegeSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTablePrivilegeSnapshot(target TablePrivilegeTarget, username string, supportedTablePrivileges []string, supportedColumnPrivileges []string, directPrivileges []TablePrivilegeGrant, effectivePrivilegesKnown bool) *TablePrivilegeSnapshot {
	this := TablePrivilegeSnapshot{}
	this.Target = target
	this.Username = username
	this.SupportedTablePrivileges = supportedTablePrivileges
	this.SupportedColumnPrivileges = supportedColumnPrivileges
	this.DirectPrivileges = directPrivileges
	this.EffectivePrivilegesKnown = effectivePrivilegesKnown
	return &this
}

// NewTablePrivilegeSnapshotWithDefaults instantiates a new TablePrivilegeSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTablePrivilegeSnapshotWithDefaults() *TablePrivilegeSnapshot {
	this := TablePrivilegeSnapshot{}
	return &this
}

// GetTarget returns the Target field value.
func (o *TablePrivilegeSnapshot) GetTarget() TablePrivilegeTarget {
	if o == nil {
		var ret TablePrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetTargetOk() (*TablePrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *TablePrivilegeSnapshot) SetTarget(v TablePrivilegeTarget) {
	o.Target = v
}

// GetUsername returns the Username field value.
func (o *TablePrivilegeSnapshot) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *TablePrivilegeSnapshot) SetUsername(v string) {
	o.Username = v
}

// GetSupportedTablePrivileges returns the SupportedTablePrivileges field value.
func (o *TablePrivilegeSnapshot) GetSupportedTablePrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedTablePrivileges
}

// GetSupportedTablePrivilegesOk returns a tuple with the SupportedTablePrivileges field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetSupportedTablePrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedTablePrivileges, true
}

// SetSupportedTablePrivileges sets field value.
func (o *TablePrivilegeSnapshot) SetSupportedTablePrivileges(v []string) {
	o.SupportedTablePrivileges = v
}

// GetSupportedColumnPrivileges returns the SupportedColumnPrivileges field value.
func (o *TablePrivilegeSnapshot) GetSupportedColumnPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedColumnPrivileges
}

// GetSupportedColumnPrivilegesOk returns a tuple with the SupportedColumnPrivileges field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetSupportedColumnPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedColumnPrivileges, true
}

// SetSupportedColumnPrivileges sets field value.
func (o *TablePrivilegeSnapshot) SetSupportedColumnPrivileges(v []string) {
	o.SupportedColumnPrivileges = v
}

// GetDirectPrivileges returns the DirectPrivileges field value.
func (o *TablePrivilegeSnapshot) GetDirectPrivileges() []TablePrivilegeGrant {
	if o == nil {
		var ret []TablePrivilegeGrant
		return ret
	}
	return o.DirectPrivileges
}

// GetDirectPrivilegesOk returns a tuple with the DirectPrivileges field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetDirectPrivilegesOk() (*[]TablePrivilegeGrant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectPrivileges, true
}

// SetDirectPrivileges sets field value.
func (o *TablePrivilegeSnapshot) SetDirectPrivileges(v []TablePrivilegeGrant) {
	o.DirectPrivileges = v
}

// GetEffectivePrivilegesKnown returns the EffectivePrivilegesKnown field value.
func (o *TablePrivilegeSnapshot) GetEffectivePrivilegesKnown() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EffectivePrivilegesKnown
}

// GetEffectivePrivilegesKnownOk returns a tuple with the EffectivePrivilegesKnown field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetEffectivePrivilegesKnownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectivePrivilegesKnown, true
}

// SetEffectivePrivilegesKnown sets field value.
func (o *TablePrivilegeSnapshot) SetEffectivePrivilegesKnown(v bool) {
	o.EffectivePrivilegesKnown = v
}

// GetEffectivePrivileges returns the EffectivePrivileges field value if set, zero value otherwise.
func (o *TablePrivilegeSnapshot) GetEffectivePrivileges() []TablePrivilegeGrant {
	if o == nil || o.EffectivePrivileges == nil {
		var ret []TablePrivilegeGrant
		return ret
	}
	return o.EffectivePrivileges
}

// GetEffectivePrivilegesOk returns a tuple with the EffectivePrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetEffectivePrivilegesOk() (*[]TablePrivilegeGrant, bool) {
	if o == nil || o.EffectivePrivileges == nil {
		return nil, false
	}
	return &o.EffectivePrivileges, true
}

// HasEffectivePrivileges returns a boolean if a field has been set.
func (o *TablePrivilegeSnapshot) HasEffectivePrivileges() bool {
	return o != nil && o.EffectivePrivileges != nil
}

// SetEffectivePrivileges gets a reference to the given []TablePrivilegeGrant and assigns it to the EffectivePrivileges field.
func (o *TablePrivilegeSnapshot) SetEffectivePrivileges(v []TablePrivilegeGrant) {
	o.EffectivePrivileges = v
}

// GetDatabaseConnect returns the DatabaseConnect field value if set, zero value otherwise.
func (o *TablePrivilegeSnapshot) GetDatabaseConnect() bool {
	if o == nil || o.DatabaseConnect == nil {
		var ret bool
		return ret
	}
	return *o.DatabaseConnect
}

// GetDatabaseConnectOk returns a tuple with the DatabaseConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetDatabaseConnectOk() (*bool, bool) {
	if o == nil || o.DatabaseConnect == nil {
		return nil, false
	}
	return o.DatabaseConnect, true
}

// HasDatabaseConnect returns a boolean if a field has been set.
func (o *TablePrivilegeSnapshot) HasDatabaseConnect() bool {
	return o != nil && o.DatabaseConnect != nil
}

// SetDatabaseConnect gets a reference to the given bool and assigns it to the DatabaseConnect field.
func (o *TablePrivilegeSnapshot) SetDatabaseConnect(v bool) {
	o.DatabaseConnect = &v
}

// GetSchemaUsage returns the SchemaUsage field value if set, zero value otherwise.
func (o *TablePrivilegeSnapshot) GetSchemaUsage() bool {
	if o == nil || o.SchemaUsage == nil {
		var ret bool
		return ret
	}
	return *o.SchemaUsage
}

// GetSchemaUsageOk returns a tuple with the SchemaUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetSchemaUsageOk() (*bool, bool) {
	if o == nil || o.SchemaUsage == nil {
		return nil, false
	}
	return o.SchemaUsage, true
}

// HasSchemaUsage returns a boolean if a field has been set.
func (o *TablePrivilegeSnapshot) HasSchemaUsage() bool {
	return o != nil && o.SchemaUsage != nil
}

// SetSchemaUsage gets a reference to the given bool and assigns it to the SchemaUsage field.
func (o *TablePrivilegeSnapshot) SetSchemaUsage(v bool) {
	o.SchemaUsage = &v
}

// GetAccountGrants returns the AccountGrants field value if set, zero value otherwise.
func (o *TablePrivilegeSnapshot) GetAccountGrants() []string {
	if o == nil || o.AccountGrants == nil {
		var ret []string
		return ret
	}
	return o.AccountGrants
}

// GetAccountGrantsOk returns a tuple with the AccountGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetAccountGrantsOk() (*[]string, bool) {
	if o == nil || o.AccountGrants == nil {
		return nil, false
	}
	return &o.AccountGrants, true
}

// HasAccountGrants returns a boolean if a field has been set.
func (o *TablePrivilegeSnapshot) HasAccountGrants() bool {
	return o != nil && o.AccountGrants != nil
}

// SetAccountGrants gets a reference to the given []string and assigns it to the AccountGrants field.
func (o *TablePrivilegeSnapshot) SetAccountGrants(v []string) {
	o.AccountGrants = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TablePrivilegeSnapshot) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSnapshot) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TablePrivilegeSnapshot) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *TablePrivilegeSnapshot) SetWarnings(v []string) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TablePrivilegeSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["username"] = o.Username
	toSerialize["supportedTablePrivileges"] = o.SupportedTablePrivileges
	toSerialize["supportedColumnPrivileges"] = o.SupportedColumnPrivileges
	toSerialize["directPrivileges"] = o.DirectPrivileges
	toSerialize["effectivePrivilegesKnown"] = o.EffectivePrivilegesKnown
	if o.EffectivePrivileges != nil {
		toSerialize["effectivePrivileges"] = o.EffectivePrivileges
	}
	if o.DatabaseConnect != nil {
		toSerialize["databaseConnect"] = o.DatabaseConnect
	}
	if o.SchemaUsage != nil {
		toSerialize["schemaUsage"] = o.SchemaUsage
	}
	if o.AccountGrants != nil {
		toSerialize["accountGrants"] = o.AccountGrants
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TablePrivilegeSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target                    *TablePrivilegeTarget  `json:"target"`
		Username                  *string                `json:"username"`
		SupportedTablePrivileges  *[]string              `json:"supportedTablePrivileges"`
		SupportedColumnPrivileges *[]string              `json:"supportedColumnPrivileges"`
		DirectPrivileges          *[]TablePrivilegeGrant `json:"directPrivileges"`
		EffectivePrivilegesKnown  *bool                  `json:"effectivePrivilegesKnown"`
		EffectivePrivileges       []TablePrivilegeGrant  `json:"effectivePrivileges,omitempty"`
		DatabaseConnect           *bool                  `json:"databaseConnect,omitempty"`
		SchemaUsage               *bool                  `json:"schemaUsage,omitempty"`
		AccountGrants             []string               `json:"accountGrants,omitempty"`
		Warnings                  []string               `json:"warnings,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	if all.SupportedTablePrivileges == nil {
		return fmt.Errorf("required field supportedTablePrivileges missing")
	}
	if all.SupportedColumnPrivileges == nil {
		return fmt.Errorf("required field supportedColumnPrivileges missing")
	}
	if all.DirectPrivileges == nil {
		return fmt.Errorf("required field directPrivileges missing")
	}
	if all.EffectivePrivilegesKnown == nil {
		return fmt.Errorf("required field effectivePrivilegesKnown missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"target", "username", "supportedTablePrivileges", "supportedColumnPrivileges", "directPrivileges", "effectivePrivilegesKnown", "effectivePrivileges", "databaseConnect", "schemaUsage", "accountGrants", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Username = *all.Username
	o.SupportedTablePrivileges = *all.SupportedTablePrivileges
	o.SupportedColumnPrivileges = *all.SupportedColumnPrivileges
	o.DirectPrivileges = *all.DirectPrivileges
	o.EffectivePrivilegesKnown = *all.EffectivePrivilegesKnown
	o.EffectivePrivileges = all.EffectivePrivileges
	o.DatabaseConnect = all.DatabaseConnect
	o.SchemaUsage = all.SchemaUsage
	o.AccountGrants = all.AccountGrants
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
