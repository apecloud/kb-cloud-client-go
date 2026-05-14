// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentSuggestedReply struct {
	Label string `json:"label"`
	// Organization-scoped single-cluster Agent execution scope.
	ScopePatch *AiAgentResourceScope `json:"scopePatch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentSuggestedReply instantiates a new AiAgentSuggestedReply object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentSuggestedReply(label string) *AiAgentSuggestedReply {
	this := AiAgentSuggestedReply{}
	this.Label = label
	return &this
}

// NewAiAgentSuggestedReplyWithDefaults instantiates a new AiAgentSuggestedReply object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentSuggestedReplyWithDefaults() *AiAgentSuggestedReply {
	this := AiAgentSuggestedReply{}
	return &this
}

// GetLabel returns the Label field value.
func (o *AiAgentSuggestedReply) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestedReply) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *AiAgentSuggestedReply) SetLabel(v string) {
	o.Label = v
}

// GetScopePatch returns the ScopePatch field value if set, zero value otherwise.
func (o *AiAgentSuggestedReply) GetScopePatch() AiAgentResourceScope {
	if o == nil || o.ScopePatch == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.ScopePatch
}

// GetScopePatchOk returns a tuple with the ScopePatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestedReply) GetScopePatchOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.ScopePatch == nil {
		return nil, false
	}
	return o.ScopePatch, true
}

// HasScopePatch returns a boolean if a field has been set.
func (o *AiAgentSuggestedReply) HasScopePatch() bool {
	return o != nil && o.ScopePatch != nil
}

// SetScopePatch gets a reference to the given AiAgentResourceScope and assigns it to the ScopePatch field.
func (o *AiAgentSuggestedReply) SetScopePatch(v AiAgentResourceScope) {
	o.ScopePatch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentSuggestedReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	if o.ScopePatch != nil {
		toSerialize["scopePatch"] = o.ScopePatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentSuggestedReply) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label      *string               `json:"label"`
		ScopePatch *AiAgentResourceScope `json:"scopePatch,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"label", "scopePatch"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Label = *all.Label
	if all.ScopePatch != nil && all.ScopePatch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScopePatch = all.ScopePatch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
