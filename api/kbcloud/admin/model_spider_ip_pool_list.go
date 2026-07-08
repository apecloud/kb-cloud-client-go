// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SpiderIPPoolList SpiderIPPool list for an environment.
type SpiderIPPoolList struct {
	// Whether Spiderpool is enabled for the environment.
	SpiderpoolEnabled bool           `json:"spiderpoolEnabled"`
	Items             []SpiderIPPool `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpiderIPPoolList instantiates a new SpiderIPPoolList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpiderIPPoolList(spiderpoolEnabled bool, items []SpiderIPPool) *SpiderIPPoolList {
	this := SpiderIPPoolList{}
	this.SpiderpoolEnabled = spiderpoolEnabled
	this.Items = items
	return &this
}

// NewSpiderIPPoolListWithDefaults instantiates a new SpiderIPPoolList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpiderIPPoolListWithDefaults() *SpiderIPPoolList {
	this := SpiderIPPoolList{}
	var spiderpoolEnabled bool = false
	this.SpiderpoolEnabled = spiderpoolEnabled
	return &this
}

// GetSpiderpoolEnabled returns the SpiderpoolEnabled field value.
func (o *SpiderIPPoolList) GetSpiderpoolEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SpiderpoolEnabled
}

// GetSpiderpoolEnabledOk returns a tuple with the SpiderpoolEnabled field value
// and a boolean to check if the value has been set.
func (o *SpiderIPPoolList) GetSpiderpoolEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpiderpoolEnabled, true
}

// SetSpiderpoolEnabled sets field value.
func (o *SpiderIPPoolList) SetSpiderpoolEnabled(v bool) {
	o.SpiderpoolEnabled = v
}

// GetItems returns the Items field value.
func (o *SpiderIPPoolList) GetItems() []SpiderIPPool {
	if o == nil {
		var ret []SpiderIPPool
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SpiderIPPoolList) GetItemsOk() (*[]SpiderIPPool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *SpiderIPPoolList) SetItems(v []SpiderIPPool) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpiderIPPoolList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["spiderpoolEnabled"] = o.SpiderpoolEnabled
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpiderIPPoolList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SpiderpoolEnabled *bool           `json:"spiderpoolEnabled"`
		Items             *[]SpiderIPPool `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SpiderpoolEnabled == nil {
		return fmt.Errorf("required field spiderpoolEnabled missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"spiderpoolEnabled", "items"})
	} else {
		return err
	}
	o.SpiderpoolEnabled = *all.SpiderpoolEnabled
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
