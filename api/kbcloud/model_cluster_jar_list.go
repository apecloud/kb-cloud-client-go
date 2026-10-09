// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterJarList struct {
	Items      []ClusterJarPackage `json:"items"`
	Generation int64               `json:"generation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterJarList instantiates a new ClusterJarList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterJarList(items []ClusterJarPackage, generation int64) *ClusterJarList {
	this := ClusterJarList{}
	this.Items = items
	this.Generation = generation
	return &this
}

// NewClusterJarListWithDefaults instantiates a new ClusterJarList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterJarListWithDefaults() *ClusterJarList {
	this := ClusterJarList{}
	return &this
}

// GetItems returns the Items field value.
func (o *ClusterJarList) GetItems() []ClusterJarPackage {
	if o == nil {
		var ret []ClusterJarPackage
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ClusterJarList) GetItemsOk() (*[]ClusterJarPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ClusterJarList) SetItems(v []ClusterJarPackage) {
	o.Items = v
}

// GetGeneration returns the Generation field value.
func (o *ClusterJarList) GetGeneration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value
// and a boolean to check if the value has been set.
func (o *ClusterJarList) GetGenerationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Generation, true
}

// SetGeneration sets field value.
func (o *ClusterJarList) SetGeneration(v int64) {
	o.Generation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterJarList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["generation"] = o.Generation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterJarList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]ClusterJarPackage `json:"items"`
		Generation *int64               `json:"generation"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Generation == nil {
		return fmt.Errorf("required field generation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "generation"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.Generation = *all.Generation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
