// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassNodeStatsList storageClassNodeStatsList is a list of storageClassNodeStats.
type StorageClassNodeStatsList struct {
	// Deduplicated apelocal statistics. Does not include other storage drivers. Missing measurements are omitted, never treated as zero.
	Summary *LocalStorageStats `json:"summary,omitempty"`
	// Resolved apelocal HostPathDefinition name.
	DefinitionName *string `json:"definitionName,omitempty"`
	// apelocal storage backend (HostDir, RawDisk or RawDiskGroup).
	BackendType *string `json:"backendType,omitempty"`
	// CSI capacity policy; unlimited does not represent physical capacity.
	CapacityMode *StorageCapacityMode `json:"capacityMode,omitempty"`
	// Capacity multiplier (1.2 means 120 percent).
	CapacityMultiplier *float64 `json:"capacityMultiplier,omitempty"`
	// Configured fixed capacity per node in GiB.
	FixedCapacity *float64 `json:"fixedCapacity,omitempty"`
	// Whether the metrics query succeeded; individual measurements can still be absent.
	MetricsAvailable *bool `json:"metricsAvailable,omitempty"`
	// the list of storage class node stats
	Items []StorageClassNodeStats `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassNodeStatsList instantiates a new StorageClassNodeStatsList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassNodeStatsList(items []StorageClassNodeStats) *StorageClassNodeStatsList {
	this := StorageClassNodeStatsList{}
	this.Items = items
	return &this
}

// NewStorageClassNodeStatsListWithDefaults instantiates a new StorageClassNodeStatsList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassNodeStatsListWithDefaults() *StorageClassNodeStatsList {
	this := StorageClassNodeStatsList{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetSummary() LocalStorageStats {
	if o == nil || o.Summary == nil {
		var ret LocalStorageStats
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetSummaryOk() (*LocalStorageStats, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given LocalStorageStats and assigns it to the Summary field.
func (o *StorageClassNodeStatsList) SetSummary(v LocalStorageStats) {
	o.Summary = &v
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetDefinitionName() string {
	if o == nil || o.DefinitionName == nil {
		var ret string
		return ret
	}
	return *o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetDefinitionNameOk() (*string, bool) {
	if o == nil || o.DefinitionName == nil {
		return nil, false
	}
	return o.DefinitionName, true
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasDefinitionName() bool {
	return o != nil && o.DefinitionName != nil
}

// SetDefinitionName gets a reference to the given string and assigns it to the DefinitionName field.
func (o *StorageClassNodeStatsList) SetDefinitionName(v string) {
	o.DefinitionName = &v
}

// GetBackendType returns the BackendType field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetBackendType() string {
	if o == nil || o.BackendType == nil {
		var ret string
		return ret
	}
	return *o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetBackendTypeOk() (*string, bool) {
	if o == nil || o.BackendType == nil {
		return nil, false
	}
	return o.BackendType, true
}

// HasBackendType returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasBackendType() bool {
	return o != nil && o.BackendType != nil
}

// SetBackendType gets a reference to the given string and assigns it to the BackendType field.
func (o *StorageClassNodeStatsList) SetBackendType(v string) {
	o.BackendType = &v
}

// GetCapacityMode returns the CapacityMode field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetCapacityMode() StorageCapacityMode {
	if o == nil || o.CapacityMode == nil {
		var ret StorageCapacityMode
		return ret
	}
	return *o.CapacityMode
}

// GetCapacityModeOk returns a tuple with the CapacityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetCapacityModeOk() (*StorageCapacityMode, bool) {
	if o == nil || o.CapacityMode == nil {
		return nil, false
	}
	return o.CapacityMode, true
}

// HasCapacityMode returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasCapacityMode() bool {
	return o != nil && o.CapacityMode != nil
}

// SetCapacityMode gets a reference to the given StorageCapacityMode and assigns it to the CapacityMode field.
func (o *StorageClassNodeStatsList) SetCapacityMode(v StorageCapacityMode) {
	o.CapacityMode = &v
}

// GetCapacityMultiplier returns the CapacityMultiplier field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetCapacityMultiplier() float64 {
	if o == nil || o.CapacityMultiplier == nil {
		var ret float64
		return ret
	}
	return *o.CapacityMultiplier
}

// GetCapacityMultiplierOk returns a tuple with the CapacityMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetCapacityMultiplierOk() (*float64, bool) {
	if o == nil || o.CapacityMultiplier == nil {
		return nil, false
	}
	return o.CapacityMultiplier, true
}

// HasCapacityMultiplier returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasCapacityMultiplier() bool {
	return o != nil && o.CapacityMultiplier != nil
}

// SetCapacityMultiplier gets a reference to the given float64 and assigns it to the CapacityMultiplier field.
func (o *StorageClassNodeStatsList) SetCapacityMultiplier(v float64) {
	o.CapacityMultiplier = &v
}

// GetFixedCapacity returns the FixedCapacity field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetFixedCapacity() float64 {
	if o == nil || o.FixedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.FixedCapacity
}

// GetFixedCapacityOk returns a tuple with the FixedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetFixedCapacityOk() (*float64, bool) {
	if o == nil || o.FixedCapacity == nil {
		return nil, false
	}
	return o.FixedCapacity, true
}

// HasFixedCapacity returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasFixedCapacity() bool {
	return o != nil && o.FixedCapacity != nil
}

// SetFixedCapacity gets a reference to the given float64 and assigns it to the FixedCapacity field.
func (o *StorageClassNodeStatsList) SetFixedCapacity(v float64) {
	o.FixedCapacity = &v
}

// GetMetricsAvailable returns the MetricsAvailable field value if set, zero value otherwise.
func (o *StorageClassNodeStatsList) GetMetricsAvailable() bool {
	if o == nil || o.MetricsAvailable == nil {
		var ret bool
		return ret
	}
	return *o.MetricsAvailable
}

// GetMetricsAvailableOk returns a tuple with the MetricsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetMetricsAvailableOk() (*bool, bool) {
	if o == nil || o.MetricsAvailable == nil {
		return nil, false
	}
	return o.MetricsAvailable, true
}

// HasMetricsAvailable returns a boolean if a field has been set.
func (o *StorageClassNodeStatsList) HasMetricsAvailable() bool {
	return o != nil && o.MetricsAvailable != nil
}

// SetMetricsAvailable gets a reference to the given bool and assigns it to the MetricsAvailable field.
func (o *StorageClassNodeStatsList) SetMetricsAvailable(v bool) {
	o.MetricsAvailable = &v
}

// GetItems returns the Items field value.
func (o *StorageClassNodeStatsList) GetItems() []StorageClassNodeStats {
	if o == nil {
		var ret []StorageClassNodeStats
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStatsList) GetItemsOk() (*[]StorageClassNodeStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *StorageClassNodeStatsList) SetItems(v []StorageClassNodeStats) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassNodeStatsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.DefinitionName != nil {
		toSerialize["definitionName"] = o.DefinitionName
	}
	if o.BackendType != nil {
		toSerialize["backendType"] = o.BackendType
	}
	if o.CapacityMode != nil {
		toSerialize["capacityMode"] = o.CapacityMode
	}
	if o.CapacityMultiplier != nil {
		toSerialize["capacityMultiplier"] = o.CapacityMultiplier
	}
	if o.FixedCapacity != nil {
		toSerialize["fixedCapacity"] = o.FixedCapacity
	}
	if o.MetricsAvailable != nil {
		toSerialize["metricsAvailable"] = o.MetricsAvailable
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassNodeStatsList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary            *LocalStorageStats       `json:"summary,omitempty"`
		DefinitionName     *string                  `json:"definitionName,omitempty"`
		BackendType        *string                  `json:"backendType,omitempty"`
		CapacityMode       *StorageCapacityMode     `json:"capacityMode,omitempty"`
		CapacityMultiplier *float64                 `json:"capacityMultiplier,omitempty"`
		FixedCapacity      *float64                 `json:"fixedCapacity,omitempty"`
		MetricsAvailable   *bool                    `json:"metricsAvailable,omitempty"`
		Items              *[]StorageClassNodeStats `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "definitionName", "backendType", "capacityMode", "capacityMultiplier", "fixedCapacity", "metricsAvailable", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.DefinitionName = all.DefinitionName
	o.BackendType = all.BackendType
	if all.CapacityMode != nil && !all.CapacityMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CapacityMode = all.CapacityMode
	}
	o.CapacityMultiplier = all.CapacityMultiplier
	o.FixedCapacity = all.FixedCapacity
	o.MetricsAvailable = all.MetricsAvailable
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
