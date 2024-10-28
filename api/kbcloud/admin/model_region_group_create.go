// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RegionGroupCreate The region group that needs to be created.

type RegionGroupCreate struct {
	// The name of the region group.
	Group string `json:"group"`
	// The Chinese name of the region group.
	GroupCn string `json:"groupCN"`
	// The English name of the region group.
	GroupEn string `json:"groupEN"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegionGroupCreate instantiates a new RegionGroupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegionGroupCreate(group string, groupCn string, groupEn string) *RegionGroupCreate {
	this := RegionGroupCreate{}
	this.Group = group
	this.GroupCn = groupCn
	this.GroupEn = groupEn
	return &this
}

// NewRegionGroupCreateWithDefaults instantiates a new RegionGroupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionGroupCreateWithDefaults() *RegionGroupCreate {
	this := RegionGroupCreate{}
	return &this
}

// GetGroup returns the Group field value.
func (o *RegionGroupCreate) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *RegionGroupCreate) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *RegionGroupCreate) SetGroup(v string) {
	o.Group = v
}

// GetGroupCn returns the GroupCn field value.
func (o *RegionGroupCreate) GetGroupCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupCn
}

// GetGroupCnOk returns a tuple with the GroupCn field value
// and a boolean to check if the value has been set.
func (o *RegionGroupCreate) GetGroupCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupCn, true
}

// SetGroupCn sets field value.
func (o *RegionGroupCreate) SetGroupCn(v string) {
	o.GroupCn = v
}

// GetGroupEn returns the GroupEn field value.
func (o *RegionGroupCreate) GetGroupEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupEn
}

// GetGroupEnOk returns a tuple with the GroupEn field value
// and a boolean to check if the value has been set.
func (o *RegionGroupCreate) GetGroupEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupEn, true
}

// SetGroupEn sets field value.
func (o *RegionGroupCreate) SetGroupEn(v string) {
	o.GroupEn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RegionGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["groupCN"] = o.GroupCn
	toSerialize["groupEN"] = o.GroupEn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RegionGroupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group   *string `json:"group"`
		GroupCn *string `json:"groupCN"`
		GroupEn *string `json:"groupEN"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.GroupCn == nil {
		return fmt.Errorf("required field groupCN missing")
	}
	if all.GroupEn == nil {
		return fmt.Errorf("required field groupEN missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"group", "groupCN", "groupEN"})
	} else {
		return err
	}
	o.Group = *all.Group
	o.GroupCn = *all.GroupCn
	o.GroupEn = *all.GroupEn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
