// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsPromote OpsPromote is the payload to promote a KubeBlocks cluster
type OpsPromote struct {
	ComponentName string        `json:"componentName"`
	InstanceName  string        `json:"instanceName"`
	Schedule      *TaskSchedule `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsPromote instantiates a new OpsPromote object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsPromote(componentName string, instanceName string) *OpsPromote {
	this := OpsPromote{}
	this.ComponentName = componentName
	this.InstanceName = instanceName
	return &this
}

// NewOpsPromoteWithDefaults instantiates a new OpsPromote object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsPromoteWithDefaults() *OpsPromote {
	this := OpsPromote{}
	return &this
}

// GetComponentName returns the ComponentName field value.
func (o *OpsPromote) GetComponentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *OpsPromote) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value.
func (o *OpsPromote) SetComponentName(v string) {
	o.ComponentName = v
}

// GetInstanceName returns the InstanceName field value.
func (o *OpsPromote) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *OpsPromote) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *OpsPromote) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsPromote) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsPromote) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsPromote) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsPromote) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsPromote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["componentName"] = o.ComponentName
	toSerialize["instanceName"] = o.InstanceName
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsPromote) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentName *string       `json:"componentName"`
		InstanceName  *string       `json:"instanceName"`
		Schedule      *TaskSchedule `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ComponentName == nil {
		return fmt.Errorf("required field componentName missing")
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instanceName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentName", "instanceName", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentName = *all.ComponentName
	o.InstanceName = *all.InstanceName
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
