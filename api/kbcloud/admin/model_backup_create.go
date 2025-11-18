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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string     `json:"name,omitempty"`
		BackupType   *BackupType `json:"backupType,omitempty"`
		BackupMethod *string     `json:"backupMethod"`
		Component    *string     `json:"component,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BackupMethod == nil {
		return fmt.Errorf("required field backupMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "backupType", "backupMethod", "component"})
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
