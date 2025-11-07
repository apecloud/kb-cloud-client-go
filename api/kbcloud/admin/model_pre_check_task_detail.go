// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PreCheckTaskDetail struct {
	TaskId      *string              `json:"taskID,omitempty"`
	TaskStatus  *PreCheckStatus      `json:"taskStatus,omitempty"`
	Progress    common.NullableInt32 `json:"progress,omitempty"`
	Checkers    []PreCheckTaskItem   `json:"checkers,omitempty"`
	Errors      []PreCheckResult     `json:"errors,omitempty"`
	Warnings    []PreCheckResult     `json:"warnings,omitempty"`
	CreatedAt   *time.Time           `json:"createdAt,omitempty"`
	CompletedAt common.NullableTime  `json:"completedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPreCheckTaskDetail instantiates a new PreCheckTaskDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPreCheckTaskDetail() *PreCheckTaskDetail {
	this := PreCheckTaskDetail{}
	return &this
}

// NewPreCheckTaskDetailWithDefaults instantiates a new PreCheckTaskDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPreCheckTaskDetailWithDefaults() *PreCheckTaskDetail {
	this := PreCheckTaskDetail{}
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *PreCheckTaskDetail) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskDetail) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *PreCheckTaskDetail) SetTaskId(v string) {
	o.TaskId = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *PreCheckTaskDetail) GetTaskStatus() PreCheckStatus {
	if o == nil || o.TaskStatus == nil {
		var ret PreCheckStatus
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskDetail) GetTaskStatusOk() (*PreCheckStatus, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasTaskStatus() bool {
	return o != nil && o.TaskStatus != nil
}

// SetTaskStatus gets a reference to the given PreCheckStatus and assigns it to the TaskStatus field.
func (o *PreCheckTaskDetail) SetTaskStatus(v PreCheckStatus) {
	o.TaskStatus = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckTaskDetail) GetProgress() int32 {
	if o == nil || o.Progress.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckTaskDetail) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasProgress() bool {
	return o != nil && o.Progress.IsSet()
}

// SetProgress gets a reference to the given common.NullableInt32 and assigns it to the Progress field.
func (o *PreCheckTaskDetail) SetProgress(v int32) {
	o.Progress.Set(&v)
}

// SetProgressNil sets the value for Progress to be an explicit nil.
func (o *PreCheckTaskDetail) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil.
func (o *PreCheckTaskDetail) UnsetProgress() {
	o.Progress.Unset()
}

// GetCheckers returns the Checkers field value if set, zero value otherwise.
func (o *PreCheckTaskDetail) GetCheckers() []PreCheckTaskItem {
	if o == nil || o.Checkers == nil {
		var ret []PreCheckTaskItem
		return ret
	}
	return o.Checkers
}

// GetCheckersOk returns a tuple with the Checkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskDetail) GetCheckersOk() (*[]PreCheckTaskItem, bool) {
	if o == nil || o.Checkers == nil {
		return nil, false
	}
	return &o.Checkers, true
}

// HasCheckers returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasCheckers() bool {
	return o != nil && o.Checkers != nil
}

// SetCheckers gets a reference to the given []PreCheckTaskItem and assigns it to the Checkers field.
func (o *PreCheckTaskDetail) SetCheckers(v []PreCheckTaskItem) {
	o.Checkers = v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckTaskDetail) GetErrors() []PreCheckResult {
	if o == nil {
		var ret []PreCheckResult
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckTaskDetail) GetErrorsOk() (*[]PreCheckResult, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []PreCheckResult and assigns it to the Errors field.
func (o *PreCheckTaskDetail) SetErrors(v []PreCheckResult) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckTaskDetail) GetWarnings() []PreCheckResult {
	if o == nil {
		var ret []PreCheckResult
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckTaskDetail) GetWarningsOk() (*[]PreCheckResult, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []PreCheckResult and assigns it to the Warnings field.
func (o *PreCheckTaskDetail) SetWarnings(v []PreCheckResult) {
	o.Warnings = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PreCheckTaskDetail) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckTaskDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PreCheckTaskDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreCheckTaskDetail) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PreCheckTaskDetail) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *PreCheckTaskDetail) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given common.NullableTime and assigns it to the CompletedAt field.
func (o *PreCheckTaskDetail) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *PreCheckTaskDetail) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *PreCheckTaskDetail) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PreCheckTaskDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TaskId != nil {
		toSerialize["taskID"] = o.TaskId
	}
	if o.TaskStatus != nil {
		toSerialize["taskStatus"] = o.TaskStatus
	}
	if o.Progress.IsSet() {
		toSerialize["progress"] = o.Progress.Get()
	}
	if o.Checkers != nil {
		toSerialize["checkers"] = o.Checkers
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completedAt"] = o.CompletedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PreCheckTaskDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId      *string              `json:"taskID,omitempty"`
		TaskStatus  *PreCheckStatus      `json:"taskStatus,omitempty"`
		Progress    common.NullableInt32 `json:"progress,omitempty"`
		Checkers    []PreCheckTaskItem   `json:"checkers,omitempty"`
		Errors      []PreCheckResult     `json:"errors,omitempty"`
		Warnings    []PreCheckResult     `json:"warnings,omitempty"`
		CreatedAt   *time.Time           `json:"createdAt,omitempty"`
		CompletedAt common.NullableTime  `json:"completedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskID", "taskStatus", "progress", "checkers", "errors", "warnings", "createdAt", "completedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = all.TaskId
	if all.TaskStatus != nil && !all.TaskStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.TaskStatus = all.TaskStatus
	}
	o.Progress = all.Progress
	o.Checkers = all.Checkers
	o.Errors = all.Errors
	o.Warnings = all.Warnings
	o.CreatedAt = all.CreatedAt
	o.CompletedAt = all.CompletedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
