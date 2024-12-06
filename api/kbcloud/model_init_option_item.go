// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

// InitOptionItem InitOptionItem is the information of init option
type InitOptionItem struct {
	// component type, refer to componentDef and support NamePrefix'
	Component  *string           `json:"component,omitempty"`
	InitParams map[string]string `json:"initParams,omitempty"`
	// config spec name
	SpecName *string `json:"specName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInitOptionItem instantiates a new InitOptionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInitOptionItem() *InitOptionItem {
	this := InitOptionItem{}
	return &this
}

// NewInitOptionItemWithDefaults instantiates a new InitOptionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInitOptionItemWithDefaults() *InitOptionItem {
	this := InitOptionItem{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *InitOptionItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *InitOptionItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *InitOptionItem) SetComponent(v string) {
	o.Component = &v
}

// GetInitParams returns the InitParams field value if set, zero value otherwise.
func (o *InitOptionItem) GetInitParams() map[string]string {
	if o == nil || o.InitParams == nil {
		var ret map[string]string
		return ret
	}
	return o.InitParams
}

// GetInitParamsOk returns a tuple with the InitParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionItem) GetInitParamsOk() (*map[string]string, bool) {
	if o == nil || o.InitParams == nil {
		return nil, false
	}
	return &o.InitParams, true
}

// HasInitParams returns a boolean if a field has been set.
func (o *InitOptionItem) HasInitParams() bool {
	return o != nil && o.InitParams != nil
}

// SetInitParams gets a reference to the given map[string]string and assigns it to the InitParams field.
func (o *InitOptionItem) SetInitParams(v map[string]string) {
	o.InitParams = v
}

// GetSpecName returns the SpecName field value if set, zero value otherwise.
func (o *InitOptionItem) GetSpecName() string {
	if o == nil || o.SpecName == nil {
		var ret string
		return ret
	}
	return *o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitOptionItem) GetSpecNameOk() (*string, bool) {
	if o == nil || o.SpecName == nil {
		return nil, false
	}
	return o.SpecName, true
}

// HasSpecName returns a boolean if a field has been set.
func (o *InitOptionItem) HasSpecName() bool {
	return o != nil && o.SpecName != nil
}

// SetSpecName gets a reference to the given string and assigns it to the SpecName field.
func (o *InitOptionItem) SetSpecName(v string) {
	o.SpecName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InitOptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
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
func (o *InitOptionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component  *string           `json:"component,omitempty"`
		InitParams map[string]string `json:"initParams,omitempty"`
		SpecName   *string           `json:"specName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "initParams", "specName"})
	} else {
		return err
	}
	o.Component = all.Component
	o.InitParams = all.InitParams
	o.SpecName = all.SpecName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
