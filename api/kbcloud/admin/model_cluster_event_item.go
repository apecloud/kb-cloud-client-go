// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterEventItem Kubernetes event for a cluster, component or InstanceSet
type ClusterEventItem struct {
	ClusterName string `json:"clusterName"`
	// Kubernetes kind of the involved object
	ObjectKind    string  `json:"objectKind"`
	ObjectName    string  `json:"objectName"`
	ObjectUid     string  `json:"objectUID"`
	ComponentName *string `json:"componentName,omitempty"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Reason        string  `json:"reason"`
	Message       string  `json:"message"`
	Source        string  `json:"source"`
	Count         int32   `json:"count"`
	// Last observation time, falling back to the event creation time
	EventTime time.Time `json:"eventTime"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterEventItem instantiates a new ClusterEventItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterEventItem(clusterName string, objectKind string, objectName string, objectUid string, name string, typeVar string, reason string, message string, source string, count int32, eventTime time.Time) *ClusterEventItem {
	this := ClusterEventItem{}
	this.ClusterName = clusterName
	this.ObjectKind = objectKind
	this.ObjectName = objectName
	this.ObjectUid = objectUid
	this.Name = name
	this.Type = typeVar
	this.Reason = reason
	this.Message = message
	this.Source = source
	this.Count = count
	this.EventTime = eventTime
	return &this
}

// NewClusterEventItemWithDefaults instantiates a new ClusterEventItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterEventItemWithDefaults() *ClusterEventItem {
	this := ClusterEventItem{}
	return &this
}

// GetClusterName returns the ClusterName field value.
func (o *ClusterEventItem) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *ClusterEventItem) SetClusterName(v string) {
	o.ClusterName = v
}

// GetObjectKind returns the ObjectKind field value.
func (o *ClusterEventItem) GetObjectKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectKind
}

// GetObjectKindOk returns a tuple with the ObjectKind field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetObjectKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectKind, true
}

// SetObjectKind sets field value.
func (o *ClusterEventItem) SetObjectKind(v string) {
	o.ObjectKind = v
}

// GetObjectName returns the ObjectName field value.
func (o *ClusterEventItem) GetObjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectName, true
}

// SetObjectName sets field value.
func (o *ClusterEventItem) SetObjectName(v string) {
	o.ObjectName = v
}

// GetObjectUid returns the ObjectUid field value.
func (o *ClusterEventItem) GetObjectUid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectUid
}

// GetObjectUidOk returns a tuple with the ObjectUid field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetObjectUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectUid, true
}

// SetObjectUid sets field value.
func (o *ClusterEventItem) SetObjectUid(v string) {
	o.ObjectUid = v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ClusterEventItem) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ClusterEventItem) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ClusterEventItem) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetName returns the Name field value.
func (o *ClusterEventItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClusterEventItem) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ClusterEventItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ClusterEventItem) SetType(v string) {
	o.Type = v
}

// GetReason returns the Reason field value.
func (o *ClusterEventItem) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ClusterEventItem) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value.
func (o *ClusterEventItem) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *ClusterEventItem) SetMessage(v string) {
	o.Message = v
}

// GetSource returns the Source field value.
func (o *ClusterEventItem) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ClusterEventItem) SetSource(v string) {
	o.Source = v
}

// GetCount returns the Count field value.
func (o *ClusterEventItem) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *ClusterEventItem) SetCount(v int32) {
	o.Count = v
}

// GetEventTime returns the EventTime field value.
func (o *ClusterEventItem) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *ClusterEventItem) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value.
func (o *ClusterEventItem) SetEventTime(v time.Time) {
	o.EventTime = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterEventItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["objectKind"] = o.ObjectKind
	toSerialize["objectName"] = o.ObjectName
	toSerialize["objectUID"] = o.ObjectUid
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["reason"] = o.Reason
	toSerialize["message"] = o.Message
	toSerialize["source"] = o.Source
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
func (o *ClusterEventItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterName   *string    `json:"clusterName"`
		ObjectKind    *string    `json:"objectKind"`
		ObjectName    *string    `json:"objectName"`
		ObjectUid     *string    `json:"objectUID"`
		ComponentName *string    `json:"componentName,omitempty"`
		Name          *string    `json:"name"`
		Type          *string    `json:"type"`
		Reason        *string    `json:"reason"`
		Message       *string    `json:"message"`
		Source        *string    `json:"source"`
		Count         *int32     `json:"count"`
		EventTime     *time.Time `json:"eventTime"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.ObjectKind == nil {
		return fmt.Errorf("required field objectKind missing")
	}
	if all.ObjectName == nil {
		return fmt.Errorf("required field objectName missing")
	}
	if all.ObjectUid == nil {
		return fmt.Errorf("required field objectUID missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.EventTime == nil {
		return fmt.Errorf("required field eventTime missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterName", "objectKind", "objectName", "objectUID", "componentName", "name", "type", "reason", "message", "source", "count", "eventTime"})
	} else {
		return err
	}
	o.ClusterName = *all.ClusterName
	o.ObjectKind = *all.ObjectKind
	o.ObjectName = *all.ObjectName
	o.ObjectUid = *all.ObjectUid
	o.ComponentName = all.ComponentName
	o.Name = *all.Name
	o.Type = *all.Type
	o.Reason = *all.Reason
	o.Message = *all.Message
	o.Source = *all.Source
	o.Count = *all.Count
	o.EventTime = *all.EventTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
