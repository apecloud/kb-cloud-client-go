// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterTaskProgresses clusterTaskProgresses is a list of task progress detail
type ClusterTaskProgresses struct {
	// Items is the list of task progress detail in the list
	Items []ClusterTaskProgress `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterTaskProgresses instantiates a new ClusterTaskProgresses object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterTaskProgresses() *ClusterTaskProgresses {
	this := ClusterTaskProgresses{}
	return &this
}

// NewClusterTaskProgressesWithDefaults instantiates a new ClusterTaskProgresses object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterTaskProgressesWithDefaults() *ClusterTaskProgresses {
	this := ClusterTaskProgresses{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterTaskProgresses) GetItems() []ClusterTaskProgress {
	if o == nil || o.Items == nil {
		var ret []ClusterTaskProgress
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgresses) GetItemsOk() (*[]ClusterTaskProgress, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterTaskProgresses) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ClusterTaskProgress and assigns it to the Items field.
func (o *ClusterTaskProgresses) SetItems(v []ClusterTaskProgress) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterTaskProgresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterTaskProgresses) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items []ClusterTaskProgress `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
