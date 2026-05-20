// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TdeRequest Request to enable transparent data encryption. Disabling TDE is not supported.
type TdeRequest struct {
	// Component type to update TDE for
	Component string `json:"component"`
	// Whether to enable TDE. Must be true because disabling TDE is not supported.
	Enable bool `json:"enable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTdeRequest instantiates a new TdeRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTdeRequest(component string, enable bool) *TdeRequest {
	this := TdeRequest{}
	this.Component = component
	this.Enable = enable
	return &this
}

// NewTdeRequestWithDefaults instantiates a new TdeRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTdeRequestWithDefaults() *TdeRequest {
	this := TdeRequest{}
	return &this
}

// GetComponent returns the Component field value.
func (o *TdeRequest) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *TdeRequest) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *TdeRequest) SetComponent(v string) {
	o.Component = v
}

// GetEnable returns the Enable field value.
func (o *TdeRequest) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *TdeRequest) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *TdeRequest) SetEnable(v bool) {
	o.Enable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TdeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["enable"] = o.Enable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TdeRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		Enable    *bool   `json:"enable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "enable"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Enable = *all.Enable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
