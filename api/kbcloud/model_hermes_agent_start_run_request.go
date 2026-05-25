// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentStartRunRequest struct {
	Message      string  `json:"message"`
	ClusterName  *string `json:"clusterName,omitempty"`
	Namespace    *string `json:"namespace,omitempty"`
	Model        *string `json:"model,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentStartRunRequest instantiates a new HermesAgentStartRunRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentStartRunRequest(message string) *HermesAgentStartRunRequest {
	this := HermesAgentStartRunRequest{}
	this.Message = message
	return &this
}

// NewHermesAgentStartRunRequestWithDefaults instantiates a new HermesAgentStartRunRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentStartRunRequestWithDefaults() *HermesAgentStartRunRequest {
	this := HermesAgentStartRunRequest{}
	return &this
}

// GetMessage returns the Message field value.
func (o *HermesAgentStartRunRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *HermesAgentStartRunRequest) SetMessage(v string) {
	o.Message = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HermesAgentStartRunRequest) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunRequest) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HermesAgentStartRunRequest) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HermesAgentStartRunRequest) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *HermesAgentStartRunRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *HermesAgentStartRunRequest) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *HermesAgentStartRunRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HermesAgentStartRunRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HermesAgentStartRunRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HermesAgentStartRunRequest) SetModel(v string) {
	o.Model = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *HermesAgentStartRunRequest) GetInstructions() string {
	if o == nil || o.Instructions == nil {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunRequest) GetInstructionsOk() (*string, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *HermesAgentStartRunRequest) HasInstructions() bool {
	return o != nil && o.Instructions != nil
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *HermesAgentStartRunRequest) SetInstructions(v string) {
	o.Instructions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentStartRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentStartRunRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message      *string `json:"message"`
		ClusterName  *string `json:"clusterName,omitempty"`
		Namespace    *string `json:"namespace,omitempty"`
		Model        *string `json:"model,omitempty"`
		Instructions *string `json:"instructions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "clusterName", "namespace", "model", "instructions"})
	} else {
		return err
	}
	o.Message = *all.Message
	o.ClusterName = all.ClusterName
	o.Namespace = all.Namespace
	o.Model = all.Model
	o.Instructions = all.Instructions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
