// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KubernetesEventList struct {
	Items []KubernetesEventItem `json:"items"`
	// Whether more matching events exist than the requested return number
	Truncated bool `json:"truncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKubernetesEventList instantiates a new KubernetesEventList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKubernetesEventList(items []KubernetesEventItem, truncated bool) *KubernetesEventList {
	this := KubernetesEventList{}
	this.Items = items
	this.Truncated = truncated
	return &this
}

// NewKubernetesEventListWithDefaults instantiates a new KubernetesEventList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKubernetesEventListWithDefaults() *KubernetesEventList {
	this := KubernetesEventList{}
	return &this
}

// GetItems returns the Items field value.
func (o *KubernetesEventList) GetItems() []KubernetesEventItem {
	if o == nil {
		var ret []KubernetesEventItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventList) GetItemsOk() (*[]KubernetesEventItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *KubernetesEventList) SetItems(v []KubernetesEventItem) {
	o.Items = v
}

// GetTruncated returns the Truncated field value.
func (o *KubernetesEventList) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventList) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *KubernetesEventList) SetTruncated(v bool) {
	o.Truncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KubernetesEventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KubernetesEventList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items     *[]KubernetesEventItem `json:"items"`
		Truncated *bool                  `json:"truncated"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "truncated"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.Truncated = *all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
