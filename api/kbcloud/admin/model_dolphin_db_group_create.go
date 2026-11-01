// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBGroupCreate struct {
	Name    string   `json:"name"`
	Members []string `json:"members,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBGroupCreate instantiates a new DolphinDBGroupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBGroupCreate(name string) *DolphinDBGroupCreate {
	this := DolphinDBGroupCreate{}
	this.Name = name
	return &this
}

// NewDolphinDBGroupCreateWithDefaults instantiates a new DolphinDBGroupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBGroupCreateWithDefaults() *DolphinDBGroupCreate {
	this := DolphinDBGroupCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *DolphinDBGroupCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DolphinDBGroupCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DolphinDBGroupCreate) SetName(v string) {
	o.Name = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DolphinDBGroupCreate) GetMembers() []string {
	if o == nil || o.Members == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBGroupCreate) GetMembersOk() (*[]string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return &o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DolphinDBGroupCreate) HasMembers() bool {
	return o != nil && o.Members != nil
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *DolphinDBGroupCreate) SetMembers(v []string) {
	o.Members = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBGroupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string  `json:"name"`
		Members []string `json:"members,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "members"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Members = all.Members

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
