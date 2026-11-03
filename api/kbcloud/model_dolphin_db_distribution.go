// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBDistribution struct {
	ObservedAt *time.Time                     `json:"observedAt,omitempty"`
	Volumes    []DolphinDBVolumeUsage         `json:"volumes,omitempty"`
	Replicas   []DolphinDBReplicaDistribution `json:"replicas,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBDistribution instantiates a new DolphinDBDistribution object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBDistribution() *DolphinDBDistribution {
	this := DolphinDBDistribution{}
	return &this
}

// NewDolphinDBDistributionWithDefaults instantiates a new DolphinDBDistribution object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBDistributionWithDefaults() *DolphinDBDistribution {
	this := DolphinDBDistribution{}
	return &this
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *DolphinDBDistribution) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBDistribution) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *DolphinDBDistribution) HasObservedAt() bool {
	return o != nil && o.ObservedAt != nil
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *DolphinDBDistribution) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *DolphinDBDistribution) GetVolumes() []DolphinDBVolumeUsage {
	if o == nil || o.Volumes == nil {
		var ret []DolphinDBVolumeUsage
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBDistribution) GetVolumesOk() (*[]DolphinDBVolumeUsage, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *DolphinDBDistribution) HasVolumes() bool {
	return o != nil && o.Volumes != nil
}

// SetVolumes gets a reference to the given []DolphinDBVolumeUsage and assigns it to the Volumes field.
func (o *DolphinDBDistribution) SetVolumes(v []DolphinDBVolumeUsage) {
	o.Volumes = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *DolphinDBDistribution) GetReplicas() []DolphinDBReplicaDistribution {
	if o == nil || o.Replicas == nil {
		var ret []DolphinDBReplicaDistribution
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBDistribution) GetReplicasOk() (*[]DolphinDBReplicaDistribution, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *DolphinDBDistribution) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given []DolphinDBReplicaDistribution and assigns it to the Replicas field.
func (o *DolphinDBDistribution) SetReplicas(v []DolphinDBReplicaDistribution) {
	o.Replicas = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObservedAt != nil {
		if o.ObservedAt.Nanosecond() == 0 {
			toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBDistribution) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObservedAt *time.Time                     `json:"observedAt,omitempty"`
		Volumes    []DolphinDBVolumeUsage         `json:"volumes,omitempty"`
		Replicas   []DolphinDBReplicaDistribution `json:"replicas,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"observedAt", "volumes", "replicas"})
	} else {
		return err
	}
	o.ObservedAt = all.ObservedAt
	o.Volumes = all.Volumes
	o.Replicas = all.Replicas

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
