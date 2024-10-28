// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ReconfigureCreate ReconfigureCreate is the payload to reconfigure a KubeBlocks cluster

type ReconfigureCreate struct {
	// component type
	Component *string `json:"component,omitempty"`
	// config file name
	ConfigFileName *string `json:"configFileName,omitempty"`
	// Specify parameters list to be updated
	Parameters map[string]string `json:"parameters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReconfigureCreate instantiates a new ReconfigureCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReconfigureCreate(parameters map[string]string) *ReconfigureCreate {
	this := ReconfigureCreate{}
	this.Parameters = parameters
	return &this
}

// NewReconfigureCreateWithDefaults instantiates a new ReconfigureCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReconfigureCreateWithDefaults() *ReconfigureCreate {
	this := ReconfigureCreate{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ReconfigureCreate) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReconfigureCreate) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ReconfigureCreate) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ReconfigureCreate) SetComponent(v string) {
	o.Component = &v
}

// GetConfigFileName returns the ConfigFileName field value if set, zero value otherwise.
func (o *ReconfigureCreate) GetConfigFileName() string {
	if o == nil || o.ConfigFileName == nil {
		var ret string
		return ret
	}
	return *o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReconfigureCreate) GetConfigFileNameOk() (*string, bool) {
	if o == nil || o.ConfigFileName == nil {
		return nil, false
	}
	return o.ConfigFileName, true
}

// HasConfigFileName returns a boolean if a field has been set.
func (o *ReconfigureCreate) HasConfigFileName() bool {
	return o != nil && o.ConfigFileName != nil
}

// SetConfigFileName gets a reference to the given string and assigns it to the ConfigFileName field.
func (o *ReconfigureCreate) SetConfigFileName(v string) {
	o.ConfigFileName = &v
}

// GetParameters returns the Parameters field value.
func (o *ReconfigureCreate) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *ReconfigureCreate) GetParametersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *ReconfigureCreate) SetParameters(v map[string]string) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReconfigureCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ConfigFileName != nil {
		toSerialize["configFileName"] = o.ConfigFileName
	}
	toSerialize["parameters"] = o.Parameters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReconfigureCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component      *string            `json:"component,omitempty"`
		ConfigFileName *string            `json:"configFileName,omitempty"`
		Parameters     *map[string]string `json:"parameters"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "configFileName", "parameters"})
	} else {
		return err
	}
	o.Component = all.Component
	o.ConfigFileName = all.ConfigFileName
	o.Parameters = *all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
