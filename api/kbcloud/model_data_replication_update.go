// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataReplicationUpdate struct {
	ClusterUpdates *DataReplicationCreate `json:"clusterUpdates,omitempty"`
	IsRestart      common.NullableBool    `json:"isRestart,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataReplicationUpdate instantiates a new DataReplicationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataReplicationUpdate() *DataReplicationUpdate {
	this := DataReplicationUpdate{}
	return &this
}

// NewDataReplicationUpdateWithDefaults instantiates a new DataReplicationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataReplicationUpdateWithDefaults() *DataReplicationUpdate {
	this := DataReplicationUpdate{}
	return &this
}

// GetClusterUpdates returns the ClusterUpdates field value if set, zero value otherwise.
func (o *DataReplicationUpdate) GetClusterUpdates() DataReplicationCreate {
	if o == nil || o.ClusterUpdates == nil {
		var ret DataReplicationCreate
		return ret
	}
	return *o.ClusterUpdates
}

// GetClusterUpdatesOk returns a tuple with the ClusterUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationUpdate) GetClusterUpdatesOk() (*DataReplicationCreate, bool) {
	if o == nil || o.ClusterUpdates == nil {
		return nil, false
	}
	return o.ClusterUpdates, true
}

// HasClusterUpdates returns a boolean if a field has been set.
func (o *DataReplicationUpdate) HasClusterUpdates() bool {
	return o != nil && o.ClusterUpdates != nil
}

// SetClusterUpdates gets a reference to the given DataReplicationCreate and assigns it to the ClusterUpdates field.
func (o *DataReplicationUpdate) SetClusterUpdates(v DataReplicationCreate) {
	o.ClusterUpdates = &v
}

// GetIsRestart returns the IsRestart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataReplicationUpdate) GetIsRestart() bool {
	if o == nil || o.IsRestart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRestart.Get()
}

// GetIsRestartOk returns a tuple with the IsRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataReplicationUpdate) GetIsRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRestart.Get(), o.IsRestart.IsSet()
}

// HasIsRestart returns a boolean if a field has been set.
func (o *DataReplicationUpdate) HasIsRestart() bool {
	return o != nil && o.IsRestart.IsSet()
}

// SetIsRestart gets a reference to the given common.NullableBool and assigns it to the IsRestart field.
func (o *DataReplicationUpdate) SetIsRestart(v bool) {
	o.IsRestart.Set(&v)
}

// SetIsRestartNil sets the value for IsRestart to be an explicit nil.
func (o *DataReplicationUpdate) SetIsRestartNil() {
	o.IsRestart.Set(nil)
}

// UnsetIsRestart ensures that no value is present for IsRestart, not even an explicit nil.
func (o *DataReplicationUpdate) UnsetIsRestart() {
	o.IsRestart.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataReplicationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterUpdates != nil {
		toSerialize["clusterUpdates"] = o.ClusterUpdates
	}
	if o.IsRestart.IsSet() {
		toSerialize["isRestart"] = o.IsRestart.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataReplicationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterUpdates *DataReplicationCreate `json:"clusterUpdates,omitempty"`
		IsRestart      common.NullableBool    `json:"isRestart,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterUpdates", "isRestart"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ClusterUpdates != nil && all.ClusterUpdates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClusterUpdates = all.ClusterUpdates
	o.IsRestart = all.IsRestart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
