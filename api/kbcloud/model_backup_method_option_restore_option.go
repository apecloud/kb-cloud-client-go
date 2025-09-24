// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupMethodOptionRestoreOption struct {
	// If this backup needs to be restored on multiple components, the names of those components must be specified.
	// If the restore is only required on the backed-up component itself, then no configuration is needed.
	//
	RestoreOnComponents []string `json:"restoreOnComponents,omitempty"`
	// If set to true, the restore will be ready after the cluster is running.
	// If set to false, the restore will be ready after the cluster is created.
	//
	DeferPostReadyUntilClusterRunning *bool `json:"deferPostReadyUntilClusterRunning,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupMethodOptionRestoreOption instantiates a new BackupMethodOptionRestoreOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupMethodOptionRestoreOption() *BackupMethodOptionRestoreOption {
	this := BackupMethodOptionRestoreOption{}
	return &this
}

// NewBackupMethodOptionRestoreOptionWithDefaults instantiates a new BackupMethodOptionRestoreOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupMethodOptionRestoreOptionWithDefaults() *BackupMethodOptionRestoreOption {
	this := BackupMethodOptionRestoreOption{}
	return &this
}

// GetRestoreOnComponents returns the RestoreOnComponents field value if set, zero value otherwise.
func (o *BackupMethodOptionRestoreOption) GetRestoreOnComponents() []string {
	if o == nil || o.RestoreOnComponents == nil {
		var ret []string
		return ret
	}
	return o.RestoreOnComponents
}

// GetRestoreOnComponentsOk returns a tuple with the RestoreOnComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupMethodOptionRestoreOption) GetRestoreOnComponentsOk() (*[]string, bool) {
	if o == nil || o.RestoreOnComponents == nil {
		return nil, false
	}
	return &o.RestoreOnComponents, true
}

// HasRestoreOnComponents returns a boolean if a field has been set.
func (o *BackupMethodOptionRestoreOption) HasRestoreOnComponents() bool {
	return o != nil && o.RestoreOnComponents != nil
}

// SetRestoreOnComponents gets a reference to the given []string and assigns it to the RestoreOnComponents field.
func (o *BackupMethodOptionRestoreOption) SetRestoreOnComponents(v []string) {
	o.RestoreOnComponents = v
}

// GetDeferPostReadyUntilClusterRunning returns the DeferPostReadyUntilClusterRunning field value if set, zero value otherwise.
func (o *BackupMethodOptionRestoreOption) GetDeferPostReadyUntilClusterRunning() bool {
	if o == nil || o.DeferPostReadyUntilClusterRunning == nil {
		var ret bool
		return ret
	}
	return *o.DeferPostReadyUntilClusterRunning
}

// GetDeferPostReadyUntilClusterRunningOk returns a tuple with the DeferPostReadyUntilClusterRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupMethodOptionRestoreOption) GetDeferPostReadyUntilClusterRunningOk() (*bool, bool) {
	if o == nil || o.DeferPostReadyUntilClusterRunning == nil {
		return nil, false
	}
	return o.DeferPostReadyUntilClusterRunning, true
}

// HasDeferPostReadyUntilClusterRunning returns a boolean if a field has been set.
func (o *BackupMethodOptionRestoreOption) HasDeferPostReadyUntilClusterRunning() bool {
	return o != nil && o.DeferPostReadyUntilClusterRunning != nil
}

// SetDeferPostReadyUntilClusterRunning gets a reference to the given bool and assigns it to the DeferPostReadyUntilClusterRunning field.
func (o *BackupMethodOptionRestoreOption) SetDeferPostReadyUntilClusterRunning(v bool) {
	o.DeferPostReadyUntilClusterRunning = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupMethodOptionRestoreOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RestoreOnComponents != nil {
		toSerialize["restoreOnComponents"] = o.RestoreOnComponents
	}
	if o.DeferPostReadyUntilClusterRunning != nil {
		toSerialize["deferPostReadyUntilClusterRunning"] = o.DeferPostReadyUntilClusterRunning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupMethodOptionRestoreOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RestoreOnComponents               []string `json:"restoreOnComponents,omitempty"`
		DeferPostReadyUntilClusterRunning *bool    `json:"deferPostReadyUntilClusterRunning,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"restoreOnComponents", "deferPostReadyUntilClusterRunning"})
	} else {
		return err
	}
	o.RestoreOnComponents = all.RestoreOnComponents
	o.DeferPostReadyUntilClusterRunning = all.DeferPostReadyUntilClusterRunning

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
