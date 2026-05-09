// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterTemplateDiffItem A parameter or raw file difference in a parameter template
type ParameterTemplateDiffItem struct {
	// The name of the configuration file
	FileName string `json:"fileName"`
	// The name of the changed parameter. Empty when rawContent is true and the whole file differs.
	ParameterName *string `json:"parameterName,omitempty"`
	// The value from the matching default template
	OldValue *string `json:"oldValue,omitempty"`
	// The value from the custom parameter template
	NewValue *string `json:"newValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplateDiffItem instantiates a new ParameterTemplateDiffItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplateDiffItem(fileName string) *ParameterTemplateDiffItem {
	this := ParameterTemplateDiffItem{}
	this.FileName = fileName
	return &this
}

// NewParameterTemplateDiffItemWithDefaults instantiates a new ParameterTemplateDiffItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateDiffItemWithDefaults() *ParameterTemplateDiffItem {
	this := ParameterTemplateDiffItem{}
	return &this
}

// GetFileName returns the FileName field value.
func (o *ParameterTemplateDiffItem) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffItem) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value.
func (o *ParameterTemplateDiffItem) SetFileName(v string) {
	o.FileName = v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *ParameterTemplateDiffItem) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffItem) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *ParameterTemplateDiffItem) HasParameterName() bool {
	return o != nil && o.ParameterName != nil
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *ParameterTemplateDiffItem) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *ParameterTemplateDiffItem) GetOldValue() string {
	if o == nil || o.OldValue == nil {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffItem) GetOldValueOk() (*string, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *ParameterTemplateDiffItem) HasOldValue() bool {
	return o != nil && o.OldValue != nil
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *ParameterTemplateDiffItem) SetOldValue(v string) {
	o.OldValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *ParameterTemplateDiffItem) GetNewValue() string {
	if o == nil || o.NewValue == nil {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffItem) GetNewValueOk() (*string, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *ParameterTemplateDiffItem) HasNewValue() bool {
	return o != nil && o.NewValue != nil
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *ParameterTemplateDiffItem) SetNewValue(v string) {
	o.NewValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplateDiffItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["fileName"] = o.FileName
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplateDiffItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileName      *string `json:"fileName"`
		ParameterName *string `json:"parameterName,omitempty"`
		OldValue      *string `json:"oldValue,omitempty"`
		NewValue      *string `json:"newValue,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.FileName == nil {
		return fmt.Errorf("required field fileName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"fileName", "parameterName", "oldValue", "newValue"})
	} else {
		return err
	}
	o.FileName = *all.FileName
	o.ParameterName = all.ParameterName
	o.OldValue = all.OldValue
	o.NewValue = all.NewValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
