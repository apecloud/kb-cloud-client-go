// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// OpsStop OpsStop is the payload to stop a KubeBlocks cluster
type OpsStop struct {
	// List of component names to stop
	ComponentNames common.NullableList[string] `json:"componentNames,omitempty"`
	Schedule       *TaskSchedule               `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsStop instantiates a new OpsStop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsStop() *OpsStop {
	this := OpsStop{}
	return &this
}

// NewOpsStopWithDefaults instantiates a new OpsStop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsStopWithDefaults() *OpsStop {
	this := OpsStop{}
	return &this
}

// GetComponentNames returns the ComponentNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsStop) GetComponentNames() []string {
	if o == nil || o.ComponentNames.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ComponentNames.Get()
}

// GetComponentNamesOk returns a tuple with the ComponentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsStop) GetComponentNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentNames.Get(), o.ComponentNames.IsSet()
}

// HasComponentNames returns a boolean if a field has been set.
func (o *OpsStop) HasComponentNames() bool {
	return o != nil && o.ComponentNames.IsSet()
}

// SetComponentNames gets a reference to the given common.NullableList[string] and assigns it to the ComponentNames field.
func (o *OpsStop) SetComponentNames(v []string) {
	o.ComponentNames.Set(&v)
}

// SetComponentNamesNil sets the value for ComponentNames to be an explicit nil.
func (o *OpsStop) SetComponentNamesNil() {
	o.ComponentNames.Set(nil)
}

// UnsetComponentNames ensures that no value is present for ComponentNames, not even an explicit nil.
func (o *OpsStop) UnsetComponentNames() {
	o.ComponentNames.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsStop) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsStop) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsStop) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsStop) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ComponentNames.IsSet() {
		toSerialize["componentNames"] = o.ComponentNames.Get()
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsStop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentNames common.NullableList[string] `json:"componentNames,omitempty"`
		Schedule       *TaskSchedule               `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentNames", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentNames = all.ComponentNames
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
