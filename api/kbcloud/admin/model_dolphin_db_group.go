// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBGroup struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBGroup instantiates a new DolphinDBGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBGroup(name string, members []string) *DolphinDBGroup {
	this := DolphinDBGroup{}
	this.Name = name
	this.Members = members
	return &this
}

// NewDolphinDBGroupWithDefaults instantiates a new DolphinDBGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBGroupWithDefaults() *DolphinDBGroup {
	this := DolphinDBGroup{}
	return &this
}

// GetName returns the Name field value.
func (o *DolphinDBGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBGroup) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value.
func (o *DolphinDBGroup) GetMembers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *DolphinDBGroup) GetMembersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value.
func (o *DolphinDBGroup) SetMembers(v []string) {
	o.Members = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["members"] = o.Members

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string   `json:"name"`
		Members *[]string `json:"members"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Members == nil {
		return fmt.Errorf("required field members missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "members"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Members = *all.Members

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
