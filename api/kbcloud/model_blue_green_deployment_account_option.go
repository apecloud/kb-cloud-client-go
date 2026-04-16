// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type BlueGreenDeploymentAccountOption struct {
	Component          *string  `json:"component,omitempty"`
	DefaultAccountName []string `json:"defaultAccountName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentAccountOption instantiates a new BlueGreenDeploymentAccountOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentAccountOption() *BlueGreenDeploymentAccountOption {
	this := BlueGreenDeploymentAccountOption{}
	return &this
}

// NewBlueGreenDeploymentAccountOptionWithDefaults instantiates a new BlueGreenDeploymentAccountOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentAccountOptionWithDefaults() *BlueGreenDeploymentAccountOption {
	this := BlueGreenDeploymentAccountOption{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *BlueGreenDeploymentAccountOption) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentAccountOption) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *BlueGreenDeploymentAccountOption) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *BlueGreenDeploymentAccountOption) SetComponent(v string) {
	o.Component = &v
}

// GetDefaultAccountName returns the DefaultAccountName field value if set, zero value otherwise.
func (o *BlueGreenDeploymentAccountOption) GetDefaultAccountName() []string {
	if o == nil || o.DefaultAccountName == nil {
		var ret []string
		return ret
	}
	return o.DefaultAccountName
}

// GetDefaultAccountNameOk returns a tuple with the DefaultAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentAccountOption) GetDefaultAccountNameOk() (*[]string, bool) {
	if o == nil || o.DefaultAccountName == nil {
		return nil, false
	}
	return &o.DefaultAccountName, true
}

// HasDefaultAccountName returns a boolean if a field has been set.
func (o *BlueGreenDeploymentAccountOption) HasDefaultAccountName() bool {
	return o != nil && o.DefaultAccountName != nil
}

// SetDefaultAccountName gets a reference to the given []string and assigns it to the DefaultAccountName field.
func (o *BlueGreenDeploymentAccountOption) SetDefaultAccountName(v []string) {
	o.DefaultAccountName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentAccountOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.DefaultAccountName != nil {
		toSerialize["defaultAccountName"] = o.DefaultAccountName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentAccountOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component          *string  `json:"component,omitempty"`
		DefaultAccountName []string `json:"defaultAccountName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "defaultAccountName"})
	} else {
		return err
	}
	o.Component = all.Component
	o.DefaultAccountName = all.DefaultAccountName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
