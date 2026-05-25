// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolConfirmationDecisionResponse struct {
	ConfirmationId string                        `json:"confirmationId"`
	Status         AiAgentToolConfirmationStatus `json:"status"`
	TurnId         string                        `json:"turnId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentToolConfirmationDecisionResponse instantiates a new AiAgentToolConfirmationDecisionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentToolConfirmationDecisionResponse(confirmationId string, status AiAgentToolConfirmationStatus, turnId string) *AiAgentToolConfirmationDecisionResponse {
	this := AiAgentToolConfirmationDecisionResponse{}
	this.ConfirmationId = confirmationId
	this.Status = status
	this.TurnId = turnId
	return &this
}

// NewAiAgentToolConfirmationDecisionResponseWithDefaults instantiates a new AiAgentToolConfirmationDecisionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentToolConfirmationDecisionResponseWithDefaults() *AiAgentToolConfirmationDecisionResponse {
	this := AiAgentToolConfirmationDecisionResponse{}
	return &this
}

// GetConfirmationId returns the ConfirmationId field value.
func (o *AiAgentToolConfirmationDecisionResponse) GetConfirmationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfirmationId
}

// GetConfirmationIdOk returns a tuple with the ConfirmationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolConfirmationDecisionResponse) GetConfirmationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationId, true
}

// SetConfirmationId sets field value.
func (o *AiAgentToolConfirmationDecisionResponse) SetConfirmationId(v string) {
	o.ConfirmationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentToolConfirmationDecisionResponse) GetStatus() AiAgentToolConfirmationStatus {
	if o == nil {
		var ret AiAgentToolConfirmationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolConfirmationDecisionResponse) GetStatusOk() (*AiAgentToolConfirmationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentToolConfirmationDecisionResponse) SetStatus(v AiAgentToolConfirmationStatus) {
	o.Status = v
}

// GetTurnId returns the TurnId field value.
func (o *AiAgentToolConfirmationDecisionResponse) GetTurnId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TurnId
}

// GetTurnIdOk returns a tuple with the TurnId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolConfirmationDecisionResponse) GetTurnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TurnId, true
}

// SetTurnId sets field value.
func (o *AiAgentToolConfirmationDecisionResponse) SetTurnId(v string) {
	o.TurnId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentToolConfirmationDecisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["confirmationId"] = o.ConfirmationId
	toSerialize["status"] = o.Status
	toSerialize["turnId"] = o.TurnId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentToolConfirmationDecisionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfirmationId *string                        `json:"confirmationId"`
		Status         *AiAgentToolConfirmationStatus `json:"status"`
		TurnId         *string                        `json:"turnId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfirmationId == nil {
		return fmt.Errorf("required field confirmationId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TurnId == nil {
		return fmt.Errorf("required field turnId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"confirmationId", "status", "turnId"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfirmationId = *all.ConfirmationId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TurnId = *all.TurnId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
