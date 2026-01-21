// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DeleteDatabaseParameterItem A database parameter item for deletion
type DeleteDatabaseParameterItem struct {
	// Name of database engine
	EngineName string `json:"engineName"`
	// Name of component
	Component string `json:"component"`
	// The name of the parameter
	Name string `json:"name"`
	// Since which version the parameter is supported
	Since *string `json:"since,omitempty"`
	// which version the parameter is deprecated
	Deprecated *string `json:"deprecated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteDatabaseParameterItem instantiates a new DeleteDatabaseParameterItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteDatabaseParameterItem(engineName string, component string, name string) *DeleteDatabaseParameterItem {
	this := DeleteDatabaseParameterItem{}
	this.EngineName = engineName
	this.Component = component
	this.Name = name
	return &this
}

// NewDeleteDatabaseParameterItemWithDefaults instantiates a new DeleteDatabaseParameterItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteDatabaseParameterItemWithDefaults() *DeleteDatabaseParameterItem {
	this := DeleteDatabaseParameterItem{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *DeleteDatabaseParameterItem) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *DeleteDatabaseParameterItem) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *DeleteDatabaseParameterItem) SetEngineName(v string) {
	o.EngineName = v
}

// GetComponent returns the Component field value.
func (o *DeleteDatabaseParameterItem) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DeleteDatabaseParameterItem) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *DeleteDatabaseParameterItem) SetComponent(v string) {
	o.Component = v
}

// GetName returns the Name field value.
func (o *DeleteDatabaseParameterItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteDatabaseParameterItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DeleteDatabaseParameterItem) SetName(v string) {
	o.Name = v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *DeleteDatabaseParameterItem) GetSince() string {
	if o == nil || o.Since == nil {
		var ret string
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDatabaseParameterItem) GetSinceOk() (*string, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *DeleteDatabaseParameterItem) HasSince() bool {
	return o != nil && o.Since != nil
}

// SetSince gets a reference to the given string and assigns it to the Since field.
func (o *DeleteDatabaseParameterItem) SetSince(v string) {
	o.Since = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *DeleteDatabaseParameterItem) GetDeprecated() string {
	if o == nil || o.Deprecated == nil {
		var ret string
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDatabaseParameterItem) GetDeprecatedOk() (*string, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *DeleteDatabaseParameterItem) HasDeprecated() bool {
	return o != nil && o.Deprecated != nil
}

// SetDeprecated gets a reference to the given string and assigns it to the Deprecated field.
func (o *DeleteDatabaseParameterItem) SetDeprecated(v string) {
	o.Deprecated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteDatabaseParameterItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteDatabaseParameterItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string `json:"engineName"`
		Component  *string `json:"component"`
		Name       *string `json:"name"`
		Since      *string `json:"since,omitempty"`
		Deprecated *string `json:"deprecated,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "component", "name", "since", "deprecated"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Component = *all.Component
	o.Name = *all.Name
	o.Since = all.Since
	o.Deprecated = all.Deprecated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
