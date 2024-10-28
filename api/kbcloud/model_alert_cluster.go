// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION AlertCluster
type AlertCluster struct {
	// NODESCRIPTION Disabled
	Disabled bool `json:"disabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertCluster instantiates a new AlertCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertCluster(disabled bool) *AlertCluster {
	this := AlertCluster{}
	this.Disabled = disabled
	return &this
}

// NewAlertClusterWithDefaults instantiates a new AlertCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertClusterWithDefaults() *AlertCluster {
	this := AlertCluster{}
	return &this
}

// GetDisabled returns the Disabled field value.
func (o *AlertCluster) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *AlertCluster) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value.
func (o *AlertCluster) SetDisabled(v bool) {
	o.Disabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["disabled"] = o.Disabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled *bool `json:"disabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Disabled == nil {
		return fmt.Errorf("required field disabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"disabled"})
	} else {
		return err
	}
	o.Disabled = *all.Disabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
