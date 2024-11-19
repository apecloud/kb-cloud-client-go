// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type InitOptionsItem struct {
	// component type, refer to componentDef and support NamePrefix'
	Component *string `json:"component,omitempty"`
	// component type, refer to componentDef and support NamePrefix, Deprecated
	ComponentDefRef *string           `json:"componentDefRef,omitempty"`
	InitParams      map[string]string `json:"initParams,omitempty"`
	// config spec name
	SpecName *string `json:"specName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInitOptionsItem instantiates a new InitOptionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInitOptionsItem() *InitOptionsItem {
	this := InitOptionsItem{}
	return &this
}

// NewInitOptionsItemWithDefaults instantiates a new InitOptionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInitOptionsItemWithDefaults() *InitOptionsItem {
	this := InitOptionsItem{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *InitOptionsItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionsItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *InitOptionsItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *InitOptionsItem) SetComponent(v string) {
	o.Component = &v
}

// GetComponentDefRef returns the ComponentDefRef field value if set, zero value otherwise.
func (o *InitOptionsItem) GetComponentDefRef() string {
	if o == nil || o.ComponentDefRef == nil {
		var ret string
		return ret
	}
	return *o.ComponentDefRef
}

// GetComponentDefRefOk returns a tuple with the ComponentDefRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionsItem) GetComponentDefRefOk() (*string, bool) {
	if o == nil || o.ComponentDefRef == nil {
		return nil, false
	}
	return o.ComponentDefRef, true
}

// HasComponentDefRef returns a boolean if a field has been set.
func (o *InitOptionsItem) HasComponentDefRef() bool {
	return o != nil && o.ComponentDefRef != nil
}

// SetComponentDefRef gets a reference to the given string and assigns it to the ComponentDefRef field.
func (o *InitOptionsItem) SetComponentDefRef(v string) {
	o.ComponentDefRef = &v
}

// GetInitParams returns the InitParams field value if set, zero value otherwise.
func (o *InitOptionsItem) GetInitParams() map[string]string {
	if o == nil || o.InitParams == nil {
		var ret map[string]string
		return ret
	}
	return o.InitParams
}

// GetInitParamsOk returns a tuple with the InitParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionsItem) GetInitParamsOk() (*map[string]string, bool) {
	if o == nil || o.InitParams == nil {
		return nil, false
	}
	return &o.InitParams, true
}

// HasInitParams returns a boolean if a field has been set.
func (o *InitOptionsItem) HasInitParams() bool {
	return o != nil && o.InitParams != nil
}

// SetInitParams gets a reference to the given map[string]string and assigns it to the InitParams field.
func (o *InitOptionsItem) SetInitParams(v map[string]string) {
	o.InitParams = v
}

// GetSpecName returns the SpecName field value if set, zero value otherwise.
func (o *InitOptionsItem) GetSpecName() string {
	if o == nil || o.SpecName == nil {
		var ret string
		return ret
	}
	return *o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionsItem) GetSpecNameOk() (*string, bool) {
	if o == nil || o.SpecName == nil {
		return nil, false
	}
	return o.SpecName, true
}

// HasSpecName returns a boolean if a field has been set.
func (o *InitOptionsItem) HasSpecName() bool {
	return o != nil && o.SpecName != nil
}

// SetSpecName gets a reference to the given string and assigns it to the SpecName field.
func (o *InitOptionsItem) SetSpecName(v string) {
	o.SpecName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InitOptionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ComponentDefRef != nil {
		toSerialize["componentDefRef"] = o.ComponentDefRef
	}
	if o.InitParams != nil {
		toSerialize["initParams"] = o.InitParams
	}
	if o.SpecName != nil {
		toSerialize["specName"] = o.SpecName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InitOptionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component       *string           `json:"component,omitempty"`
		ComponentDefRef *string           `json:"componentDefRef,omitempty"`
		InitParams      map[string]string `json:"initParams,omitempty"`
		SpecName        *string           `json:"specName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "componentDefRef", "initParams", "specName"})
	} else {
		return err
	}
	o.Component = all.Component
	o.ComponentDefRef = all.ComponentDefRef
	o.InitParams = all.InitParams
	o.SpecName = all.SpecName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
