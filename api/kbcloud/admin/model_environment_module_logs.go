// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleLogs Environment module pod logs
type EnvironmentModuleLogs struct {
	// Environment ID
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Environment module name
	ModuleName *string `json:"module_name,omitempty"`
	// Pod name
	PodName *string `json:"pod_name,omitempty"`
	// Container name
	ContainerName *string    `json:"container_name,omitempty"`
	Logs          []LogEntry `json:"logs,omitempty"`
	// Next timestamp for pagination
	NextTimestamp *time.Time `json:"next_timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleLogs instantiates a new EnvironmentModuleLogs object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleLogs() *EnvironmentModuleLogs {
	this := EnvironmentModuleLogs{}
	return &this
}

// NewEnvironmentModuleLogsWithDefaults instantiates a new EnvironmentModuleLogs object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleLogsWithDefaults() *EnvironmentModuleLogs {
	this := EnvironmentModuleLogs{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *EnvironmentModuleLogs) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasModuleName() bool {
	return o != nil && o.ModuleName != nil
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *EnvironmentModuleLogs) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasPodName() bool {
	return o != nil && o.PodName != nil
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *EnvironmentModuleLogs) SetPodName(v string) {
	o.PodName = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetContainerName() string {
	if o == nil || o.ContainerName == nil {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetContainerNameOk() (*string, bool) {
	if o == nil || o.ContainerName == nil {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasContainerName() bool {
	return o != nil && o.ContainerName != nil
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *EnvironmentModuleLogs) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetLogs() []LogEntry {
	if o == nil || o.Logs == nil {
		var ret []LogEntry
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetLogsOk() (*[]LogEntry, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return &o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasLogs() bool {
	return o != nil && o.Logs != nil
}

// SetLogs gets a reference to the given []LogEntry and assigns it to the Logs field.
func (o *EnvironmentModuleLogs) SetLogs(v []LogEntry) {
	o.Logs = v
}

// GetNextTimestamp returns the NextTimestamp field value if set, zero value otherwise.
func (o *EnvironmentModuleLogs) GetNextTimestamp() time.Time {
	if o == nil || o.NextTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.NextTimestamp
}

// GetNextTimestampOk returns a tuple with the NextTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleLogs) GetNextTimestampOk() (*time.Time, bool) {
	if o == nil || o.NextTimestamp == nil {
		return nil, false
	}
	return o.NextTimestamp, true
}

// HasNextTimestamp returns a boolean if a field has been set.
func (o *EnvironmentModuleLogs) HasNextTimestamp() bool {
	return o != nil && o.NextTimestamp != nil
}

// SetNextTimestamp gets a reference to the given time.Time and assigns it to the NextTimestamp field.
func (o *EnvironmentModuleLogs) SetNextTimestamp(v time.Time) {
	o.NextTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.ModuleName != nil {
		toSerialize["module_name"] = o.ModuleName
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
func (o *EnvironmentModuleLogs) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId *string    `json:"environment_id,omitempty"`
		ModuleName    *string    `json:"module_name,omitempty"`
		PodName       *string    `json:"pod_name,omitempty"`
		ContainerName *string    `json:"container_name,omitempty"`
		Logs          []LogEntry `json:"logs,omitempty"`
		NextTimestamp *time.Time `json:"next_timestamp,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment_id", "module_name", "pod_name", "container_name", "logs", "next_timestamp"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.ModuleName = all.ModuleName
	o.PodName = all.PodName
	o.ContainerName = all.ContainerName
	o.Logs = all.Logs
	o.NextTimestamp = all.NextTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
