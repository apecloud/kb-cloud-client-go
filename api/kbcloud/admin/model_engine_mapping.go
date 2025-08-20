// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineMapping struct {
	Source              *string                    `json:"source,omitempty"`
	Target              *string                    `json:"target,omitempty"`
	Modules             []string                   `json:"modules,omitempty"`
	Events              [][]EventObject            `json:"events,omitempty"`
	ReplicationMetadata *ReplicationMetadataObject `json:"replicationMetadata,omitempty"`
	Descriptions        *MappingDescription        `json:"descriptions,omitempty"`
	PreCheckers         []string                   `json:"preCheckers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineMapping instantiates a new EngineMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineMapping() *EngineMapping {
	this := EngineMapping{}
	return &this
}

// NewEngineMappingWithDefaults instantiates a new EngineMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineMappingWithDefaults() *EngineMapping {
	this := EngineMapping{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EngineMapping) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EngineMapping) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *EngineMapping) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *EngineMapping) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *EngineMapping) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *EngineMapping) SetTarget(v string) {
	o.Target = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *EngineMapping) GetModules() []string {
	if o == nil || o.Modules == nil {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetModulesOk() (*[]string, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *EngineMapping) HasModules() bool {
	return o != nil && o.Modules != nil
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *EngineMapping) SetModules(v []string) {
	o.Modules = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EngineMapping) GetEvents() [][]EventObject {
	if o == nil || o.Events == nil {
		var ret [][]EventObject
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetEventsOk() (*[][]EventObject, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EngineMapping) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given [][]EventObject and assigns it to the Events field.
func (o *EngineMapping) SetEvents(v [][]EventObject) {
	o.Events = v
}

// GetReplicationMetadata returns the ReplicationMetadata field value if set, zero value otherwise.
func (o *EngineMapping) GetReplicationMetadata() ReplicationMetadataObject {
	if o == nil || o.ReplicationMetadata == nil {
		var ret ReplicationMetadataObject
		return ret
	}
	return *o.ReplicationMetadata
}

// GetReplicationMetadataOk returns a tuple with the ReplicationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetReplicationMetadataOk() (*ReplicationMetadataObject, bool) {
	if o == nil || o.ReplicationMetadata == nil {
		return nil, false
	}
	return o.ReplicationMetadata, true
}

// HasReplicationMetadata returns a boolean if a field has been set.
func (o *EngineMapping) HasReplicationMetadata() bool {
	return o != nil && o.ReplicationMetadata != nil
}

// SetReplicationMetadata gets a reference to the given ReplicationMetadataObject and assigns it to the ReplicationMetadata field.
func (o *EngineMapping) SetReplicationMetadata(v ReplicationMetadataObject) {
	o.ReplicationMetadata = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *EngineMapping) GetDescriptions() MappingDescription {
	if o == nil || o.Descriptions == nil {
		var ret MappingDescription
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetDescriptionsOk() (*MappingDescription, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *EngineMapping) HasDescriptions() bool {
	return o != nil && o.Descriptions != nil
}

// SetDescriptions gets a reference to the given MappingDescription and assigns it to the Descriptions field.
func (o *EngineMapping) SetDescriptions(v MappingDescription) {
	o.Descriptions = &v
}

// GetPreCheckers returns the PreCheckers field value if set, zero value otherwise.
func (o *EngineMapping) GetPreCheckers() []string {
	if o == nil || o.PreCheckers == nil {
		var ret []string
		return ret
	}
	return o.PreCheckers
}

// GetPreCheckersOk returns a tuple with the PreCheckers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineMapping) GetPreCheckersOk() (*[]string, bool) {
	if o == nil || o.PreCheckers == nil {
		return nil, false
	}
	return &o.PreCheckers, true
}

// HasPreCheckers returns a boolean if a field has been set.
func (o *EngineMapping) HasPreCheckers() bool {
	return o != nil && o.PreCheckers != nil
}

// SetPreCheckers gets a reference to the given []string and assigns it to the PreCheckers field.
func (o *EngineMapping) SetPreCheckers(v []string) {
	o.PreCheckers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.ReplicationMetadata != nil {
		toSerialize["replicationMetadata"] = o.ReplicationMetadata
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
	}
	if o.PreCheckers != nil {
		toSerialize["preCheckers"] = o.PreCheckers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source              *string                    `json:"source,omitempty"`
		Target              *string                    `json:"target,omitempty"`
		Modules             []string                   `json:"modules,omitempty"`
		Events              [][]EventObject            `json:"events,omitempty"`
		ReplicationMetadata *ReplicationMetadataObject `json:"replicationMetadata,omitempty"`
		Descriptions        *MappingDescription        `json:"descriptions,omitempty"`
		PreCheckers         []string                   `json:"preCheckers,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "target", "modules", "events", "replicationMetadata", "descriptions", "preCheckers"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Source = all.Source
	o.Target = all.Target
	o.Modules = all.Modules
	o.Events = all.Events
	if all.ReplicationMetadata != nil && all.ReplicationMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationMetadata = all.ReplicationMetadata
	if all.Descriptions != nil && all.Descriptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Descriptions = all.Descriptions
	o.PreCheckers = all.PreCheckers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
