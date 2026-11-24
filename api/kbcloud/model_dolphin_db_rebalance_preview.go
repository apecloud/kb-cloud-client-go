// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBRebalancePreview struct {
	ObservedAt        *time.Time           `json:"observedAt,omitempty"`
	UpdatedBeforeDays *float64             `json:"updatedBeforeDays,omitempty"`
	Total             *int64               `json:"total,omitempty"`
	Offset            *int64               `json:"offset,omitempty"`
	Entries           []DolphinDBMigration `json:"entries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBRebalancePreview instantiates a new DolphinDBRebalancePreview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBRebalancePreview() *DolphinDBRebalancePreview {
	this := DolphinDBRebalancePreview{}
	return &this
}

// NewDolphinDBRebalancePreviewWithDefaults instantiates a new DolphinDBRebalancePreview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBRebalancePreviewWithDefaults() *DolphinDBRebalancePreview {
	this := DolphinDBRebalancePreview{}
	return &this
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *DolphinDBRebalancePreview) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalancePreview) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *DolphinDBRebalancePreview) HasObservedAt() bool {
	return o != nil && o.ObservedAt != nil
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *DolphinDBRebalancePreview) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// GetUpdatedBeforeDays returns the UpdatedBeforeDays field value if set, zero value otherwise.
func (o *DolphinDBRebalancePreview) GetUpdatedBeforeDays() float64 {
	if o == nil || o.UpdatedBeforeDays == nil {
		var ret float64
		return ret
	}
	return *o.UpdatedBeforeDays
}

// GetUpdatedBeforeDaysOk returns a tuple with the UpdatedBeforeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalancePreview) GetUpdatedBeforeDaysOk() (*float64, bool) {
	if o == nil || o.UpdatedBeforeDays == nil {
		return nil, false
	}
	return o.UpdatedBeforeDays, true
}

// HasUpdatedBeforeDays returns a boolean if a field has been set.
func (o *DolphinDBRebalancePreview) HasUpdatedBeforeDays() bool {
	return o != nil && o.UpdatedBeforeDays != nil
}

// SetUpdatedBeforeDays gets a reference to the given float64 and assigns it to the UpdatedBeforeDays field.
func (o *DolphinDBRebalancePreview) SetUpdatedBeforeDays(v float64) {
	o.UpdatedBeforeDays = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DolphinDBRebalancePreview) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalancePreview) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DolphinDBRebalancePreview) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *DolphinDBRebalancePreview) SetTotal(v int64) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *DolphinDBRebalancePreview) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalancePreview) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *DolphinDBRebalancePreview) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *DolphinDBRebalancePreview) SetOffset(v int64) {
	o.Offset = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *DolphinDBRebalancePreview) GetEntries() []DolphinDBMigration {
	if o == nil || o.Entries == nil {
		var ret []DolphinDBMigration
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalancePreview) GetEntriesOk() (*[]DolphinDBMigration, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *DolphinDBRebalancePreview) HasEntries() bool {
	return o != nil && o.Entries != nil
}

// SetEntries gets a reference to the given []DolphinDBMigration and assigns it to the Entries field.
func (o *DolphinDBRebalancePreview) SetEntries(v []DolphinDBMigration) {
	o.Entries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBRebalancePreview) MarshalJSON() ([]byte, error) {
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
	if o.UpdatedBeforeDays != nil {
		toSerialize["updatedBeforeDays"] = o.UpdatedBeforeDays
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBRebalancePreview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObservedAt        *time.Time           `json:"observedAt,omitempty"`
		UpdatedBeforeDays *float64             `json:"updatedBeforeDays,omitempty"`
		Total             *int64               `json:"total,omitempty"`
		Offset            *int64               `json:"offset,omitempty"`
		Entries           []DolphinDBMigration `json:"entries,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"observedAt", "updatedBeforeDays", "total", "offset", "entries"})
	} else {
		return err
	}
	o.ObservedAt = all.ObservedAt
	o.UpdatedBeforeDays = all.UpdatedBeforeDays
	o.Total = all.Total
	o.Offset = all.Offset
	o.Entries = all.Entries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
