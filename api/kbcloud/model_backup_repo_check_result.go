// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupRepoCheckResult Result of the platform storage access check, independent of the BackupRepo CR phase.
type BackupRepoCheckResult struct {
	Status    BackupRepoCheckStatus `json:"status"`
	Message   *string               `json:"message,omitempty"`
	CheckedAt *time.Time            `json:"checkedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoCheckResult instantiates a new BackupRepoCheckResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoCheckResult(status BackupRepoCheckStatus) *BackupRepoCheckResult {
	this := BackupRepoCheckResult{}
	this.Status = status
	return &this
}

// NewBackupRepoCheckResultWithDefaults instantiates a new BackupRepoCheckResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoCheckResultWithDefaults() *BackupRepoCheckResult {
	this := BackupRepoCheckResult{}
	return &this
}

// GetStatus returns the Status field value.
func (o *BackupRepoCheckResult) GetStatus() BackupRepoCheckStatus {
	if o == nil {
		var ret BackupRepoCheckStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BackupRepoCheckResult) GetStatusOk() (*BackupRepoCheckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *BackupRepoCheckResult) SetStatus(v BackupRepoCheckStatus) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BackupRepoCheckResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCheckResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BackupRepoCheckResult) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BackupRepoCheckResult) SetMessage(v string) {
	o.Message = &v
}

// GetCheckedAt returns the CheckedAt field value if set, zero value otherwise.
func (o *BackupRepoCheckResult) GetCheckedAt() time.Time {
	if o == nil || o.CheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCheckResult) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.CheckedAt == nil {
		return nil, false
	}
	return o.CheckedAt, true
}

// HasCheckedAt returns a boolean if a field has been set.
func (o *BackupRepoCheckResult) HasCheckedAt() bool {
	return o != nil && o.CheckedAt != nil
}

// SetCheckedAt gets a reference to the given time.Time and assigns it to the CheckedAt field.
func (o *BackupRepoCheckResult) SetCheckedAt(v time.Time) {
	o.CheckedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.CheckedAt != nil {
		if o.CheckedAt.Nanosecond() == 0 {
			toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status    *BackupRepoCheckStatus `json:"status"`
		Message   *string                `json:"message,omitempty"`
		CheckedAt *time.Time             `json:"checkedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "message", "checkedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Message = all.Message
	o.CheckedAt = all.CheckedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
