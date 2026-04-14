// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsArchitectureChange OpsArchitectureChange is the payload to upgrade a standalone cluster to primary-replica/HA mode in-place.
type OpsArchitectureChange struct {
	// Target engine mode (e.g. "replication")
	TargetMode string `json:"targetMode"`
	// Number of replicas in target mode (minimum 2)
	Replicas int32         `json:"replicas"`
	Schedule *TaskSchedule `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsArchitectureChange instantiates a new OpsArchitectureChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsArchitectureChange(targetMode string, replicas int32) *OpsArchitectureChange {
	this := OpsArchitectureChange{}
	this.TargetMode = targetMode
	this.Replicas = replicas
	return &this
}

// NewOpsArchitectureChangeWithDefaults instantiates a new OpsArchitectureChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsArchitectureChangeWithDefaults() *OpsArchitectureChange {
	this := OpsArchitectureChange{}
	return &this
}

// GetTargetMode returns the TargetMode field value.
func (o *OpsArchitectureChange) GetTargetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetMode
}

// GetTargetModeOk returns a tuple with the TargetMode field value
// and a boolean to check if the value has been set.
func (o *OpsArchitectureChange) GetTargetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetMode, true
}

// SetTargetMode sets field value.
func (o *OpsArchitectureChange) SetTargetMode(v string) {
	o.TargetMode = v
}

// GetReplicas returns the Replicas field value.
func (o *OpsArchitectureChange) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *OpsArchitectureChange) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *OpsArchitectureChange) SetReplicas(v int32) {
	o.Replicas = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsArchitectureChange) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsArchitectureChange) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsArchitectureChange) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsArchitectureChange) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsArchitectureChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["targetMode"] = o.TargetMode
	toSerialize["replicas"] = o.Replicas
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsArchitectureChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TargetMode *string       `json:"targetMode"`
		Replicas   *int32        `json:"replicas"`
		Schedule   *TaskSchedule `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TargetMode == nil {
		return fmt.Errorf("required field targetMode missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"targetMode", "replicas", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TargetMode = *all.TargetMode
	o.Replicas = *all.Replicas
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
