// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ViewInvolvedObjectsSelector A label selector is a label query over a set of resources. 
type ViewInvolvedObjectsSelector struct {
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []ViewInvolvedObjectsSelectorMatchExpressionsItem `json:"matchExpressions,omitempty"`
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewViewInvolvedObjectsSelector instantiates a new ViewInvolvedObjectsSelector object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewViewInvolvedObjectsSelector() *ViewInvolvedObjectsSelector {
	this := ViewInvolvedObjectsSelector{}
	return &this
}

// NewViewInvolvedObjectsSelectorWithDefaults instantiates a new ViewInvolvedObjectsSelector object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewViewInvolvedObjectsSelectorWithDefaults() *ViewInvolvedObjectsSelector {
	this := ViewInvolvedObjectsSelector{}
	return &this
}
// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *ViewInvolvedObjectsSelector) GetMatchExpressions() []ViewInvolvedObjectsSelectorMatchExpressionsItem {
	if o == nil || o.MatchExpressions == nil {
		var ret []ViewInvolvedObjectsSelectorMatchExpressionsItem
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewInvolvedObjectsSelector) GetMatchExpressionsOk() (*[]ViewInvolvedObjectsSelectorMatchExpressionsItem, bool) {
	if o == nil || o.MatchExpressions == nil {
		return nil, false
	}
	return &o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *ViewInvolvedObjectsSelector) HasMatchExpressions() bool {
	return o != nil && o.MatchExpressions != nil
}

// SetMatchExpressions gets a reference to the given []ViewInvolvedObjectsSelectorMatchExpressionsItem and assigns it to the MatchExpressions field.
func (o *ViewInvolvedObjectsSelector) SetMatchExpressions(v []ViewInvolvedObjectsSelectorMatchExpressionsItem) {
	o.MatchExpressions = v
}


// GetMatchLabels returns the MatchLabels field value if set, zero value otherwise.
func (o *ViewInvolvedObjectsSelector) GetMatchLabels() map[string]string {
	if o == nil || o.MatchLabels == nil {
		var ret map[string]string
		return ret
	}
	return o.MatchLabels
}

// GetMatchLabelsOk returns a tuple with the MatchLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewInvolvedObjectsSelector) GetMatchLabelsOk() (*map[string]string, bool) {
	if o == nil || o.MatchLabels == nil {
		return nil, false
	}
	return &o.MatchLabels, true
}

// HasMatchLabels returns a boolean if a field has been set.
func (o *ViewInvolvedObjectsSelector) HasMatchLabels() bool {
	return o != nil && o.MatchLabels != nil
}

// SetMatchLabels gets a reference to the given map[string]string and assigns it to the MatchLabels field.
func (o *ViewInvolvedObjectsSelector) SetMatchLabels(v map[string]string) {
	o.MatchLabels = v
}



// MarshalJSON serializes the struct using spec logic.
func (o ViewInvolvedObjectsSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MatchExpressions != nil {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if o.MatchLabels != nil {
		toSerialize["matchLabels"] = o.MatchLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ViewInvolvedObjectsSelector) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchExpressions []ViewInvolvedObjectsSelectorMatchExpressionsItem `json:"matchExpressions,omitempty"`
		MatchLabels map[string]string `json:"matchLabels,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "matchExpressions", "matchLabels",  })
	} else {
		return err
	}
	o.MatchExpressions = all.MatchExpressions
	o.MatchLabels = all.MatchLabels

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
