// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchRecoveryList struct {
	Recoveries  []ElasticsearchShardRecovery `json:"recoveries"`
	CollectedAt string                       `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchRecoveryList instantiates a new ElasticsearchRecoveryList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchRecoveryList(recoveries []ElasticsearchShardRecovery, collectedAt string) *ElasticsearchRecoveryList {
	this := ElasticsearchRecoveryList{}
	this.Recoveries = recoveries
	this.CollectedAt = collectedAt
	return &this
}

// NewElasticsearchRecoveryListWithDefaults instantiates a new ElasticsearchRecoveryList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchRecoveryListWithDefaults() *ElasticsearchRecoveryList {
	this := ElasticsearchRecoveryList{}
	return &this
}

// GetRecoveries returns the Recoveries field value.
func (o *ElasticsearchRecoveryList) GetRecoveries() []ElasticsearchShardRecovery {
	if o == nil {
		var ret []ElasticsearchShardRecovery
		return ret
	}
	return o.Recoveries
}

// GetRecoveriesOk returns a tuple with the Recoveries field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchRecoveryList) GetRecoveriesOk() (*[]ElasticsearchShardRecovery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recoveries, true
}

// SetRecoveries sets field value.
func (o *ElasticsearchRecoveryList) SetRecoveries(v []ElasticsearchShardRecovery) {
	o.Recoveries = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ElasticsearchRecoveryList) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchRecoveryList) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ElasticsearchRecoveryList) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchRecoveryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["recoveries"] = o.Recoveries
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchRecoveryList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Recoveries  *[]ElasticsearchShardRecovery `json:"recoveries"`
		CollectedAt *string                       `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Recoveries == nil {
		return fmt.Errorf("required field recoveries missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"recoveries", "collectedAt"})
	} else {
		return err
	}
	o.Recoveries = *all.Recoveries
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
