// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DataCheckCreate Exactly one entry point must be used. Provide channelID with checkType to create from an existing channel, or provide environmentID, project, source, target, replicationObjects, and checkType for direct creation.
type DataCheckCreate struct {
	// Type of check to run. A standalone check runs exactly one type.
	CheckType DataCheckType `json:"checkType"`
	// Existing data replication channel ID used only as a create-time shortcut.
	ChannelId *string `json:"channelID,omitempty"`
	// Environment UUID for direct check creation.
	EnvironmentId *string `json:"environmentID,omitempty"`
	// Kubernetes namespace for direct check creation.
	Project            *string                    `json:"project,omitempty"`
	Source             *DataChannelEndpointCreate `json:"source,omitempty"`
	Target             *DataChannelEndpointCreate `json:"target,omitempty"`
	ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckCreate instantiates a new DataCheckCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckCreate(checkType DataCheckType) *DataCheckCreate {
	this := DataCheckCreate{}
	this.CheckType = checkType
	return &this
}

// NewDataCheckCreateWithDefaults instantiates a new DataCheckCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckCreateWithDefaults() *DataCheckCreate {
	this := DataCheckCreate{}
	return &this
}

// GetCheckType returns the CheckType field value.
func (o *DataCheckCreate) GetCheckType() DataCheckType {
	if o == nil {
		var ret DataCheckType
		return ret
	}
	return o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetCheckTypeOk() (*DataCheckType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckType, true
}

// SetCheckType sets field value.
func (o *DataCheckCreate) SetCheckType(v DataCheckType) {
	o.CheckType = v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *DataCheckCreate) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *DataCheckCreate) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *DataCheckCreate) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DataCheckCreate) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DataCheckCreate) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DataCheckCreate) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DataCheckCreate) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DataCheckCreate) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DataCheckCreate) SetProject(v string) {
	o.Project = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DataCheckCreate) GetSource() DataChannelEndpointCreate {
	if o == nil || o.Source == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetSourceOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DataCheckCreate) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given DataChannelEndpointCreate and assigns it to the Source field.
func (o *DataCheckCreate) SetSource(v DataChannelEndpointCreate) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DataCheckCreate) GetTarget() DataChannelEndpointCreate {
	if o == nil || o.Target == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetTargetOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DataCheckCreate) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given DataChannelEndpointCreate and assigns it to the Target field.
func (o *DataCheckCreate) SetTarget(v DataChannelEndpointCreate) {
	o.Target = &v
}

// GetReplicationObjects returns the ReplicationObjects field value if set, zero value otherwise.
func (o *DataCheckCreate) GetReplicationObjects() DataChannelObject {
	if o == nil || o.ReplicationObjects == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ReplicationObjects
}

// GetReplicationObjectsOk returns a tuple with the ReplicationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckCreate) GetReplicationObjectsOk() (*DataChannelObject, bool) {
	if o == nil || o.ReplicationObjects == nil {
		return nil, false
	}
	return o.ReplicationObjects, true
}

// HasReplicationObjects returns a boolean if a field has been set.
func (o *DataCheckCreate) HasReplicationObjects() bool {
	return o != nil && o.ReplicationObjects != nil
}

// SetReplicationObjects gets a reference to the given DataChannelObject and assigns it to the ReplicationObjects field.
func (o *DataCheckCreate) SetReplicationObjects(v DataChannelObject) {
	o.ReplicationObjects = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["checkType"] = o.CheckType
	if o.ChannelId != nil {
		toSerialize["channelID"] = o.ChannelId
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckType          *DataCheckType             `json:"checkType"`
		ChannelId          *string                    `json:"channelID,omitempty"`
		EnvironmentId      *string                    `json:"environmentID,omitempty"`
		Project            *string                    `json:"project,omitempty"`
		Source             *DataChannelEndpointCreate `json:"source,omitempty"`
		Target             *DataChannelEndpointCreate `json:"target,omitempty"`
		ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CheckType == nil {
		return fmt.Errorf("required field checkType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkType", "channelID", "environmentID", "project", "source", "target", "replicationObjects"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CheckType.IsValid() {
		hasInvalidField = true
	} else {
		o.CheckType = *all.CheckType
	}
	o.ChannelId = all.ChannelId
	o.EnvironmentId = all.EnvironmentId
	o.Project = all.Project
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
