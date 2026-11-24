// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Org Org info
type Org struct {
	Status         OrgStatus           `json:"status"`
	ExpiredAt      common.NullableTime `json:"expiredAt"`
	PurgeError     *string             `json:"purgeError,omitempty"`
	AllowedActions []string            `json:"allowedActions"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists
	CreatedAt time.Time `json:"createdAt"`
	// The description of the organization
	Description *string `json:"description,omitempty"`
	// The display name of the organization
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the organization
	Id *string `json:"id,omitempty"`
	// The full, unique name of this Object in the format organizations/{name}, set during creation. name must be a valid RFC 1123 compliant DNS label
	Name string `json:"name"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists
	UpdatedAt time.Time `json:"updatedAt"`
	// return true if the organization is enabled
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrg instantiates a new Org object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrg(status OrgStatus, expiredAt common.NullableTime, allowedActions []string, createdAt time.Time, name string, updatedAt time.Time, enabled bool) *Org {
	this := Org{}
	this.Status = status
	this.ExpiredAt = expiredAt
	this.AllowedActions = allowedActions
	this.CreatedAt = createdAt
	this.Name = name
	this.UpdatedAt = updatedAt
	this.Enabled = enabled
	return &this
}

// NewOrgWithDefaults instantiates a new Org object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgWithDefaults() *Org {
	this := Org{}
	return &this
}

// GetStatus returns the Status field value.
func (o *Org) GetStatus() OrgStatus {
	if o == nil {
		var ret OrgStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Org) GetStatusOk() (*OrgStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *Org) SetStatus(v OrgStatus) {
	o.Status = v
}

// GetExpiredAt returns the ExpiredAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *Org) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Org) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// SetExpiredAt sets field value.
func (o *Org) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}

// GetPurgeError returns the PurgeError field value if set, zero value otherwise.
func (o *Org) GetPurgeError() string {
	if o == nil || o.PurgeError == nil {
		var ret string
		return ret
	}
	return *o.PurgeError
}

// GetPurgeErrorOk returns a tuple with the PurgeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetPurgeErrorOk() (*string, bool) {
	if o == nil || o.PurgeError == nil {
		return nil, false
	}
	return o.PurgeError, true
}

// HasPurgeError returns a boolean if a field has been set.
func (o *Org) HasPurgeError() bool {
	return o != nil && o.PurgeError != nil
}

// SetPurgeError gets a reference to the given string and assigns it to the PurgeError field.
func (o *Org) SetPurgeError(v string) {
	o.PurgeError = &v
}

// GetAllowedActions returns the AllowedActions field value.
func (o *Org) GetAllowedActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value
// and a boolean to check if the value has been set.
func (o *Org) GetAllowedActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedActions, true
}

// SetAllowedActions sets field value.
func (o *Org) SetAllowedActions(v []string) {
	o.AllowedActions = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Org) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Org) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Org) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Org) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Org) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Org) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Org) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Org) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Org) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Org) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Org) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Org) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *Org) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Org) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Org) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Org) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Org) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Org) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetEnabled returns the Enabled field value.
func (o *Org) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Org) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *Org) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Org) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Org) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status         *OrgStatus          `json:"status"`
		ExpiredAt      common.NullableTime `json:"expiredAt"`
		PurgeError     *string             `json:"purgeError,omitempty"`
		AllowedActions *[]string           `json:"allowedActions"`
		CreatedAt      *time.Time          `json:"createdAt"`
		Description    *string             `json:"description,omitempty"`
		DisplayName    *string             `json:"displayName,omitempty"`
		Id             *string             `json:"id,omitempty"`
		Name           *string             `json:"name"`
		UpdatedAt      *time.Time          `json:"updatedAt"`
		Enabled        *bool               `json:"enabled"`
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
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "expiredAt", "purgeError", "allowedActions", "createdAt", "description", "displayName", "id", "name", "updatedAt", "enabled"})
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
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.Id = all.Id
	o.Name = *all.Name
	o.UpdatedAt = *all.UpdatedAt
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
