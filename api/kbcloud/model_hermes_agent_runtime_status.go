// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentRuntimeStatus struct {
	RuntimeStatus      HermesAgentRuntimeStatusCode `json:"runtimeStatus"`
	ModelConfigStatus  HermesAgentModelConfigStatus `json:"modelConfigStatus"`
	EndpointConfigured bool                         `json:"endpointConfigured"`
	Model              *string                      `json:"model,omitempty"`
	Message            *string                      `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentRuntimeStatus instantiates a new HermesAgentRuntimeStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentRuntimeStatus(runtimeStatus HermesAgentRuntimeStatusCode, modelConfigStatus HermesAgentModelConfigStatus, endpointConfigured bool) *HermesAgentRuntimeStatus {
	this := HermesAgentRuntimeStatus{}
	this.RuntimeStatus = runtimeStatus
	this.ModelConfigStatus = modelConfigStatus
	this.EndpointConfigured = endpointConfigured
	return &this
}

// NewHermesAgentRuntimeStatusWithDefaults instantiates a new HermesAgentRuntimeStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentRuntimeStatusWithDefaults() *HermesAgentRuntimeStatus {
	this := HermesAgentRuntimeStatus{}
	return &this
}

// GetRuntimeStatus returns the RuntimeStatus field value.
func (o *HermesAgentRuntimeStatus) GetRuntimeStatus() HermesAgentRuntimeStatusCode {
	if o == nil {
		var ret HermesAgentRuntimeStatusCode
		return ret
	}
	return o.RuntimeStatus
}

// GetRuntimeStatusOk returns a tuple with the RuntimeStatus field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRuntimeStatus) GetRuntimeStatusOk() (*HermesAgentRuntimeStatusCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuntimeStatus, true
}

// SetRuntimeStatus sets field value.
func (o *HermesAgentRuntimeStatus) SetRuntimeStatus(v HermesAgentRuntimeStatusCode) {
	o.RuntimeStatus = v
}

// GetModelConfigStatus returns the ModelConfigStatus field value.
func (o *HermesAgentRuntimeStatus) GetModelConfigStatus() HermesAgentModelConfigStatus {
	if o == nil {
		var ret HermesAgentModelConfigStatus
		return ret
	}
	return o.ModelConfigStatus
}

// GetModelConfigStatusOk returns a tuple with the ModelConfigStatus field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRuntimeStatus) GetModelConfigStatusOk() (*HermesAgentModelConfigStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelConfigStatus, true
}

// SetModelConfigStatus sets field value.
func (o *HermesAgentRuntimeStatus) SetModelConfigStatus(v HermesAgentModelConfigStatus) {
	o.ModelConfigStatus = v
}

// GetEndpointConfigured returns the EndpointConfigured field value.
func (o *HermesAgentRuntimeStatus) GetEndpointConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EndpointConfigured
}

// GetEndpointConfiguredOk returns a tuple with the EndpointConfigured field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRuntimeStatus) GetEndpointConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointConfigured, true
}

// SetEndpointConfigured sets field value.
func (o *HermesAgentRuntimeStatus) SetEndpointConfigured(v bool) {
	o.EndpointConfigured = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HermesAgentRuntimeStatus) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRuntimeStatus) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HermesAgentRuntimeStatus) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HermesAgentRuntimeStatus) SetModel(v string) {
	o.Model = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HermesAgentRuntimeStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRuntimeStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HermesAgentRuntimeStatus) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HermesAgentRuntimeStatus) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentRuntimeStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runtimeStatus"] = o.RuntimeStatus
	toSerialize["modelConfigStatus"] = o.ModelConfigStatus
	toSerialize["endpointConfigured"] = o.EndpointConfigured
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentRuntimeStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuntimeStatus      *HermesAgentRuntimeStatusCode `json:"runtimeStatus"`
		ModelConfigStatus  *HermesAgentModelConfigStatus `json:"modelConfigStatus"`
		EndpointConfigured *bool                         `json:"endpointConfigured"`
		Model              *string                       `json:"model,omitempty"`
		Message            *string                       `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RuntimeStatus == nil {
		return fmt.Errorf("required field runtimeStatus missing")
	}
	if all.ModelConfigStatus == nil {
		return fmt.Errorf("required field modelConfigStatus missing")
	}
	if all.EndpointConfigured == nil {
		return fmt.Errorf("required field endpointConfigured missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runtimeStatus", "modelConfigStatus", "endpointConfigured", "model", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.RuntimeStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.RuntimeStatus = *all.RuntimeStatus
	}
	if !all.ModelConfigStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ModelConfigStatus = *all.ModelConfigStatus
	}
	o.EndpointConfigured = *all.EndpointConfigured
	o.Model = all.Model
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
