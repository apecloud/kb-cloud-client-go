// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RequiredSchedulingRule Specifies the hard affinity/anti-affinity rules for pod scheduling.
type RequiredSchedulingRule struct {
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms,omitempty"`
	// Represents a label selector for matching pods based on their labels.
	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`
	// The key used to denote how Pods are grouped together (e.g., by node).
	TopologyKey *string `json:"topologyKey,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRequiredSchedulingRule instantiates a new RequiredSchedulingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRequiredSchedulingRule() *RequiredSchedulingRule {
	this := RequiredSchedulingRule{}
	return &this
}

// NewRequiredSchedulingRuleWithDefaults instantiates a new RequiredSchedulingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRequiredSchedulingRuleWithDefaults() *RequiredSchedulingRule {
	this := RequiredSchedulingRule{}
	return &this
}

// GetNodeSelectorTerms returns the NodeSelectorTerms field value if set, zero value otherwise.
func (o *RequiredSchedulingRule) GetNodeSelectorTerms() []NodeSelectorTerm {
	if o == nil || o.NodeSelectorTerms == nil {
		var ret []NodeSelectorTerm
		return ret
	}
	return o.NodeSelectorTerms
}

// GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredSchedulingRule) GetNodeSelectorTermsOk() (*[]NodeSelectorTerm, bool) {
	if o == nil || o.NodeSelectorTerms == nil {
		return nil, false
	}
	return &o.NodeSelectorTerms, true
}

// HasNodeSelectorTerms returns a boolean if a field has been set.
func (o *RequiredSchedulingRule) HasNodeSelectorTerms() bool {
	return o != nil && o.NodeSelectorTerms != nil
}

// SetNodeSelectorTerms gets a reference to the given []NodeSelectorTerm and assigns it to the NodeSelectorTerms field.
func (o *RequiredSchedulingRule) SetNodeSelectorTerms(v []NodeSelectorTerm) {
	o.NodeSelectorTerms = v
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *RequiredSchedulingRule) GetLabelSelector() LabelSelector {
	if o == nil || o.LabelSelector == nil {
		var ret LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredSchedulingRule) GetLabelSelectorOk() (*LabelSelector, bool) {
	if o == nil || o.LabelSelector == nil {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *RequiredSchedulingRule) HasLabelSelector() bool {
	return o != nil && o.LabelSelector != nil
}

// SetLabelSelector gets a reference to the given LabelSelector and assigns it to the LabelSelector field.
func (o *RequiredSchedulingRule) SetLabelSelector(v LabelSelector) {
	o.LabelSelector = &v
}

// GetTopologyKey returns the TopologyKey field value if set, zero value otherwise.
func (o *RequiredSchedulingRule) GetTopologyKey() string {
	if o == nil || o.TopologyKey == nil {
		var ret string
		return ret
	}
	return *o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredSchedulingRule) GetTopologyKeyOk() (*string, bool) {
	if o == nil || o.TopologyKey == nil {
		return nil, false
	}
	return o.TopologyKey, true
}

// HasTopologyKey returns a boolean if a field has been set.
func (o *RequiredSchedulingRule) HasTopologyKey() bool {
	return o != nil && o.TopologyKey != nil
}

// SetTopologyKey gets a reference to the given string and assigns it to the TopologyKey field.
func (o *RequiredSchedulingRule) SetTopologyKey(v string) {
	o.TopologyKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RequiredSchedulingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeSelectorTerms != nil {
		toSerialize["nodeSelectorTerms"] = o.NodeSelectorTerms
	}
	if o.LabelSelector != nil {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if o.TopologyKey != nil {
		toSerialize["topologyKey"] = o.TopologyKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RequiredSchedulingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms,omitempty"`
		LabelSelector     *LabelSelector     `json:"labelSelector,omitempty"`
		TopologyKey       *string            `json:"topologyKey,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeSelectorTerms", "labelSelector", "topologyKey"})
	} else {
		return err
	}

	hasInvalidField := false
	o.NodeSelectorTerms = all.NodeSelectorTerms
	if all.LabelSelector != nil && all.LabelSelector.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LabelSelector = all.LabelSelector
	o.TopologyKey = all.TopologyKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
