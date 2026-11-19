// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentOpsOptionRollback Horizontal scaling rollback capability. Omitted directions preserve existing cancellation behavior; other cancellation restrictions still apply.
type ComponentOpsOptionRollback struct {
	// Whether execution of scale-in can be rolled back. Omitted means allowed.
	ScaleIn *bool `json:"scaleIn,omitempty"`
	// Whether execution of scale-out can be rolled back. Omitted means allowed.
	ScaleOut *bool `json:"scaleOut,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionRollback instantiates a new ComponentOpsOptionRollback object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionRollback() *ComponentOpsOptionRollback {
	this := ComponentOpsOptionRollback{}
	return &this
}

// NewComponentOpsOptionRollbackWithDefaults instantiates a new ComponentOpsOptionRollback object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionRollbackWithDefaults() *ComponentOpsOptionRollback {
	this := ComponentOpsOptionRollback{}
	return &this
}

// GetScaleIn returns the ScaleIn field value if set, zero value otherwise.
func (o *ComponentOpsOptionRollback) GetScaleIn() bool {
	if o == nil || o.ScaleIn == nil {
		var ret bool
		return ret
	}
	return *o.ScaleIn
}

// GetScaleInOk returns a tuple with the ScaleIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionRollback) GetScaleInOk() (*bool, bool) {
	if o == nil || o.ScaleIn == nil {
		return nil, false
	}
	return o.ScaleIn, true
}

// HasScaleIn returns a boolean if a field has been set.
func (o *ComponentOpsOptionRollback) HasScaleIn() bool {
	return o != nil && o.ScaleIn != nil
}

// SetScaleIn gets a reference to the given bool and assigns it to the ScaleIn field.
func (o *ComponentOpsOptionRollback) SetScaleIn(v bool) {
	o.ScaleIn = &v
}

// GetScaleOut returns the ScaleOut field value if set, zero value otherwise.
func (o *ComponentOpsOptionRollback) GetScaleOut() bool {
	if o == nil || o.ScaleOut == nil {
		var ret bool
		return ret
	}
	return *o.ScaleOut
}

// GetScaleOutOk returns a tuple with the ScaleOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionRollback) GetScaleOutOk() (*bool, bool) {
	if o == nil || o.ScaleOut == nil {
		return nil, false
	}
	return o.ScaleOut, true
}

// HasScaleOut returns a boolean if a field has been set.
func (o *ComponentOpsOptionRollback) HasScaleOut() bool {
	return o != nil && o.ScaleOut != nil
}

// SetScaleOut gets a reference to the given bool and assigns it to the ScaleOut field.
func (o *ComponentOpsOptionRollback) SetScaleOut(v bool) {
	o.ScaleOut = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionRollback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ScaleIn != nil {
		toSerialize["scaleIn"] = o.ScaleIn
	}
	if o.ScaleOut != nil {
		toSerialize["scaleOut"] = o.ScaleOut
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionRollback) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScaleIn  *bool `json:"scaleIn,omitempty"`
		ScaleOut *bool `json:"scaleOut,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scaleIn", "scaleOut"})
	} else {
		return err
	}
	o.ScaleIn = all.ScaleIn
	o.ScaleOut = all.ScaleOut

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
