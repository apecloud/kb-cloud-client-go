// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EnvironmentStatus Environment status
type EnvironmentStatus struct {
	// Current service state of environment.
	Conditions []EnvironmentCondition `json:"conditions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStatus instantiates a new EnvironmentStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStatus() *EnvironmentStatus {
	this := EnvironmentStatus{}
	return &this
}

// NewEnvironmentStatusWithDefaults instantiates a new EnvironmentStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStatusWithDefaults() *EnvironmentStatus {
	this := EnvironmentStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *EnvironmentStatus) GetConditions() []EnvironmentCondition {
	if o == nil || o.Conditions == nil {
		var ret []EnvironmentCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStatus) GetConditionsOk() (*[]EnvironmentCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *EnvironmentStatus) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []EnvironmentCondition and assigns it to the Conditions field.
func (o *EnvironmentStatus) SetConditions(v []EnvironmentCondition) {
	o.Conditions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions []EnvironmentCondition `json:"conditions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conditions"})
	} else {
		return err
	}
	o.Conditions = all.Conditions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
