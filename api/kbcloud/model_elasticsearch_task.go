// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchTask struct {
	TaskId            string  `json:"taskId"`
	NodeId            string  `json:"nodeId"`
	NodeName          string  `json:"nodeName"`
	Action            string  `json:"action"`
	Description       *string `json:"description,omitempty"`
	Type              *string `json:"type,omitempty"`
	StartTime         *string `json:"startTime,omitempty"`
	RunningTimeMillis int64   `json:"runningTimeMillis"`
	Cancellable       bool    `json:"cancellable"`
	Cancelled         *bool   `json:"cancelled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchTask instantiates a new ElasticsearchTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchTask(taskId string, nodeId string, nodeName string, action string, runningTimeMillis int64, cancellable bool) *ElasticsearchTask {
	this := ElasticsearchTask{}
	this.TaskId = taskId
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.Action = action
	this.RunningTimeMillis = runningTimeMillis
	this.Cancellable = cancellable
	return &this
}

// NewElasticsearchTaskWithDefaults instantiates a new ElasticsearchTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchTaskWithDefaults() *ElasticsearchTask {
	this := ElasticsearchTask{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *ElasticsearchTask) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *ElasticsearchTask) SetTaskId(v string) {
	o.TaskId = v
}

// GetNodeId returns the NodeId field value.
func (o *ElasticsearchTask) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value.
func (o *ElasticsearchTask) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value.
func (o *ElasticsearchTask) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value.
func (o *ElasticsearchTask) SetNodeName(v string) {
	o.NodeName = v
}

// GetAction returns the Action field value.
func (o *ElasticsearchTask) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ElasticsearchTask) SetAction(v string) {
	o.Action = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ElasticsearchTask) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ElasticsearchTask) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ElasticsearchTask) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ElasticsearchTask) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ElasticsearchTask) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ElasticsearchTask) SetType(v string) {
	o.Type = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ElasticsearchTask) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ElasticsearchTask) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ElasticsearchTask) SetStartTime(v string) {
	o.StartTime = &v
}

// GetRunningTimeMillis returns the RunningTimeMillis field value.
func (o *ElasticsearchTask) GetRunningTimeMillis() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RunningTimeMillis
}

// GetRunningTimeMillisOk returns a tuple with the RunningTimeMillis field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetRunningTimeMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunningTimeMillis, true
}

// SetRunningTimeMillis sets field value.
func (o *ElasticsearchTask) SetRunningTimeMillis(v int64) {
	o.RunningTimeMillis = v
}

// GetCancellable returns the Cancellable field value.
func (o *ElasticsearchTask) GetCancellable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Cancellable
}

// GetCancellableOk returns a tuple with the Cancellable field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetCancellableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cancellable, true
}

// SetCancellable sets field value.
func (o *ElasticsearchTask) SetCancellable(v bool) {
	o.Cancellable = v
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *ElasticsearchTask) GetCancelled() bool {
	if o == nil || o.Cancelled == nil {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchTask) GetCancelledOk() (*bool, bool) {
	if o == nil || o.Cancelled == nil {
		return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *ElasticsearchTask) HasCancelled() bool {
	return o != nil && o.Cancelled != nil
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *ElasticsearchTask) SetCancelled(v bool) {
	o.Cancelled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["nodeId"] = o.NodeId
	toSerialize["nodeName"] = o.NodeName
	toSerialize["action"] = o.Action
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	toSerialize["runningTimeMillis"] = o.RunningTimeMillis
	toSerialize["cancellable"] = o.Cancellable
	if o.Cancelled != nil {
		toSerialize["cancelled"] = o.Cancelled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId            *string `json:"taskId"`
		NodeId            *string `json:"nodeId"`
		NodeName          *string `json:"nodeName"`
		Action            *string `json:"action"`
		Description       *string `json:"description,omitempty"`
		Type              *string `json:"type,omitempty"`
		StartTime         *string `json:"startTime,omitempty"`
		RunningTimeMillis *int64  `json:"runningTimeMillis"`
		Cancellable       *bool   `json:"cancellable"`
		Cancelled         *bool   `json:"cancelled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.NodeId == nil {
		return fmt.Errorf("required field nodeId missing")
	}
	if all.NodeName == nil {
		return fmt.Errorf("required field nodeName missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.RunningTimeMillis == nil {
		return fmt.Errorf("required field runningTimeMillis missing")
	}
	if all.Cancellable == nil {
		return fmt.Errorf("required field cancellable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "nodeId", "nodeName", "action", "description", "type", "startTime", "runningTimeMillis", "cancellable", "cancelled"})
	} else {
		return err
	}
	o.TaskId = *all.TaskId
	o.NodeId = *all.NodeId
	o.NodeName = *all.NodeName
	o.Action = *all.Action
	o.Description = all.Description
	o.Type = all.Type
	o.StartTime = all.StartTime
	o.RunningTimeMillis = *all.RunningTimeMillis
	o.Cancellable = *all.Cancellable
	o.Cancelled = all.Cancelled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
