// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataReplicationOption struct {
	ModuleDefinitions []ModuleDefinition   `json:"moduleDefinitions,omitempty"`
	EngineDefinitions []EngineDefinition   `json:"engineDefinitions,omitempty"`
	Mappings          []EngineMapping      `json:"mappings,omitempty"`
	Standards         []StandardDefinition `json:"standards,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataReplicationOption instantiates a new DataReplicationOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataReplicationOption() *DataReplicationOption {
	this := DataReplicationOption{}
	return &this
}

// NewDataReplicationOptionWithDefaults instantiates a new DataReplicationOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataReplicationOptionWithDefaults() *DataReplicationOption {
	this := DataReplicationOption{}
	return &this
}

// GetModuleDefinitions returns the ModuleDefinitions field value if set, zero value otherwise.
func (o *DataReplicationOption) GetModuleDefinitions() []ModuleDefinition {
	if o == nil || o.ModuleDefinitions == nil {
		var ret []ModuleDefinition
		return ret
	}
	return o.ModuleDefinitions
}

// GetModuleDefinitionsOk returns a tuple with the ModuleDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationOption) GetModuleDefinitionsOk() (*[]ModuleDefinition, bool) {
	if o == nil || o.ModuleDefinitions == nil {
		return nil, false
	}
	return &o.ModuleDefinitions, true
}

// HasModuleDefinitions returns a boolean if a field has been set.
func (o *DataReplicationOption) HasModuleDefinitions() bool {
	return o != nil && o.ModuleDefinitions != nil
}

// SetModuleDefinitions gets a reference to the given []ModuleDefinition and assigns it to the ModuleDefinitions field.
func (o *DataReplicationOption) SetModuleDefinitions(v []ModuleDefinition) {
	o.ModuleDefinitions = v
}

// GetEngineDefinitions returns the EngineDefinitions field value if set, zero value otherwise.
func (o *DataReplicationOption) GetEngineDefinitions() []EngineDefinition {
	if o == nil || o.EngineDefinitions == nil {
		var ret []EngineDefinition
		return ret
	}
	return o.EngineDefinitions
}

// GetEngineDefinitionsOk returns a tuple with the EngineDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationOption) GetEngineDefinitionsOk() (*[]EngineDefinition, bool) {
	if o == nil || o.EngineDefinitions == nil {
		return nil, false
	}
	return &o.EngineDefinitions, true
}

// HasEngineDefinitions returns a boolean if a field has been set.
func (o *DataReplicationOption) HasEngineDefinitions() bool {
	return o != nil && o.EngineDefinitions != nil
}

// SetEngineDefinitions gets a reference to the given []EngineDefinition and assigns it to the EngineDefinitions field.
func (o *DataReplicationOption) SetEngineDefinitions(v []EngineDefinition) {
	o.EngineDefinitions = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *DataReplicationOption) GetMappings() []EngineMapping {
	if o == nil || o.Mappings == nil {
		var ret []EngineMapping
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationOption) GetMappingsOk() (*[]EngineMapping, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *DataReplicationOption) HasMappings() bool {
	return o != nil && o.Mappings != nil
}

// SetMappings gets a reference to the given []EngineMapping and assigns it to the Mappings field.
func (o *DataReplicationOption) SetMappings(v []EngineMapping) {
	o.Mappings = v
}

// GetStandards returns the Standards field value if set, zero value otherwise.
func (o *DataReplicationOption) GetStandards() []StandardDefinition {
	if o == nil || o.Standards == nil {
		var ret []StandardDefinition
		return ret
	}
	return o.Standards
}

// GetStandardsOk returns a tuple with the Standards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationOption) GetStandardsOk() (*[]StandardDefinition, bool) {
	if o == nil || o.Standards == nil {
		return nil, false
	}
	return &o.Standards, true
}

// HasStandards returns a boolean if a field has been set.
func (o *DataReplicationOption) HasStandards() bool {
	return o != nil && o.Standards != nil
}

// SetStandards gets a reference to the given []StandardDefinition and assigns it to the Standards field.
func (o *DataReplicationOption) SetStandards(v []StandardDefinition) {
	o.Standards = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataReplicationOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ModuleDefinitions != nil {
		toSerialize["moduleDefinitions"] = o.ModuleDefinitions
	}
	if o.EngineDefinitions != nil {
		toSerialize["engineDefinitions"] = o.EngineDefinitions
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Standards != nil {
		toSerialize["standards"] = o.Standards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataReplicationOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModuleDefinitions []ModuleDefinition   `json:"moduleDefinitions,omitempty"`
		EngineDefinitions []EngineDefinition   `json:"engineDefinitions,omitempty"`
		Mappings          []EngineMapping      `json:"mappings,omitempty"`
		Standards         []StandardDefinition `json:"standards,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"moduleDefinitions", "engineDefinitions", "mappings", "standards"})
	} else {
		return err
	}
	o.ModuleDefinitions = all.ModuleDefinitions
	o.EngineDefinitions = all.EngineDefinitions
	o.Mappings = all.Mappings
	o.Standards = all.Standards

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
