// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InitOptionObject struct {
	ConfigSpecName string                 `json:"configSpecName"`
	Label          LocalizedDescription   `json:"label"`
	Type           InitOptionType         `json:"type"`
	Description    *LocalizedDescription  `json:"description,omitempty"`
	Options        []EngineInitOptionItem `json:"options,omitempty"`
	Value          *string                `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInitOptionObject instantiates a new InitOptionObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInitOptionObject(configSpecName string, label LocalizedDescription, typeVar InitOptionType) *InitOptionObject {
	this := InitOptionObject{}
	this.ConfigSpecName = configSpecName
	this.Label = label
	this.Type = typeVar
	return &this
}

// NewInitOptionObjectWithDefaults instantiates a new InitOptionObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInitOptionObjectWithDefaults() *InitOptionObject {
	this := InitOptionObject{}
	return &this
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *InitOptionObject) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *InitOptionObject) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetLabel returns the Label field value.
func (o *InitOptionObject) GetLabel() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetLabelOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *InitOptionObject) SetLabel(v LocalizedDescription) {
	o.Label = v
}

// GetType returns the Type field value.
func (o *InitOptionObject) GetType() InitOptionType {
	if o == nil {
		var ret InitOptionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetTypeOk() (*InitOptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InitOptionObject) SetType(v InitOptionType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InitOptionObject) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InitOptionObject) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *InitOptionObject) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InitOptionObject) GetOptions() []EngineInitOptionItem {
	if o == nil || o.Options == nil {
		var ret []EngineInitOptionItem
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetOptionsOk() (*[]EngineInitOptionItem, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InitOptionObject) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []EngineInitOptionItem and assigns it to the Options field.
func (o *InitOptionObject) SetOptions(v []EngineInitOptionItem) {
	o.Options = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InitOptionObject) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionObject) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InitOptionObject) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InitOptionObject) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InitOptionObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InitOptionObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpecName *string                `json:"configSpecName"`
		Label          *LocalizedDescription  `json:"label"`
		Type           *InitOptionType        `json:"type"`
		Description    *LocalizedDescription  `json:"description,omitempty"`
		Options        []EngineInitOptionItem `json:"options,omitempty"`
		Value          *string                `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configSpecName", "label", "type", "description", "options", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfigSpecName = *all.ConfigSpecName
	if all.Label.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Label = *all.Label
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.Options = all.Options
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
