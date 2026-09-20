// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateReplicas HScale assignment for one instance template. name and replicas only.
type InstanceTemplateReplicas struct {
	// Instance template name declared on the engine option.
	Name string `json:"name"`
	// Target replica count for this instance template.
	Replicas int32 `json:"replicas"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateReplicas instantiates a new InstanceTemplateReplicas object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateReplicas(name string, replicas int32) *InstanceTemplateReplicas {
	this := InstanceTemplateReplicas{}
	this.Name = name
	this.Replicas = replicas
	return &this
}

// NewInstanceTemplateReplicasWithDefaults instantiates a new InstanceTemplateReplicas object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateReplicasWithDefaults() *InstanceTemplateReplicas {
	this := InstanceTemplateReplicas{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateReplicas) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateReplicas) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateReplicas) SetName(v string) {
	o.Name = v
}

// GetReplicas returns the Replicas field value.
func (o *InstanceTemplateReplicas) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateReplicas) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *InstanceTemplateReplicas) SetReplicas(v int32) {
	o.Replicas = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateReplicas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["replicas"] = o.Replicas

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateReplicas) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name"`
		Replicas *int32  `json:"replicas"`
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "replicas"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Replicas = *all.Replicas

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
