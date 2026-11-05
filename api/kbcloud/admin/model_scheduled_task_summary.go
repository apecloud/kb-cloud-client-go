// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ScheduledTaskSummary struct {
	// ID of the scheduled task
	ScheduledTaskId string `json:"scheduledTaskId"`
	// Stable name of the scheduled task
	Name string `json:"name"`
	// Scheduler provider
	Provider *ScheduledTaskProvider `json:"provider,omitempty"`
	// Cron expression used to schedule the task
	CronExpression *string `json:"cronExpression,omitempty"`
	// Most recently observable execution status
	LastExecutionStatus *ScheduledTaskLastExecutionStatus `json:"lastExecutionStatus,omitempty"`
	// Whether the schedule is immutable, disabled, or enabled
	Enabled *ScheduledTaskEnableStatus `json:"enabled,omitempty"`
	// Scope affected by the task, such as platform, organization, environment, or cluster
	Scope *string `json:"scope,omitempty"`
	// Current normalized status of the scheduled task
	Status ScheduledTaskStatus `json:"status"`
	// Most recent execution start time as a Unix timestamp in seconds
	LastRunAt *string `json:"lastRunAt,omitempty"`
	// Next scheduled execution time as a Unix timestamp in seconds
	NextRunAt *string `json:"nextRunAt,omitempty"`
	// Creation time of the scheduled task record
	CreatedAt time.Time `json:"createdAt"`
	// Last update time of the scheduled task record
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduledTaskSummary instantiates a new ScheduledTaskSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduledTaskSummary(scheduledTaskId string, name string, status ScheduledTaskStatus, createdAt time.Time, updatedAt time.Time) *ScheduledTaskSummary {
	this := ScheduledTaskSummary{}
	this.ScheduledTaskId = scheduledTaskId
	this.Name = name
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewScheduledTaskSummaryWithDefaults instantiates a new ScheduledTaskSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduledTaskSummaryWithDefaults() *ScheduledTaskSummary {
	this := ScheduledTaskSummary{}
	return &this
}

// GetScheduledTaskId returns the ScheduledTaskId field value.
func (o *ScheduledTaskSummary) GetScheduledTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScheduledTaskId
}

// GetScheduledTaskIdOk returns a tuple with the ScheduledTaskId field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetScheduledTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledTaskId, true
}

// SetScheduledTaskId sets field value.
func (o *ScheduledTaskSummary) SetScheduledTaskId(v string) {
	o.ScheduledTaskId = v
}

