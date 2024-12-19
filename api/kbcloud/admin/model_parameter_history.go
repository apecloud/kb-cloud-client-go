// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"
)

// ParameterHistory The history of a parameter
type ParameterHistory struct {
	// The name of the parameter
	ParameterName string `json:"parameterName"`
	// The old value of the parameter
	OldValue string `json:"oldValue"`
	// The new value of the parameter
	NewValue string `json:"newValue"`
	// The date and time the parameter was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterHistory instantiates a new ParameterHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterHistory(parameterName string, oldValue string, newValue string, updatedAt time.Time) *ParameterHistory {
	this := ParameterHistory{}
	this.ParameterName = parameterName
	this.OldValue = oldValue
	this.NewValue = newValue
	this.UpdatedAt = updatedAt
	return &this
}

// NewParameterHistoryWithDefaults instantiates a new ParameterHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterHistoryWithDefaults() *ParameterHistory {
	this := ParameterHistory{}
	return &this
}

// GetParameterName returns the ParameterName field value.
func (o *ParameterHistory) GetParameterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetParameterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterName, true
}

// SetParameterName sets field value.
func (o *ParameterHistory) SetParameterName(v string) {
	o.ParameterName = v
}

// GetOldValue returns the OldValue field value.
func (o *ParameterHistory) GetOldValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value.
func (o *ParameterHistory) SetOldValue(v string) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value.
func (o *ParameterHistory) GetNewValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value.
func (o *ParameterHistory) SetNewValue(v string) {
	o.NewValue = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ParameterHistory) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ParameterHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ParameterHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["parameterName"] = o.ParameterName
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParameterName *string    `json:"parameterName"`
		OldValue      *string    `json:"oldValue"`
		NewValue      *string    `json:"newValue"`
		UpdatedAt     *time.Time `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ParameterName == nil {
		return fmt.Errorf("required field parameterName missing")
	}
	if all.OldValue == nil {
		return fmt.Errorf("required field oldValue missing")
	}
	if all.NewValue == nil {
		return fmt.Errorf("required field newValue missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"parameterName", "oldValue", "newValue", "updatedAt"})
	} else {
		return err
	}
	o.ParameterName = *all.ParameterName
	o.OldValue = *all.OldValue
	o.NewValue = *all.NewValue
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
