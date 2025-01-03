// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentInfo Component information in an environment
type ComponentInfo struct {
	// Environment ID
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Last update time of the data
	LastUpdated *time.Time                    `json:"last_updated,omitempty"`
	Components  []ComponentInfoComponentsItem `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentInfo instantiates a new ComponentInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentInfo() *ComponentInfo {
	this := ComponentInfo{}
	return &this
}

// NewComponentInfoWithDefaults instantiates a new ComponentInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentInfoWithDefaults() *ComponentInfo {
	this := ComponentInfo{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ComponentInfo) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfo) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ComponentInfo) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ComponentInfo) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ComponentInfo) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfo) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ComponentInfo) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ComponentInfo) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ComponentInfo) GetComponents() []ComponentInfoComponentsItem {
	if o == nil || o.Components == nil {
		var ret []ComponentInfoComponentsItem
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentInfo) GetComponentsOk() (*[]ComponentInfoComponentsItem, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ComponentInfo) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentInfoComponentsItem and assigns it to the Components field.
func (o *ComponentInfo) SetComponents(v []ComponentInfoComponentsItem) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId *string                       `json:"environment_id,omitempty"`
		LastUpdated   *time.Time                    `json:"last_updated,omitempty"`
		Components    []ComponentInfoComponentsItem `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment_id", "last_updated", "components"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.LastUpdated = all.LastUpdated
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
