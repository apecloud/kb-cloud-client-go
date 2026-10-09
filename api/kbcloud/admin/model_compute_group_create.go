// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComputeGroupCreate struct {
	Name      string `json:"name"`
	Replicas  int64  `json:"replicas"`
	ClassCode string `json:"classCode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComputeGroupCreate instantiates a new ComputeGroupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupCreate(name string, replicas int64, classCode string) *ComputeGroupCreate {
	this := ComputeGroupCreate{}
	this.Name = name
	this.Replicas = replicas
	this.ClassCode = classCode
	return &this
}

// NewComputeGroupCreateWithDefaults instantiates a new ComputeGroupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupCreateWithDefaults() *ComputeGroupCreate {
	this := ComputeGroupCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *ComputeGroupCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ComputeGroupCreate) SetName(v string) {
	o.Name = v
}

// GetReplicas returns the Replicas field value.
func (o *ComputeGroupCreate) GetReplicas() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupCreate) GetReplicasOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *ComputeGroupCreate) SetReplicas(v int64) {
	o.Replicas = v
}

// GetClassCode returns the ClassCode field value.
func (o *ComputeGroupCreate) GetClassCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupCreate) GetClassCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassCode, true
}

// SetClassCode sets field value.
func (o *ComputeGroupCreate) SetClassCode(v string) {
	o.ClassCode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["replicas"] = o.Replicas
	toSerialize["classCode"] = o.ClassCode
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		Replicas  *int64  `json:"replicas"`
		ClassCode *string `json:"classCode"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	if all.ClassCode == nil {
		return fmt.Errorf("required field classCode missing")
	}
	o.Name = *all.Name
	o.Replicas = *all.Replicas
	o.ClassCode = *all.ClassCode

	return nil
}
