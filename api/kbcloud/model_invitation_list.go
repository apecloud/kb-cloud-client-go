// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InvitationList InvitationList is a list of invitations
type InvitationList struct {
	// Items is the list of Invitation objects in the list
	Items []Invitation `json:"items"`
	// PageResult info
	PageResult *PageResult `json:"pageResult,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInvitationList instantiates a new InvitationList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInvitationList(items []Invitation) *InvitationList {
	this := InvitationList{}
	this.Items = items
	return &this
}

// NewInvitationListWithDefaults instantiates a new InvitationList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInvitationListWithDefaults() *InvitationList {
	this := InvitationList{}
	return &this
}

// GetItems returns the Items field value.
func (o *InvitationList) GetItems() []Invitation {
	if o == nil {
		var ret []Invitation
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InvitationList) GetItemsOk() (*[]Invitation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *InvitationList) SetItems(v []Invitation) {
	o.Items = v
}

// GetPageResult returns the PageResult field value if set, zero value otherwise.
func (o *InvitationList) GetPageResult() PageResult {
	if o == nil || o.PageResult == nil {
		var ret PageResult
		return ret
	}
	return *o.PageResult
}

// GetPageResultOk returns a tuple with the PageResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationList) GetPageResultOk() (*PageResult, bool) {
	if o == nil || o.PageResult == nil {
		return nil, false
	}
	return o.PageResult, true
}

// HasPageResult returns a boolean if a field has been set.
func (o *InvitationList) HasPageResult() bool {
	return o != nil && o.PageResult != nil
}

// SetPageResult gets a reference to the given PageResult and assigns it to the PageResult field.
func (o *InvitationList) SetPageResult(v PageResult) {
	o.PageResult = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InvitationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.PageResult != nil {
		toSerialize["pageResult"] = o.PageResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InvitationList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]Invitation `json:"items"`
		PageResult *PageResult   `json:"pageResult,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "pageResult"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	if all.PageResult != nil && all.PageResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PageResult = all.PageResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
