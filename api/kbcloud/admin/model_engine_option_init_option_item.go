// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOption_initOptionItem struct {
	Value string                `json:"value"`
	Label *LocalizedDescription `json:"label,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOption_initOptionItem instantiates a new EngineOption_initOptionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOption_initOptionItem(value string) *EngineOption_initOptionItem {
	this := EngineOption_initOptionItem{}
	this.Value = value
	return &this
}

// NewEngineOption_initOptionItemWithDefaults instantiates a new EngineOption_initOptionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOption_initOptionItemWithDefaults() *EngineOption_initOptionItem {
	this := EngineOption_initOptionItem{}
	return &this
}

// GetValue returns the Value field value.
func (o *EngineOption_initOptionItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EngineOption_initOptionItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *EngineOption_initOptionItem) SetValue(v string) {
	o.Value = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EngineOption_initOptionItem) GetLabel() LocalizedDescription {
	if o == nil || o.Label == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOption_initOptionItem) GetLabelOk() (*LocalizedDescription, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EngineOption_initOptionItem) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given LocalizedDescription and assigns it to the Label field.
func (o *EngineOption_initOptionItem) SetLabel(v LocalizedDescription) {
	o.Label = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOption_initOptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["value"] = o.Value
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOption_initOptionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value *string               `json:"value"`
		Label *LocalizedDescription `json:"label,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "label"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Value = *all.Value
	if all.Label != nil && all.Label.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Label = all.Label

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
