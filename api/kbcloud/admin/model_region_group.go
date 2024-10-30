// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// RegionGroup The region group 
type RegionGroup struct {
	// The name of the region group
	Group string `json:"group"`
	// The Chinese name of the region group.
	GroupCn string `json:"groupCN"`
	// The English name of the region group.
	GroupEn string `json:"groupEN"`
	// The name of the provider.
	Provider string `json:"provider"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewRegionGroup instantiates a new RegionGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegionGroup(group string, groupCn string, groupEn string, provider string) *RegionGroup {
	this := RegionGroup{}
	this.Group = group
	this.GroupCn = groupCn
	this.GroupEn = groupEn
	this.Provider = provider
	return &this
}

// NewRegionGroupWithDefaults instantiates a new RegionGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionGroupWithDefaults() *RegionGroup {
	this := RegionGroup{}
	return &this
}
// GetGroup returns the Group field value.
func (o *RegionGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *RegionGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *RegionGroup) SetGroup(v string) {
	o.Group = v
}


// GetGroupCn returns the GroupCn field value.
func (o *RegionGroup) GetGroupCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupCn
}

// GetGroupCnOk returns a tuple with the GroupCn field value
// and a boolean to check if the value has been set.
func (o *RegionGroup) GetGroupCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupCn, true
}

// SetGroupCn sets field value.
func (o *RegionGroup) SetGroupCn(v string) {
	o.GroupCn = v
}


// GetGroupEn returns the GroupEn field value.
func (o *RegionGroup) GetGroupEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupEn
}

// GetGroupEnOk returns a tuple with the GroupEn field value
// and a boolean to check if the value has been set.
func (o *RegionGroup) GetGroupEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupEn, true
}

// SetGroupEn sets field value.
func (o *RegionGroup) SetGroupEn(v string) {
	o.GroupEn = v
}


// GetProvider returns the Provider field value.
func (o *RegionGroup) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *RegionGroup) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *RegionGroup) SetProvider(v string) {
	o.Provider = v
}



// MarshalJSON serializes the struct using spec logic.
func (o RegionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["groupCN"] = o.GroupCn
	toSerialize["groupEN"] = o.GroupEn
	toSerialize["provider"] = o.Provider

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RegionGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group *string `json:"group"`
		GroupCn *string `json:"groupCN"`
		GroupEn *string `json:"groupEN"`
		Provider *string `json:"provider"`
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
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "group", "groupCN", "groupEN", "provider",  })
	} else {
		return err
	}
	o.Group = *all.Group
	o.GroupCn = *all.GroupCn
	o.GroupEn = *all.GroupEn
	o.Provider = *all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
