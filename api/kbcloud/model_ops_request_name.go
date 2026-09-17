// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsRequestName OpsRequestName is the name of a KubeBlocks OpsRequest response
type OpsRequestName struct {
	OpsRequestName *string `json:"opsRequestName,omitempty"`
	TaskId         string  `json:"taskId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsRequestName instantiates a new OpsRequestName object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRequestName(taskId string) *OpsRequestName {
	this := OpsRequestName{}
	this.TaskId = taskId
	return &this
}

// NewOpsRequestNameWithDefaults instantiates a new OpsRequestName object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRequestNameWithDefaults() *OpsRequestName {
	this := OpsRequestName{}
	return &this
}

// GetOpsRequestName returns the OpsRequestName field value if set, zero value otherwise.
func (o *OpsRequestName) GetOpsRequestName() string {
	if o == nil || o.OpsRequestName == nil {
		var ret string
		return ret
	}
	return *o.OpsRequestName
}

// GetOpsRequestNameOk returns a tuple with the OpsRequestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRequestName) GetOpsRequestNameOk() (*string, bool) {
	if o == nil || o.OpsRequestName == nil {
		return nil, false
	}
	return o.OpsRequestName, true
}

// HasOpsRequestName returns a boolean if a field has been set.
func (o *OpsRequestName) HasOpsRequestName() bool {
	return o != nil && o.OpsRequestName != nil
}

// SetOpsRequestName gets a reference to the given string and assigns it to the OpsRequestName field.
func (o *OpsRequestName) SetOpsRequestName(v string) {
	o.OpsRequestName = &v
}

// GetTaskId returns the TaskId field value.
func (o *OpsRequestName) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *OpsRequestName) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *OpsRequestName) SetTaskId(v string) {
	o.TaskId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsRequestName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OpsRequestName != nil {
		toSerialize["opsRequestName"] = o.OpsRequestName
	}
	toSerialize["taskId"] = o.TaskId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRequestName) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsRequestName *string `json:"opsRequestName,omitempty"`
		TaskId         *string `json:"taskId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"opsRequestName", "taskId"})
	} else {
		return err
	}
	o.OpsRequestName = all.OpsRequestName
	o.TaskId = *all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
