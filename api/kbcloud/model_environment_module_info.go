// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleInfo Environment module information in an environment
type EnvironmentModuleInfo struct {
	// Environment ID
	EnvironmentId *string `json:"environmentId,omitempty"`
	// Last update time of the data
	LastUpdated        *time.Time          `json:"lastUpdated,omitempty"`
	EnvironmentModules []EnvironmentModule `json:"environmentModules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleInfo instantiates a new EnvironmentModuleInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleInfo() *EnvironmentModuleInfo {
	this := EnvironmentModuleInfo{}
	return &this
}

// NewEnvironmentModuleInfoWithDefaults instantiates a new EnvironmentModuleInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleInfoWithDefaults() *EnvironmentModuleInfo {
	this := EnvironmentModuleInfo{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *EnvironmentModuleInfo) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInfo) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *EnvironmentModuleInfo) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *EnvironmentModuleInfo) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EnvironmentModuleInfo) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInfo) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EnvironmentModuleInfo) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EnvironmentModuleInfo) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetEnvironmentModules returns the EnvironmentModules field value if set, zero value otherwise.
func (o *EnvironmentModuleInfo) GetEnvironmentModules() []EnvironmentModule {
	if o == nil || o.EnvironmentModules == nil {
		var ret []EnvironmentModule
		return ret
	}
	return o.EnvironmentModules
}

// GetEnvironmentModulesOk returns a tuple with the EnvironmentModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInfo) GetEnvironmentModulesOk() (*[]EnvironmentModule, bool) {
	if o == nil || o.EnvironmentModules == nil {
		return nil, false
	}
	return &o.EnvironmentModules, true
}

// HasEnvironmentModules returns a boolean if a field has been set.
func (o *EnvironmentModuleInfo) HasEnvironmentModules() bool {
	return o != nil && o.EnvironmentModules != nil
}

// SetEnvironmentModules gets a reference to the given []EnvironmentModule and assigns it to the EnvironmentModules field.
func (o *EnvironmentModuleInfo) SetEnvironmentModules(v []EnvironmentModule) {
	o.EnvironmentModules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["lastUpdated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastUpdated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EnvironmentModules != nil {
		toSerialize["environmentModules"] = o.EnvironmentModules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId      *string             `json:"environmentId,omitempty"`
		LastUpdated        *time.Time          `json:"lastUpdated,omitempty"`
		EnvironmentModules []EnvironmentModule `json:"environmentModules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environmentId", "lastUpdated", "environmentModules"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.LastUpdated = all.LastUpdated
	o.EnvironmentModules = all.EnvironmentModules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
