// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DatabaseOption struct {
	Enabled       bool                              `json:"enabled"`
	CreateOptions []DatabaseOptionCreateOptionsItem `json:"createOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseOption instantiates a new DatabaseOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseOption(enabled bool) *DatabaseOption {
	this := DatabaseOption{}
	this.Enabled = enabled
	return &this
}

// NewDatabaseOptionWithDefaults instantiates a new DatabaseOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseOptionWithDefaults() *DatabaseOption {
	this := DatabaseOption{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *DatabaseOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *DatabaseOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreateOptions returns the CreateOptions field value if set, zero value otherwise.
func (o *DatabaseOption) GetCreateOptions() []DatabaseOptionCreateOptionsItem {
	if o == nil || o.CreateOptions == nil {
		var ret []DatabaseOptionCreateOptionsItem
		return ret
	}
	return o.CreateOptions
}

// GetCreateOptionsOk returns a tuple with the CreateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetCreateOptionsOk() (*[]DatabaseOptionCreateOptionsItem, bool) {
	if o == nil || o.CreateOptions == nil {
		return nil, false
	}
	return &o.CreateOptions, true
}

// HasCreateOptions returns a boolean if a field has been set.
func (o *DatabaseOption) HasCreateOptions() bool {
	return o != nil && o.CreateOptions != nil
}

// SetCreateOptions gets a reference to the given []DatabaseOptionCreateOptionsItem and assigns it to the CreateOptions field.
func (o *DatabaseOption) SetCreateOptions(v []DatabaseOptionCreateOptionsItem) {
	o.CreateOptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.CreateOptions != nil {
		toSerialize["createOptions"] = o.CreateOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool                             `json:"enabled"`
		CreateOptions []DatabaseOptionCreateOptionsItem `json:"createOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "createOptions"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.CreateOptions = all.CreateOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
