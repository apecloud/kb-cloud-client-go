// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AutohealingListItem struct {
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// component to be rebuilt
	ComponentName *string `json:"componentName,omitempty"`
	// pod to be rebuilt
	PodName *string `json:"podName,omitempty"`
	// rebuild start time
	StartTime *time.Time `json:"startTime,omitempty"`
	// current phase of rebuild job
	Phase *string `json:"phase,omitempty"`
	// whether the job is started
	Started      *bool   `json:"started,omitempty"`
	StartMessage *string `json:"startMessage,omitempty"`
	// whether the job is finished
	Finished                      *bool   `json:"finished,omitempty"`
	FinishMessage                 *string `json:"finishMessage,omitempty"`
	BackupName                    *string `json:"backupName,omitempty"`
	RebuildInstanceOpsRequestName *string `json:"rebuildInstanceOpsRequestName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutohealingListItem instantiates a new AutohealingListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutohealingListItem() *AutohealingListItem {
	this := AutohealingListItem{}
	return &this
}

// NewAutohealingListItemWithDefaults instantiates a new AutohealingListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutohealingListItemWithDefaults() *AutohealingListItem {
	this := AutohealingListItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutohealingListItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutohealingListItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutohealingListItem) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutohealingListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutohealingListItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutohealingListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AutohealingListItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AutohealingListItem) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AutohealingListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *AutohealingListItem) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *AutohealingListItem) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *AutohealingListItem) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *AutohealingListItem) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *AutohealingListItem) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *AutohealingListItem) SetPodName(v string) {
	o.PodName = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AutohealingListItem) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AutohealingListItem) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AutohealingListItem) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *AutohealingListItem) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *AutohealingListItem) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *AutohealingListItem) SetPhase(v string) {
	o.Phase = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AutohealingListItem) GetStarted() bool {
	if o == nil || o.Started == nil {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetStartedOk() (*bool, bool) {
	if o == nil || o.Started == nil {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *AutohealingListItem) HasStarted() bool {
	return o != nil && o.Started != nil
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *AutohealingListItem) SetStarted(v bool) {
	o.Started = &v
}

// GetStartMessage returns the StartMessage field value if set, zero value otherwise.
func (o *AutohealingListItem) GetStartMessage() string {
	if o == nil || o.StartMessage == nil {
		var ret string
		return ret
	}
	return *o.StartMessage
}

// GetStartMessageOk returns a tuple with the StartMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetStartMessageOk() (*string, bool) {
	if o == nil || o.StartMessage == nil {
		return nil, false
	}
	return o.StartMessage, true
}

// HasStartMessage returns a boolean if a field has been set.
func (o *AutohealingListItem) HasStartMessage() bool {
	return o != nil && o.StartMessage != nil
}

// SetStartMessage gets a reference to the given string and assigns it to the StartMessage field.
func (o *AutohealingListItem) SetStartMessage(v string) {
	o.StartMessage = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *AutohealingListItem) GetFinished() bool {
	if o == nil || o.Finished == nil {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetFinishedOk() (*bool, bool) {
	if o == nil || o.Finished == nil {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *AutohealingListItem) HasFinished() bool {
	return o != nil && o.Finished != nil
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *AutohealingListItem) SetFinished(v bool) {
	o.Finished = &v
}

// GetFinishMessage returns the FinishMessage field value if set, zero value otherwise.
func (o *AutohealingListItem) GetFinishMessage() string {
	if o == nil || o.FinishMessage == nil {
		var ret string
		return ret
	}
	return *o.FinishMessage
}

// GetFinishMessageOk returns a tuple with the FinishMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetFinishMessageOk() (*string, bool) {
	if o == nil || o.FinishMessage == nil {
		return nil, false
	}
	return o.FinishMessage, true
}

// HasFinishMessage returns a boolean if a field has been set.
func (o *AutohealingListItem) HasFinishMessage() bool {
	return o != nil && o.FinishMessage != nil
}

// SetFinishMessage gets a reference to the given string and assigns it to the FinishMessage field.
func (o *AutohealingListItem) SetFinishMessage(v string) {
	o.FinishMessage = &v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise.
func (o *AutohealingListItem) GetBackupName() string {
	if o == nil || o.BackupName == nil {
		var ret string
		return ret
	}
	return *o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetBackupNameOk() (*string, bool) {
	if o == nil || o.BackupName == nil {
		return nil, false
	}
	return o.BackupName, true
}

// HasBackupName returns a boolean if a field has been set.
func (o *AutohealingListItem) HasBackupName() bool {
	return o != nil && o.BackupName != nil
}

// SetBackupName gets a reference to the given string and assigns it to the BackupName field.
func (o *AutohealingListItem) SetBackupName(v string) {
	o.BackupName = &v
}

// GetRebuildInstanceOpsRequestName returns the RebuildInstanceOpsRequestName field value if set, zero value otherwise.
func (o *AutohealingListItem) GetRebuildInstanceOpsRequestName() string {
	if o == nil || o.RebuildInstanceOpsRequestName == nil {
		var ret string
		return ret
	}
	return *o.RebuildInstanceOpsRequestName
}

// GetRebuildInstanceOpsRequestNameOk returns a tuple with the RebuildInstanceOpsRequestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutohealingListItem) GetRebuildInstanceOpsRequestNameOk() (*string, bool) {
	if o == nil || o.RebuildInstanceOpsRequestName == nil {
		return nil, false
	}
	return o.RebuildInstanceOpsRequestName, true
}

// HasRebuildInstanceOpsRequestName returns a boolean if a field has been set.
func (o *AutohealingListItem) HasRebuildInstanceOpsRequestName() bool {
	return o != nil && o.RebuildInstanceOpsRequestName != nil
}

// SetRebuildInstanceOpsRequestName gets a reference to the given string and assigns it to the RebuildInstanceOpsRequestName field.
func (o *AutohealingListItem) SetRebuildInstanceOpsRequestName(v string) {
	o.RebuildInstanceOpsRequestName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutohealingListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.PodName != nil {
		toSerialize["podName"] = o.PodName
	}
	if o.StartTime != nil {
		if o.StartTime.Nanosecond() == 0 {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Started != nil {
		toSerialize["started"] = o.Started
	}
	if o.StartMessage != nil {
		toSerialize["startMessage"] = o.StartMessage
	}
	if o.Finished != nil {
		toSerialize["finished"] = o.Finished
	}
	if o.FinishMessage != nil {
		toSerialize["finishMessage"] = o.FinishMessage
	}
	if o.BackupName != nil {
		toSerialize["backupName"] = o.BackupName
	}
	if o.RebuildInstanceOpsRequestName != nil {
		toSerialize["rebuildInstanceOpsRequestName"] = o.RebuildInstanceOpsRequestName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutohealingListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                          *string    `json:"name,omitempty"`
		CreatedAt                     *time.Time `json:"createdAt,omitempty"`
		UpdatedAt                     *time.Time `json:"updatedAt,omitempty"`
		ComponentName                 *string    `json:"componentName,omitempty"`
		PodName                       *string    `json:"podName,omitempty"`
		StartTime                     *time.Time `json:"startTime,omitempty"`
		Phase                         *string    `json:"phase,omitempty"`
		Started                       *bool      `json:"started,omitempty"`
		StartMessage                  *string    `json:"startMessage,omitempty"`
		Finished                      *bool      `json:"finished,omitempty"`
		FinishMessage                 *string    `json:"finishMessage,omitempty"`
		BackupName                    *string    `json:"backupName,omitempty"`
		RebuildInstanceOpsRequestName *string    `json:"rebuildInstanceOpsRequestName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "createdAt", "updatedAt", "componentName", "podName", "startTime", "phase", "started", "startMessage", "finished", "finishMessage", "backupName", "rebuildInstanceOpsRequestName"})
	} else {
		return err
	}
	o.Name = all.Name
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.ComponentName = all.ComponentName
	o.PodName = all.PodName
	o.StartTime = all.StartTime
	o.Phase = all.Phase
	o.Started = all.Started
	o.StartMessage = all.StartMessage
	o.Finished = all.Finished
	o.FinishMessage = all.FinishMessage
	o.BackupName = all.BackupName
	o.RebuildInstanceOpsRequestName = all.RebuildInstanceOpsRequestName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
