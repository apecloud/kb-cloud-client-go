// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterEventList struct {
	Items []ClusterEventItem `json:"items"`
	// Whether more matching events exist than the limit
	Truncated bool `json:"truncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterEventList instantiates a new ClusterEventList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterEventList(items []ClusterEventItem, truncated bool) *ClusterEventList {
	this := ClusterEventList{}
	this.Items = items
	this.Truncated = truncated
	return &this
}

// NewClusterEventListWithDefaults instantiates a new ClusterEventList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterEventListWithDefaults() *ClusterEventList {
	this := ClusterEventList{}
	return &this
}

// GetItems returns the Items field value.
func (o *ClusterEventList) GetItems() []ClusterEventItem {
	if o == nil {
		var ret []ClusterEventItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ClusterEventList) GetItemsOk() (*[]ClusterEventItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ClusterEventList) SetItems(v []ClusterEventItem) {
	o.Items = v
}

// GetTruncated returns the Truncated field value.
func (o *ClusterEventList) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *ClusterEventList) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *ClusterEventList) SetTruncated(v bool) {
	o.Truncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterEventList) MarshalJSON() ([]byte, error) {
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
func (o *ClusterEventList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items     *[]ClusterEventItem `json:"items"`
		Truncated *bool               `json:"truncated"`
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
