// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// IpPoolProviderList Pod IP pools and selection policy reported by one provider.
type IpPoolProviderList struct {
	// KBE provider used for Pod IP pool discovery and explicit selection.
	Name IpPoolProvider `json:"name"`
	// Per-address-family requirements for explicit IP pool selection.
	SelectionPolicy IpPoolSelectionPolicy `json:"selectionPolicy"`
	Items           []IpPool              `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIpPoolProviderList instantiates a new IpPoolProviderList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIpPoolProviderList(name IpPoolProvider, selectionPolicy IpPoolSelectionPolicy, items []IpPool) *IpPoolProviderList {
	this := IpPoolProviderList{}
	this.Name = name
	this.SelectionPolicy = selectionPolicy
	this.Items = items
	return &this
}

// NewIpPoolProviderListWithDefaults instantiates a new IpPoolProviderList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIpPoolProviderListWithDefaults() *IpPoolProviderList {
	this := IpPoolProviderList{}
	return &this
}

// GetName returns the Name field value.
func (o *IpPoolProviderList) GetName() IpPoolProvider {
	if o == nil {
		var ret IpPoolProvider
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpPoolProviderList) GetNameOk() (*IpPoolProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IpPoolProviderList) SetName(v IpPoolProvider) {
	o.Name = v
}

// GetSelectionPolicy returns the SelectionPolicy field value.
func (o *IpPoolProviderList) GetSelectionPolicy() IpPoolSelectionPolicy {
	if o == nil {
		var ret IpPoolSelectionPolicy
		return ret
	}
	return o.SelectionPolicy
}

// GetSelectionPolicyOk returns a tuple with the SelectionPolicy field value
// and a boolean to check if the value has been set.
func (o *IpPoolProviderList) GetSelectionPolicyOk() (*IpPoolSelectionPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectionPolicy, true
}

// SetSelectionPolicy sets field value.
func (o *IpPoolProviderList) SetSelectionPolicy(v IpPoolSelectionPolicy) {
	o.SelectionPolicy = v
}

// GetItems returns the Items field value.
func (o *IpPoolProviderList) GetItems() []IpPool {
	if o == nil {
		var ret []IpPool
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *IpPoolProviderList) GetItemsOk() (*[]IpPool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *IpPoolProviderList) SetItems(v []IpPool) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IpPoolProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["selectionPolicy"] = o.SelectionPolicy
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IpPoolProviderList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *IpPoolProvider        `json:"name"`
		SelectionPolicy *IpPoolSelectionPolicy `json:"selectionPolicy"`
		Items           *[]IpPool              `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SelectionPolicy == nil {
		return fmt.Errorf("required field selectionPolicy missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "selectionPolicy", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}
	if all.SelectionPolicy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectionPolicy = *all.SelectionPolicy
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
