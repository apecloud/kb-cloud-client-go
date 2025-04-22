// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NodeAffinity Rules for scheduling pods based on node labels.
type NodeAffinity struct {
	// Specifies the hard affinity/anti-affinity rules for pod scheduling.
	RequiredDuringSchedulingIgnoredDuringExecution *RequiredSchedulingRule `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeAffinity instantiates a new NodeAffinity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeAffinity() *NodeAffinity {
	this := NodeAffinity{}
	return &this
}

// NewNodeAffinityWithDefaults instantiates a new NodeAffinity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeAffinityWithDefaults() *NodeAffinity {
	this := NodeAffinity{}
	return &this
}

// GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() RequiredSchedulingRule {
	if o == nil || o.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		var ret RequiredSchedulingRule
		return ret
	}
	return *o.RequiredDuringSchedulingIgnoredDuringExecution
}

// GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*RequiredSchedulingRule, bool) {
	if o == nil || o.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		return nil, false
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution, true
}

// HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *NodeAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool {
	return o != nil && o.RequiredDuringSchedulingIgnoredDuringExecution != nil
}

// SetRequiredDuringSchedulingIgnoredDuringExecution gets a reference to the given RequiredSchedulingRule and assigns it to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (o *NodeAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v RequiredSchedulingRule) {
	o.RequiredDuringSchedulingIgnoredDuringExecution = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeAffinity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		toSerialize["requiredDuringSchedulingIgnoredDuringExecution"] = o.RequiredDuringSchedulingIgnoredDuringExecution
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeAffinity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequiredDuringSchedulingIgnoredDuringExecution *RequiredSchedulingRule `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"requiredDuringSchedulingIgnoredDuringExecution"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RequiredDuringSchedulingIgnoredDuringExecution != nil && all.RequiredDuringSchedulingIgnoredDuringExecution.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RequiredDuringSchedulingIgnoredDuringExecution = all.RequiredDuringSchedulingIgnoredDuringExecution

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
