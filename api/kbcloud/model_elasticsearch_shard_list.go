// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchShardList struct {
	Shards      []ElasticsearchShard `json:"shards"`
	CollectedAt string               `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchShardList instantiates a new ElasticsearchShardList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchShardList(shards []ElasticsearchShard, collectedAt string) *ElasticsearchShardList {
	this := ElasticsearchShardList{}
	this.Shards = shards
	this.CollectedAt = collectedAt
	return &this
}

// NewElasticsearchShardListWithDefaults instantiates a new ElasticsearchShardList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchShardListWithDefaults() *ElasticsearchShardList {
	this := ElasticsearchShardList{}
	return &this
}

// GetShards returns the Shards field value.
func (o *ElasticsearchShardList) GetShards() []ElasticsearchShard {
	if o == nil {
		var ret []ElasticsearchShard
		return ret
	}
	return o.Shards
}

// GetShardsOk returns a tuple with the Shards field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardList) GetShardsOk() (*[]ElasticsearchShard, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shards, true
}

// SetShards sets field value.
func (o *ElasticsearchShardList) SetShards(v []ElasticsearchShard) {
	o.Shards = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ElasticsearchShardList) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchShardList) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ElasticsearchShardList) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchShardList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["shards"] = o.Shards
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchShardList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Shards      *[]ElasticsearchShard `json:"shards"`
		CollectedAt *string               `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Shards == nil {
		return fmt.Errorf("required field shards missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"shards", "collectedAt"})
	} else {
		return err
	}
	o.Shards = *all.Shards
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
