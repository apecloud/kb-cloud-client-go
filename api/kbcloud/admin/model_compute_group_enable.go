// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComputeGroupEnable struct {
	// Existing compute component to enable.
	ComponentName string `json:"componentName"`
	// Target compute group name. An enabled component cannot be renamed.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComputeGroupEnable instantiates a new ComputeGroupEnable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupEnable(componentName string, name string) *ComputeGroupEnable {
	this := ComputeGroupEnable{}
	this.ComponentName = componentName
	this.Name = name
	return &this
}

// NewComputeGroupEnableWithDefaults instantiates a new ComputeGroupEnable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupEnableWithDefaults() *ComputeGroupEnable {
	this := ComputeGroupEnable{}
	return &this
}

// GetComponentName returns the ComponentName field value.
func (o *ComputeGroupEnable) GetComponentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupEnable) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value.
func (o *ComputeGroupEnable) SetComponentName(v string) {
	o.ComponentName = v
}

// GetName returns the Name field value.
func (o *ComputeGroupEnable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupEnable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ComputeGroupEnable) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupEnable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["componentName"] = o.ComponentName
	toSerialize["name"] = o.Name
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupEnable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentName *string `json:"componentName"`
		Name          *string `json:"name"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ComponentName == nil {
		return fmt.Errorf("required field componentName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.ComponentName = *all.ComponentName
	o.Name = *all.Name

	return nil
}
