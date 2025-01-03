// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentLogs Component pod logs
type ComponentLogs struct {
	// Environment ID
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Component name
	ComponentName *string `json:"component_name,omitempty"`
	// Pod name
	PodName *string `json:"pod_name,omitempty"`
	// Container name
	ContainerName *string                 `json:"container_name,omitempty"`
	Logs          []ComponentLogsLogsItem `json:"logs,omitempty"`
	// Next timestamp for pagination
	NextTimestamp *time.Time `json:"next_timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentLogs instantiates a new ComponentLogs object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentLogs() *ComponentLogs {
	this := ComponentLogs{}
	return &this
}

// NewComponentLogsWithDefaults instantiates a new ComponentLogs object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentLogsWithDefaults() *ComponentLogs {
	this := ComponentLogs{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ComponentLogs) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ComponentLogs) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ComponentLogs) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ComponentLogs) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ComponentLogs) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ComponentLogs) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *ComponentLogs) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *ComponentLogs) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *ComponentLogs) SetPodName(v string) {
	o.PodName = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *ComponentLogs) GetContainerName() string {
	if o == nil || o.ContainerName == nil {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetContainerNameOk() (*string, bool) {
	if o == nil || o.ContainerName == nil {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *ComponentLogs) HasContainerName() bool {
	return o != nil && o.ContainerName != nil
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *ComponentLogs) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ComponentLogs) GetLogs() []ComponentLogsLogsItem {
	if o == nil || o.Logs == nil {
		var ret []ComponentLogsLogsItem
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetLogsOk() (*[]ComponentLogsLogsItem, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return &o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ComponentLogs) HasLogs() bool {
	return o != nil && o.Logs != nil
}

// SetLogs gets a reference to the given []ComponentLogsLogsItem and assigns it to the Logs field.
func (o *ComponentLogs) SetLogs(v []ComponentLogsLogsItem) {
	o.Logs = v
}

// GetNextTimestamp returns the NextTimestamp field value if set, zero value otherwise.
func (o *ComponentLogs) GetNextTimestamp() time.Time {
	if o == nil || o.NextTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.NextTimestamp
}

// GetNextTimestampOk returns a tuple with the NextTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentLogs) GetNextTimestampOk() (*time.Time, bool) {
	if o == nil || o.NextTimestamp == nil {
		return nil, false
	}
	return o.NextTimestamp, true
}

// HasNextTimestamp returns a boolean if a field has been set.
func (o *ComponentLogs) HasNextTimestamp() bool {
	return o != nil && o.NextTimestamp != nil
}

// SetNextTimestamp gets a reference to the given time.Time and assigns it to the NextTimestamp field.
func (o *ComponentLogs) SetNextTimestamp(v time.Time) {
	o.NextTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.ComponentName != nil {
		toSerialize["component_name"] = o.ComponentName
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
	}
	if o.ContainerName != nil {
		toSerialize["container_name"] = o.ContainerName
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	if o.NextTimestamp != nil {
		if o.NextTimestamp.Nanosecond() == 0 {
			toSerialize["next_timestamp"] = o.NextTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["next_timestamp"] = o.NextTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentLogs) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId *string                 `json:"environment_id,omitempty"`
		ComponentName *string                 `json:"component_name,omitempty"`
		PodName       *string                 `json:"pod_name,omitempty"`
		ContainerName *string                 `json:"container_name,omitempty"`
		Logs          []ComponentLogsLogsItem `json:"logs,omitempty"`
		NextTimestamp *time.Time              `json:"next_timestamp,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment_id", "component_name", "pod_name", "container_name", "logs", "next_timestamp"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.ComponentName = all.ComponentName
	o.PodName = all.PodName
	o.ContainerName = all.ContainerName
	o.Logs = all.Logs
	o.NextTimestamp = all.NextTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
