// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

// ParamTplsItem the item of the parameter template
type ParamTplsItem struct {
	// component type, refer to componentDef and support NamePrefix, Deprecated.
	ComponentDefRef *string `json:"componentDefRef,omitempty"`
	// component type, refer to componentDef and support NamePrefix
	Component *string `json:"component,omitempty"`
	// name of assigned parameter template
	ParamTplName      *string            `json:"paramTplName,omitempty"`
	ParamTplPartition *ParamTplPartition `json:"paramTplPartition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplsItem instantiates a new ParamTplsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplsItem() *ParamTplsItem {
	this := ParamTplsItem{}
	return &this
}

// NewParamTplsItemWithDefaults instantiates a new ParamTplsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplsItemWithDefaults() *ParamTplsItem {
	this := ParamTplsItem{}
	return &this
}

// GetComponentDefRef returns the ComponentDefRef field value if set, zero value otherwise.
func (o *ParamTplsItem) GetComponentDefRef() string {
	if o == nil || o.ComponentDefRef == nil {
		var ret string
		return ret
	}
	return *o.ComponentDefRef
}

// GetComponentDefRefOk returns a tuple with the ComponentDefRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetComponentDefRefOk() (*string, bool) {
	if o == nil || o.ComponentDefRef == nil {
		return nil, false
	}
	return o.ComponentDefRef, true
}

// HasComponentDefRef returns a boolean if a field has been set.
func (o *ParamTplsItem) HasComponentDefRef() bool {
	return o != nil && o.ComponentDefRef != nil
}

// SetComponentDefRef gets a reference to the given string and assigns it to the ComponentDefRef field.
func (o *ParamTplsItem) SetComponentDefRef(v string) {
	o.ComponentDefRef = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParamTplsItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParamTplsItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ParamTplsItem) SetComponent(v string) {
	o.Component = &v
}

// GetParamTplName returns the ParamTplName field value if set, zero value otherwise.
func (o *ParamTplsItem) GetParamTplName() string {
	if o == nil || o.ParamTplName == nil {
		var ret string
		return ret
	}
	return *o.ParamTplName
}

// GetParamTplNameOk returns a tuple with the ParamTplName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetParamTplNameOk() (*string, bool) {
	if o == nil || o.ParamTplName == nil {
		return nil, false
	}
	return o.ParamTplName, true
}

// HasParamTplName returns a boolean if a field has been set.
func (o *ParamTplsItem) HasParamTplName() bool {
	return o != nil && o.ParamTplName != nil
}

// SetParamTplName gets a reference to the given string and assigns it to the ParamTplName field.
func (o *ParamTplsItem) SetParamTplName(v string) {
	o.ParamTplName = &v
}

// GetParamTplPartition returns the ParamTplPartition field value if set, zero value otherwise.
func (o *ParamTplsItem) GetParamTplPartition() ParamTplPartition {
	if o == nil || o.ParamTplPartition == nil {
		var ret ParamTplPartition
		return ret
	}
	return *o.ParamTplPartition
}

// GetParamTplPartitionOk returns a tuple with the ParamTplPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetParamTplPartitionOk() (*ParamTplPartition, bool) {
	if o == nil || o.ParamTplPartition == nil {
		return nil, false
	}
	return o.ParamTplPartition, true
}

// HasParamTplPartition returns a boolean if a field has been set.
func (o *ParamTplsItem) HasParamTplPartition() bool {
	return o != nil && o.ParamTplPartition != nil
}

// SetParamTplPartition gets a reference to the given ParamTplPartition and assigns it to the ParamTplPartition field.
func (o *ParamTplsItem) SetParamTplPartition(v ParamTplPartition) {
	o.ParamTplPartition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ComponentDefRef != nil {
		toSerialize["componentDefRef"] = o.ComponentDefRef
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ParamTplName != nil {
		toSerialize["paramTplName"] = o.ParamTplName
	}
	if o.ParamTplPartition != nil {
		toSerialize["paramTplPartition"] = o.ParamTplPartition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentDefRef   *string            `json:"componentDefRef,omitempty"`
		Component         *string            `json:"component,omitempty"`
		ParamTplName      *string            `json:"paramTplName,omitempty"`
		ParamTplPartition *ParamTplPartition `json:"paramTplPartition,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentDefRef", "component", "paramTplName", "paramTplPartition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentDefRef = all.ComponentDefRef
	o.Component = all.Component
	o.ParamTplName = all.ParamTplName
	if all.ParamTplPartition != nil && !all.ParamTplPartition.IsValid() {
		hasInvalidField = true
	} else {
		o.ParamTplPartition = all.ParamTplPartition
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
