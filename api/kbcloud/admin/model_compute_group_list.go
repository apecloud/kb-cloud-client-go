// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComputeGroupList struct {
	Items               []ComputeGroup `json:"items"`
	ManagementSupported *bool          `json:"managementSupported,omitempty"`
	UnsupportedReason   *string        `json:"unsupportedReason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComputeGroupList instantiates a new ComputeGroupList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupList(items []ComputeGroup) *ComputeGroupList {
	this := ComputeGroupList{}
	this.Items = items
	return &this
}

// NewComputeGroupListWithDefaults instantiates a new ComputeGroupList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupListWithDefaults() *ComputeGroupList {
	this := ComputeGroupList{}
	return &this
}

// GetItems returns the Items field value.
func (o *ComputeGroupList) GetItems() []ComputeGroup {
	if o == nil {
		var ret []ComputeGroup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ComputeGroupList) GetItemsOk() (*[]ComputeGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ComputeGroupList) SetItems(v []ComputeGroup) {
	o.Items = v
}

// GetManagementSupported returns the ManagementSupported field value if set, zero value otherwise.
func (o *ComputeGroupList) GetManagementSupported() bool {
	if o == nil || o.ManagementSupported == nil {
		var ret bool
		return ret
	}
	return *o.ManagementSupported
}

// GetManagementSupportedOk returns a tuple with the ManagementSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupList) GetManagementSupportedOk() (*bool, bool) {
	if o == nil || o.ManagementSupported == nil {
		return nil, false
	}
	return o.ManagementSupported, true
}

// HasManagementSupported returns a boolean if a field has been set.
func (o *ComputeGroupList) HasManagementSupported() bool {
	return o != nil && o.ManagementSupported != nil
}

// SetManagementSupported gets a reference to the given bool and assigns it to the ManagementSupported field.
func (o *ComputeGroupList) SetManagementSupported(v bool) {
	o.ManagementSupported = &v
}

// GetUnsupportedReason returns the UnsupportedReason field value if set, zero value otherwise.
func (o *ComputeGroupList) GetUnsupportedReason() string {
	if o == nil || o.UnsupportedReason == nil {
		var ret string
		return ret
	}
	return *o.UnsupportedReason
}

// GetUnsupportedReasonOk returns a tuple with the UnsupportedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupList) GetUnsupportedReasonOk() (*string, bool) {
	if o == nil || o.UnsupportedReason == nil {
		return nil, false
	}
	return o.UnsupportedReason, true
}

// HasUnsupportedReason returns a boolean if a field has been set.
func (o *ComputeGroupList) HasUnsupportedReason() bool {
	return o != nil && o.UnsupportedReason != nil
}

// SetUnsupportedReason gets a reference to the given string and assigns it to the UnsupportedReason field.
func (o *ComputeGroupList) SetUnsupportedReason(v string) {
	o.UnsupportedReason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.ManagementSupported != nil {
		toSerialize["managementSupported"] = o.ManagementSupported
	}
	if o.UnsupportedReason != nil {
		toSerialize["unsupportedReason"] = o.UnsupportedReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items               *[]ComputeGroup `json:"items"`
		ManagementSupported *bool           `json:"managementSupported,omitempty"`
		UnsupportedReason   *string         `json:"unsupportedReason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "managementSupported", "unsupportedReason"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.ManagementSupported = all.ManagementSupported
	o.UnsupportedReason = all.UnsupportedReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
