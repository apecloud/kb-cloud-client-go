// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type DmsTaskList struct {
	Tasks []DmsTaskInfo `json:"tasks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewDmsTaskList instantiates a new DmsTaskList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTaskList() *DmsTaskList {
	this := DmsTaskList{}
	return &this
}

// NewDmsTaskListWithDefaults instantiates a new DmsTaskList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTaskListWithDefaults() *DmsTaskList {
	this := DmsTaskList{}
	return &this
}
// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *DmsTaskList) GetTasks() []DmsTaskInfo {
	if o == nil || o.Tasks == nil {
		var ret []DmsTaskInfo
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTaskList) GetTasksOk() (*[]DmsTaskInfo, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *DmsTaskList) HasTasks() bool {
	return o != nil && o.Tasks != nil
}

// SetTasks gets a reference to the given []DmsTaskInfo and assigns it to the Tasks field.
func (o *DmsTaskList) SetTasks(v []DmsTaskInfo) {
	o.Tasks = v
}



// MarshalJSON serializes the struct using spec logic.
func (o DmsTaskList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTaskList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tasks []DmsTaskInfo `json:"tasks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "tasks",  })
	} else {
		return err
	}
	o.Tasks = all.Tasks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
