// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMissingRequirement struct {
	RequirementId string               `json:"requirementId"`
	MissingReason AiAgentMissingReason `json:"missingReason"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	Display *AiAgentStatusDisplay `json:"display,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMissingRequirement instantiates a new AiAgentMissingRequirement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMissingRequirement(requirementId string, missingReason AiAgentMissingReason) *AiAgentMissingRequirement {
	this := AiAgentMissingRequirement{}
	this.RequirementId = requirementId
	this.MissingReason = missingReason
	return &this
}

// NewAiAgentMissingRequirementWithDefaults instantiates a new AiAgentMissingRequirement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMissingRequirementWithDefaults() *AiAgentMissingRequirement {
	this := AiAgentMissingRequirement{}
	return &this
}

// GetRequirementId returns the RequirementId field value.
func (o *AiAgentMissingRequirement) GetRequirementId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequirementId
}

// GetRequirementIdOk returns a tuple with the RequirementId field value
// and a boolean to check if the value has been set.
func (o *AiAgentMissingRequirement) GetRequirementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementId, true
}

// SetRequirementId sets field value.
func (o *AiAgentMissingRequirement) SetRequirementId(v string) {
	o.RequirementId = v
}

// GetMissingReason returns the MissingReason field value.
func (o *AiAgentMissingRequirement) GetMissingReason() AiAgentMissingReason {
	if o == nil {
		var ret AiAgentMissingReason
		return ret
	}
	return o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value
// and a boolean to check if the value has been set.
func (o *AiAgentMissingRequirement) GetMissingReasonOk() (*AiAgentMissingReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MissingReason, true
}

// SetMissingReason sets field value.
func (o *AiAgentMissingRequirement) SetMissingReason(v AiAgentMissingReason) {
	o.MissingReason = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *AiAgentMissingRequirement) GetDisplay() AiAgentStatusDisplay {
	if o == nil || o.Display == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMissingRequirement) GetDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *AiAgentMissingRequirement) HasDisplay() bool {
	return o != nil && o.Display != nil
}

// SetDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the Display field.
func (o *AiAgentMissingRequirement) SetDisplay(v AiAgentStatusDisplay) {
	o.Display = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMissingRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["requirementId"] = o.RequirementId
	toSerialize["missingReason"] = o.MissingReason
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMissingRequirement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequirementId *string               `json:"requirementId"`
		MissingReason *AiAgentMissingReason `json:"missingReason"`
		Display       *AiAgentStatusDisplay `json:"display,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RequirementId == nil {
		return fmt.Errorf("required field requirementId missing")
	}
	if all.MissingReason == nil {
		return fmt.Errorf("required field missingReason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"requirementId", "missingReason", "display"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RequirementId = *all.RequirementId
	if !all.MissingReason.IsValid() {
		hasInvalidField = true
	} else {
		o.MissingReason = *all.MissingReason
	}
	if all.Display != nil && all.Display.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Display = all.Display

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
