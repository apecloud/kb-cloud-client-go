// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateVScale VScale assignment for one instance template. name and classCode are required.
type InstanceTemplateVScale struct {
	// Instance template name declared on the engine option.
	Name string `json:"name"`
	// Class code to apply to this instance template.
	ClassCode string `json:"classCode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateVScale instantiates a new InstanceTemplateVScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateVScale(name string, classCode string) *InstanceTemplateVScale {
	this := InstanceTemplateVScale{}
	this.Name = name
	this.ClassCode = classCode
	return &this
}

// NewInstanceTemplateVScaleWithDefaults instantiates a new InstanceTemplateVScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateVScaleWithDefaults() *InstanceTemplateVScale {
	this := InstanceTemplateVScale{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateVScale) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateVScale) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateVScale) SetName(v string) {
	o.Name = v
}

// GetClassCode returns the ClassCode field value.
func (o *InstanceTemplateVScale) GetClassCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateVScale) GetClassCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassCode, true
}

// SetClassCode sets field value.
func (o *InstanceTemplateVScale) SetClassCode(v string) {
	o.ClassCode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateVScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["classCode"] = o.ClassCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateVScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		ClassCode *string `json:"classCode"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ClassCode == nil {
		return fmt.Errorf("required field classCode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "classCode"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.ClassCode = *all.ClassCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
