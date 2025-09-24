// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PreCheckCreate struct {
	Modules            []string                   `json:"modules,omitempty"`
	EnvironmentId      *string                    `json:"environmentID,omitempty"`
	Source             *DataChannelEndpointCreate `json:"source,omitempty"`
	Target             *DataChannelEndpointCreate `json:"target,omitempty"`
	ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPreCheckCreate instantiates a new PreCheckCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPreCheckCreate() *PreCheckCreate {
	this := PreCheckCreate{}
	return &this
}

// NewPreCheckCreateWithDefaults instantiates a new PreCheckCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPreCheckCreateWithDefaults() *PreCheckCreate {
	this := PreCheckCreate{}
	return &this
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *PreCheckCreate) GetModules() []string {
	if o == nil || o.Modules == nil {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckCreate) GetModulesOk() (*[]string, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *PreCheckCreate) HasModules() bool {
	return o != nil && o.Modules != nil
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *PreCheckCreate) SetModules(v []string) {
	o.Modules = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *PreCheckCreate) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckCreate) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *PreCheckCreate) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *PreCheckCreate) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PreCheckCreate) GetSource() DataChannelEndpointCreate {
	if o == nil || o.Source == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckCreate) GetSourceOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PreCheckCreate) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given DataChannelEndpointCreate and assigns it to the Source field.
func (o *PreCheckCreate) SetSource(v DataChannelEndpointCreate) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PreCheckCreate) GetTarget() DataChannelEndpointCreate {
	if o == nil || o.Target == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckCreate) GetTargetOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PreCheckCreate) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given DataChannelEndpointCreate and assigns it to the Target field.
func (o *PreCheckCreate) SetTarget(v DataChannelEndpointCreate) {
	o.Target = &v
}

// GetReplicationObjects returns the ReplicationObjects field value if set, zero value otherwise.
func (o *PreCheckCreate) GetReplicationObjects() DataChannelObject {
	if o == nil || o.ReplicationObjects == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ReplicationObjects
}

// GetReplicationObjectsOk returns a tuple with the ReplicationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreCheckCreate) GetReplicationObjectsOk() (*DataChannelObject, bool) {
	if o == nil || o.ReplicationObjects == nil {
		return nil, false
	}
	return o.ReplicationObjects, true
}

// HasReplicationObjects returns a boolean if a field has been set.
func (o *PreCheckCreate) HasReplicationObjects() bool {
	return o != nil && o.ReplicationObjects != nil
}

// SetReplicationObjects gets a reference to the given DataChannelObject and assigns it to the ReplicationObjects field.
func (o *PreCheckCreate) SetReplicationObjects(v DataChannelObject) {
	o.ReplicationObjects = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PreCheckCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
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
func (o *PreCheckCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Modules            []string                   `json:"modules,omitempty"`
		EnvironmentId      *string                    `json:"environmentID,omitempty"`
		Source             *DataChannelEndpointCreate `json:"source,omitempty"`
		Target             *DataChannelEndpointCreate `json:"target,omitempty"`
		ReplicationObjects *DataChannelObject         `json:"replicationObjects,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"modules", "environmentID", "source", "target", "replicationObjects"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Modules = all.Modules
	o.EnvironmentId = all.EnvironmentId
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
