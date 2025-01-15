// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOptionHistory struct {
	ModifierId    string       `json:"modifierId"`
	ModifierEmail string       `json:"modifierEmail"`
	Option        EngineOption `json:"option"`
	CreatedAt     time.Time    `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionHistory instantiates a new EngineOptionHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionHistory(modifierId string, modifierEmail string, option EngineOption, createdAt time.Time) *EngineOptionHistory {
	this := EngineOptionHistory{}
	this.ModifierId = modifierId
	this.ModifierEmail = modifierEmail
	this.Option = option
	this.CreatedAt = createdAt
	return &this
}

// NewEngineOptionHistoryWithDefaults instantiates a new EngineOptionHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionHistoryWithDefaults() *EngineOptionHistory {
	this := EngineOptionHistory{}
	return &this
}

// GetModifierId returns the ModifierId field value.
func (o *EngineOptionHistory) GetModifierId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifierId
}

// GetModifierIdOk returns a tuple with the ModifierId field value
// and a boolean to check if the value has been set.
func (o *EngineOptionHistory) GetModifierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifierId, true
}

// SetModifierId sets field value.
func (o *EngineOptionHistory) SetModifierId(v string) {
	o.ModifierId = v
}

// GetModifierEmail returns the ModifierEmail field value.
func (o *EngineOptionHistory) GetModifierEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifierEmail
}

// GetModifierEmailOk returns a tuple with the ModifierEmail field value
// and a boolean to check if the value has been set.
func (o *EngineOptionHistory) GetModifierEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifierEmail, true
}

// SetModifierEmail sets field value.
func (o *EngineOptionHistory) SetModifierEmail(v string) {
	o.ModifierEmail = v
}

// GetOption returns the Option field value.
func (o *EngineOptionHistory) GetOption() EngineOption {
	if o == nil {
		var ret EngineOption
		return ret
	}
	return o.Option
}

// GetOptionOk returns a tuple with the Option field value
// and a boolean to check if the value has been set.
func (o *EngineOptionHistory) GetOptionOk() (*EngineOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Option, true
}

// SetOption sets field value.
func (o *EngineOptionHistory) SetOption(v EngineOption) {
	o.Option = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *EngineOptionHistory) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EngineOptionHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *EngineOptionHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["modifierId"] = o.ModifierId
	toSerialize["modifierEmail"] = o.ModifierEmail
	toSerialize["option"] = o.Option
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModifierId    *string       `json:"modifierId"`
		ModifierEmail *string       `json:"modifierEmail"`
		Option        *EngineOption `json:"option"`
		CreatedAt     *time.Time    `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ModifierId == nil {
		return fmt.Errorf("required field modifierId missing")
	}
	if all.ModifierEmail == nil {
		return fmt.Errorf("required field modifierEmail missing")
	}
	if all.Option == nil {
		return fmt.Errorf("required field option missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"modifierId", "modifierEmail", "option", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ModifierId = *all.ModifierId
	o.ModifierEmail = *all.ModifierEmail
	if all.Option.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Option = *all.Option
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
