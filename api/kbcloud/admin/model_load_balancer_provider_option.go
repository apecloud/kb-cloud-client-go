// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerProviderOption A provider that can be selected for an environment load balancer.
type LoadBalancerProviderOption struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	// Provider-neutral annotations applied to load balancer Services by exposure type.
	Annotations LoadBalancerAnnotations `json:"annotations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancerProviderOption instantiates a new LoadBalancerProviderOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancerProviderOption(name string, enabled bool, annotations LoadBalancerAnnotations) *LoadBalancerProviderOption {
	this := LoadBalancerProviderOption{}
	this.Name = name
	this.Enabled = enabled
	this.Annotations = annotations
	return &this
}

// NewLoadBalancerProviderOptionWithDefaults instantiates a new LoadBalancerProviderOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerProviderOptionWithDefaults() *LoadBalancerProviderOption {
	this := LoadBalancerProviderOption{}
	return &this
}

// GetName returns the Name field value.
func (o *LoadBalancerProviderOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerProviderOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LoadBalancerProviderOption) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value.
func (o *LoadBalancerProviderOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerProviderOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *LoadBalancerProviderOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAnnotations returns the Annotations field value.
func (o *LoadBalancerProviderOption) GetAnnotations() LoadBalancerAnnotations {
	if o == nil {
		var ret LoadBalancerAnnotations
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerProviderOption) GetAnnotationsOk() (*LoadBalancerAnnotations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LoadBalancerProviderOption) SetAnnotations(v LoadBalancerAnnotations) {
	o.Annotations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancerProviderOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	toSerialize["annotations"] = o.Annotations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancerProviderOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                  `json:"name"`
		Enabled     *bool                    `json:"enabled"`
		Annotations *LoadBalancerAnnotations `json:"annotations"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "enabled", "annotations"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Enabled = *all.Enabled
	if all.Annotations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Annotations = *all.Annotations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
