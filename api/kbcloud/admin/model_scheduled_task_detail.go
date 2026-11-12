// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ScheduledTaskDetail struct {
	// ID of the scheduled task. TaskFlow IDs may be numeric for backward compatibility; other providers use provider-qualified IDs such as cr4w:123
	ScheduledTaskId *string `json:"scheduledTaskId,omitempty"`
	// Stable name of the scheduled task
	Name *string `json:"name,omitempty"`
	// Scheduler provider
	Provider *ScheduledTaskProvider `json:"provider,omitempty"`
	// Cron expression used to schedule the task
	CronExpression *string `json:"cronExpression,omitempty"`
	// Timezone used to evaluate the cron expression
	Timezone *string `json:"timezone,omitempty"`
	// Most recently observable execution status
	LastExecutionStatus *ScheduledTaskLastExecutionStatus `json:"lastExecutionStatus,omitempty"`
	// Whether the schedule is immutable, disabled, or enabled
	Enabled *ScheduledTaskEnableStatus `json:"enabled,omitempty"`
	// Scope affected by the task, such as platform, organization, environment, or cluster
	Scope *string `json:"scope,omitempty"`
	// Current normalized status of the scheduled task
	Status *ScheduledTaskStatus `json:"status,omitempty"`
	// Most recent execution start time as a Unix timestamp in seconds
	LastRunAt *string `json:"lastRunAt,omitempty"`
	// Next scheduled execution time as a Unix timestamp in seconds
	NextRunAt *string `json:"nextRunAt,omitempty"`
	// Creation time of the scheduled task record
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Stable task identity within the scheduler provider
	ProviderTaskKey *string `json:"providerTaskKey,omitempty"`
	// Current task status or error message
	Message *string `json:"message,omitempty"`
	// Provider-specific, non-sensitive task details
	ProviderDetails map[string]interface{} `json:"providerDetails,omitempty"`
	// Last time the provider state was observed
	LastObservedAt *time.Time `json:"lastObservedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduledTaskDetail instantiates a new ScheduledTaskDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduledTaskDetail() *ScheduledTaskDetail {
	this := ScheduledTaskDetail{}
	return &this
}

// GetScheduledTaskId returns the ScheduledTaskId field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetScheduledTaskId() string {
	if o == nil || o.ScheduledTaskId == nil {
		var ret string
		return ret
	}
	return *o.ScheduledTaskId
}

// GetScheduledTaskIdOk returns a tuple with the ScheduledTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetScheduledTaskIdOk() (*string, bool) {
	if o == nil || o.ScheduledTaskId == nil {
		return nil, false
	}
	return o.ScheduledTaskId, true
}

// HasScheduledTaskId returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasScheduledTaskId() bool {
	return o != nil && o.ScheduledTaskId != nil
}

// SetScheduledTaskId gets a reference to the given string and assigns it to the ScheduledTaskId field.
func (o *ScheduledTaskDetail) SetScheduledTaskId(v string) {
	o.ScheduledTaskId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScheduledTaskDetail) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetProvider() ScheduledTaskProvider {
	if o == nil || o.Provider == nil {
		var ret ScheduledTaskProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetProviderOk() (*ScheduledTaskProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given ScheduledTaskProvider and assigns it to the Provider field.
func (o *ScheduledTaskDetail) SetProvider(v ScheduledTaskProvider) {
	o.Provider = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *ScheduledTaskDetail) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ScheduledTaskDetail) SetTimezone(v string) {
	o.Timezone = &v
}

// GetLastExecutionStatus returns the LastExecutionStatus field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetLastExecutionStatus() ScheduledTaskLastExecutionStatus {
	if o == nil || o.LastExecutionStatus == nil {
		var ret ScheduledTaskLastExecutionStatus
		return ret
	}
	return *o.LastExecutionStatus
}

// GetLastExecutionStatusOk returns a tuple with the LastExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetLastExecutionStatusOk() (*ScheduledTaskLastExecutionStatus, bool) {
	if o == nil || o.LastExecutionStatus == nil {
		return nil, false
	}
	return o.LastExecutionStatus, true
}

// HasLastExecutionStatus returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasLastExecutionStatus() bool {
	return o != nil && o.LastExecutionStatus != nil
}

// SetLastExecutionStatus gets a reference to the given ScheduledTaskLastExecutionStatus and assigns it to the LastExecutionStatus field.
func (o *ScheduledTaskDetail) SetLastExecutionStatus(v ScheduledTaskLastExecutionStatus) {
	o.LastExecutionStatus = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetEnabled() ScheduledTaskEnableStatus {
	if o == nil || o.Enabled == nil {
		var ret ScheduledTaskEnableStatus
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetEnabledOk() (*ScheduledTaskEnableStatus, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given ScheduledTaskEnableStatus and assigns it to the Enabled field.
func (o *ScheduledTaskDetail) SetEnabled(v ScheduledTaskEnableStatus) {
	o.Enabled = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ScheduledTaskDetail) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetStatus() ScheduledTaskStatus {
	if o == nil || o.Status == nil {
		var ret ScheduledTaskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetStatusOk() (*ScheduledTaskStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given ScheduledTaskStatus and assigns it to the Status field.
func (o *ScheduledTaskDetail) SetStatus(v ScheduledTaskStatus) {
	o.Status = &v
}

// GetLastRunAt returns the LastRunAt field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetLastRunAt() string {
	if o == nil || o.LastRunAt == nil {
		var ret string
		return ret
	}
	return *o.LastRunAt
}

// GetLastRunAtOk returns a tuple with the LastRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetLastRunAtOk() (*string, bool) {
	if o == nil || o.LastRunAt == nil {
		return nil, false
	}
	return o.LastRunAt, true
}

// HasLastRunAt returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasLastRunAt() bool {
	return o != nil && o.LastRunAt != nil
}

// SetLastRunAt gets a reference to the given string and assigns it to the LastRunAt field.
func (o *ScheduledTaskDetail) SetLastRunAt(v string) {
	o.LastRunAt = &v
}

// GetNextRunAt returns the NextRunAt field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetNextRunAt() string {
	if o == nil || o.NextRunAt == nil {
		var ret string
		return ret
	}
	return *o.NextRunAt
}

// GetNextRunAtOk returns a tuple with the NextRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetNextRunAtOk() (*string, bool) {
	if o == nil || o.NextRunAt == nil {
		return nil, false
	}
	return o.NextRunAt, true
}

// HasNextRunAt returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasNextRunAt() bool {
	return o != nil && o.NextRunAt != nil
}

// SetNextRunAt gets a reference to the given string and assigns it to the NextRunAt field.
func (o *ScheduledTaskDetail) SetNextRunAt(v string) {
	o.NextRunAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ScheduledTaskDetail) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetProviderTaskKey returns the ProviderTaskKey field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetProviderTaskKey() string {
	if o == nil || o.ProviderTaskKey == nil {
		var ret string
		return ret
	}
	return *o.ProviderTaskKey
}

// GetProviderTaskKeyOk returns a tuple with the ProviderTaskKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetProviderTaskKeyOk() (*string, bool) {
	if o == nil || o.ProviderTaskKey == nil {
		return nil, false
	}
	return o.ProviderTaskKey, true
}

// HasProviderTaskKey returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasProviderTaskKey() bool {
	return o != nil && o.ProviderTaskKey != nil
}

// SetProviderTaskKey gets a reference to the given string and assigns it to the ProviderTaskKey field.
func (o *ScheduledTaskDetail) SetProviderTaskKey(v string) {
	o.ProviderTaskKey = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ScheduledTaskDetail) SetMessage(v string) {
	o.Message = &v
}

// GetProviderDetails returns the ProviderDetails field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetProviderDetails() map[string]interface{} {
	if o == nil || o.ProviderDetails == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ProviderDetails
}

// GetProviderDetailsOk returns a tuple with the ProviderDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetProviderDetailsOk() (*map[string]interface{}, bool) {
	if o == nil || o.ProviderDetails == nil {
		return nil, false
	}
	return &o.ProviderDetails, true
}

// HasProviderDetails returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasProviderDetails() bool {
	return o != nil && o.ProviderDetails != nil
}

// SetProviderDetails gets a reference to the given map[string]interface{} and assigns it to the ProviderDetails field.
func (o *ScheduledTaskDetail) SetProviderDetails(v map[string]interface{}) {
	o.ProviderDetails = v
}

// GetLastObservedAt returns the LastObservedAt field value if set, zero value otherwise.
func (o *ScheduledTaskDetail) GetLastObservedAt() time.Time {
	if o == nil || o.LastObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastObservedAt
}

// GetLastObservedAtOk returns a tuple with the LastObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskDetail) GetLastObservedAtOk() (*time.Time, bool) {
	if o == nil || o.LastObservedAt == nil {
		return nil, false
	}
	return o.LastObservedAt, true
}

// HasLastObservedAt returns a boolean if a field has been set.
func (o *ScheduledTaskDetail) HasLastObservedAt() bool {
	return o != nil && o.LastObservedAt != nil
}

// SetLastObservedAt gets a reference to the given time.Time and assigns it to the LastObservedAt field.
func (o *ScheduledTaskDetail) SetLastObservedAt(v time.Time) {
	o.LastObservedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduledTaskDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ScheduledTaskId != nil {
		toSerialize["scheduledTaskId"] = o.ScheduledTaskId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.LastExecutionStatus != nil {
		toSerialize["lastExecutionStatus"] = o.LastExecutionStatus
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastRunAt != nil {
		toSerialize["lastRunAt"] = o.LastRunAt
	}
	if o.NextRunAt != nil {
		toSerialize["nextRunAt"] = o.NextRunAt
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ProviderTaskKey != nil {
		toSerialize["providerTaskKey"] = o.ProviderTaskKey
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ProviderDetails != nil {
		toSerialize["providerDetails"] = o.ProviderDetails
	}
	if o.LastObservedAt != nil {
		if o.LastObservedAt.Nanosecond() == 0 {
			toSerialize["lastObservedAt"] = o.LastObservedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastObservedAt"] = o.LastObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduledTaskDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScheduledTaskId     *string                           `json:"scheduledTaskId,omitempty"`
		Name                *string                           `json:"name,omitempty"`
		Provider            *ScheduledTaskProvider            `json:"provider,omitempty"`
		CronExpression      *string                           `json:"cronExpression,omitempty"`
		Timezone            *string                           `json:"timezone,omitempty"`
		LastExecutionStatus *ScheduledTaskLastExecutionStatus `json:"lastExecutionStatus,omitempty"`
		Enabled             *ScheduledTaskEnableStatus        `json:"enabled,omitempty"`
		Scope               *string                           `json:"scope,omitempty"`
		Status              *ScheduledTaskStatus              `json:"status,omitempty"`
		LastRunAt           *string                           `json:"lastRunAt,omitempty"`
		NextRunAt           *string                           `json:"nextRunAt,omitempty"`
		CreatedAt           *time.Time                        `json:"createdAt,omitempty"`
		ProviderTaskKey     *string                           `json:"providerTaskKey,omitempty"`
		Message             *string                           `json:"message,omitempty"`
		ProviderDetails     map[string]interface{}            `json:"providerDetails,omitempty"`
		LastObservedAt      *time.Time                        `json:"lastObservedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scheduledTaskId", "name", "provider", "cronExpression", "timezone", "lastExecutionStatus", "enabled", "scope", "status", "lastRunAt", "nextRunAt", "createdAt", "providerTaskKey", "message", "providerDetails", "lastObservedAt"})
	} else {
		return err
	}
	hasInvalidField := false
	o.ScheduledTaskId = all.ScheduledTaskId
	o.Name = all.Name
	if all.Provider != nil && !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = all.Provider
	}
	o.CronExpression = all.CronExpression
	o.Timezone = all.Timezone
	if all.LastExecutionStatus != nil && !all.LastExecutionStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.LastExecutionStatus = all.LastExecutionStatus
	}
	if all.Enabled != nil && !all.Enabled.IsValid() {
		hasInvalidField = true
	} else {
		o.Enabled = all.Enabled
	}
	o.Scope = all.Scope
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.LastRunAt = all.LastRunAt
	o.NextRunAt = all.NextRunAt
	o.CreatedAt = all.CreatedAt
	o.ProviderTaskKey = all.ProviderTaskKey
	o.Message = all.Message
	o.ProviderDetails = all.ProviderDetails
	o.LastObservedAt = all.LastObservedAt
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
