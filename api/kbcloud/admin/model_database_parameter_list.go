// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseParameterList A list of database parameters
type DatabaseParameterList struct {
	// Name of database engine
	EngineName string `json:"engineName"`
	// Name of component
	Component string                      `json:"component"`
	Items     []DatabaseParameterListItem `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseParameterList instantiates a new DatabaseParameterList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseParameterList(engineName string, component string) *DatabaseParameterList {
	this := DatabaseParameterList{}
	this.EngineName = engineName
	this.Component = component
	return &this
}

// NewDatabaseParameterListWithDefaults instantiates a new DatabaseParameterList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseParameterListWithDefaults() *DatabaseParameterList {
	this := DatabaseParameterList{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *DatabaseParameterList) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterList) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *DatabaseParameterList) SetEngineName(v string) {
	o.EngineName = v
}

// GetComponent returns the Component field value.
func (o *DatabaseParameterList) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterList) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *DatabaseParameterList) SetComponent(v string) {
	o.Component = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DatabaseParameterList) GetItems() []DatabaseParameterListItem {
	if o == nil || o.Items == nil {
		var ret []DatabaseParameterListItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterList) GetItemsOk() (*[]DatabaseParameterListItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DatabaseParameterList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []DatabaseParameterListItem and assigns it to the Items field.
func (o *DatabaseParameterList) SetItems(v []DatabaseParameterListItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseParameterList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["component"] = o.Component
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseParameterList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string                     `json:"engineName"`
		Component  *string                     `json:"component"`
		Items      []DatabaseParameterListItem `json:"items,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "component", "items"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Component = *all.Component
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
