// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsUpgrade Upgrade component and/or selected instance templates. Only explicitly specified target fields are updated. An empty serviceVersion on a template in the Cluster CR inherits the component version and may change when that component is upgraded; this is distinct from omitting a template from this request.
type OpsUpgrade struct {
	Version           common.NullableString     `json:"version,omitempty"`
	InstanceTemplates []InstanceTemplateUpgrade `json:"instanceTemplates,omitempty"`
	Component         string                    `json:"component"`
	Schedule          *TaskSchedule             `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsUpgrade instantiates a new OpsUpgrade object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsUpgrade(component string) *OpsUpgrade {
	this := OpsUpgrade{}
	this.Component = component
	return &this
}

// NewOpsUpgradeWithDefaults instantiates a new OpsUpgrade object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsUpgradeWithDefaults() *OpsUpgrade {
	this := OpsUpgrade{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsUpgrade) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsUpgrade) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *OpsUpgrade) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given common.NullableString and assigns it to the Version field.
func (o *OpsUpgrade) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *OpsUpgrade) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *OpsUpgrade) UnsetVersion() {
	o.Version.Unset()
}

// GetInstanceTemplates returns the InstanceTemplates field value if set, zero value otherwise.
func (o *OpsUpgrade) GetInstanceTemplates() []InstanceTemplateUpgrade {
	if o == nil || o.InstanceTemplates == nil {
		var ret []InstanceTemplateUpgrade
		return ret
	}
	return o.InstanceTemplates
}

// GetInstanceTemplatesOk returns a tuple with the InstanceTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsUpgrade) GetInstanceTemplatesOk() (*[]InstanceTemplateUpgrade, bool) {
	if o == nil || o.InstanceTemplates == nil {
		return nil, false
	}
	return &o.InstanceTemplates, true
}

// HasInstanceTemplates returns a boolean if a field has been set.
func (o *OpsUpgrade) HasInstanceTemplates() bool {
	return o != nil && o.InstanceTemplates != nil
}

// SetInstanceTemplates gets a reference to the given []InstanceTemplateUpgrade and assigns it to the InstanceTemplates field.
func (o *OpsUpgrade) SetInstanceTemplates(v []InstanceTemplateUpgrade) {
	o.InstanceTemplates = v
}

// GetComponent returns the Component field value.
func (o *OpsUpgrade) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsUpgrade) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsUpgrade) SetComponent(v string) {
	o.Component = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsUpgrade) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsUpgrade) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsUpgrade) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsUpgrade) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.InstanceTemplates != nil {
		toSerialize["instanceTemplates"] = o.InstanceTemplates
	}
	toSerialize["component"] = o.Component
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Version           common.NullableString     `json:"version,omitempty"`
		InstanceTemplates []InstanceTemplateUpgrade `json:"instanceTemplates,omitempty"`
		Component         *string                   `json:"component"`
		Schedule          *TaskSchedule             `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"version", "instanceTemplates", "component", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Version = all.Version
	o.InstanceTemplates = all.InstanceTemplates
	o.Component = *all.Component
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
