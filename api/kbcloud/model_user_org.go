// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// UserOrg UserOrg is the payload for user organization
type UserOrg struct {
	Status         OrgStatus           `json:"status"`
	ExpiredAt      common.NullableTime `json:"expiredAt"`
	PurgeError     *string             `json:"purgeError,omitempty"`
	AllowedActions []string            `json:"allowedActions"`
	// The ID of the organization
	OrgId string `json:"orgId"`
	// The role of the user in the organization
	Role string `json:"role"`
	// The display name of the organization
	DisplayName *string `json:"displayName,omitempty"`
	// The name of the organization
	Name string `json:"name"`
	// The description of the organization
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserOrg instantiates a new UserOrg object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserOrg(status OrgStatus, expiredAt common.NullableTime, allowedActions []string, orgId string, role string, name string) *UserOrg {
	this := UserOrg{}
	this.Status = status
	this.ExpiredAt = expiredAt
	this.AllowedActions = allowedActions
	this.OrgId = orgId
	this.Role = role
	this.Name = name
	return &this
}

// NewUserOrgWithDefaults instantiates a new UserOrg object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserOrgWithDefaults() *UserOrg {
	this := UserOrg{}
	return &this
}

// GetStatus returns the Status field value.
func (o *UserOrg) GetStatus() OrgStatus {
	if o == nil {
		var ret OrgStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserOrg) GetStatusOk() (*OrgStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *UserOrg) SetStatus(v OrgStatus) {
	o.Status = v
}

// GetExpiredAt returns the ExpiredAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *UserOrg) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserOrg) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// SetExpiredAt sets field value.
func (o *UserOrg) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}

// GetPurgeError returns the PurgeError field value if set, zero value otherwise.
func (o *UserOrg) GetPurgeError() string {
	if o == nil || o.PurgeError == nil {
		var ret string
		return ret
	}
	return *o.PurgeError
}

// GetPurgeErrorOk returns a tuple with the PurgeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOrg) GetPurgeErrorOk() (*string, bool) {
	if o == nil || o.PurgeError == nil {
		return nil, false
	}
	return o.PurgeError, true
}

// HasPurgeError returns a boolean if a field has been set.
func (o *UserOrg) HasPurgeError() bool {
	return o != nil && o.PurgeError != nil
}

// SetPurgeError gets a reference to the given string and assigns it to the PurgeError field.
func (o *UserOrg) SetPurgeError(v string) {
	o.PurgeError = &v
}

// GetAllowedActions returns the AllowedActions field value.
func (o *UserOrg) GetAllowedActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value
// and a boolean to check if the value has been set.
func (o *UserOrg) GetAllowedActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedActions, true
}

// SetAllowedActions sets field value.
func (o *UserOrg) SetAllowedActions(v []string) {
	o.AllowedActions = v
}

// GetOrgId returns the OrgId field value.
func (o *UserOrg) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UserOrg) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *UserOrg) SetOrgId(v string) {
	o.OrgId = v
}

// GetRole returns the Role field value.
func (o *UserOrg) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserOrg) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *UserOrg) SetRole(v string) {
	o.Role = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserOrg) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOrg) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserOrg) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserOrg) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value.
func (o *UserOrg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserOrg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UserOrg) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserOrg) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOrg) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserOrg) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserOrg) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	toSerialize["expiredAt"] = o.ExpiredAt.Get()
	if o.PurgeError != nil {
		toSerialize["purgeError"] = o.PurgeError
	}
	toSerialize["allowedActions"] = o.AllowedActions
	toSerialize["orgId"] = o.OrgId
	toSerialize["role"] = o.Role
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserOrg) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status         *OrgStatus          `json:"status"`
		ExpiredAt      common.NullableTime `json:"expiredAt"`
		PurgeError     *string             `json:"purgeError,omitempty"`
		AllowedActions *[]string           `json:"allowedActions"`
		OrgId          *string             `json:"orgId"`
		Role           *string             `json:"role"`
		DisplayName    *string             `json:"displayName,omitempty"`
		Name           *string             `json:"name"`
		Description    *string             `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.ExpiredAt.IsSet() {
		return fmt.Errorf("required field expiredAt missing")
	}
	if all.AllowedActions == nil {
		return fmt.Errorf("required field allowedActions missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgId missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "expiredAt", "purgeError", "allowedActions", "orgId", "role", "displayName", "name", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.ExpiredAt = all.ExpiredAt
	o.PurgeError = all.PurgeError
	o.AllowedActions = *all.AllowedActions
	o.OrgId = *all.OrgId
	o.Role = *all.Role
	o.DisplayName = all.DisplayName
	o.Name = *all.Name
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
