// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterTaskDetail taskCondition is the information of the task condition
type ClusterTaskDetail struct {
	// reason of the task condition
	Reason *string `json:"reason,omitempty"`
	// type of the task condition
	Type *string `json:"type,omitempty"`
	// status of the task condition
	Status *string `json:"status,omitempty"`
	// message of the task condition
	Message *string `json:"message,omitempty"`
	// last transition time of the task condition
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterTaskDetail instantiates a new ClusterTaskDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterTaskDetail() *ClusterTaskDetail {
	this := ClusterTaskDetail{}
	return &this
}

// NewClusterTaskDetailWithDefaults instantiates a new ClusterTaskDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterTaskDetailWithDefaults() *ClusterTaskDetail {
	this := ClusterTaskDetail{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ClusterTaskDetail) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskDetail) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ClusterTaskDetail) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ClusterTaskDetail) SetReason(v string) {
	o.Reason = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterTaskDetail) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskDetail) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterTaskDetail) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterTaskDetail) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterTaskDetail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskDetail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterTaskDetail) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterTaskDetail) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterTaskDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterTaskDetail) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterTaskDetail) SetMessage(v string) {
	o.Message = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *ClusterTaskDetail) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskDetail) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *ClusterTaskDetail) HasLastTransitionTime() bool {
	return o != nil && o.LastTransitionTime != nil
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *ClusterTaskDetail) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterTaskDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.LastTransitionTime != nil {
		if o.LastTransitionTime.Nanosecond() == 0 {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterTaskDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason             *string    `json:"reason,omitempty"`
		Type               *string    `json:"type,omitempty"`
		Status             *string    `json:"status,omitempty"`
		Message            *string    `json:"message,omitempty"`
		LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reason", "type", "status", "message", "lastTransitionTime"})
	} else {
		return err
	}
	o.Reason = all.Reason
	o.Type = all.Type
	o.Status = all.Status
	o.Message = all.Message
	o.LastTransitionTime = all.LastTransitionTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
