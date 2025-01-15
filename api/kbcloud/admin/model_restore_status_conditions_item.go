// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RestoreStatusConditionsItem struct {
	Message *string `json:"message,omitempty"`
	Reason  *string `json:"reason,omitempty"`
	// conditionType
	Type               *string    `json:"type,omitempty"`
	ObservedGeneration *int64     `json:"observedGeneration,omitempty"`
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Status             *string    `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreStatusConditionsItem instantiates a new RestoreStatusConditionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreStatusConditionsItem() *RestoreStatusConditionsItem {
	this := RestoreStatusConditionsItem{}
	return &this
}

// NewRestoreStatusConditionsItemWithDefaults instantiates a new RestoreStatusConditionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreStatusConditionsItemWithDefaults() *RestoreStatusConditionsItem {
	this := RestoreStatusConditionsItem{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RestoreStatusConditionsItem) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RestoreStatusConditionsItem) SetReason(v string) {
	o.Reason = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RestoreStatusConditionsItem) SetType(v string) {
	o.Type = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasObservedGeneration() bool {
	return o != nil && o.ObservedGeneration != nil
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *RestoreStatusConditionsItem) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasLastTransitionTime() bool {
	return o != nil && o.LastTransitionTime != nil
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *RestoreStatusConditionsItem) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestoreStatusConditionsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusConditionsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestoreStatusConditionsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RestoreStatusConditionsItem) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreStatusConditionsItem) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ObservedGeneration != nil {
		toSerialize["observedGeneration"] = o.ObservedGeneration
	}
	if o.LastTransitionTime != nil {
		if o.LastTransitionTime.Nanosecond() == 0 {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreStatusConditionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message            *string    `json:"message,omitempty"`
		Reason             *string    `json:"reason,omitempty"`
		Type               *string    `json:"type,omitempty"`
		ObservedGeneration *int64     `json:"observedGeneration,omitempty"`
		LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
		Status             *string    `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "reason", "type", "observedGeneration", "lastTransitionTime", "status"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Reason = all.Reason
	o.Type = all.Type
	o.ObservedGeneration = all.ObservedGeneration
	o.LastTransitionTime = all.LastTransitionTime
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
