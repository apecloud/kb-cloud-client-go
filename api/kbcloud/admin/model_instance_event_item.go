// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceEventItem Instance event
type InstanceEventItem struct {
	// Cluster name of the event
	ClusterName string `json:"clusterName"`
	// Instance name of the event
	InstanceName string `json:"instanceName"`
	// Name of the event
	Name string `json:"name"`
	// Reason of the event
	Reason string `json:"reason"`
	// Message of the event
	Message string `json:"message"`
	// Type of the event
	Type string `json:"type"`
	// Source of the event
	Source *string `json:"source,omitempty"`
	// Count of the event
	Count int32 `json:"count"`
	// Last timestamp of the event
	EventTime time.Time `json:"eventTime"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceEventItem instantiates a new InstanceEventItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceEventItem(clusterName string, instanceName string, name string, reason string, message string, typeVar string, count int32, eventTime time.Time) *InstanceEventItem {
	this := InstanceEventItem{}
	this.ClusterName = clusterName
	this.InstanceName = instanceName
	this.Name = name
	this.Reason = reason
	this.Message = message
	this.Type = typeVar
	this.Count = count
	this.EventTime = eventTime
	return &this
}

// NewInstanceEventItemWithDefaults instantiates a new InstanceEventItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceEventItemWithDefaults() *InstanceEventItem {
	this := InstanceEventItem{}
	return &this
}

// GetClusterName returns the ClusterName field value.
func (o *InstanceEventItem) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *InstanceEventItem) SetClusterName(v string) {
	o.ClusterName = v
}

// GetInstanceName returns the InstanceName field value.
func (o *InstanceEventItem) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *InstanceEventItem) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetName returns the Name field value.
func (o *InstanceEventItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceEventItem) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value.
func (o *InstanceEventItem) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *InstanceEventItem) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value.
func (o *InstanceEventItem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *InstanceEventItem) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value.
func (o *InstanceEventItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InstanceEventItem) SetType(v string) {
	o.Type = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InstanceEventItem) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InstanceEventItem) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InstanceEventItem) SetSource(v string) {
	o.Source = &v
}

// GetCount returns the Count field value.
func (o *InstanceEventItem) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *InstanceEventItem) SetCount(v int32) {
	o.Count = v
}

// GetEventTime returns the EventTime field value.
func (o *InstanceEventItem) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *InstanceEventItem) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value.
func (o *InstanceEventItem) SetEventTime(v time.Time) {
	o.EventTime = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceEventItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["instanceName"] = o.InstanceName
	toSerialize["name"] = o.Name
	toSerialize["reason"] = o.Reason
	toSerialize["message"] = o.Message
	toSerialize["type"] = o.Type
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	toSerialize["count"] = o.Count
	if o.EventTime.Nanosecond() == 0 {
		toSerialize["eventTime"] = o.EventTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["eventTime"] = o.EventTime.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceEventItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterName  *string    `json:"clusterName"`
		InstanceName *string    `json:"instanceName"`
		Name         *string    `json:"name"`
		Reason       *string    `json:"reason"`
		Message      *string    `json:"message"`
		Type         *string    `json:"type"`
		Source       *string    `json:"source,omitempty"`
		Count        *int32     `json:"count"`
		EventTime    *time.Time `json:"eventTime"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instanceName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.EventTime == nil {
		return fmt.Errorf("required field eventTime missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterName", "instanceName", "name", "reason", "message", "type", "source", "count", "eventTime"})
	} else {
		return err
	}
	o.ClusterName = *all.ClusterName
	o.InstanceName = *all.InstanceName
	o.Name = *all.Name
	o.Reason = *all.Reason
	o.Message = *all.Message
	o.Type = *all.Type
	o.Source = all.Source
	o.Count = *all.Count
	o.EventTime = *all.EventTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
