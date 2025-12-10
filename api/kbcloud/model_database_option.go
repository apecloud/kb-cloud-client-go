// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DatabaseOption struct {
	Enabled bool  `json:"enabled"`
	Update  *bool `json:"update,omitempty"`
	// max length of database name.
	// If not set, use default value.
	//
	MaxLen *int32 `json:"maxLen,omitempty"`
	// min length of database name.
	// If not set, use default value.
	//
	MinLen              *int32  `json:"minLen,omitempty"`
	DatabaseNamePattern *string `json:"databaseNamePattern,omitempty"`
	// The database option information that will be displayed when listing databases
	//
	ListOption []string `json:"listOption,omitempty"`
	// The database option cloud be set when creating databases
	//
	AvailableOptions []string `json:"availableOptions,omitempty"`
	// The database option name and type cloud be updated when updating databases
	//
	AvailbaleUpdateOptions []DatabaseOptionAvailbaleUpdateOptionsItem `json:"availbaleUpdateOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseOption instantiates a new DatabaseOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseOption(enabled bool) *DatabaseOption {
	this := DatabaseOption{}
	this.Enabled = enabled
	var maxLen int32 = 64
	this.MaxLen = &maxLen
	var minLen int32 = 3
	this.MinLen = &minLen
	return &this
}

// NewDatabaseOptionWithDefaults instantiates a new DatabaseOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseOptionWithDefaults() *DatabaseOption {
	this := DatabaseOption{}
	var maxLen int32 = 64
	this.MaxLen = &maxLen
	var minLen int32 = 3
	this.MinLen = &minLen
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *DatabaseOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *DatabaseOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DatabaseOption) GetUpdate() bool {
	if o == nil || o.Update == nil {
		var ret bool
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetUpdateOk() (*bool, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DatabaseOption) HasUpdate() bool {
	return o != nil && o.Update != nil
}

// SetUpdate gets a reference to the given bool and assigns it to the Update field.
func (o *DatabaseOption) SetUpdate(v bool) {
	o.Update = &v
}

// GetMaxLen returns the MaxLen field value if set, zero value otherwise.
func (o *DatabaseOption) GetMaxLen() int32 {
	if o == nil || o.MaxLen == nil {
		var ret int32
		return ret
	}
	return *o.MaxLen
}

// GetMaxLenOk returns a tuple with the MaxLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetMaxLenOk() (*int32, bool) {
	if o == nil || o.MaxLen == nil {
		return nil, false
	}
	return o.MaxLen, true
}

// HasMaxLen returns a boolean if a field has been set.
func (o *DatabaseOption) HasMaxLen() bool {
	return o != nil && o.MaxLen != nil
}

// SetMaxLen gets a reference to the given int32 and assigns it to the MaxLen field.
func (o *DatabaseOption) SetMaxLen(v int32) {
	o.MaxLen = &v
}

// GetMinLen returns the MinLen field value if set, zero value otherwise.
func (o *DatabaseOption) GetMinLen() int32 {
	if o == nil || o.MinLen == nil {
		var ret int32
		return ret
	}
	return *o.MinLen
}

// GetMinLenOk returns a tuple with the MinLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetMinLenOk() (*int32, bool) {
	if o == nil || o.MinLen == nil {
		return nil, false
	}
	return o.MinLen, true
}

// HasMinLen returns a boolean if a field has been set.
func (o *DatabaseOption) HasMinLen() bool {
	return o != nil && o.MinLen != nil
}

// SetMinLen gets a reference to the given int32 and assigns it to the MinLen field.
func (o *DatabaseOption) SetMinLen(v int32) {
	o.MinLen = &v
}

// GetDatabaseNamePattern returns the DatabaseNamePattern field value if set, zero value otherwise.
func (o *DatabaseOption) GetDatabaseNamePattern() string {
	if o == nil || o.DatabaseNamePattern == nil {
		var ret string
		return ret
	}
	return *o.DatabaseNamePattern
}

// GetDatabaseNamePatternOk returns a tuple with the DatabaseNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetDatabaseNamePatternOk() (*string, bool) {
	if o == nil || o.DatabaseNamePattern == nil {
		return nil, false
	}
	return o.DatabaseNamePattern, true
}

// HasDatabaseNamePattern returns a boolean if a field has been set.
func (o *DatabaseOption) HasDatabaseNamePattern() bool {
	return o != nil && o.DatabaseNamePattern != nil
}

// SetDatabaseNamePattern gets a reference to the given string and assigns it to the DatabaseNamePattern field.
func (o *DatabaseOption) SetDatabaseNamePattern(v string) {
	o.DatabaseNamePattern = &v
}

// GetListOption returns the ListOption field value if set, zero value otherwise.
func (o *DatabaseOption) GetListOption() []string {
	if o == nil || o.ListOption == nil {
		var ret []string
		return ret
	}
	return o.ListOption
}

// GetListOptionOk returns a tuple with the ListOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetListOptionOk() (*[]string, bool) {
	if o == nil || o.ListOption == nil {
		return nil, false
	}
	return &o.ListOption, true
}

// HasListOption returns a boolean if a field has been set.
func (o *DatabaseOption) HasListOption() bool {
	return o != nil && o.ListOption != nil
}

// SetListOption gets a reference to the given []string and assigns it to the ListOption field.
func (o *DatabaseOption) SetListOption(v []string) {
	o.ListOption = v
}

// GetAvailableOptions returns the AvailableOptions field value if set, zero value otherwise.
func (o *DatabaseOption) GetAvailableOptions() []string {
	if o == nil || o.AvailableOptions == nil {
		var ret []string
		return ret
	}
	return o.AvailableOptions
}

// GetAvailableOptionsOk returns a tuple with the AvailableOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetAvailableOptionsOk() (*[]string, bool) {
	if o == nil || o.AvailableOptions == nil {
		return nil, false
	}
	return &o.AvailableOptions, true
}

// HasAvailableOptions returns a boolean if a field has been set.
func (o *DatabaseOption) HasAvailableOptions() bool {
	return o != nil && o.AvailableOptions != nil
}

// SetAvailableOptions gets a reference to the given []string and assigns it to the AvailableOptions field.
func (o *DatabaseOption) SetAvailableOptions(v []string) {
	o.AvailableOptions = v
}

// GetAvailbaleUpdateOptions returns the AvailbaleUpdateOptions field value if set, zero value otherwise.
func (o *DatabaseOption) GetAvailbaleUpdateOptions() []DatabaseOptionAvailbaleUpdateOptionsItem {
	if o == nil || o.AvailbaleUpdateOptions == nil {
		var ret []DatabaseOptionAvailbaleUpdateOptionsItem
		return ret
	}
	return o.AvailbaleUpdateOptions
}

// GetAvailbaleUpdateOptionsOk returns a tuple with the AvailbaleUpdateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetAvailbaleUpdateOptionsOk() (*[]DatabaseOptionAvailbaleUpdateOptionsItem, bool) {
	if o == nil || o.AvailbaleUpdateOptions == nil {
		return nil, false
	}
	return &o.AvailbaleUpdateOptions, true
}

// HasAvailbaleUpdateOptions returns a boolean if a field has been set.
func (o *DatabaseOption) HasAvailbaleUpdateOptions() bool {
	return o != nil && o.AvailbaleUpdateOptions != nil
}

// SetAvailbaleUpdateOptions gets a reference to the given []DatabaseOptionAvailbaleUpdateOptionsItem and assigns it to the AvailbaleUpdateOptions field.
func (o *DatabaseOption) SetAvailbaleUpdateOptions(v []DatabaseOptionAvailbaleUpdateOptionsItem) {
	o.AvailbaleUpdateOptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	if o.MaxLen != nil {
		toSerialize["maxLen"] = o.MaxLen
	}
	if o.MinLen != nil {
		toSerialize["minLen"] = o.MinLen
	}
	if o.DatabaseNamePattern != nil {
		toSerialize["databaseNamePattern"] = o.DatabaseNamePattern
	}
	if o.ListOption != nil {
		toSerialize["listOption"] = o.ListOption
	}
	if o.AvailableOptions != nil {
		toSerialize["availableOptions"] = o.AvailableOptions
	}
	if o.AvailbaleUpdateOptions != nil {
		toSerialize["availbaleUpdateOptions"] = o.AvailbaleUpdateOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled                *bool                                      `json:"enabled"`
		Update                 *bool                                      `json:"update,omitempty"`
		MaxLen                 *int32                                     `json:"maxLen,omitempty"`
		MinLen                 *int32                                     `json:"minLen,omitempty"`
		DatabaseNamePattern    *string                                    `json:"databaseNamePattern,omitempty"`
		ListOption             []string                                   `json:"listOption,omitempty"`
		AvailableOptions       []string                                   `json:"availableOptions,omitempty"`
		AvailbaleUpdateOptions []DatabaseOptionAvailbaleUpdateOptionsItem `json:"availbaleUpdateOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "update", "maxLen", "minLen", "databaseNamePattern", "listOption", "availableOptions", "availbaleUpdateOptions"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Update = all.Update
	o.MaxLen = all.MaxLen
	o.MinLen = all.MinLen
	o.DatabaseNamePattern = all.DatabaseNamePattern
	o.ListOption = all.ListOption
	o.AvailableOptions = all.AvailableOptions
	o.AvailbaleUpdateOptions = all.AvailbaleUpdateOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
