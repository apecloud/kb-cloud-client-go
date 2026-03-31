// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupSchedule backup schedule for database table backup
type BackupSchedule struct {
	// the backup schedule id
	Id *string `json:"id,omitempty"`
	// the name of backup schedule
	ScheduleName *string `json:"scheduleName,omitempty"`
	// the cluster id
	ClusterId *string `json:"clusterId,omitempty"`
	// the backup method
	BackupMethod *string `json:"backupMethod,omitempty"`
	// backup parameters
	Parameters []BackupScheduleParametersItem `json:"parameters,omitempty"`
	// backup selected objects
	SelectedObjects []SelectiveObjectTreeNode `json:"selectedObjects,omitempty"`
	// the creation time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// the last update time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// the cron expression for backup schedule
	CronExpression *string `json:"cronExpression,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupSchedule instantiates a new BackupSchedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupSchedule() *BackupSchedule {
	this := BackupSchedule{}
	return &this
}

// NewBackupScheduleWithDefaults instantiates a new BackupSchedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupScheduleWithDefaults() *BackupSchedule {
	this := BackupSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupSchedule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupSchedule) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupSchedule) SetId(v string) {
	o.Id = &v
}

// GetScheduleName returns the ScheduleName field value if set, zero value otherwise.
func (o *BackupSchedule) GetScheduleName() string {
	if o == nil || o.ScheduleName == nil {
		var ret string
		return ret
	}
	return *o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetScheduleNameOk() (*string, bool) {
	if o == nil || o.ScheduleName == nil {
		return nil, false
	}
	return o.ScheduleName, true
}

// HasScheduleName returns a boolean if a field has been set.
func (o *BackupSchedule) HasScheduleName() bool {
	return o != nil && o.ScheduleName != nil
}

// SetScheduleName gets a reference to the given string and assigns it to the ScheduleName field.
func (o *BackupSchedule) SetScheduleName(v string) {
	o.ScheduleName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *BackupSchedule) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *BackupSchedule) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *BackupSchedule) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetBackupMethod returns the BackupMethod field value if set, zero value otherwise.
func (o *BackupSchedule) GetBackupMethod() string {
	if o == nil || o.BackupMethod == nil {
		var ret string
		return ret
	}
	return *o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetBackupMethodOk() (*string, bool) {
	if o == nil || o.BackupMethod == nil {
		return nil, false
	}
	return o.BackupMethod, true
}

// HasBackupMethod returns a boolean if a field has been set.
func (o *BackupSchedule) HasBackupMethod() bool {
	return o != nil && o.BackupMethod != nil
}

// SetBackupMethod gets a reference to the given string and assigns it to the BackupMethod field.
func (o *BackupSchedule) SetBackupMethod(v string) {
	o.BackupMethod = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BackupSchedule) GetParameters() []BackupScheduleParametersItem {
	if o == nil || o.Parameters == nil {
		var ret []BackupScheduleParametersItem
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetParametersOk() (*[]BackupScheduleParametersItem, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BackupSchedule) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []BackupScheduleParametersItem and assigns it to the Parameters field.
func (o *BackupSchedule) SetParameters(v []BackupScheduleParametersItem) {
	o.Parameters = v
}

// GetSelectedObjects returns the SelectedObjects field value if set, zero value otherwise.
func (o *BackupSchedule) GetSelectedObjects() []SelectiveObjectTreeNode {
	if o == nil || o.SelectedObjects == nil {
		var ret []SelectiveObjectTreeNode
		return ret
	}
	return o.SelectedObjects
}

// GetSelectedObjectsOk returns a tuple with the SelectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetSelectedObjectsOk() (*[]SelectiveObjectTreeNode, bool) {
	if o == nil || o.SelectedObjects == nil {
		return nil, false
	}
	return &o.SelectedObjects, true
}

// HasSelectedObjects returns a boolean if a field has been set.
func (o *BackupSchedule) HasSelectedObjects() bool {
	return o != nil && o.SelectedObjects != nil
}

// SetSelectedObjects gets a reference to the given []SelectiveObjectTreeNode and assigns it to the SelectedObjects field.
func (o *BackupSchedule) SetSelectedObjects(v []SelectiveObjectTreeNode) {
	o.SelectedObjects = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BackupSchedule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BackupSchedule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BackupSchedule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BackupSchedule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BackupSchedule) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BackupSchedule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *BackupSchedule) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *BackupSchedule) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *BackupSchedule) SetCronExpression(v string) {
	o.CronExpression = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ScheduleName != nil {
		toSerialize["scheduleName"] = o.ScheduleName
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.BackupMethod != nil {
		toSerialize["backupMethod"] = o.BackupMethod
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.SelectedObjects != nil {
		toSerialize["selectedObjects"] = o.SelectedObjects
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
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupSchedule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *string                        `json:"id,omitempty"`
		ScheduleName    *string                        `json:"scheduleName,omitempty"`
		ClusterId       *string                        `json:"clusterId,omitempty"`
		BackupMethod    *string                        `json:"backupMethod,omitempty"`
		Parameters      []BackupScheduleParametersItem `json:"parameters,omitempty"`
		SelectedObjects []SelectiveObjectTreeNode      `json:"selectedObjects,omitempty"`
		CreatedAt       *time.Time                     `json:"createdAt,omitempty"`
		UpdatedAt       *time.Time                     `json:"updatedAt,omitempty"`
		CronExpression  *string                        `json:"cronExpression,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "scheduleName", "clusterId", "backupMethod", "parameters", "selectedObjects", "createdAt", "updatedAt", "cronExpression"})
	} else {
		return err
	}
	o.Id = all.Id
	o.ScheduleName = all.ScheduleName
	o.ClusterId = all.ClusterId
	o.BackupMethod = all.BackupMethod
	o.Parameters = all.Parameters
	o.SelectedObjects = all.SelectedObjects
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.CronExpression = all.CronExpression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
