// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgConfigurationItem struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConfigurationItem instantiates a new OrgConfigurationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConfigurationItem(id string, displayName string) *OrgConfigurationItem {
	this := OrgConfigurationItem{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewOrgConfigurationItemWithDefaults instantiates a new OrgConfigurationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConfigurationItemWithDefaults() *OrgConfigurationItem {
	this := OrgConfigurationItem{}
	return &this
}

// GetId returns the Id field value.
func (o *OrgConfigurationItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OrgConfigurationItem) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value.
func (o *OrgConfigurationItem) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationItem) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *OrgConfigurationItem) SetDisplayName(v string) {
	o.DisplayName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConfigurationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConfigurationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string `json:"id"`
		DisplayName *string `json:"displayName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "displayName"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.DisplayName = *all.DisplayName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
