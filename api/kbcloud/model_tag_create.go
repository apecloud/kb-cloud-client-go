// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// TagCreate Tag create 
type TagCreate struct {
	// The cluster id corresponding to the tag
	ClusterId string `json:"clusterId"`
	// The key value pair of the tag needed to create
	Items []TagCreateItemsItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewTagCreate instantiates a new TagCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagCreate(clusterId string, items []TagCreateItemsItem) *TagCreate {
	this := TagCreate{}
	this.ClusterId = clusterId
	this.Items = items
	return &this
}

// NewTagCreateWithDefaults instantiates a new TagCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagCreateWithDefaults() *TagCreate {
	this := TagCreate{}
	return &this
}
// GetClusterId returns the ClusterId field value.
func (o *TagCreate) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *TagCreate) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *TagCreate) SetClusterId(v string) {
	o.ClusterId = v
}


// GetItems returns the Items field value.
func (o *TagCreate) GetItems() []TagCreateItemsItem {
	if o == nil {
		var ret []TagCreateItemsItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TagCreate) GetItemsOk() (*[]TagCreateItemsItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *TagCreate) SetItems(v []TagCreateItemsItem) {
	o.Items = v
}



// MarshalJSON serializes the struct using spec logic.
func (o TagCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterId"] = o.ClusterId
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId *string `json:"clusterId"`
		Items *[]TagCreateItemsItem `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterId missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "clusterId", "items",  })
	} else {
		return err
	}
	o.ClusterId = *all.ClusterId
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
