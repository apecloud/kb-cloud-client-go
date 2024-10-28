// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EnvironmentAdded struct {
	// whether the environment has ever been added
	AddedBefore *bool `json:"addedBefore,omitempty"`
	// whether the environment is still in use
	StillInUse *bool `json:"stillInUse,omitempty"`
	// Environment info
	Environment *Environment `json:"environment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentAdded instantiates a new EnvironmentAdded object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentAdded() *EnvironmentAdded {
	this := EnvironmentAdded{}
	return &this
}

// NewEnvironmentAddedWithDefaults instantiates a new EnvironmentAdded object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentAddedWithDefaults() *EnvironmentAdded {
	this := EnvironmentAdded{}
	return &this
}

// GetAddedBefore returns the AddedBefore field value if set, zero value otherwise.
func (o *EnvironmentAdded) GetAddedBefore() bool {
	if o == nil || o.AddedBefore == nil {
		var ret bool
		return ret
	}
	return *o.AddedBefore
}

// GetAddedBeforeOk returns a tuple with the AddedBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAdded) GetAddedBeforeOk() (*bool, bool) {
	if o == nil || o.AddedBefore == nil {
		return nil, false
	}
	return o.AddedBefore, true
}

// HasAddedBefore returns a boolean if a field has been set.
func (o *EnvironmentAdded) HasAddedBefore() bool {
	return o != nil && o.AddedBefore != nil
}

// SetAddedBefore gets a reference to the given bool and assigns it to the AddedBefore field.
func (o *EnvironmentAdded) SetAddedBefore(v bool) {
	o.AddedBefore = &v
}

// GetStillInUse returns the StillInUse field value if set, zero value otherwise.
func (o *EnvironmentAdded) GetStillInUse() bool {
	if o == nil || o.StillInUse == nil {
		var ret bool
		return ret
	}
	return *o.StillInUse
}

// GetStillInUseOk returns a tuple with the StillInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAdded) GetStillInUseOk() (*bool, bool) {
	if o == nil || o.StillInUse == nil {
		return nil, false
	}
	return o.StillInUse, true
}

// HasStillInUse returns a boolean if a field has been set.
func (o *EnvironmentAdded) HasStillInUse() bool {
	return o != nil && o.StillInUse != nil
}

// SetStillInUse gets a reference to the given bool and assigns it to the StillInUse field.
func (o *EnvironmentAdded) SetStillInUse(v bool) {
	o.StillInUse = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EnvironmentAdded) GetEnvironment() Environment {
	if o == nil || o.Environment == nil {
		var ret Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAdded) GetEnvironmentOk() (*Environment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentAdded) HasEnvironment() bool {
	return o != nil && o.Environment != nil
}

// SetEnvironment gets a reference to the given Environment and assigns it to the Environment field.
func (o *EnvironmentAdded) SetEnvironment(v Environment) {
	o.Environment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentAdded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AddedBefore != nil {
		toSerialize["addedBefore"] = o.AddedBefore
	}
	if o.StillInUse != nil {
		toSerialize["stillInUse"] = o.StillInUse
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentAdded) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddedBefore *bool        `json:"addedBefore,omitempty"`
		StillInUse  *bool        `json:"stillInUse,omitempty"`
		Environment *Environment `json:"environment,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"addedBefore", "stillInUse", "environment"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AddedBefore = all.AddedBefore
	o.StillInUse = all.StillInUse
	if all.Environment != nil && all.Environment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Environment = all.Environment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
