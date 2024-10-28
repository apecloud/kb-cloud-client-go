// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NodeResourceStatsList NodeResourceStatsList is a list of NodeResourceStats

type NodeResourceStatsList struct {
	// Items is the list of NodeResourceStats objects in the list
	Items []NodeResourceStats `json:"items"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeResourceStatsList instantiates a new NodeResourceStatsList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeResourceStatsList(items []NodeResourceStats, updatedAt time.Time) *NodeResourceStatsList {
	this := NodeResourceStatsList{}
	this.Items = items
	this.UpdatedAt = updatedAt
	return &this
}

// NewNodeResourceStatsListWithDefaults instantiates a new NodeResourceStatsList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeResourceStatsListWithDefaults() *NodeResourceStatsList {
	this := NodeResourceStatsList{}
	return &this
}

// GetItems returns the Items field value.
func (o *NodeResourceStatsList) GetItems() []NodeResourceStats {
	if o == nil {
		var ret []NodeResourceStats
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NodeResourceStatsList) GetItemsOk() (*[]NodeResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *NodeResourceStatsList) SetItems(v []NodeResourceStats) {
	o.Items = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *NodeResourceStatsList) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NodeResourceStatsList) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *NodeResourceStatsList) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeResourceStatsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeResourceStatsList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items     *[]NodeResourceStats `json:"items"`
		UpdatedAt *time.Time           `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "updatedAt"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
