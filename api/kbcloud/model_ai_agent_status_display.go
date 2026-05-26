// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AiAgentStatusDisplay Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
type AiAgentStatusDisplay struct {
	// Stable machine-readable status code
	Code string `json:"code"`
	// Default display label
	DisplayLabel string `json:"displayLabel"`
	// Default display description
	DisplayDescription *string `json:"displayDescription,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentStatusDisplay instantiates a new AiAgentStatusDisplay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentStatusDisplay(code string, displayLabel string) *AiAgentStatusDisplay {
	this := AiAgentStatusDisplay{}
	this.Code = code
	this.DisplayLabel = displayLabel
	return &this
}

// NewAiAgentStatusDisplayWithDefaults instantiates a new AiAgentStatusDisplay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentStatusDisplayWithDefaults() *AiAgentStatusDisplay {
	this := AiAgentStatusDisplay{}
	return &this
}

// GetCode returns the Code field value.
func (o *AiAgentStatusDisplay) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AiAgentStatusDisplay) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *AiAgentStatusDisplay) SetCode(v string) {
	o.Code = v
}

// GetDisplayLabel returns the DisplayLabel field value.
func (o *AiAgentStatusDisplay) GetDisplayLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value
// and a boolean to check if the value has been set.
func (o *AiAgentStatusDisplay) GetDisplayLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayLabel, true
}

// SetDisplayLabel sets field value.
func (o *AiAgentStatusDisplay) SetDisplayLabel(v string) {
	o.DisplayLabel = v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *AiAgentStatusDisplay) GetDisplayDescription() string {
	if o == nil || o.DisplayDescription == nil {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStatusDisplay) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || o.DisplayDescription == nil {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *AiAgentStatusDisplay) HasDisplayDescription() bool {
	return o != nil && o.DisplayDescription != nil
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *AiAgentStatusDisplay) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentStatusDisplay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["displayLabel"] = o.DisplayLabel
	if o.DisplayDescription != nil {
		toSerialize["displayDescription"] = o.DisplayDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentStatusDisplay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code               *string `json:"code"`
		DisplayLabel       *string `json:"displayLabel"`
		DisplayDescription *string `json:"displayDescription,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.DisplayLabel == nil {
		return fmt.Errorf("required field displayLabel missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "displayLabel", "displayDescription"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.DisplayLabel = *all.DisplayLabel
	o.DisplayDescription = all.DisplayDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
