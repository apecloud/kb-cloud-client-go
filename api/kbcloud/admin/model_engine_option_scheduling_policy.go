// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineOptionSchedulingPolicy struct {
	// If set to true, resource distribution is at the cluster level.
	//
	ClusterScopeAntiAffinity *bool `json:"clusterScopeAntiAffinity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionSchedulingPolicy instantiates a new EngineOptionSchedulingPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionSchedulingPolicy() *EngineOptionSchedulingPolicy {
	this := EngineOptionSchedulingPolicy{}
	return &this
}

// NewEngineOptionSchedulingPolicyWithDefaults instantiates a new EngineOptionSchedulingPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionSchedulingPolicyWithDefaults() *EngineOptionSchedulingPolicy {
	this := EngineOptionSchedulingPolicy{}
	return &this
}

// GetClusterScopeAntiAffinity returns the ClusterScopeAntiAffinity field value if set, zero value otherwise.
func (o *EngineOptionSchedulingPolicy) GetClusterScopeAntiAffinity() bool {
	if o == nil || o.ClusterScopeAntiAffinity == nil {
		var ret bool
		return ret
	}
	return *o.ClusterScopeAntiAffinity
}

// GetClusterScopeAntiAffinityOk returns a tuple with the ClusterScopeAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionSchedulingPolicy) GetClusterScopeAntiAffinityOk() (*bool, bool) {
	if o == nil || o.ClusterScopeAntiAffinity == nil {
		return nil, false
	}
	return o.ClusterScopeAntiAffinity, true
}

// HasClusterScopeAntiAffinity returns a boolean if a field has been set.
func (o *EngineOptionSchedulingPolicy) HasClusterScopeAntiAffinity() bool {
	return o != nil && o.ClusterScopeAntiAffinity != nil
}

// SetClusterScopeAntiAffinity gets a reference to the given bool and assigns it to the ClusterScopeAntiAffinity field.
func (o *EngineOptionSchedulingPolicy) SetClusterScopeAntiAffinity(v bool) {
	o.ClusterScopeAntiAffinity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionSchedulingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterScopeAntiAffinity != nil {
		toSerialize["clusterScopeAntiAffinity"] = o.ClusterScopeAntiAffinity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionSchedulingPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterScopeAntiAffinity *bool `json:"clusterScopeAntiAffinity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterScopeAntiAffinity"})
	} else {
		return err
	}
	o.ClusterScopeAntiAffinity = all.ClusterScopeAntiAffinity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
