// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBHealthCheck struct {
	State  DolphinDBHealthState `json:"state"`
	Source string               `json:"source"`
	Reason *string              `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBHealthCheck instantiates a new DolphinDBHealthCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBHealthCheck(state DolphinDBHealthState, source string) *DolphinDBHealthCheck {
	this := DolphinDBHealthCheck{}
	this.State = state
	this.Source = source
	return &this
}

// NewDolphinDBHealthCheckWithDefaults instantiates a new DolphinDBHealthCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBHealthCheckWithDefaults() *DolphinDBHealthCheck {
	this := DolphinDBHealthCheck{}
	return &this
}

// GetState returns the State field value.
func (o *DolphinDBHealthCheck) GetState() DolphinDBHealthState {
	if o == nil {
		var ret DolphinDBHealthState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthCheck) GetStateOk() (*DolphinDBHealthState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DolphinDBHealthCheck) SetState(v DolphinDBHealthState) {
	o.State = v
}

// GetSource returns the Source field value.
func (o *DolphinDBHealthCheck) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthCheck) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DolphinDBHealthCheck) SetSource(v string) {
	o.Source = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DolphinDBHealthCheck) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthCheck) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DolphinDBHealthCheck) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DolphinDBHealthCheck) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State
	toSerialize["source"] = o.Source
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBHealthCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State  *DolphinDBHealthState `json:"state"`
		Source *string               `json:"source"`
		Reason *string               `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"state", "source", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	o.Source = *all.Source
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
