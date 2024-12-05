// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Feature struct {
	// Name of a feature
	Name string `json:"name"`
	// The enablement state for the feature
	Enabled bool `json:"enabled"`
	// Indicates the maturity level of a feature
	PreRelease FeaturePreReleaseType `json:"preRelease"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeature instantiates a new Feature object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeature(name string, enabled bool, preRelease FeaturePreReleaseType) *Feature {
	this := Feature{}
	this.Name = name
	this.Enabled = enabled
	this.PreRelease = preRelease
	return &this
}

// NewFeatureWithDefaults instantiates a new Feature object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureWithDefaults() *Feature {
	this := Feature{}
	return &this
}

// GetName returns the Name field value.
func (o *Feature) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Feature) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Feature) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value.
func (o *Feature) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Feature) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *Feature) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPreRelease returns the PreRelease field value.
func (o *Feature) GetPreRelease() FeaturePreReleaseType {
	if o == nil {
		var ret FeaturePreReleaseType
		return ret
	}
	return o.PreRelease
}

// GetPreReleaseOk returns a tuple with the PreRelease field value
// and a boolean to check if the value has been set.
func (o *Feature) GetPreReleaseOk() (*FeaturePreReleaseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreRelease, true
}

// SetPreRelease sets field value.
func (o *Feature) SetPreRelease(v FeaturePreReleaseType) {
	o.PreRelease = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Feature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	toSerialize["preRelease"] = o.PreRelease

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Feature) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string                `json:"name"`
		Enabled    *bool                  `json:"enabled"`
		PreRelease *FeaturePreReleaseType `json:"preRelease"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.PreRelease == nil {
		return fmt.Errorf("required field preRelease missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "enabled", "preRelease"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Enabled = *all.Enabled
	if !all.PreRelease.IsValid() {
		hasInvalidField = true
	} else {
		o.PreRelease = *all.PreRelease
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
