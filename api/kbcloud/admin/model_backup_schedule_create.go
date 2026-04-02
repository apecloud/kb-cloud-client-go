// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupScheduleCreate payload to create a backup schedule
type BackupScheduleCreate struct {
	// the name of backup schedule
	ScheduleName string `json:"scheduleName"`
	// the backup method
	BackupMethod string `json:"backupMethod"`
	// backup parameters
	Parameters map[string]string `json:"parameters,omitempty"`
	// backup selected objects
	SelectedObjects []SelectiveObjectTreeNode `json:"selectedObjects,omitempty"`
	// the cron expression for backup schedule
	CronExpression *string `json:"cronExpression,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupScheduleCreate instantiates a new BackupScheduleCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupScheduleCreate(scheduleName string, backupMethod string) *BackupScheduleCreate {
	this := BackupScheduleCreate{}
	this.ScheduleName = scheduleName
	this.BackupMethod = backupMethod
	return &this
}

// NewBackupScheduleCreateWithDefaults instantiates a new BackupScheduleCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupScheduleCreateWithDefaults() *BackupScheduleCreate {
	this := BackupScheduleCreate{}
	return &this
}

// GetScheduleName returns the ScheduleName field value.
func (o *BackupScheduleCreate) GetScheduleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScheduleName
}

// GetScheduleNameOk returns a tuple with the ScheduleName field value
// and a boolean to check if the value has been set.
func (o *BackupScheduleCreate) GetScheduleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleName, true
}

// SetScheduleName sets field value.
func (o *BackupScheduleCreate) SetScheduleName(v string) {
	o.ScheduleName = v
}

// GetBackupMethod returns the BackupMethod field value.
func (o *BackupScheduleCreate) GetBackupMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value
// and a boolean to check if the value has been set.
func (o *BackupScheduleCreate) GetBackupMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupMethod, true
}

// SetBackupMethod sets field value.
func (o *BackupScheduleCreate) SetBackupMethod(v string) {
	o.BackupMethod = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BackupScheduleCreate) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleCreate) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BackupScheduleCreate) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *BackupScheduleCreate) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetSelectedObjects returns the SelectedObjects field value if set, zero value otherwise.
func (o *BackupScheduleCreate) GetSelectedObjects() []SelectiveObjectTreeNode {
	if o == nil || o.SelectedObjects == nil {
		var ret []SelectiveObjectTreeNode
		return ret
	}
	return o.SelectedObjects
}

// GetSelectedObjectsOk returns a tuple with the SelectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleCreate) GetSelectedObjectsOk() (*[]SelectiveObjectTreeNode, bool) {
	if o == nil || o.SelectedObjects == nil {
		return nil, false
	}
	return &o.SelectedObjects, true
}

// HasSelectedObjects returns a boolean if a field has been set.
func (o *BackupScheduleCreate) HasSelectedObjects() bool {
	return o != nil && o.SelectedObjects != nil
}

// SetSelectedObjects gets a reference to the given []SelectiveObjectTreeNode and assigns it to the SelectedObjects field.
func (o *BackupScheduleCreate) SetSelectedObjects(v []SelectiveObjectTreeNode) {
	o.SelectedObjects = v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *BackupScheduleCreate) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleCreate) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *BackupScheduleCreate) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *BackupScheduleCreate) SetCronExpression(v string) {
	o.CronExpression = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupScheduleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["scheduleName"] = o.ScheduleName
	toSerialize["backupMethod"] = o.BackupMethod
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.SelectedObjects != nil {
		toSerialize["selectedObjects"] = o.SelectedObjects
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
func (o *BackupScheduleCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScheduleName    *string                   `json:"scheduleName"`
		BackupMethod    *string                   `json:"backupMethod"`
		Parameters      map[string]string         `json:"parameters,omitempty"`
		SelectedObjects []SelectiveObjectTreeNode `json:"selectedObjects,omitempty"`
		CronExpression  *string                   `json:"cronExpression,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ScheduleName == nil {
		return fmt.Errorf("required field scheduleName missing")
	}
	if all.BackupMethod == nil {
		return fmt.Errorf("required field backupMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scheduleName", "backupMethod", "parameters", "selectedObjects", "cronExpression"})
	} else {
		return err
	}
	o.ScheduleName = *all.ScheduleName
	o.BackupMethod = *all.BackupMethod
	o.Parameters = all.Parameters
	o.SelectedObjects = all.SelectedObjects
	o.CronExpression = all.CronExpression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
