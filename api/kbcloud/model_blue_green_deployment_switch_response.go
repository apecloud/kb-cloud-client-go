// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentSwitchResponse blueGreenDeploymentSwitchResponse represents the response when switching a blue-green deployment.
type BlueGreenDeploymentSwitchResponse struct {
	// The ID of the blue-green deployment.
	DeploymentId *string `json:"deploymentID,omitempty"`
	// Task ID.
	TaskId string `json:"taskId"`
	// The steps involved in the blue-green deployment switch.
	Steps []BlueGreenDeploymentSwitchStep `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentSwitchResponse instantiates a new BlueGreenDeploymentSwitchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentSwitchResponse(taskId string) *BlueGreenDeploymentSwitchResponse {
	this := BlueGreenDeploymentSwitchResponse{}
	this.TaskId = taskId
	return &this
}

// NewBlueGreenDeploymentSwitchResponseWithDefaults instantiates a new BlueGreenDeploymentSwitchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentSwitchResponseWithDefaults() *BlueGreenDeploymentSwitchResponse {
	this := BlueGreenDeploymentSwitchResponse{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchResponse) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchResponse) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchResponse) HasDeploymentId() bool {
	return o != nil && o.DeploymentId != nil
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *BlueGreenDeploymentSwitchResponse) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetTaskId returns the TaskId field value.
func (o *BlueGreenDeploymentSwitchResponse) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchResponse) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *BlueGreenDeploymentSwitchResponse) SetTaskId(v string) {
	o.TaskId = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchResponse) GetSteps() []BlueGreenDeploymentSwitchStep {
	if o == nil || o.Steps == nil {
		var ret []BlueGreenDeploymentSwitchStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchResponse) GetStepsOk() (*[]BlueGreenDeploymentSwitchStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchResponse) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []BlueGreenDeploymentSwitchStep and assigns it to the Steps field.
func (o *BlueGreenDeploymentSwitchResponse) SetSteps(v []BlueGreenDeploymentSwitchStep) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentSwitchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DeploymentId != nil {
		toSerialize["deploymentID"] = o.DeploymentId
	}
	toSerialize["taskId"] = o.TaskId
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentSwitchResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentId *string                         `json:"deploymentID,omitempty"`
		TaskId       *string                         `json:"taskId"`
		Steps        []BlueGreenDeploymentSwitchStep `json:"steps,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"deploymentID", "taskId", "steps"})
	} else {
		return err
	}
	o.DeploymentId = all.DeploymentId
	o.TaskId = *all.TaskId
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
