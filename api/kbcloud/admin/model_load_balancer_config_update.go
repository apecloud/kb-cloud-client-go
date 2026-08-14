// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerConfigUpdate Complete editable load balancer annotations for an environment.
type LoadBalancerConfigUpdate struct {
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
func NewLoadBalancerConfigUpdate(annotations LoadBalancerAnnotations) *LoadBalancerConfigUpdate {
	this := LoadBalancerConfigUpdate{}
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
	toSerialize["annotations"] = o.Annotations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancerConfigUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations *LoadBalancerAnnotations `json:"annotations"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"annotations"})
	} else {
		return err
	}

	hasInvalidField := false
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
