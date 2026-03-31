// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupCreate BackupCreate is the payload to create a KubeBlocks cluster backup
type BackupCreate struct {
	// name of the backup, if not specified, a name will be generated automatically
	Name *string `json:"name,omitempty"`
	// the type of backup
	BackupType *BackupType `json:"backupType,omitempty"`
	// specified the backup method
	BackupMethod string `json:"backupMethod"`
	// component of the cluster to back up
	Component *string `json:"component,omitempty"`
	// specify the retention period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// parameters for the backup
	Parameters map[string]string `json:"parameters,omitempty"`
	// selected objects for the backup
	SelectedObjects []SelectiveObjectTreeNode `json:"selectedObjects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupCreate instantiates a new BackupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupCreate(backupMethod string) *BackupCreate {
	this := BackupCreate{}
	this.BackupMethod = backupMethod
	return &this
}

// NewBackupCreateWithDefaults instantiates a new BackupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupCreateWithDefaults() *BackupCreate {
	this := BackupCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupCreate) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupCreate) SetName(v string) {
	o.Name = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *BackupCreate) GetBackupType() BackupType {
	if o == nil || o.BackupType == nil {
		var ret BackupType
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetBackupTypeOk() (*BackupType, bool) {
	if o == nil || o.BackupType == nil {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *BackupCreate) HasBackupType() bool {
	return o != nil && o.BackupType != nil
}

// SetBackupType gets a reference to the given BackupType and assigns it to the BackupType field.
func (o *BackupCreate) SetBackupType(v BackupType) {
	o.BackupType = &v
}

// GetBackupMethod returns the BackupMethod field value.
func (o *BackupCreate) GetBackupMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetBackupMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupMethod, true
}

// SetBackupMethod sets field value.
func (o *BackupCreate) SetBackupMethod(v string) {
	o.BackupMethod = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *BackupCreate) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *BackupCreate) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *BackupCreate) SetComponent(v string) {
	o.Component = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *BackupCreate) GetRetentionPeriod() string {
	if o == nil || o.RetentionPeriod == nil {
		var ret string
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetRetentionPeriodOk() (*string, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *BackupCreate) HasRetentionPeriod() bool {
	return o != nil && o.RetentionPeriod != nil
}

// SetRetentionPeriod gets a reference to the given string and assigns it to the RetentionPeriod field.
func (o *BackupCreate) SetRetentionPeriod(v string) {
	o.RetentionPeriod = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BackupCreate) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BackupCreate) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *BackupCreate) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetSelectedObjects returns the SelectedObjects field value if set, zero value otherwise.
func (o *BackupCreate) GetSelectedObjects() []SelectiveObjectTreeNode {
	if o == nil || o.SelectedObjects == nil {
		var ret []SelectiveObjectTreeNode
		return ret
	}
	return o.SelectedObjects
}

// GetSelectedObjectsOk returns a tuple with the SelectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetSelectedObjectsOk() (*[]SelectiveObjectTreeNode, bool) {
	if o == nil || o.SelectedObjects == nil {
		return nil, false
	}
	return &o.SelectedObjects, true
}

// HasSelectedObjects returns a boolean if a field has been set.
func (o *BackupCreate) HasSelectedObjects() bool {
	return o != nil && o.SelectedObjects != nil
}

// SetSelectedObjects gets a reference to the given []SelectiveObjectTreeNode and assigns it to the SelectedObjects field.
func (o *BackupCreate) SetSelectedObjects(v []SelectiveObjectTreeNode) {
	o.SelectedObjects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BackupType != nil {
		toSerialize["backupType"] = o.BackupType
	}
	toSerialize["backupMethod"] = o.BackupMethod
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.SelectedObjects != nil {
		toSerialize["selectedObjects"] = o.SelectedObjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string                   `json:"name,omitempty"`
		BackupType      *BackupType               `json:"backupType,omitempty"`
		BackupMethod    *string                   `json:"backupMethod"`
		Component       *string                   `json:"component,omitempty"`
		RetentionPeriod *string                   `json:"retentionPeriod,omitempty"`
		Parameters      map[string]string         `json:"parameters,omitempty"`
		SelectedObjects []SelectiveObjectTreeNode `json:"selectedObjects,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BackupMethod == nil {
		return fmt.Errorf("required field backupMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "backupType", "backupMethod", "component", "retentionPeriod", "parameters", "selectedObjects"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.BackupType != nil && !all.BackupType.IsValid() {
		hasInvalidField = true
	} else {
		o.BackupType = all.BackupType
	}
	o.BackupMethod = *all.BackupMethod
	o.Component = all.Component
	o.RetentionPeriod = all.RetentionPeriod
	o.Parameters = all.Parameters
	o.SelectedObjects = all.SelectedObjects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
