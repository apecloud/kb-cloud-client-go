// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DashboardConfig struct {
	// Frontend-defined platform-wide key. Organization identifiers in this key do not enforce tenant isolation.
	Key string `json:"key"`
	// Stored JSON value, or null when the key has no configuration.
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardConfig instantiates a new DashboardConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardConfig(key string) *DashboardConfig {
	this := DashboardConfig{}
	this.Key = key
	return &this
}

// NewDashboardConfigWithDefaults instantiates a new DashboardConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardConfigWithDefaults() *DashboardConfig {
	this := DashboardConfig{}
	return &this
}

// GetKey returns the Key field value.
func (o *DashboardConfig) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DashboardConfig) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *DashboardConfig) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DashboardConfig) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardConfig) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DashboardConfig) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *DashboardConfig) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string     `json:"key"`
		Value interface{} `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "value"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
