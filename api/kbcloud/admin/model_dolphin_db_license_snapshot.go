// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBLicenseSnapshot struct {
	CheckedAt   time.Time              `json:"checkedAt"`
	Inventory   DolphinDBHealthCheck   `json:"inventory"`
	Consistency DolphinDBHealthCheck   `json:"consistency"`
	Nodes       []DolphinDBNodeLicense `json:"nodes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBLicenseSnapshot instantiates a new DolphinDBLicenseSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBLicenseSnapshot(checkedAt time.Time, inventory DolphinDBHealthCheck, consistency DolphinDBHealthCheck, nodes []DolphinDBNodeLicense) *DolphinDBLicenseSnapshot {
	this := DolphinDBLicenseSnapshot{}
	this.CheckedAt = checkedAt
	this.Inventory = inventory
	this.Consistency = consistency
	this.Nodes = nodes
	return &this
}

// NewDolphinDBLicenseSnapshotWithDefaults instantiates a new DolphinDBLicenseSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBLicenseSnapshotWithDefaults() *DolphinDBLicenseSnapshot {
	this := DolphinDBLicenseSnapshot{}
	return &this
}

// GetCheckedAt returns the CheckedAt field value.
func (o *DolphinDBLicenseSnapshot) GetCheckedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseSnapshot) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckedAt, true
}

// SetCheckedAt sets field value.
func (o *DolphinDBLicenseSnapshot) SetCheckedAt(v time.Time) {
	o.CheckedAt = v
}

// GetInventory returns the Inventory field value.
func (o *DolphinDBLicenseSnapshot) GetInventory() DolphinDBHealthCheck {
	if o == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseSnapshot) GetInventoryOk() (*DolphinDBHealthCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inventory, true
}

// SetInventory sets field value.
func (o *DolphinDBLicenseSnapshot) SetInventory(v DolphinDBHealthCheck) {
	o.Inventory = v
}

// GetConsistency returns the Consistency field value.
func (o *DolphinDBLicenseSnapshot) GetConsistency() DolphinDBHealthCheck {
	if o == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseSnapshot) GetConsistencyOk() (*DolphinDBHealthCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consistency, true
}

// SetConsistency sets field value.
func (o *DolphinDBLicenseSnapshot) SetConsistency(v DolphinDBHealthCheck) {
	o.Consistency = v
}

// GetNodes returns the Nodes field value.
func (o *DolphinDBLicenseSnapshot) GetNodes() []DolphinDBNodeLicense {
	if o == nil {
		var ret []DolphinDBNodeLicense
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseSnapshot) GetNodesOk() (*[]DolphinDBNodeLicense, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *DolphinDBLicenseSnapshot) SetNodes(v []DolphinDBNodeLicense) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBLicenseSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CheckedAt.Nanosecond() == 0 {
		toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["inventory"] = o.Inventory
	toSerialize["consistency"] = o.Consistency
	toSerialize["nodes"] = o.Nodes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBLicenseSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckedAt   *time.Time              `json:"checkedAt"`
		Inventory   *DolphinDBHealthCheck   `json:"inventory"`
		Consistency *DolphinDBHealthCheck   `json:"consistency"`
		Nodes       *[]DolphinDBNodeLicense `json:"nodes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CheckedAt == nil {
		return fmt.Errorf("required field checkedAt missing")
	}
	if all.Inventory == nil {
		return fmt.Errorf("required field inventory missing")
	}
	if all.Consistency == nil {
		return fmt.Errorf("required field consistency missing")
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkedAt", "inventory", "consistency", "nodes"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckedAt = *all.CheckedAt
	if all.Inventory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Inventory = *all.Inventory
	if all.Consistency.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Consistency = *all.Consistency
	o.Nodes = *all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
