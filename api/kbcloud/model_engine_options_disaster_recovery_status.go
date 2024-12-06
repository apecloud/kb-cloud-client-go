// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type EngineOptionsDisasterRecoveryStatus struct {
	Delay            *EngineOptionsDisasterRecoverySource `json:"delay,omitempty"`
	ReplicationPoint *EngineOptionsDisasterRecoverySource `json:"replicationPoint,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionsDisasterRecoveryStatus instantiates a new EngineOptionsDisasterRecoveryStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionsDisasterRecoveryStatus() *EngineOptionsDisasterRecoveryStatus {
	this := EngineOptionsDisasterRecoveryStatus{}
	return &this
}

// NewEngineOptionsDisasterRecoveryStatusWithDefaults instantiates a new EngineOptionsDisasterRecoveryStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionsDisasterRecoveryStatusWithDefaults() *EngineOptionsDisasterRecoveryStatus {
	this := EngineOptionsDisasterRecoveryStatus{}
	return &this
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *EngineOptionsDisasterRecoveryStatus) GetDelay() EngineOptionsDisasterRecoverySource {
	if o == nil || o.Delay == nil {
		var ret EngineOptionsDisasterRecoverySource
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionsDisasterRecoveryStatus) GetDelayOk() (*EngineOptionsDisasterRecoverySource, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *EngineOptionsDisasterRecoveryStatus) HasDelay() bool {
	return o != nil && o.Delay != nil
}

// SetDelay gets a reference to the given EngineOptionsDisasterRecoverySource and assigns it to the Delay field.
func (o *EngineOptionsDisasterRecoveryStatus) SetDelay(v EngineOptionsDisasterRecoverySource) {
	o.Delay = &v
}

// GetReplicationPoint returns the ReplicationPoint field value if set, zero value otherwise.
func (o *EngineOptionsDisasterRecoveryStatus) GetReplicationPoint() EngineOptionsDisasterRecoverySource {
	if o == nil || o.ReplicationPoint == nil {
		var ret EngineOptionsDisasterRecoverySource
		return ret
	}
	return *o.ReplicationPoint
}

// GetReplicationPointOk returns a tuple with the ReplicationPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionsDisasterRecoveryStatus) GetReplicationPointOk() (*EngineOptionsDisasterRecoverySource, bool) {
	if o == nil || o.ReplicationPoint == nil {
		return nil, false
	}
	return o.ReplicationPoint, true
}

// HasReplicationPoint returns a boolean if a field has been set.
func (o *EngineOptionsDisasterRecoveryStatus) HasReplicationPoint() bool {
	return o != nil && o.ReplicationPoint != nil
}

// SetReplicationPoint gets a reference to the given EngineOptionsDisasterRecoverySource and assigns it to the ReplicationPoint field.
func (o *EngineOptionsDisasterRecoveryStatus) SetReplicationPoint(v EngineOptionsDisasterRecoverySource) {
	o.ReplicationPoint = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionsDisasterRecoveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	if o.ReplicationPoint != nil {
		toSerialize["replicationPoint"] = o.ReplicationPoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionsDisasterRecoveryStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delay            *EngineOptionsDisasterRecoverySource `json:"delay,omitempty"`
		ReplicationPoint *EngineOptionsDisasterRecoverySource `json:"replicationPoint,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"delay", "replicationPoint"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Delay != nil && all.Delay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Delay = all.Delay
	if all.ReplicationPoint != nil && all.ReplicationPoint.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationPoint = all.ReplicationPoint

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
