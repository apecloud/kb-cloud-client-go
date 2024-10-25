// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ProviderList The list of cloud providers.

type ProviderList struct {
	// The list of cloud providers.
	Items []Provider `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProviderList instantiates a new ProviderList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProviderList() *ProviderList {
	this := ProviderList{}
	return &this
}

// NewProviderListWithDefaults instantiates a new ProviderList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProviderListWithDefaults() *ProviderList {
	this := ProviderList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ProviderList) GetItems() []Provider {
	if o == nil || o.Items == nil {
		var ret []Provider
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderList) GetItemsOk() (*[]Provider, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ProviderList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []Provider and assigns it to the Items field.
func (o *ProviderList) SetItems(v []Provider) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProviderList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items []Provider `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
