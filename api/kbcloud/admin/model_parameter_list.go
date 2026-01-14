// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterList A list of cluster parameter
type ParameterList struct {
	// Major version of database engine, eg: 8.0
	MajorVersion string `json:"majorVersion"`
	// Name of database engine
	Engine string `json:"engine"`
	// Name of component
	Component string          `json:"component"`
	Items     []ParameterItem `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterList instantiates a new ParameterList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterList(majorVersion string, engine string, component string) *ParameterList {
	this := ParameterList{}
	this.MajorVersion = majorVersion
	this.Engine = engine
	this.Component = component
	return &this
}

// NewParameterListWithDefaults instantiates a new ParameterList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterListWithDefaults() *ParameterList {
	this := ParameterList{}
	return &this
}

// GetMajorVersion returns the MajorVersion field value.
func (o *ParameterList) GetMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value
// and a boolean to check if the value has been set.
func (o *ParameterList) GetMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersion, true
}

// SetMajorVersion sets field value.
func (o *ParameterList) SetMajorVersion(v string) {
	o.MajorVersion = v
}

// GetEngine returns the Engine field value.
func (o *ParameterList) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ParameterList) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ParameterList) SetEngine(v string) {
	o.Engine = v
}

// GetComponent returns the Component field value.
func (o *ParameterList) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParameterList) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParameterList) SetComponent(v string) {
	o.Component = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ParameterList) GetItems() []ParameterItem {
	if o == nil || o.Items == nil {
		var ret []ParameterItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterList) GetItemsOk() (*[]ParameterItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ParameterList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ParameterItem and assigns it to the Items field.
func (o *ParameterList) SetItems(v []ParameterItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["majorVersion"] = o.MajorVersion
	toSerialize["engine"] = o.Engine
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
func (o *ParameterList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MajorVersion *string         `json:"majorVersion"`
		Engine       *string         `json:"engine"`
		Component    *string         `json:"component"`
		Items        []ParameterItem `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MajorVersion == nil {
		return fmt.Errorf("required field majorVersion missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"majorVersion", "engine", "component", "items"})
	} else {
		return err
	}
	o.MajorVersion = *all.MajorVersion
	o.Engine = *all.Engine
	o.Component = *all.Component
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
