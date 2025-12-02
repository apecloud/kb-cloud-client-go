// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DisasterRecoveryPromote the Promote object for disasterRecovery instance
type DisasterRecoveryPromote struct {
	// the mode of cluster after promotion. default is the primary cluster's mode
	Mode common.NullableString `json:"mode,omitempty"`
	// the reason for promoting
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryPromote instantiates a new DisasterRecoveryPromote object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryPromote() *DisasterRecoveryPromote {
	this := DisasterRecoveryPromote{}
	return &this
}

// NewDisasterRecoveryPromoteWithDefaults instantiates a new DisasterRecoveryPromote object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryPromoteWithDefaults() *DisasterRecoveryPromote {
	this := DisasterRecoveryPromote{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryPromote) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryPromote) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *DisasterRecoveryPromote) HasMode() bool {
	return o != nil && o.Mode.IsSet()
}

// SetMode gets a reference to the given common.NullableString and assigns it to the Mode field.
func (o *DisasterRecoveryPromote) SetMode(v string) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil.
func (o *DisasterRecoveryPromote) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil.
func (o *DisasterRecoveryPromote) UnsetMode() {
	o.Mode.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DisasterRecoveryPromote) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryPromote) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DisasterRecoveryPromote) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DisasterRecoveryPromote) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryPromote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryPromote) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode   common.NullableString `json:"mode,omitempty"`
		Reason *string               `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "reason"})
	} else {
		return err
	}
	o.Mode = all.Mode
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
