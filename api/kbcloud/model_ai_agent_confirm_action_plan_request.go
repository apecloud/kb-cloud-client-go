// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentConfirmActionPlanRequest struct {
	ConfirmationText *string `json:"confirmationText,omitempty"`
	ClientNonce      *string `json:"clientNonce,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentConfirmActionPlanRequest instantiates a new AiAgentConfirmActionPlanRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentConfirmActionPlanRequest() *AiAgentConfirmActionPlanRequest {
	this := AiAgentConfirmActionPlanRequest{}
	return &this
}

// NewAiAgentConfirmActionPlanRequestWithDefaults instantiates a new AiAgentConfirmActionPlanRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentConfirmActionPlanRequestWithDefaults() *AiAgentConfirmActionPlanRequest {
	this := AiAgentConfirmActionPlanRequest{}
	return &this
}

// GetConfirmationText returns the ConfirmationText field value if set, zero value otherwise.
func (o *AiAgentConfirmActionPlanRequest) GetConfirmationText() string {
	if o == nil || o.ConfirmationText == nil {
		var ret string
		return ret
	}
	return *o.ConfirmationText
}

// GetConfirmationTextOk returns a tuple with the ConfirmationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanRequest) GetConfirmationTextOk() (*string, bool) {
	if o == nil || o.ConfirmationText == nil {
		return nil, false
	}
	return o.ConfirmationText, true
}

// HasConfirmationText returns a boolean if a field has been set.
func (o *AiAgentConfirmActionPlanRequest) HasConfirmationText() bool {
	return o != nil && o.ConfirmationText != nil
}

// SetConfirmationText gets a reference to the given string and assigns it to the ConfirmationText field.
func (o *AiAgentConfirmActionPlanRequest) SetConfirmationText(v string) {
	o.ConfirmationText = &v
}

// GetClientNonce returns the ClientNonce field value if set, zero value otherwise.
func (o *AiAgentConfirmActionPlanRequest) GetClientNonce() string {
	if o == nil || o.ClientNonce == nil {
		var ret string
		return ret
	}
	return *o.ClientNonce
}

// GetClientNonceOk returns a tuple with the ClientNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanRequest) GetClientNonceOk() (*string, bool) {
	if o == nil || o.ClientNonce == nil {
		return nil, false
	}
	return o.ClientNonce, true
}

// HasClientNonce returns a boolean if a field has been set.
func (o *AiAgentConfirmActionPlanRequest) HasClientNonce() bool {
	return o != nil && o.ClientNonce != nil
}

// SetClientNonce gets a reference to the given string and assigns it to the ClientNonce field.
func (o *AiAgentConfirmActionPlanRequest) SetClientNonce(v string) {
	o.ClientNonce = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentConfirmActionPlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ConfirmationText != nil {
		toSerialize["confirmationText"] = o.ConfirmationText
	}
	if o.ClientNonce != nil {
		toSerialize["clientNonce"] = o.ClientNonce
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentConfirmActionPlanRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfirmationText *string `json:"confirmationText,omitempty"`
		ClientNonce      *string `json:"clientNonce,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"confirmationText", "clientNonce"})
	} else {
		return err
	}
	o.ConfirmationText = all.ConfirmationText
	o.ClientNonce = all.ClientNonce

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
