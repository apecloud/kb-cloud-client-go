// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ResourceConstraintUpdate struct {
	Replicas *IntegerOption  `json:"replicas,omitempty"`
	Shards   *IntegerOption  `json:"shards,omitempty"`
	Volumes  []StorageOption `json:"volumes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResourceConstraintUpdate instantiates a new ResourceConstraintUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResourceConstraintUpdate() *ResourceConstraintUpdate {
	this := ResourceConstraintUpdate{}
	return &this
}

// NewResourceConstraintUpdateWithDefaults instantiates a new ResourceConstraintUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResourceConstraintUpdateWithDefaults() *ResourceConstraintUpdate {
	this := ResourceConstraintUpdate{}
	return &this
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ResourceConstraintUpdate) GetReplicas() IntegerOption {
	if o == nil || o.Replicas == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraintUpdate) GetReplicasOk() (*IntegerOption, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ResourceConstraintUpdate) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given IntegerOption and assigns it to the Replicas field.
func (o *ResourceConstraintUpdate) SetReplicas(v IntegerOption) {
	o.Replicas = &v
}

// GetShards returns the Shards field value if set, zero value otherwise.
func (o *ResourceConstraintUpdate) GetShards() IntegerOption {
	if o == nil || o.Shards == nil {
		var ret IntegerOption
		return ret
	}
	return *o.Shards
}

// GetShardsOk returns a tuple with the Shards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraintUpdate) GetShardsOk() (*IntegerOption, bool) {
	if o == nil || o.Shards == nil {
		return nil, false
	}
	return o.Shards, true
}

// HasShards returns a boolean if a field has been set.
func (o *ResourceConstraintUpdate) HasShards() bool {
	return o != nil && o.Shards != nil
}

// SetShards gets a reference to the given IntegerOption and assigns it to the Shards field.
func (o *ResourceConstraintUpdate) SetShards(v IntegerOption) {
	o.Shards = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ResourceConstraintUpdate) GetVolumes() []StorageOption {
	if o == nil || o.Volumes == nil {
		var ret []StorageOption
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConstraintUpdate) GetVolumesOk() (*[]StorageOption, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ResourceConstraintUpdate) HasVolumes() bool {
	return o != nil && o.Volumes != nil
}

// SetVolumes gets a reference to the given []StorageOption and assigns it to the Volumes field.
func (o *ResourceConstraintUpdate) SetVolumes(v []StorageOption) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResourceConstraintUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Shards != nil {
		toSerialize["shards"] = o.Shards
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResourceConstraintUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Replicas *IntegerOption  `json:"replicas,omitempty"`
		Shards   *IntegerOption  `json:"shards,omitempty"`
		Volumes  []StorageOption `json:"volumes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"replicas", "shards", "volumes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Replicas != nil && all.Replicas.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Replicas = all.Replicas
	if all.Shards != nil && all.Shards.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Shards = all.Shards
	o.Volumes = all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
