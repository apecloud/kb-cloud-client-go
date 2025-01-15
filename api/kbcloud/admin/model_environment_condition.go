// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentCondition EnvironmentCondition contains details for the current condition of this environment.
type EnvironmentCondition struct {
	// Human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status is the status of the condition. Can be True, False, Unknown.
	Status EnvironmentConditionStatus `json:"status"`
	// Type is the type of the condition.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCondition instantiates a new EnvironmentCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCondition(status EnvironmentConditionStatus, typeVar string) *EnvironmentCondition {
	this := EnvironmentCondition{}
	this.Status = status
	this.Type = typeVar
	return &this
}

// NewEnvironmentConditionWithDefaults instantiates a new EnvironmentCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentConditionWithDefaults() *EnvironmentCondition {
	this := EnvironmentCondition{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EnvironmentCondition) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCondition) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EnvironmentCondition) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EnvironmentCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EnvironmentCondition) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCondition) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EnvironmentCondition) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EnvironmentCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value.
func (o *EnvironmentCondition) GetStatus() EnvironmentConditionStatus {
	if o == nil {
		var ret EnvironmentConditionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCondition) GetStatusOk() (*EnvironmentConditionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *EnvironmentCondition) SetStatus(v EnvironmentConditionStatus) {
	o.Status = v
}

// GetType returns the Type field value.
func (o *EnvironmentCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EnvironmentCondition) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string                     `json:"message,omitempty"`
		Reason  *string                     `json:"reason,omitempty"`
		Status  *EnvironmentConditionStatus `json:"status"`
		Type    *string                     `json:"type"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "reason", "status", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = all.Message
	o.Reason = all.Reason
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
