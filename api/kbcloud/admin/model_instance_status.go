// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceStatus Status for instance

type InstanceStatus struct {
	// The current phase of the cluster
	Phase string `json:"phase"`
	// A brief reason for the cluster's current phase
	Reason *string `json:"reason,omitempty"`
	// A human-readable message indicating details about the cluster's current phase
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceStatus instantiates a new InstanceStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceStatus(phase string) *InstanceStatus {
	this := InstanceStatus{}
	this.Phase = phase
	return &this
}

// NewInstanceStatusWithDefaults instantiates a new InstanceStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceStatusWithDefaults() *InstanceStatus {
	this := InstanceStatus{}
	return &this
}

// GetPhase returns the Phase field value.
func (o *InstanceStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *InstanceStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value.
func (o *InstanceStatus) SetPhase(v string) {
	o.Phase = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InstanceStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InstanceStatus) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InstanceStatus) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InstanceStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InstanceStatus) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InstanceStatus) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["phase"] = o.Phase
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Phase   *string `json:"phase"`
		Reason  *string `json:"reason,omitempty"`
		Message *string `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Phase == nil {
		return fmt.Errorf("required field phase missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"phase", "reason", "message"})
	} else {
		return err
	}
	o.Phase = *all.Phase
	o.Reason = all.Reason
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
