// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type DisasterRecoveryOption struct {
	Enabled       *bool  `json:"enabled,omitempty"`
	InstanceLimit *int32 `json:"instanceLimit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryOption instantiates a new DisasterRecoveryOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryOption() *DisasterRecoveryOption {
	this := DisasterRecoveryOption{}
	var enabled bool = false
	this.Enabled = &enabled
	var instanceLimit int32 = 8
	this.InstanceLimit = &instanceLimit
	return &this
}

// NewDisasterRecoveryOptionWithDefaults instantiates a new DisasterRecoveryOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryOptionWithDefaults() *DisasterRecoveryOption {
	this := DisasterRecoveryOption{}
	var enabled bool = false
	this.Enabled = &enabled
	var instanceLimit int32 = 8
	this.InstanceLimit = &instanceLimit
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DisasterRecoveryOption) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryOption) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DisasterRecoveryOption) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DisasterRecoveryOption) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInstanceLimit returns the InstanceLimit field value if set, zero value otherwise.
func (o *DisasterRecoveryOption) GetInstanceLimit() int32 {
	if o == nil || o.InstanceLimit == nil {
		var ret int32
		return ret
	}
	return *o.InstanceLimit
}

// GetInstanceLimitOk returns a tuple with the InstanceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryOption) GetInstanceLimitOk() (*int32, bool) {
	if o == nil || o.InstanceLimit == nil {
		return nil, false
	}
	return o.InstanceLimit, true
}

// HasInstanceLimit returns a boolean if a field has been set.
func (o *DisasterRecoveryOption) HasInstanceLimit() bool {
	return o != nil && o.InstanceLimit != nil
}

// SetInstanceLimit gets a reference to the given int32 and assigns it to the InstanceLimit field.
func (o *DisasterRecoveryOption) SetInstanceLimit(v int32) {
	o.InstanceLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.InstanceLimit != nil {
		toSerialize["instanceLimit"] = o.InstanceLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool  `json:"enabled,omitempty"`
		InstanceLimit *int32 `json:"instanceLimit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "instanceLimit"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.InstanceLimit = all.InstanceLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
