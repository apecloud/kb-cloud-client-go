// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContinueReportRequest struct {
	Message string                           `json:"message"`
	Mode    AiAgentContinueReportRequestMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentContinueReportRequest instantiates a new AiAgentContinueReportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentContinueReportRequest(message string, mode AiAgentContinueReportRequestMode) *AiAgentContinueReportRequest {
	this := AiAgentContinueReportRequest{}
	this.Message = message
	this.Mode = mode
	return &this
}

// NewAiAgentContinueReportRequestWithDefaults instantiates a new AiAgentContinueReportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentContinueReportRequestWithDefaults() *AiAgentContinueReportRequest {
	this := AiAgentContinueReportRequest{}
	return &this
}

// GetMessage returns the Message field value.
func (o *AiAgentContinueReportRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueReportRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AiAgentContinueReportRequest) SetMessage(v string) {
	o.Message = v
}

// GetMode returns the Mode field value.
func (o *AiAgentContinueReportRequest) GetMode() AiAgentContinueReportRequestMode {
	if o == nil {
		var ret AiAgentContinueReportRequestMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueReportRequest) GetModeOk() (*AiAgentContinueReportRequestMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiAgentContinueReportRequest) SetMode(v AiAgentContinueReportRequestMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentContinueReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentContinueReportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string                           `json:"message"`
		Mode    *AiAgentContinueReportRequestMode `json:"mode"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = *all.Message
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
