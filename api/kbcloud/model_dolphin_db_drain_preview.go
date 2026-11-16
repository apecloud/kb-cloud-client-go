// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBDrainPreview struct {
	ObservedAt     time.Time                 `json:"observedAt"`
	Source         string                    `json:"source"`
	RemainingNodes []string                  `json:"remainingNodes"`
	Total          int32                     `json:"total"`
	Offset         int32                     `json:"offset"`
	Entries        []DolphinDBDrainMigration `json:"entries"`
	// Always notValidated. This preview is not proof of sufficient physical capacity or permission to delete the node.
	CapacityCheck string `json:"capacityCheck"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBDrainPreview instantiates a new DolphinDBDrainPreview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBDrainPreview(observedAt time.Time, source string, remainingNodes []string, total int32, offset int32, entries []DolphinDBDrainMigration, capacityCheck string) *DolphinDBDrainPreview {
	this := DolphinDBDrainPreview{}
	this.ObservedAt = observedAt
	this.Source = source
	this.RemainingNodes = remainingNodes
	this.Total = total
	this.Offset = offset
	this.Entries = entries
	this.CapacityCheck = capacityCheck
	return &this
}

// NewDolphinDBDrainPreviewWithDefaults instantiates a new DolphinDBDrainPreview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBDrainPreviewWithDefaults() *DolphinDBDrainPreview {
	this := DolphinDBDrainPreview{}
	return &this
}

// GetObservedAt returns the ObservedAt field value.
func (o *DolphinDBDrainPreview) GetObservedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetObservedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedAt, true
}

// SetObservedAt sets field value.
func (o *DolphinDBDrainPreview) SetObservedAt(v time.Time) {
	o.ObservedAt = v
}

// GetSource returns the Source field value.
func (o *DolphinDBDrainPreview) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DolphinDBDrainPreview) SetSource(v string) {
	o.Source = v
}

// GetRemainingNodes returns the RemainingNodes field value.
func (o *DolphinDBDrainPreview) GetRemainingNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RemainingNodes
}

// GetRemainingNodesOk returns a tuple with the RemainingNodes field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetRemainingNodesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingNodes, true
}

// SetRemainingNodes sets field value.
func (o *DolphinDBDrainPreview) SetRemainingNodes(v []string) {
	o.RemainingNodes = v
}

// GetTotal returns the Total field value.
func (o *DolphinDBDrainPreview) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *DolphinDBDrainPreview) SetTotal(v int32) {
	o.Total = v
}

// GetOffset returns the Offset field value.
func (o *DolphinDBDrainPreview) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *DolphinDBDrainPreview) SetOffset(v int32) {
	o.Offset = v
}

// GetEntries returns the Entries field value.
func (o *DolphinDBDrainPreview) GetEntries() []DolphinDBDrainMigration {
	if o == nil {
		var ret []DolphinDBDrainMigration
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetEntriesOk() (*[]DolphinDBDrainMigration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entries, true
}

// SetEntries sets field value.
func (o *DolphinDBDrainPreview) SetEntries(v []DolphinDBDrainMigration) {
	o.Entries = v
}

// GetCapacityCheck returns the CapacityCheck field value.
func (o *DolphinDBDrainPreview) GetCapacityCheck() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapacityCheck
}

// GetCapacityCheckOk returns a tuple with the CapacityCheck field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainPreview) GetCapacityCheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityCheck, true
}

// SetCapacityCheck sets field value.
func (o *DolphinDBDrainPreview) SetCapacityCheck(v string) {
	o.CapacityCheck = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBDrainPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObservedAt.Nanosecond() == 0 {
		toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["source"] = o.Source
	toSerialize["remainingNodes"] = o.RemainingNodes
	toSerialize["total"] = o.Total
	toSerialize["offset"] = o.Offset
	toSerialize["entries"] = o.Entries
	toSerialize["capacityCheck"] = o.CapacityCheck

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBDrainPreview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObservedAt     *time.Time                 `json:"observedAt"`
		Source         *string                    `json:"source"`
		RemainingNodes *[]string                  `json:"remainingNodes"`
		Total          *int32                     `json:"total"`
		Offset         *int32                     `json:"offset"`
		Entries        *[]DolphinDBDrainMigration `json:"entries"`
		CapacityCheck  *string                    `json:"capacityCheck"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ObservedAt == nil {
		return fmt.Errorf("required field observedAt missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.RemainingNodes == nil {
		return fmt.Errorf("required field remainingNodes missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Offset == nil {
		return fmt.Errorf("required field offset missing")
	}
	if all.Entries == nil {
		return fmt.Errorf("required field entries missing")
	}
	if all.CapacityCheck == nil {
		return fmt.Errorf("required field capacityCheck missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"observedAt", "source", "remainingNodes", "total", "offset", "entries", "capacityCheck"})
	} else {
		return err
	}
	o.ObservedAt = *all.ObservedAt
	o.Source = *all.Source
	o.RemainingNodes = *all.RemainingNodes
	o.Total = *all.Total
	o.Offset = *all.Offset
	o.Entries = *all.Entries
	o.CapacityCheck = *all.CapacityCheck

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