// GetName returns the Name field value.
func (o *ScheduledTaskSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ScheduledTaskSummary) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetProvider() ScheduledTaskProvider {
	if o == nil || o.Provider == nil {
		var ret ScheduledTaskProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetProviderOk() (*ScheduledTaskProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given ScheduledTaskProvider and assigns it to the Provider field.
func (o *ScheduledTaskSummary) SetProvider(v ScheduledTaskProvider) {
	o.Provider = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *ScheduledTaskSummary) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetLastExecutionStatus returns the LastExecutionStatus field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetLastExecutionStatus() ScheduledTaskLastExecutionStatus {
	if o == nil || o.LastExecutionStatus == nil {
		var ret ScheduledTaskLastExecutionStatus
		return ret
	}
	return *o.LastExecutionStatus
}

// GetLastExecutionStatusOk returns a tuple with the LastExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetLastExecutionStatusOk() (*ScheduledTaskLastExecutionStatus, bool) {
	if o == nil || o.LastExecutionStatus == nil {
		return nil, false
	}
	return o.LastExecutionStatus, true
}

// HasLastExecutionStatus returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasLastExecutionStatus() bool {
	return o != nil && o.LastExecutionStatus != nil
}

// SetLastExecutionStatus gets a reference to the given ScheduledTaskLastExecutionStatus and assigns it to the LastExecutionStatus field.
func (o *ScheduledTaskSummary) SetLastExecutionStatus(v ScheduledTaskLastExecutionStatus) {
	o.LastExecutionStatus = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetEnabled() ScheduledTaskEnableStatus {
	if o == nil || o.Enabled == nil {
		var ret ScheduledTaskEnableStatus
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetEnabledOk() (*ScheduledTaskEnableStatus, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given ScheduledTaskEnableStatus and assigns it to the Enabled field.
func (o *ScheduledTaskSummary) SetEnabled(v ScheduledTaskEnableStatus) {
	o.Enabled = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ScheduledTaskSummary) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value.
func (o *ScheduledTaskSummary) GetStatus() ScheduledTaskStatus {
	if o == nil {
		var ret ScheduledTaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetStatusOk() (*ScheduledTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ScheduledTaskSummary) SetStatus(v ScheduledTaskStatus) {
	o.Status = v
}

// GetLastRunAt returns the LastRunAt field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetLastRunAt() string {
	if o == nil || o.LastRunAt == nil {
		var ret string
		return ret
	}
	return *o.LastRunAt
}

// GetLastRunAtOk returns a tuple with the LastRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetLastRunAtOk() (*string, bool) {
	if o == nil || o.LastRunAt == nil {
		return nil, false
	}
	return o.LastRunAt, true
}

// HasLastRunAt returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasLastRunAt() bool {
	return o != nil && o.LastRunAt != nil
}

// SetLastRunAt gets a reference to the given string and assigns it to the LastRunAt field.
func (o *ScheduledTaskSummary) SetLastRunAt(v string) {
	o.LastRunAt = &v
}

// GetNextRunAt returns the NextRunAt field value if set, zero value otherwise.
func (o *ScheduledTaskSummary) GetNextRunAt() string {
	if o == nil || o.NextRunAt == nil {
		var ret string
		return ret
	}
	return *o.NextRunAt
}

// GetNextRunAtOk returns a tuple with the NextRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetNextRunAtOk() (*string, bool) {
	if o == nil || o.NextRunAt == nil {
		return nil, false
	}
	return o.NextRunAt, true
}

// HasNextRunAt returns a boolean if a field has been set.
func (o *ScheduledTaskSummary) HasNextRunAt() bool {
	return o != nil && o.NextRunAt != nil
}

// SetNextRunAt gets a reference to the given string and assigns it to the NextRunAt field.
func (o *ScheduledTaskSummary) SetNextRunAt(v string) {
	o.NextRunAt = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ScheduledTaskSummary) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ScheduledTaskSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ScheduledTaskSummary) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ScheduledTaskSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduledTaskSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["scheduledTaskId"] = o.ScheduledTaskId
	toSerialize["name"] = o.Name
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
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
	toSerialize["status"] = o.Status
	if o.LastRunAt != nil {
		toSerialize["lastRunAt"] = o.LastRunAt
	}
	if o.NextRunAt != nil {
		toSerialize["nextRunAt"] = o.NextRunAt
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduledTaskSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScheduledTaskId     *string                           `json:"scheduledTaskId"`
		Name                *string                           `json:"name"`
		Provider            *ScheduledTaskProvider            `json:"provider,omitempty"`
		CronExpression      *string                           `json:"cronExpression,omitempty"`
		LastExecutionStatus *ScheduledTaskLastExecutionStatus `json:"lastExecutionStatus,omitempty"`
		Enabled             *ScheduledTaskEnableStatus        `json:"enabled,omitempty"`
		Scope               *string                           `json:"scope,omitempty"`
		Status              *ScheduledTaskStatus              `json:"status"`
		LastRunAt           *string                           `json:"lastRunAt,omitempty"`
		NextRunAt           *string                           `json:"nextRunAt,omitempty"`
		CreatedAt           *time.Time                        `json:"createdAt"`
		UpdatedAt           *time.Time                        `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ScheduledTaskId == nil {
		return fmt.Errorf("required field scheduledTaskId missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scheduledTaskId", "name", "provider", "cronExpression", "lastExecutionStatus", "enabled", "scope", "status", "lastRunAt", "nextRunAt", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ScheduledTaskId = *all.ScheduledTaskId
	o.Name = *all.Name
	if all.Provider != nil && !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = all.Provider
	}
	o.CronExpression = all.CronExpression
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
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.LastRunAt = all.LastRunAt
	o.NextRunAt = all.NextRunAt
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
