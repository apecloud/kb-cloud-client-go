// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type InfluxDBBucketRetentionRule struct {
	Type *InfluxDBRetentionRuleType `json:"type,omitempty"`
	// 0 means forever. Positive retention must be at least 3600 seconds. Omit on PATCH to preserve retention.
	EverySeconds *int64 `json:"everySeconds,omitempty"`
	// 0 selects the native default. Omit on PATCH to preserve the existing duration.
	ShardGroupDurationSeconds *int64 `json:"shardGroupDurationSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBBucketRetentionRule instantiates a new InfluxDBBucketRetentionRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBBucketRetentionRule() *InfluxDBBucketRetentionRule {
	this := InfluxDBBucketRetentionRule{}
	return &this
}

// NewInfluxDBBucketRetentionRuleWithDefaults instantiates a new InfluxDBBucketRetentionRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBBucketRetentionRuleWithDefaults() *InfluxDBBucketRetentionRule {
	this := InfluxDBBucketRetentionRule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InfluxDBBucketRetentionRule) GetType() InfluxDBRetentionRuleType {
	if o == nil || o.Type == nil {
		var ret InfluxDBRetentionRuleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketRetentionRule) GetTypeOk() (*InfluxDBRetentionRuleType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InfluxDBBucketRetentionRule) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given InfluxDBRetentionRuleType and assigns it to the Type field.
func (o *InfluxDBBucketRetentionRule) SetType(v InfluxDBRetentionRuleType) {
	o.Type = &v
}

// GetEverySeconds returns the EverySeconds field value if set, zero value otherwise.
func (o *InfluxDBBucketRetentionRule) GetEverySeconds() int64 {
	if o == nil || o.EverySeconds == nil {
		var ret int64
		return ret
	}
	return *o.EverySeconds
}

// GetEverySecondsOk returns a tuple with the EverySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketRetentionRule) GetEverySecondsOk() (*int64, bool) {
	if o == nil || o.EverySeconds == nil {
		return nil, false
	}
	return o.EverySeconds, true
}

// HasEverySeconds returns a boolean if a field has been set.
func (o *InfluxDBBucketRetentionRule) HasEverySeconds() bool {
	return o != nil && o.EverySeconds != nil
}

// SetEverySeconds gets a reference to the given int64 and assigns it to the EverySeconds field.
func (o *InfluxDBBucketRetentionRule) SetEverySeconds(v int64) {
	o.EverySeconds = &v
}

// GetShardGroupDurationSeconds returns the ShardGroupDurationSeconds field value if set, zero value otherwise.
func (o *InfluxDBBucketRetentionRule) GetShardGroupDurationSeconds() int64 {
	if o == nil || o.ShardGroupDurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.ShardGroupDurationSeconds
}

// GetShardGroupDurationSecondsOk returns a tuple with the ShardGroupDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketRetentionRule) GetShardGroupDurationSecondsOk() (*int64, bool) {
	if o == nil || o.ShardGroupDurationSeconds == nil {
		return nil, false
	}
	return o.ShardGroupDurationSeconds, true
}

// HasShardGroupDurationSeconds returns a boolean if a field has been set.
func (o *InfluxDBBucketRetentionRule) HasShardGroupDurationSeconds() bool {
	return o != nil && o.ShardGroupDurationSeconds != nil
}

// SetShardGroupDurationSeconds gets a reference to the given int64 and assigns it to the ShardGroupDurationSeconds field.
func (o *InfluxDBBucketRetentionRule) SetShardGroupDurationSeconds(v int64) {
	o.ShardGroupDurationSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBBucketRetentionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.EverySeconds != nil {
		toSerialize["everySeconds"] = o.EverySeconds
	}
	if o.ShardGroupDurationSeconds != nil {
		toSerialize["shardGroupDurationSeconds"] = o.ShardGroupDurationSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBBucketRetentionRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type                      *InfluxDBRetentionRuleType `json:"type,omitempty"`
		EverySeconds              *int64                     `json:"everySeconds,omitempty"`
		ShardGroupDurationSeconds *int64                     `json:"shardGroupDurationSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "everySeconds", "shardGroupDurationSeconds"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.EverySeconds = all.EverySeconds
	o.ShardGroupDurationSeconds = all.ShardGroupDurationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
