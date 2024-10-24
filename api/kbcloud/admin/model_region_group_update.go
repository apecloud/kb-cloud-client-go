// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RegionGroupUpdate The region group that needs to be updated.
// NODESCRIPTION RegionGroupUpdate
//
// Deprecated: This model is deprecated.
type RegionGroupUpdate struct {
	// The Chinese name of the region group.
	GroupCn *string `json:"groupCN,omitempty"`
	// The English name of the region group.
	GroupEn *string `json:"groupEN,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegionGroupUpdate instantiates a new RegionGroupUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegionGroupUpdate() *RegionGroupUpdate {
	this := RegionGroupUpdate{}
	return &this
}

// NewRegionGroupUpdateWithDefaults instantiates a new RegionGroupUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionGroupUpdateWithDefaults() *RegionGroupUpdate {
	this := RegionGroupUpdate{}
	return &this
}

// GetGroupCn returns the GroupCn field value if set, zero value otherwise.
func (o *RegionGroupUpdate) GetGroupCn() string {
	if o == nil || o.GroupCn == nil {
		var ret string
		return ret
	}
	return *o.GroupCn
}

// GetGroupCnOk returns a tuple with the GroupCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionGroupUpdate) GetGroupCnOk() (*string, bool) {
	if o == nil || o.GroupCn == nil {
		return nil, false
	}
	return o.GroupCn, true
}

// HasGroupCn returns a boolean if a field has been set.
func (o *RegionGroupUpdate) HasGroupCn() bool {
	return o != nil && o.GroupCn != nil
}

// SetGroupCn gets a reference to the given string and assigns it to the GroupCn field.
func (o *RegionGroupUpdate) SetGroupCn(v string) {
	o.GroupCn = &v
}

// GetGroupEn returns the GroupEn field value if set, zero value otherwise.
func (o *RegionGroupUpdate) GetGroupEn() string {
	if o == nil || o.GroupEn == nil {
		var ret string
		return ret
	}
	return *o.GroupEn
}

// GetGroupEnOk returns a tuple with the GroupEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionGroupUpdate) GetGroupEnOk() (*string, bool) {
	if o == nil || o.GroupEn == nil {
		return nil, false
	}
	return o.GroupEn, true
}

// HasGroupEn returns a boolean if a field has been set.
func (o *RegionGroupUpdate) HasGroupEn() bool {
	return o != nil && o.GroupEn != nil
}

// SetGroupEn gets a reference to the given string and assigns it to the GroupEn field.
func (o *RegionGroupUpdate) SetGroupEn(v string) {
	o.GroupEn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RegionGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GroupCn != nil {
		toSerialize["groupCN"] = o.GroupCn
	}
	if o.GroupEn != nil {
		toSerialize["groupEN"] = o.GroupEn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RegionGroupUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupCn *string `json:"groupCN,omitempty"`
		GroupEn *string `json:"groupEN,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"groupCN", "groupEN"})
	} else {
		return err
	}
	o.GroupCn = all.GroupCn
	o.GroupEn = all.GroupEn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
