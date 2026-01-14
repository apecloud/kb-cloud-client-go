// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataChannelModuleProgress struct {
	ModuleName *string                `json:"moduleName,omitempty"`
	Progress   common.NullableFloat64 `json:"progress,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Message    common.NullableString  `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelModuleProgress instantiates a new DataChannelModuleProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelModuleProgress() *DataChannelModuleProgress {
	this := DataChannelModuleProgress{}
	return &this
}

// NewDataChannelModuleProgressWithDefaults instantiates a new DataChannelModuleProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelModuleProgressWithDefaults() *DataChannelModuleProgress {
	this := DataChannelModuleProgress{}
	return &this
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *DataChannelModuleProgress) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelModuleProgress) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *DataChannelModuleProgress) HasModuleName() bool {
	return o != nil && o.ModuleName != nil
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *DataChannelModuleProgress) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelModuleProgress) GetProgress() float64 {
	if o == nil || o.Progress.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Progress.Get()
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelModuleProgress) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress.Get(), o.Progress.IsSet()
}

// HasProgress returns a boolean if a field has been set.
func (o *DataChannelModuleProgress) HasProgress() bool {
	return o != nil && o.Progress.IsSet()
}

// SetProgress gets a reference to the given common.NullableFloat64 and assigns it to the Progress field.
func (o *DataChannelModuleProgress) SetProgress(v float64) {
	o.Progress.Set(&v)
}

// SetProgressNil sets the value for Progress to be an explicit nil.
func (o *DataChannelModuleProgress) SetProgressNil() {
	o.Progress.Set(nil)
}

// UnsetProgress ensures that no value is present for Progress, not even an explicit nil.
func (o *DataChannelModuleProgress) UnsetProgress() {
	o.Progress.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataChannelModuleProgress) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelModuleProgress) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataChannelModuleProgress) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DataChannelModuleProgress) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataChannelModuleProgress) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataChannelModuleProgress) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DataChannelModuleProgress) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given common.NullableString and assigns it to the Message field.
func (o *DataChannelModuleProgress) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *DataChannelModuleProgress) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *DataChannelModuleProgress) UnsetMessage() {
	o.Message.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelModuleProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ModuleName != nil {
		toSerialize["moduleName"] = o.ModuleName
	}
	if o.Progress.IsSet() {
		toSerialize["progress"] = o.Progress.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelModuleProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModuleName *string                `json:"moduleName,omitempty"`
		Progress   common.NullableFloat64 `json:"progress,omitempty"`
		Status     *string                `json:"status,omitempty"`
		Message    common.NullableString  `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"moduleName", "progress", "status", "message"})
	} else {
		return err
	}
	o.ModuleName = all.ModuleName
	o.Progress = all.Progress
	o.Status = all.Status
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
