// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerConfigUpdate Complete editable load balancer settings for an environment.
type LoadBalancerConfigUpdate struct {
	// Provider selected for the load balancer.
	Provider string `json:"provider"`
	// Follow admin_environment.provider instead of pinning the selected provider.
	ProviderInherited bool `json:"providerInherited"`
	// Provider-neutral annotations applied to load balancer Services by exposure type.
	Annotations LoadBalancerAnnotations `json:"annotations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancerConfigUpdate instantiates a new LoadBalancerConfigUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancerConfigUpdate(provider string, providerInherited bool, annotations LoadBalancerAnnotations) *LoadBalancerConfigUpdate {
	this := LoadBalancerConfigUpdate{}
	this.Provider = provider
	this.ProviderInherited = providerInherited
	this.Annotations = annotations
	return &this
}

// NewLoadBalancerConfigUpdateWithDefaults instantiates a new LoadBalancerConfigUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerConfigUpdateWithDefaults() *LoadBalancerConfigUpdate {
	this := LoadBalancerConfigUpdate{}
	return &this
}

// GetProvider returns the Provider field value.
func (o *LoadBalancerConfigUpdate) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerConfigUpdate) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *LoadBalancerConfigUpdate) SetProvider(v string) {
	o.Provider = v
}

// GetProviderInherited returns the ProviderInherited field value.
func (o *LoadBalancerConfigUpdate) GetProviderInherited() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ProviderInherited
}

// GetProviderInheritedOk returns a tuple with the ProviderInherited field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerConfigUpdate) GetProviderInheritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderInherited, true
}

// SetProviderInherited sets field value.
func (o *LoadBalancerConfigUpdate) SetProviderInherited(v bool) {
	o.ProviderInherited = v
}

// GetAnnotations returns the Annotations field value.
func (o *LoadBalancerConfigUpdate) GetAnnotations() LoadBalancerAnnotations {
	if o == nil {
		var ret LoadBalancerAnnotations
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerConfigUpdate) GetAnnotationsOk() (*LoadBalancerAnnotations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LoadBalancerConfigUpdate) SetAnnotations(v LoadBalancerAnnotations) {
	o.Annotations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancerConfigUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provider"] = o.Provider
	toSerialize["providerInherited"] = o.ProviderInherited
	toSerialize["annotations"] = o.Annotations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancerConfigUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider          *string                  `json:"provider"`
		ProviderInherited *bool                    `json:"providerInherited"`
		Annotations       *LoadBalancerAnnotations `json:"annotations"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.ProviderInherited == nil {
		return fmt.Errorf("required field providerInherited missing")
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provider", "providerInherited", "annotations"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Provider = *all.Provider
	o.ProviderInherited = *all.ProviderInherited
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
