// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TdeResponse Response after enabling transparent data encryption
type TdeResponse struct {
	// Whether TDE is enabled after the operation
	Enable bool `json:"enable"`
	// Task ID if the operation is asynchronous (e.g. reconfigure)
	TaskId *string `json:"taskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTdeResponse instantiates a new TdeResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTdeResponse(enable bool) *TdeResponse {
	this := TdeResponse{}
	this.Enable = enable
	return &this
}

// NewTdeResponseWithDefaults instantiates a new TdeResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTdeResponseWithDefaults() *TdeResponse {
	this := TdeResponse{}
	return &this
}

// GetEnable returns the Enable field value.
func (o *TdeResponse) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *TdeResponse) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *TdeResponse) SetEnable(v bool) {
	o.Enable = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TdeResponse) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TdeResponse) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TdeResponse) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *TdeResponse) SetTaskId(v string) {
	o.TaskId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TdeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enable"] = o.Enable
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TdeResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enable *bool   `json:"enable"`
		TaskId *string `json:"taskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enable", "taskId"})
	} else {
		return err
	}
	o.Enable = *all.Enable
	o.TaskId = all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
