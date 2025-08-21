// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataReplicationCreate struct {
	ChannelName        *string                    `json:"channelName,omitempty"`
	Modules            []string                   `json:"modules,omitempty"`
	EnvironmentId      *string                    `json:"environmentID,omitempty"`
	Project            *string                    `json:"project,omitempty"`
	StandardDefinition *string                    `json:"standardDefinition,omitempty"`
	Source             *DataChannelEndpointCreate `json:"source,omitempty"`
	Target             *DataChannelEndpointCreate `json:"target,omitempty"`
	ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
	Events             []EventObject              `json:"events,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataReplicationCreate instantiates a new DataReplicationCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataReplicationCreate() *DataReplicationCreate {
	this := DataReplicationCreate{}
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	return &this
}

// NewDataReplicationCreateWithDefaults instantiates a new DataReplicationCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataReplicationCreateWithDefaults() *DataReplicationCreate {
	this := DataReplicationCreate{}
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	return &this
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *DataReplicationCreate) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetModules() []string {
	if o == nil || o.Modules == nil {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetModulesOk() (*[]string, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasModules() bool {
	return o != nil && o.Modules != nil
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *DataReplicationCreate) SetModules(v []string) {
	o.Modules = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DataReplicationCreate) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DataReplicationCreate) SetProject(v string) {
	o.Project = &v
}

// GetStandardDefinition returns the StandardDefinition field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetStandardDefinition() string {
	if o == nil || o.StandardDefinition == nil {
		var ret string
		return ret
	}
	return *o.StandardDefinition
}

// GetStandardDefinitionOk returns a tuple with the StandardDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetStandardDefinitionOk() (*string, bool) {
	if o == nil || o.StandardDefinition == nil {
		return nil, false
	}
	return o.StandardDefinition, true
}

// HasStandardDefinition returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasStandardDefinition() bool {
	return o != nil && o.StandardDefinition != nil
}

// SetStandardDefinition gets a reference to the given string and assigns it to the StandardDefinition field.
func (o *DataReplicationCreate) SetStandardDefinition(v string) {
	o.StandardDefinition = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetSource() DataChannelEndpointCreate {
	if o == nil || o.Source == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetSourceOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given DataChannelEndpointCreate and assigns it to the Source field.
func (o *DataReplicationCreate) SetSource(v DataChannelEndpointCreate) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetTarget() DataChannelEndpointCreate {
	if o == nil || o.Target == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetTargetOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given DataChannelEndpointCreate and assigns it to the Target field.
func (o *DataReplicationCreate) SetTarget(v DataChannelEndpointCreate) {
	o.Target = &v
}

// GetReplicationObjects returns the ReplicationObjects field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetReplicationObjects() DataChannelObject {
	if o == nil || o.ReplicationObjects == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ReplicationObjects
}

// GetReplicationObjectsOk returns a tuple with the ReplicationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetReplicationObjectsOk() (*DataChannelObject, bool) {
	if o == nil || o.ReplicationObjects == nil {
		return nil, false
	}
	return o.ReplicationObjects, true
}

// HasReplicationObjects returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasReplicationObjects() bool {
	return o != nil && o.ReplicationObjects != nil
}

// SetReplicationObjects gets a reference to the given DataChannelObject and assigns it to the ReplicationObjects field.
func (o *DataReplicationCreate) SetReplicationObjects(v DataChannelObject) {
	o.ReplicationObjects = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *DataReplicationCreate) GetEvents() []EventObject {
	if o == nil || o.Events == nil {
		var ret []EventObject
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationCreate) GetEventsOk() (*[]EventObject, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *DataReplicationCreate) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []EventObject and assigns it to the Events field.
func (o *DataReplicationCreate) SetEvents(v []EventObject) {
	o.Events = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataReplicationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ChannelName != nil {
		toSerialize["channelName"] = o.ChannelName
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.StandardDefinition != nil {
		toSerialize["standardDefinition"] = o.StandardDefinition
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.ReplicationObjects != nil {
		toSerialize["replicationObjects"] = o.ReplicationObjects
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataReplicationCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelName        *string                    `json:"channelName,omitempty"`
		Modules            []string                   `json:"modules,omitempty"`
		EnvironmentId      *string                    `json:"environmentID,omitempty"`
		Project            *string                    `json:"project,omitempty"`
		StandardDefinition *string                    `json:"standardDefinition,omitempty"`
		Source             *DataChannelEndpointCreate `json:"source,omitempty"`
		Target             *DataChannelEndpointCreate `json:"target,omitempty"`
		ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
		Events             []EventObject              `json:"events,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channelName", "modules", "environmentID", "project", "standardDefinition", "source", "target", "replicationObjects", "events"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChannelName = all.ChannelName
	o.Modules = all.Modules
	o.EnvironmentId = all.EnvironmentId
	o.Project = all.Project
	o.StandardDefinition = all.StandardDefinition
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	if all.ReplicationObjects != nil && all.ReplicationObjects.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationObjects = all.ReplicationObjects
	o.Events = all.Events

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
