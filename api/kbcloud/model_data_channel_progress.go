// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelProgress struct {
	Progress       common.NullableFloat64      `json:"progress,omitempty"`
	ModuleProgress []DataChannelModuleProgress `json:"moduleProgress,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelProgress instantiates a new DataChannelProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelProgress() *DataChannelProgress {
	this := DataChannelProgress{}
	return &this
}

// NewDataChannelProgressWithDefaults instantiates a new DataChannelProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelProgressWithDefaults() *DataChannelProgress {
	this := DataChannelProgress{}
	return &this
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelProgress) GetProgress() float64 {
	if o == nil || o.Progress.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelProgress) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *DataChannelProgress) HasProgress() bool {
	return o != nil && o.Progress.IsSet()
}

// SetProgress gets a reference to the given common.NullableFloat64 and assigns it to the Progress field.
func (o *DataChannelProgress) SetProgress(v float64) {
	o.Progress.Set(&v)
}

// SetProgressNil sets the value for Progress to be an explicit nil.
func (o *DataChannelProgress) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil.
func (o *DataChannelProgress) UnsetProgress() {
	o.Progress.Unset()
}

// GetModuleProgress returns the ModuleProgress field value if set, zero value otherwise.
func (o *DataChannelProgress) GetModuleProgress() []DataChannelModuleProgress {
	if o == nil || o.ModuleProgress == nil {
		var ret []DataChannelModuleProgress
		return ret
	}
	return o.ModuleProgress
}

// GetModuleProgressOk returns a tuple with the ModuleProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelProgress) GetModuleProgressOk() (*[]DataChannelModuleProgress, bool) {
	if o == nil || o.ModuleProgress == nil {
		return nil, false
	}
	return &o.ModuleProgress, true
}

// HasModuleProgress returns a boolean if a field has been set.
func (o *DataChannelProgress) HasModuleProgress() bool {
	return o != nil && o.ModuleProgress != nil
}

// SetModuleProgress gets a reference to the given []DataChannelModuleProgress and assigns it to the ModuleProgress field.
func (o *DataChannelProgress) SetModuleProgress(v []DataChannelModuleProgress) {
	o.ModuleProgress = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Progress.IsSet() {
		toSerialize["progress"] = o.Progress.Get()
	}
	if o.ModuleProgress != nil {
		toSerialize["moduleProgress"] = o.ModuleProgress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Progress       common.NullableFloat64      `json:"progress,omitempty"`
		ModuleProgress []DataChannelModuleProgress `json:"moduleProgress,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"progress", "moduleProgress"})
	} else {
		return err
	}
	o.Progress = all.Progress
	o.ModuleProgress = all.ModuleProgress

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
