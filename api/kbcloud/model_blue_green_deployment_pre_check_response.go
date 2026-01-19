// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentPreCheckResponse blueGreenDeploymentPreCheckResponse is the response for retrieving the precheck task of a blue-green deployment.
type BlueGreenDeploymentPreCheckResponse struct {
	// Task ID.
	TaskId string `json:"taskId"`
	// The checkers of the precheck task.
	Checkers []BlueGreenDeploymentPrecheckItem `json:"checkers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentPreCheckResponse instantiates a new BlueGreenDeploymentPreCheckResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentPreCheckResponse(taskId string) *BlueGreenDeploymentPreCheckResponse {
	this := BlueGreenDeploymentPreCheckResponse{}
	this.TaskId = taskId
	return &this
}

// NewBlueGreenDeploymentPreCheckResponseWithDefaults instantiates a new BlueGreenDeploymentPreCheckResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentPreCheckResponseWithDefaults() *BlueGreenDeploymentPreCheckResponse {
	this := BlueGreenDeploymentPreCheckResponse{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *BlueGreenDeploymentPreCheckResponse) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPreCheckResponse) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *BlueGreenDeploymentPreCheckResponse) SetTaskId(v string) {
	o.TaskId = v
}

// GetCheckers returns the Checkers field value if set, zero value otherwise.
func (o *BlueGreenDeploymentPreCheckResponse) GetCheckers() []BlueGreenDeploymentPrecheckItem {
	if o == nil || o.Checkers == nil {
		var ret []BlueGreenDeploymentPrecheckItem
		return ret
	}
	return o.Checkers
}

// GetCheckersOk returns a tuple with the Checkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPreCheckResponse) GetCheckersOk() (*[]BlueGreenDeploymentPrecheckItem, bool) {
	if o == nil || o.Checkers == nil {
		return nil, false
	}
	return &o.Checkers, true
}

// HasCheckers returns a boolean if a field has been set.
func (o *BlueGreenDeploymentPreCheckResponse) HasCheckers() bool {
	return o != nil && o.Checkers != nil
}

// SetCheckers gets a reference to the given []BlueGreenDeploymentPrecheckItem and assigns it to the Checkers field.
func (o *BlueGreenDeploymentPreCheckResponse) SetCheckers(v []BlueGreenDeploymentPrecheckItem) {
	o.Checkers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentPreCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	if o.Checkers != nil {
		toSerialize["checkers"] = o.Checkers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentPreCheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId   *string                           `json:"taskId"`
		Checkers []BlueGreenDeploymentPrecheckItem `json:"checkers,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "checkers"})
	} else {
		return err
	}
	o.TaskId = *all.TaskId
	o.Checkers = all.Checkers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
