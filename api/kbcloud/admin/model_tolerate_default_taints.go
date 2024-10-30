// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// TolerateDefaultTaints When creating a cluster, add the default tolerations from the bootstrap node to the pods 
type TolerateDefaultTaints struct {
	Enabled *bool `json:"enabled,omitempty"`
	Taints []string `json:"taints,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewTolerateDefaultTaints instantiates a new TolerateDefaultTaints object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTolerateDefaultTaints() *TolerateDefaultTaints {
	this := TolerateDefaultTaints{}
	return &this
}

// NewTolerateDefaultTaintsWithDefaults instantiates a new TolerateDefaultTaints object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTolerateDefaultTaintsWithDefaults() *TolerateDefaultTaints {
	this := TolerateDefaultTaints{}
	return &this
}
// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TolerateDefaultTaints) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TolerateDefaultTaints) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TolerateDefaultTaints) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TolerateDefaultTaints) SetEnabled(v bool) {
	o.Enabled = &v
}


// GetTaints returns the Taints field value if set, zero value otherwise.
func (o *TolerateDefaultTaints) GetTaints() []string {
	if o == nil || o.Taints == nil {
		var ret []string
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TolerateDefaultTaints) GetTaintsOk() (*[]string, bool) {
	if o == nil || o.Taints == nil {
		return nil, false
	}
	return &o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *TolerateDefaultTaints) HasTaints() bool {
	return o != nil && o.Taints != nil
}

// SetTaints gets a reference to the given []string and assigns it to the Taints field.
func (o *TolerateDefaultTaints) SetTaints(v []string) {
	o.Taints = v
}



// MarshalJSON serializes the struct using spec logic.
func (o TolerateDefaultTaints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Taints != nil {
		toSerialize["taints"] = o.Taints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TolerateDefaultTaints) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool `json:"enabled,omitempty"`
		Taints []string `json:"taints,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "enabled", "taints",  })
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Taints = all.Taints

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
