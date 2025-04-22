// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// LabelSelector Represents a label selector for matching pods based on their labels.
type LabelSelector struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLabelSelector instantiates a new LabelSelector object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLabelSelector() *LabelSelector {
	this := LabelSelector{}
	return &this
}

// NewLabelSelectorWithDefaults instantiates a new LabelSelector object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLabelSelectorWithDefaults() *LabelSelector {
	this := LabelSelector{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *LabelSelector) GetMatchExpressions() []NodeSelectorRequirement {
	if o == nil || o.MatchExpressions == nil {
		var ret []NodeSelectorRequirement
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelSelector) GetMatchExpressionsOk() (*[]NodeSelectorRequirement, bool) {
	if o == nil || o.MatchExpressions == nil {
		return nil, false
	}
	return &o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *LabelSelector) HasMatchExpressions() bool {
	return o != nil && o.MatchExpressions != nil
}

// SetMatchExpressions gets a reference to the given []NodeSelectorRequirement and assigns it to the MatchExpressions field.
func (o *LabelSelector) SetMatchExpressions(v []NodeSelectorRequirement) {
	o.MatchExpressions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LabelSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MatchExpressions != nil {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LabelSelector) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"matchExpressions"})
	} else {
		return err
	}
	o.MatchExpressions = all.MatchExpressions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
