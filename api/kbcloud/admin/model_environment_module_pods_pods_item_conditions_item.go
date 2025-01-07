// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentModulePodsPodsItemConditionsItem struct {
	// Condition type
	Type *string `json:"type,omitempty"`
	// Condition status
	Status *string `json:"status,omitempty"`
	// Last transition time
	LastTransitionTime *time.Time `json:"last_transition_time,omitempty"`
	// Condition reason
	Reason *string `json:"reason,omitempty"`
	// Condition message
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModulePodsPodsItemConditionsItem instantiates a new EnvironmentModulePodsPodsItemConditionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModulePodsPodsItemConditionsItem() *EnvironmentModulePodsPodsItemConditionsItem {
	this := EnvironmentModulePodsPodsItemConditionsItem{}
	return &this
}

// NewEnvironmentModulePodsPodsItemConditionsItemWithDefaults instantiates a new EnvironmentModulePodsPodsItemConditionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModulePodsPodsItemConditionsItemWithDefaults() *EnvironmentModulePodsPodsItemConditionsItem {
	this := EnvironmentModulePodsPodsItemConditionsItem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentModulePodsPodsItemConditionsItem) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnvironmentModulePodsPodsItemConditionsItem) SetStatus(v string) {
	o.Status = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) HasLastTransitionTime() bool {
	return o != nil && o.LastTransitionTime != nil
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *EnvironmentModulePodsPodsItemConditionsItem) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EnvironmentModulePodsPodsItemConditionsItem) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EnvironmentModulePodsPodsItemConditionsItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EnvironmentModulePodsPodsItemConditionsItem) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModulePodsPodsItemConditionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastTransitionTime != nil {
		if o.LastTransitionTime.Nanosecond() == 0 {
			toSerialize["last_transition_time"] = o.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_transition_time"] = o.LastTransitionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
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
func (o *EnvironmentModulePodsPodsItemConditionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type               *string    `json:"type,omitempty"`
		Status             *string    `json:"status,omitempty"`
		LastTransitionTime *time.Time `json:"last_transition_time,omitempty"`
		Reason             *string    `json:"reason,omitempty"`
		Message            *string    `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "status", "last_transition_time", "reason", "message"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Status = all.Status
	o.LastTransitionTime = all.LastTransitionTime
	o.Reason = all.Reason
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
