// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleDetails details information for an environment module
type EnvironmentModuleDetails struct {
	// Environment name
	EnvironmentName *string `json:"environmentName,omitempty"`
	// Single environment module information
	ModuleInfo *EnvironmentModule `json:"moduleInfo,omitempty"`
	// Last update time of the data
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	ModulePods  []Pod      `json:"modulePods,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleDetails instantiates a new EnvironmentModuleDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleDetails() *EnvironmentModuleDetails {
	this := EnvironmentModuleDetails{}
	return &this
}

// NewEnvironmentModuleDetailsWithDefaults instantiates a new EnvironmentModuleDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleDetailsWithDefaults() *EnvironmentModuleDetails {
	this := EnvironmentModuleDetails{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *EnvironmentModuleDetails) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleDetails) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *EnvironmentModuleDetails) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *EnvironmentModuleDetails) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetModuleInfo returns the ModuleInfo field value if set, zero value otherwise.
func (o *EnvironmentModuleDetails) GetModuleInfo() EnvironmentModule {
	if o == nil || o.ModuleInfo == nil {
		var ret EnvironmentModule
		return ret
	}
	return *o.ModuleInfo
}

// GetModuleInfoOk returns a tuple with the ModuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleDetails) GetModuleInfoOk() (*EnvironmentModule, bool) {
	if o == nil || o.ModuleInfo == nil {
		return nil, false
	}
	return o.ModuleInfo, true
}

// HasModuleInfo returns a boolean if a field has been set.
func (o *EnvironmentModuleDetails) HasModuleInfo() bool {
	return o != nil && o.ModuleInfo != nil
}

// SetModuleInfo gets a reference to the given EnvironmentModule and assigns it to the ModuleInfo field.
func (o *EnvironmentModuleDetails) SetModuleInfo(v EnvironmentModule) {
	o.ModuleInfo = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EnvironmentModuleDetails) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleDetails) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EnvironmentModuleDetails) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EnvironmentModuleDetails) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetModulePods returns the ModulePods field value if set, zero value otherwise.
func (o *EnvironmentModuleDetails) GetModulePods() []Pod {
	if o == nil || o.ModulePods == nil {
		var ret []Pod
		return ret
	}
	return o.ModulePods
}

// GetModulePodsOk returns a tuple with the ModulePods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleDetails) GetModulePodsOk() (*[]Pod, bool) {
	if o == nil || o.ModulePods == nil {
		return nil, false
	}
	return &o.ModulePods, true
}

// HasModulePods returns a boolean if a field has been set.
func (o *EnvironmentModuleDetails) HasModulePods() bool {
	return o != nil && o.ModulePods != nil
}

// SetModulePods gets a reference to the given []Pod and assigns it to the ModulePods field.
func (o *EnvironmentModuleDetails) SetModulePods(v []Pod) {
	o.ModulePods = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.ModuleInfo != nil {
		toSerialize["moduleInfo"] = o.ModuleInfo
	}
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["lastUpdated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastUpdated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModulePods != nil {
		toSerialize["modulePods"] = o.ModulePods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentName *string            `json:"environmentName,omitempty"`
		ModuleInfo      *EnvironmentModule `json:"moduleInfo,omitempty"`
		LastUpdated     *time.Time         `json:"lastUpdated,omitempty"`
		ModulePods      []Pod              `json:"modulePods,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environmentName", "moduleInfo", "lastUpdated", "modulePods"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EnvironmentName = all.EnvironmentName
	if all.ModuleInfo != nil && all.ModuleInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModuleInfo = all.ModuleInfo
	o.LastUpdated = all.LastUpdated
	o.ModulePods = all.ModulePods

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
