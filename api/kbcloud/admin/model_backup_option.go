// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupOption If present, it must be set defaultMethod and fullMethod
type BackupOption struct {
	DefaultMethod string `json:"defaultMethod"`
	// If set, the default backup method will be applied to the specified component.
	// If not set, the default backup method will be applied to all components.
	//
	DefaultComponent *string `json:"defaultComponent,omitempty"`
	// selector for the default backup policy template. this is necessary when referencing other addon components
	DefaultBptSelector map[string]string          `json:"defaultBPTSelector,omitempty"`
	RestoreOption      *BackupOptionRestoreOption `json:"restoreOption,omitempty"`
	FullMethod         []BackupMethodOption       `json:"fullMethod"`
	IncrementalMethod  []BackupMethodOption       `json:"incrementalMethod,omitempty"`
	ContinuousMethod   []BackupMethodOption       `json:"continuousMethod,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupOption instantiates a new BackupOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupOption(defaultMethod string, fullMethod []BackupMethodOption) *BackupOption {
	this := BackupOption{}
	this.DefaultMethod = defaultMethod
	this.FullMethod = fullMethod
	return &this
}

// NewBackupOptionWithDefaults instantiates a new BackupOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupOptionWithDefaults() *BackupOption {
	this := BackupOption{}
	return &this
}

// GetDefaultMethod returns the DefaultMethod field value.
func (o *BackupOption) GetDefaultMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultMethod
}

// GetDefaultMethodOk returns a tuple with the DefaultMethod field value
// and a boolean to check if the value has been set.
func (o *BackupOption) GetDefaultMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMethod, true
}

// SetDefaultMethod sets field value.
func (o *BackupOption) SetDefaultMethod(v string) {
	o.DefaultMethod = v
}

// GetDefaultComponent returns the DefaultComponent field value if set, zero value otherwise.
func (o *BackupOption) GetDefaultComponent() string {
	if o == nil || o.DefaultComponent == nil {
		var ret string
		return ret
	}
	return *o.DefaultComponent
}

// GetDefaultComponentOk returns a tuple with the DefaultComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetDefaultComponentOk() (*string, bool) {
	if o == nil || o.DefaultComponent == nil {
		return nil, false
	}
	return o.DefaultComponent, true
}

// HasDefaultComponent returns a boolean if a field has been set.
func (o *BackupOption) HasDefaultComponent() bool {
	return o != nil && o.DefaultComponent != nil
}

// SetDefaultComponent gets a reference to the given string and assigns it to the DefaultComponent field.
func (o *BackupOption) SetDefaultComponent(v string) {
	o.DefaultComponent = &v
}

// GetDefaultBptSelector returns the DefaultBptSelector field value if set, zero value otherwise.
func (o *BackupOption) GetDefaultBptSelector() map[string]string {
	if o == nil || o.DefaultBptSelector == nil {
		var ret map[string]string
		return ret
	}
	return o.DefaultBptSelector
}

// GetDefaultBptSelectorOk returns a tuple with the DefaultBptSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetDefaultBptSelectorOk() (*map[string]string, bool) {
	if o == nil || o.DefaultBptSelector == nil {
		return nil, false
	}
	return &o.DefaultBptSelector, true
}

// HasDefaultBptSelector returns a boolean if a field has been set.
func (o *BackupOption) HasDefaultBptSelector() bool {
	return o != nil && o.DefaultBptSelector != nil
}

// SetDefaultBptSelector gets a reference to the given map[string]string and assigns it to the DefaultBptSelector field.
func (o *BackupOption) SetDefaultBptSelector(v map[string]string) {
	o.DefaultBptSelector = v
}

// GetRestoreOption returns the RestoreOption field value if set, zero value otherwise.
func (o *BackupOption) GetRestoreOption() BackupOptionRestoreOption {
	if o == nil || o.RestoreOption == nil {
		var ret BackupOptionRestoreOption
		return ret
	}
	return *o.RestoreOption
}

// GetRestoreOptionOk returns a tuple with the RestoreOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetRestoreOptionOk() (*BackupOptionRestoreOption, bool) {
	if o == nil || o.RestoreOption == nil {
		return nil, false
	}
	return o.RestoreOption, true
}

// HasRestoreOption returns a boolean if a field has been set.
func (o *BackupOption) HasRestoreOption() bool {
	return o != nil && o.RestoreOption != nil
}

// SetRestoreOption gets a reference to the given BackupOptionRestoreOption and assigns it to the RestoreOption field.
func (o *BackupOption) SetRestoreOption(v BackupOptionRestoreOption) {
	o.RestoreOption = &v
}

// GetFullMethod returns the FullMethod field value.
func (o *BackupOption) GetFullMethod() []BackupMethodOption {
	if o == nil {
		var ret []BackupMethodOption
		return ret
	}
	return o.FullMethod
}

// GetFullMethodOk returns a tuple with the FullMethod field value
// and a boolean to check if the value has been set.
func (o *BackupOption) GetFullMethodOk() (*[]BackupMethodOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullMethod, true
}

// SetFullMethod sets field value.
func (o *BackupOption) SetFullMethod(v []BackupMethodOption) {
	o.FullMethod = v
}

// GetIncrementalMethod returns the IncrementalMethod field value if set, zero value otherwise.
func (o *BackupOption) GetIncrementalMethod() []BackupMethodOption {
	if o == nil || o.IncrementalMethod == nil {
		var ret []BackupMethodOption
		return ret
	}
	return o.IncrementalMethod
}

// GetIncrementalMethodOk returns a tuple with the IncrementalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetIncrementalMethodOk() (*[]BackupMethodOption, bool) {
	if o == nil || o.IncrementalMethod == nil {
		return nil, false
	}
	return &o.IncrementalMethod, true
}

// HasIncrementalMethod returns a boolean if a field has been set.
func (o *BackupOption) HasIncrementalMethod() bool {
	return o != nil && o.IncrementalMethod != nil
}

// SetIncrementalMethod gets a reference to the given []BackupMethodOption and assigns it to the IncrementalMethod field.
func (o *BackupOption) SetIncrementalMethod(v []BackupMethodOption) {
	o.IncrementalMethod = v
}

// GetContinuousMethod returns the ContinuousMethod field value if set, zero value otherwise.
func (o *BackupOption) GetContinuousMethod() []BackupMethodOption {
	if o == nil || o.ContinuousMethod == nil {
		var ret []BackupMethodOption
		return ret
	}
	return o.ContinuousMethod
}

// GetContinuousMethodOk returns a tuple with the ContinuousMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOption) GetContinuousMethodOk() (*[]BackupMethodOption, bool) {
	if o == nil || o.ContinuousMethod == nil {
		return nil, false
	}
	return &o.ContinuousMethod, true
}

// HasContinuousMethod returns a boolean if a field has been set.
func (o *BackupOption) HasContinuousMethod() bool {
	return o != nil && o.ContinuousMethod != nil
}

// SetContinuousMethod gets a reference to the given []BackupMethodOption and assigns it to the ContinuousMethod field.
func (o *BackupOption) SetContinuousMethod(v []BackupMethodOption) {
	o.ContinuousMethod = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["defaultMethod"] = o.DefaultMethod
	if o.DefaultComponent != nil {
		toSerialize["defaultComponent"] = o.DefaultComponent
	}
	if o.DefaultBptSelector != nil {
		toSerialize["defaultBPTSelector"] = o.DefaultBptSelector
	}
	if o.RestoreOption != nil {
		toSerialize["restoreOption"] = o.RestoreOption
	}
	toSerialize["fullMethod"] = o.FullMethod
	if o.IncrementalMethod != nil {
		toSerialize["incrementalMethod"] = o.IncrementalMethod
	}
	if o.ContinuousMethod != nil {
		toSerialize["continuousMethod"] = o.ContinuousMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultMethod      *string                    `json:"defaultMethod"`
		DefaultComponent   *string                    `json:"defaultComponent,omitempty"`
		DefaultBptSelector map[string]string          `json:"defaultBPTSelector,omitempty"`
		RestoreOption      *BackupOptionRestoreOption `json:"restoreOption,omitempty"`
		FullMethod         *[]BackupMethodOption      `json:"fullMethod"`
		IncrementalMethod  []BackupMethodOption       `json:"incrementalMethod,omitempty"`
		ContinuousMethod   []BackupMethodOption       `json:"continuousMethod,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DefaultMethod == nil {
		return fmt.Errorf("required field defaultMethod missing")
	}
	if all.FullMethod == nil {
		return fmt.Errorf("required field fullMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"defaultMethod", "defaultComponent", "defaultBPTSelector", "restoreOption", "fullMethod", "incrementalMethod", "continuousMethod"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefaultMethod = *all.DefaultMethod
	o.DefaultComponent = all.DefaultComponent
	o.DefaultBptSelector = all.DefaultBptSelector
	if all.RestoreOption != nil && all.RestoreOption.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RestoreOption = all.RestoreOption
	o.FullMethod = *all.FullMethod
	o.IncrementalMethod = all.IncrementalMethod
	o.ContinuousMethod = all.ContinuousMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
