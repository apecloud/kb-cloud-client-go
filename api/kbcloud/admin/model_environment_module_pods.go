// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModulePods Pods information for an environment module
type EnvironmentModulePods struct {
	// Environment name
	EnvironmentName *string `json:"environment_name,omitempty"`
	// Environment module name
	ModuleName *string `json:"module_name,omitempty"`
	// Last update time of the data
	LastUpdated *time.Time  `json:"last_updated,omitempty"`
	Pods        []ModulePod `json:"pods,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModulePods instantiates a new EnvironmentModulePods object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModulePods() *EnvironmentModulePods {
	this := EnvironmentModulePods{}
	return &this
}

// NewEnvironmentModulePodsWithDefaults instantiates a new EnvironmentModulePods object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModulePodsWithDefaults() *EnvironmentModulePods {
	this := EnvironmentModulePods{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *EnvironmentModulePods) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePods) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *EnvironmentModulePods) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *EnvironmentModulePods) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *EnvironmentModulePods) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePods) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *EnvironmentModulePods) HasModuleName() bool {
	return o != nil && o.ModuleName != nil
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *EnvironmentModulePods) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EnvironmentModulePods) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePods) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EnvironmentModulePods) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EnvironmentModulePods) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *EnvironmentModulePods) GetPods() []ModulePod {
	if o == nil || o.Pods == nil {
		var ret []ModulePod
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModulePods) GetPodsOk() (*[]ModulePod, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return &o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *EnvironmentModulePods) HasPods() bool {
	return o != nil && o.Pods != nil
}

// SetPods gets a reference to the given []ModulePod and assigns it to the Pods field.
func (o *EnvironmentModulePods) SetPods(v []ModulePod) {
	o.Pods = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModulePods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentName != nil {
		toSerialize["environment_name"] = o.EnvironmentName
	}
	if o.ModuleName != nil {
		toSerialize["module_name"] = o.ModuleName
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
func (o *EnvironmentModulePods) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentName *string     `json:"environment_name,omitempty"`
		ModuleName      *string     `json:"module_name,omitempty"`
		LastUpdated     *time.Time  `json:"last_updated,omitempty"`
		Pods            []ModulePod `json:"pods,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment_name", "module_name", "last_updated", "pods"})
	} else {
		return err
	}
	o.EnvironmentName = all.EnvironmentName
	o.ModuleName = all.ModuleName
	o.LastUpdated = all.LastUpdated
	o.Pods = all.Pods

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
