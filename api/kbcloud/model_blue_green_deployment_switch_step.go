// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BlueGreenDeploymentSwitchStep blueGreenDeploymentSwitchStep details a step in the blue-green deployment switch process.
type BlueGreenDeploymentSwitchStep struct {
	// The name of the step.
	StepName *string `json:"stepName,omitempty"`
	// The status of a step in a blue-green deployment switch.
	StepStatus *BlueGreenDeploymentSwitchStepStatus `json:"stepStatus,omitempty"`
	// The message of the step.
	StepMessage *string `json:"stepMessage,omitempty"`
	// Additional details of the step.
	StepDetail interface{} `json:"stepDetail,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentSwitchStep instantiates a new BlueGreenDeploymentSwitchStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentSwitchStep() *BlueGreenDeploymentSwitchStep {
	this := BlueGreenDeploymentSwitchStep{}
	return &this
}

// NewBlueGreenDeploymentSwitchStepWithDefaults instantiates a new BlueGreenDeploymentSwitchStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentSwitchStepWithDefaults() *BlueGreenDeploymentSwitchStep {
	this := BlueGreenDeploymentSwitchStep{}
	return &this
}

// GetStepName returns the StepName field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchStep) GetStepName() string {
	if o == nil || o.StepName == nil {
		var ret string
		return ret
	}
	return *o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchStep) GetStepNameOk() (*string, bool) {
	if o == nil || o.StepName == nil {
		return nil, false
	}
	return o.StepName, true
}

// HasStepName returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchStep) HasStepName() bool {
	return o != nil && o.StepName != nil
}

// SetStepName gets a reference to the given string and assigns it to the StepName field.
func (o *BlueGreenDeploymentSwitchStep) SetStepName(v string) {
	o.StepName = &v
}

// GetStepStatus returns the StepStatus field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchStep) GetStepStatus() BlueGreenDeploymentSwitchStepStatus {
	if o == nil || o.StepStatus == nil {
		var ret BlueGreenDeploymentSwitchStepStatus
		return ret
	}
	return *o.StepStatus
}

// GetStepStatusOk returns a tuple with the StepStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchStep) GetStepStatusOk() (*BlueGreenDeploymentSwitchStepStatus, bool) {
	if o == nil || o.StepStatus == nil {
		return nil, false
	}
	return o.StepStatus, true
}

// HasStepStatus returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchStep) HasStepStatus() bool {
	return o != nil && o.StepStatus != nil
}

// SetStepStatus gets a reference to the given BlueGreenDeploymentSwitchStepStatus and assigns it to the StepStatus field.
func (o *BlueGreenDeploymentSwitchStep) SetStepStatus(v BlueGreenDeploymentSwitchStepStatus) {
	o.StepStatus = &v
}

// GetStepMessage returns the StepMessage field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchStep) GetStepMessage() string {
	if o == nil || o.StepMessage == nil {
		var ret string
		return ret
	}
	return *o.StepMessage
}

// GetStepMessageOk returns a tuple with the StepMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchStep) GetStepMessageOk() (*string, bool) {
	if o == nil || o.StepMessage == nil {
		return nil, false
	}
	return o.StepMessage, true
}

// HasStepMessage returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchStep) HasStepMessage() bool {
	return o != nil && o.StepMessage != nil
}

// SetStepMessage gets a reference to the given string and assigns it to the StepMessage field.
func (o *BlueGreenDeploymentSwitchStep) SetStepMessage(v string) {
	o.StepMessage = &v
}

// GetStepDetail returns the StepDetail field value if set, zero value otherwise.
func (o *BlueGreenDeploymentSwitchStep) GetStepDetail() interface{} {
	if o == nil || o.StepDetail == nil {
		var ret interface{}
		return ret
	}
	return o.StepDetail
}

// GetStepDetailOk returns a tuple with the StepDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentSwitchStep) GetStepDetailOk() (*interface{}, bool) {
	if o == nil || o.StepDetail == nil {
		return nil, false
	}
	return &o.StepDetail, true
}

// HasStepDetail returns a boolean if a field has been set.
func (o *BlueGreenDeploymentSwitchStep) HasStepDetail() bool {
	return o != nil && o.StepDetail != nil
}

// SetStepDetail gets a reference to the given interface{} and assigns it to the StepDetail field.
func (o *BlueGreenDeploymentSwitchStep) SetStepDetail(v interface{}) {
	o.StepDetail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentSwitchStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StepName != nil {
		toSerialize["stepName"] = o.StepName
	}
	if o.StepStatus != nil {
		toSerialize["stepStatus"] = o.StepStatus
	}
	if o.StepMessage != nil {
		toSerialize["stepMessage"] = o.StepMessage
	}
	if o.StepDetail != nil {
		toSerialize["stepDetail"] = o.StepDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentSwitchStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StepName    *string                              `json:"stepName,omitempty"`
		StepStatus  *BlueGreenDeploymentSwitchStepStatus `json:"stepStatus,omitempty"`
		StepMessage *string                              `json:"stepMessage,omitempty"`
		StepDetail  interface{}                          `json:"stepDetail,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"stepName", "stepStatus", "stepMessage", "stepDetail"})
	} else {
		return err
	}

	hasInvalidField := false
	o.StepName = all.StepName
	if all.StepStatus != nil && !all.StepStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.StepStatus = all.StepStatus
	}
	o.StepMessage = all.StepMessage
	o.StepDetail = all.StepDetail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
