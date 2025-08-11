// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineDefinition struct {
	Name       *string                 `json:"name,omitempty"`
	EngineName *string                 `json:"engineName,omitempty"`
	Source     *EngineDefinitionDetail `json:"source,omitempty"`
	Target     *EngineDefinitionDetail `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineDefinition instantiates a new EngineDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineDefinition() *EngineDefinition {
	this := EngineDefinition{}
	return &this
}

// NewEngineDefinitionWithDefaults instantiates a new EngineDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineDefinitionWithDefaults() *EngineDefinition {
	this := EngineDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EngineDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EngineDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EngineDefinition) SetName(v string) {
	o.Name = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *EngineDefinition) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinition) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *EngineDefinition) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *EngineDefinition) SetEngineName(v string) {
	o.EngineName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EngineDefinition) GetSource() EngineDefinitionDetail {
	if o == nil || o.Source == nil {
		var ret EngineDefinitionDetail
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinition) GetSourceOk() (*EngineDefinitionDetail, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EngineDefinition) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given EngineDefinitionDetail and assigns it to the Source field.
func (o *EngineDefinition) SetSource(v EngineDefinitionDetail) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *EngineDefinition) GetTarget() EngineDefinitionDetail {
	if o == nil || o.Target == nil {
		var ret EngineDefinitionDetail
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinition) GetTargetOk() (*EngineDefinitionDetail, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *EngineDefinition) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given EngineDefinitionDetail and assigns it to the Target field.
func (o *EngineDefinition) SetTarget(v EngineDefinitionDetail) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string                 `json:"name,omitempty"`
		EngineName *string                 `json:"engineName,omitempty"`
		Source     *EngineDefinitionDetail `json:"source,omitempty"`
		Target     *EngineDefinitionDetail `json:"target,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "engineName", "source", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.EngineName = all.EngineName
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
