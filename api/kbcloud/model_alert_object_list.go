// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// AlertObjectList AlertObjectList is a list of alert object
type AlertObjectList struct {
	// Items is the list of alert objects in the list
	Items []AlertObject `json:"items,omitempty"`
	// PageResult info
	PageResult *PageResult `json:"pageResult,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertObjectList instantiates a new AlertObjectList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertObjectList() *AlertObjectList {
	this := AlertObjectList{}
	return &this
}

// NewAlertObjectListWithDefaults instantiates a new AlertObjectList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertObjectListWithDefaults() *AlertObjectList {
	this := AlertObjectList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AlertObjectList) GetItems() []AlertObject {
	if o == nil || o.Items == nil {
		var ret []AlertObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertObjectList) GetItemsOk() (*[]AlertObject, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AlertObjectList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []AlertObject and assigns it to the Items field.
func (o *AlertObjectList) SetItems(v []AlertObject) {
	o.Items = v
}

// GetPageResult returns the PageResult field value if set, zero value otherwise.
func (o *AlertObjectList) GetPageResult() PageResult {
	if o == nil || o.PageResult == nil {
		var ret PageResult
		return ret
	}
	return *o.PageResult
}

// GetPageResultOk returns a tuple with the PageResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertObjectList) GetPageResultOk() (*PageResult, bool) {
	if o == nil || o.PageResult == nil {
		return nil, false
	}
	return o.PageResult, true
}

// HasPageResult returns a boolean if a field has been set.
func (o *AlertObjectList) HasPageResult() bool {
	return o != nil && o.PageResult != nil
}

// SetPageResult gets a reference to the given PageResult and assigns it to the PageResult field.
func (o *AlertObjectList) SetPageResult(v PageResult) {
	o.PageResult = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertObjectList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.PageResult != nil {
		toSerialize["pageResult"] = o.PageResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertObjectList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      []AlertObject `json:"items,omitempty"`
		PageResult *PageResult   `json:"pageResult,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "pageResult"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = all.Items
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
