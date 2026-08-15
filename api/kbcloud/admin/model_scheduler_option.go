// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SchedulerOption struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSchedulerOption instantiates a new SchedulerOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSchedulerOption() *SchedulerOption {
	this := SchedulerOption{}
	return &this
}

// NewSchedulerOptionWithDefaults instantiates a new SchedulerOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSchedulerOptionWithDefaults() *SchedulerOption {
	this := SchedulerOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchedulerOption) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerOption) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchedulerOption) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchedulerOption) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchedulerOption) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerOption) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchedulerOption) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchedulerOption) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *SchedulerOption) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerOption) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *SchedulerOption) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *SchedulerOption) SetDefault(v bool) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SchedulerOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SchedulerOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string `json:"name,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		Default     *bool   `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "displayName", "default"})
	} else {
		return err
	}
	o.Name = all.Name
	o.DisplayName = all.DisplayName
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
