// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentPods Pods information for a component
type ComponentPods struct {
	// Environment ID
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Component name
	ComponentName *string `json:"component_name,omitempty"`
	// Last update time of the data
	LastUpdated *time.Time              `json:"last_updated,omitempty"`
	Pods        []ComponentPodsPodsItem `json:"pods,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentPods instantiates a new ComponentPods object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentPods() *ComponentPods {
	this := ComponentPods{}
	return &this
}

// NewComponentPodsWithDefaults instantiates a new ComponentPods object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentPodsWithDefaults() *ComponentPods {
	this := ComponentPods{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ComponentPods) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPods) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ComponentPods) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ComponentPods) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ComponentPods) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPods) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ComponentPods) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ComponentPods) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ComponentPods) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPods) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ComponentPods) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ComponentPods) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *ComponentPods) GetPods() []ComponentPodsPodsItem {
	if o == nil || o.Pods == nil {
		var ret []ComponentPodsPodsItem
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPods) GetPodsOk() (*[]ComponentPodsPodsItem, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return &o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *ComponentPods) HasPods() bool {
	return o != nil && o.Pods != nil
}

// SetPods gets a reference to the given []ComponentPodsPodsItem and assigns it to the Pods field.
func (o *ComponentPods) SetPods(v []ComponentPodsPodsItem) {
	o.Pods = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentPods) MarshalJSON() ([]byte, error) {
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
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentPods) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId *string                 `json:"environment_id,omitempty"`
		ComponentName *string                 `json:"component_name,omitempty"`
		LastUpdated   *time.Time              `json:"last_updated,omitempty"`
		Pods          []ComponentPodsPodsItem `json:"pods,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment_id", "component_name", "last_updated", "pods"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.ComponentName = all.ComponentName
	o.LastUpdated = all.LastUpdated
	o.Pods = all.Pods

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
