// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelItem struct {
	ChannelId          *string                  `json:"channelID,omitempty"`
	ChannelName        *string                  `json:"channelName,omitempty"`
	ChannelStatus      *ChannelStatus           `json:"channelStatus,omitempty"`
	EnvironmentId      *string                  `json:"environmentID,omitempty"`
	EnvironmentName    *string                  `json:"environmentName,omitempty"`
	Project            *string                  `json:"project,omitempty"`
	StandardDefinition *string                  `json:"standardDefinition,omitempty"`
	Source             *DataChannelListEndpoint `json:"source,omitempty"`
	Target             *DataChannelListEndpoint `json:"target,omitempty"`
	ReplicationObjects *DataChannelObject       `json:"replicationObjects,omitempty"`
	Modules            []string                 `json:"modules,omitempty"`
	Events             []EventObject            `json:"events,omitempty"`
	CreatedAt          *time.Time               `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelItem instantiates a new DataChannelItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelItem() *DataChannelItem {
	this := DataChannelItem{}
	return &this
}

// NewDataChannelItemWithDefaults instantiates a new DataChannelItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelItemWithDefaults() *DataChannelItem {
	this := DataChannelItem{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *DataChannelItem) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *DataChannelItem) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *DataChannelItem) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *DataChannelItem) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *DataChannelItem) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *DataChannelItem) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetChannelStatus returns the ChannelStatus field value if set, zero value otherwise.
func (o *DataChannelItem) GetChannelStatus() ChannelStatus {
	if o == nil || o.ChannelStatus == nil {
		var ret ChannelStatus
		return ret
	}
	return *o.ChannelStatus
}

// GetChannelStatusOk returns a tuple with the ChannelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetChannelStatusOk() (*ChannelStatus, bool) {
	if o == nil || o.ChannelStatus == nil {
		return nil, false
	}
	return o.ChannelStatus, true
}

// HasChannelStatus returns a boolean if a field has been set.
func (o *DataChannelItem) HasChannelStatus() bool {
	return o != nil && o.ChannelStatus != nil
}

// SetChannelStatus gets a reference to the given ChannelStatus and assigns it to the ChannelStatus field.
func (o *DataChannelItem) SetChannelStatus(v ChannelStatus) {
	o.ChannelStatus = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DataChannelItem) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DataChannelItem) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DataChannelItem) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *DataChannelItem) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *DataChannelItem) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *DataChannelItem) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DataChannelItem) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DataChannelItem) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DataChannelItem) SetProject(v string) {
	o.Project = &v
}

// GetStandardDefinition returns the StandardDefinition field value if set, zero value otherwise.
func (o *DataChannelItem) GetStandardDefinition() string {
	if o == nil || o.StandardDefinition == nil {
		var ret string
		return ret
	}
	return *o.StandardDefinition
}

// GetStandardDefinitionOk returns a tuple with the StandardDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetStandardDefinitionOk() (*string, bool) {
	if o == nil || o.StandardDefinition == nil {
		return nil, false
	}
	return o.StandardDefinition, true
}

// HasStandardDefinition returns a boolean if a field has been set.
func (o *DataChannelItem) HasStandardDefinition() bool {
	return o != nil && o.StandardDefinition != nil
}

// SetStandardDefinition gets a reference to the given string and assigns it to the StandardDefinition field.
func (o *DataChannelItem) SetStandardDefinition(v string) {
	o.StandardDefinition = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DataChannelItem) GetSource() DataChannelListEndpoint {
	if o == nil || o.Source == nil {
		var ret DataChannelListEndpoint
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetSourceOk() (*DataChannelListEndpoint, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DataChannelItem) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given DataChannelListEndpoint and assigns it to the Source field.
func (o *DataChannelItem) SetSource(v DataChannelListEndpoint) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DataChannelItem) GetTarget() DataChannelListEndpoint {
	if o == nil || o.Target == nil {
		var ret DataChannelListEndpoint
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetTargetOk() (*DataChannelListEndpoint, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DataChannelItem) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given DataChannelListEndpoint and assigns it to the Target field.
func (o *DataChannelItem) SetTarget(v DataChannelListEndpoint) {
	o.Target = &v
}

// GetReplicationObjects returns the ReplicationObjects field value if set, zero value otherwise.
func (o *DataChannelItem) GetReplicationObjects() DataChannelObject {
	if o == nil || o.ReplicationObjects == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ReplicationObjects
}

// GetReplicationObjectsOk returns a tuple with the ReplicationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetReplicationObjectsOk() (*DataChannelObject, bool) {
	if o == nil || o.ReplicationObjects == nil {
		return nil, false
	}
	return o.ReplicationObjects, true
}

// HasReplicationObjects returns a boolean if a field has been set.
func (o *DataChannelItem) HasReplicationObjects() bool {
	return o != nil && o.ReplicationObjects != nil
}

// SetReplicationObjects gets a reference to the given DataChannelObject and assigns it to the ReplicationObjects field.
func (o *DataChannelItem) SetReplicationObjects(v DataChannelObject) {
	o.ReplicationObjects = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *DataChannelItem) GetModules() []string {
	if o == nil || o.Modules == nil {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetModulesOk() (*[]string, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *DataChannelItem) HasModules() bool {
	return o != nil && o.Modules != nil
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *DataChannelItem) SetModules(v []string) {
	o.Modules = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *DataChannelItem) GetEvents() []EventObject {
	if o == nil || o.Events == nil {
		var ret []EventObject
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetEventsOk() (*[]EventObject, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *DataChannelItem) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []EventObject and assigns it to the Events field.
func (o *DataChannelItem) SetEvents(v []EventObject) {
	o.Events = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataChannelItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataChannelItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DataChannelItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ChannelId != nil {
		toSerialize["channelID"] = o.ChannelId
	}
	if o.ChannelName != nil {
		toSerialize["channelName"] = o.ChannelName
	}
	if o.ChannelStatus != nil {
		toSerialize["channelStatus"] = o.ChannelStatus
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
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
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId          *string                  `json:"channelID,omitempty"`
		ChannelName        *string                  `json:"channelName,omitempty"`
		ChannelStatus      *ChannelStatus           `json:"channelStatus,omitempty"`
		EnvironmentId      *string                  `json:"environmentID,omitempty"`
		EnvironmentName    *string                  `json:"environmentName,omitempty"`
		Project            *string                  `json:"project,omitempty"`
		StandardDefinition *string                  `json:"standardDefinition,omitempty"`
		Source             *DataChannelListEndpoint `json:"source,omitempty"`
		Target             *DataChannelListEndpoint `json:"target,omitempty"`
		ReplicationObjects *DataChannelObject       `json:"replicationObjects,omitempty"`
		Modules            []string                 `json:"modules,omitempty"`
		Events             []EventObject            `json:"events,omitempty"`
		CreatedAt          *time.Time               `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channelID", "channelName", "channelStatus", "environmentID", "environmentName", "project", "standardDefinition", "source", "target", "replicationObjects", "modules", "events", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChannelId = all.ChannelId
	o.ChannelName = all.ChannelName
	if all.ChannelStatus != nil && !all.ChannelStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ChannelStatus = all.ChannelStatus
	}
	o.EnvironmentId = all.EnvironmentId
	o.EnvironmentName = all.EnvironmentName
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
	o.Modules = all.Modules
	o.Events = all.Events
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
