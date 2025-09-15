// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentPricingList list of information with environment pricing
type EnvironmentPricingList struct {
	// Items is the list of environment pricing information
	Items []EnvironmentPricing `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentPricingList instantiates a new EnvironmentPricingList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentPricingList(items []EnvironmentPricing) *EnvironmentPricingList {
	this := EnvironmentPricingList{}
	this.Items = items
	return &this
}

// NewEnvironmentPricingListWithDefaults instantiates a new EnvironmentPricingList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentPricingListWithDefaults() *EnvironmentPricingList {
	this := EnvironmentPricingList{}
	return &this
}

// GetItems returns the Items field value.
func (o *EnvironmentPricingList) GetItems() []EnvironmentPricing {
	if o == nil {
		var ret []EnvironmentPricing
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EnvironmentPricingList) GetItemsOk() (*[]EnvironmentPricing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *EnvironmentPricingList) SetItems(v []EnvironmentPricing) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentPricingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentPricingList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]EnvironmentPricing `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
