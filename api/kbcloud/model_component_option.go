// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComponentOption struct {
	// component type
	Name string `json:"name"`
	// Determine whether the componentDef of kb-cluster belongs to this component type through this matching regularization.
	// if not set, componentDef must be equal to component type.
	//
	MatchRegex *string                 `json:"matchRegex,omitempty"`
	Title      LocalizedDescription    `json:"title"`
	Order      int32                   `json:"order"`
	RoleOrder  []string                `json:"roleOrder,omitempty"`
	Version    *ComponentOptionVersion `json:"version,omitempty"`
	// Main component flag
	Main *bool `json:"main,omitempty"`
	// whether the component supports custom secret
	CustomSecret *bool `json:"customSecret,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOption instantiates a new ComponentOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOption(name string, title LocalizedDescription, order int32) *ComponentOption {
	this := ComponentOption{}
	this.Name = name
	this.Title = title
	this.Order = order
	return &this
}

// NewComponentOptionWithDefaults instantiates a new ComponentOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionWithDefaults() *ComponentOption {
	this := ComponentOption{}
	return &this
}

// GetName returns the Name field value.
func (o *ComponentOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ComponentOption) SetName(v string) {
	o.Name = v
}

// GetMatchRegex returns the MatchRegex field value if set, zero value otherwise.
func (o *ComponentOption) GetMatchRegex() string {
	if o == nil || o.MatchRegex == nil {
		var ret string
		return ret
	}
	return *o.MatchRegex
}

// GetMatchRegexOk returns a tuple with the MatchRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetMatchRegexOk() (*string, bool) {
	if o == nil || o.MatchRegex == nil {
		return nil, false
	}
	return o.MatchRegex, true
}

// HasMatchRegex returns a boolean if a field has been set.
func (o *ComponentOption) HasMatchRegex() bool {
	return o != nil && o.MatchRegex != nil
}

// SetMatchRegex gets a reference to the given string and assigns it to the MatchRegex field.
func (o *ComponentOption) SetMatchRegex(v string) {
	o.MatchRegex = &v
}

// GetTitle returns the Title field value.
func (o *ComponentOption) GetTitle() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ComponentOption) SetTitle(v LocalizedDescription) {
	o.Title = v
}

// GetOrder returns the Order field value.
func (o *ComponentOption) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *ComponentOption) SetOrder(v int32) {
	o.Order = v
}

// GetRoleOrder returns the RoleOrder field value if set, zero value otherwise.
func (o *ComponentOption) GetRoleOrder() []string {
	if o == nil || o.RoleOrder == nil {
		var ret []string
		return ret
	}
	return o.RoleOrder
}

// GetRoleOrderOk returns a tuple with the RoleOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetRoleOrderOk() (*[]string, bool) {
	if o == nil || o.RoleOrder == nil {
		return nil, false
	}
	return &o.RoleOrder, true
}

// HasRoleOrder returns a boolean if a field has been set.
func (o *ComponentOption) HasRoleOrder() bool {
	return o != nil && o.RoleOrder != nil
}

// SetRoleOrder gets a reference to the given []string and assigns it to the RoleOrder field.
func (o *ComponentOption) SetRoleOrder(v []string) {
	o.RoleOrder = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComponentOption) GetVersion() ComponentOptionVersion {
	if o == nil || o.Version == nil {
		var ret ComponentOptionVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetVersionOk() (*ComponentOptionVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComponentOption) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given ComponentOptionVersion and assigns it to the Version field.
func (o *ComponentOption) SetVersion(v ComponentOptionVersion) {
	o.Version = &v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *ComponentOption) GetMain() bool {
	if o == nil || o.Main == nil {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetMainOk() (*bool, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *ComponentOption) HasMain() bool {
	return o != nil && o.Main != nil
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *ComponentOption) SetMain(v bool) {
	o.Main = &v
}

// GetCustomSecret returns the CustomSecret field value if set, zero value otherwise.
func (o *ComponentOption) GetCustomSecret() bool {
	if o == nil || o.CustomSecret == nil {
		var ret bool
		return ret
	}
	return *o.CustomSecret
}

// GetCustomSecretOk returns a tuple with the CustomSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOption) GetCustomSecretOk() (*bool, bool) {
	if o == nil || o.CustomSecret == nil {
		return nil, false
	}
	return o.CustomSecret, true
}

// HasCustomSecret returns a boolean if a field has been set.
func (o *ComponentOption) HasCustomSecret() bool {
	return o != nil && o.CustomSecret != nil
}

// SetCustomSecret gets a reference to the given bool and assigns it to the CustomSecret field.
func (o *ComponentOption) SetCustomSecret(v bool) {
	o.CustomSecret = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.MatchRegex != nil {
		toSerialize["matchRegex"] = o.MatchRegex
	}
	toSerialize["title"] = o.Title
	toSerialize["order"] = o.Order
	if o.RoleOrder != nil {
		toSerialize["roleOrder"] = o.RoleOrder
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Main != nil {
		toSerialize["main"] = o.Main
	}
	if o.CustomSecret != nil {
		toSerialize["customSecret"] = o.CustomSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                 `json:"name"`
		MatchRegex   *string                 `json:"matchRegex,omitempty"`
		Title        *LocalizedDescription   `json:"title"`
		Order        *int32                  `json:"order"`
		RoleOrder    []string                `json:"roleOrder,omitempty"`
		Version      *ComponentOptionVersion `json:"version,omitempty"`
		Main         *bool                   `json:"main,omitempty"`
		CustomSecret *bool                   `json:"customSecret,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "matchRegex", "title", "order", "roleOrder", "version", "main", "customSecret"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.MatchRegex = all.MatchRegex
	if all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = *all.Title
	o.Order = *all.Order
	o.RoleOrder = all.RoleOrder
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version
	o.Main = all.Main
	o.CustomSecret = all.CustomSecret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
