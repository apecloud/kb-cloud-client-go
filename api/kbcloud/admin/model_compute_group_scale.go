// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComputeGroupScale struct {
	// Optional target name when enabling the default compute component; defaults to default_compute_group. Cannot rename an enabled group.
	Name *string `json:"name,omitempty"`
	// Set true to enable the default compute component as a compute group. Cannot be combined with replicas or classCode.
	ComputeGroupEnabled *bool   `json:"computeGroupEnabled,omitempty"`
	Replicas            *int64  `json:"replicas,omitempty"`
	ClassCode           *string `json:"classCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewComputeGroupScale instantiates a new ComputeGroupScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupScale() *ComputeGroupScale {
	this := ComputeGroupScale{}
	return &this
}

// NewComputeGroupScaleWithDefaults instantiates a new ComputeGroupScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupScaleWithDefaults() *ComputeGroupScale {
	this := ComputeGroupScale{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeGroupScale) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupScale) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeGroupScale) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeGroupScale) SetName(v string) {
	o.Name = &v
}

// GetComputeGroupEnabled returns the ComputeGroupEnabled field value if set, zero value otherwise.
func (o *ComputeGroupScale) GetComputeGroupEnabled() bool {
	if o == nil || o.ComputeGroupEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ComputeGroupEnabled
}

// GetComputeGroupEnabledOk returns a tuple with the ComputeGroupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupScale) GetComputeGroupEnabledOk() (*bool, bool) {
	if o == nil || o.ComputeGroupEnabled == nil {
		return nil, false
	}
	return o.ComputeGroupEnabled, true
}

// HasComputeGroupEnabled returns a boolean if a field has been set.
func (o *ComputeGroupScale) HasComputeGroupEnabled() bool {
	return o != nil && o.ComputeGroupEnabled != nil
}

// SetComputeGroupEnabled gets a reference to the given bool and assigns it to the ComputeGroupEnabled field.
func (o *ComputeGroupScale) SetComputeGroupEnabled(v bool) {
	o.ComputeGroupEnabled = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ComputeGroupScale) GetReplicas() int64 {
	if o == nil || o.Replicas == nil {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupScale) GetReplicasOk() (*int64, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ComputeGroupScale) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *ComputeGroupScale) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ComputeGroupScale) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupScale) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ComputeGroupScale) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ComputeGroupScale) SetClassCode(v string) {
	o.ClassCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ComputeGroupEnabled != nil {
		toSerialize["computeGroupEnabled"] = o.ComputeGroupEnabled
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                *string `json:"name,omitempty"`
		ComputeGroupEnabled *bool   `json:"computeGroupEnabled,omitempty"`
		Replicas            *int64  `json:"replicas,omitempty"`
		ClassCode           *string `json:"classCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	o.Name = all.Name
	o.ComputeGroupEnabled = all.ComputeGroupEnabled
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode

	return nil
}
