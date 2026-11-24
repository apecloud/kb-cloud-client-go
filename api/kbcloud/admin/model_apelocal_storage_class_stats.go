// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ApelocalStorageClassStats apelocal-specific StorageClass capacity policy and measurements.
type ApelocalStorageClassStats struct {
	// Deduplicated apelocal statistics. Does not include other storage drivers. Missing measurements are omitted, never treated as zero.
	Summary *ApelocalCapacitySummary `json:"summary,omitempty"`
	// Resolved apelocal HostPathDefinition name.
	DefinitionName *string `json:"definitionName,omitempty"`
	// apelocal storage backend (HostDir, RawDisk or RawDiskGroup).
	BackendType *string `json:"backendType,omitempty"`
	// CSI capacity policy; unlimited does not represent physical capacity.
	CapacityMode *ApelocalCapacityMode `json:"capacityMode,omitempty"`
	// Capacity multiplier (1.2 means 120 percent).
	CapacityMultiplier *float64 `json:"capacityMultiplier,omitempty"`
	// Configured fixed capacity per node in GiB.
	FixedCapacity *float64 `json:"fixedCapacity,omitempty"`
	// Whether the metrics query succeeded; individual measurements can still be absent.
	MetricsAvailable *bool `json:"metricsAvailable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApelocalStorageClassStats instantiates a new ApelocalStorageClassStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApelocalStorageClassStats() *ApelocalStorageClassStats {
	this := ApelocalStorageClassStats{}
	return &this
}

// NewApelocalStorageClassStatsWithDefaults instantiates a new ApelocalStorageClassStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApelocalStorageClassStatsWithDefaults() *ApelocalStorageClassStats {
	this := ApelocalStorageClassStats{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetSummary() ApelocalCapacitySummary {
	if o == nil || o.Summary == nil {
		var ret ApelocalCapacitySummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetSummaryOk() (*ApelocalCapacitySummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ApelocalCapacitySummary and assigns it to the Summary field.
func (o *ApelocalStorageClassStats) SetSummary(v ApelocalCapacitySummary) {
	o.Summary = &v
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetDefinitionName() string {
	if o == nil || o.DefinitionName == nil {
		var ret string
		return ret
	}
	return *o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetDefinitionNameOk() (*string, bool) {
	if o == nil || o.DefinitionName == nil {
		return nil, false
	}
	return o.DefinitionName, true
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasDefinitionName() bool {
	return o != nil && o.DefinitionName != nil
}

// SetDefinitionName gets a reference to the given string and assigns it to the DefinitionName field.
func (o *ApelocalStorageClassStats) SetDefinitionName(v string) {
	o.DefinitionName = &v
}

// GetBackendType returns the BackendType field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetBackendType() string {
	if o == nil || o.BackendType == nil {
		var ret string
		return ret
	}
	return *o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetBackendTypeOk() (*string, bool) {
	if o == nil || o.BackendType == nil {
		return nil, false
	}
	return o.BackendType, true
}

// HasBackendType returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasBackendType() bool {
	return o != nil && o.BackendType != nil
}

// SetBackendType gets a reference to the given string and assigns it to the BackendType field.
func (o *ApelocalStorageClassStats) SetBackendType(v string) {
	o.BackendType = &v
}

// GetCapacityMode returns the CapacityMode field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetCapacityMode() ApelocalCapacityMode {
	if o == nil || o.CapacityMode == nil {
		var ret ApelocalCapacityMode
		return ret
	}
	return *o.CapacityMode
}

// GetCapacityModeOk returns a tuple with the CapacityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetCapacityModeOk() (*ApelocalCapacityMode, bool) {
	if o == nil || o.CapacityMode == nil {
		return nil, false
	}
	return o.CapacityMode, true
}

// HasCapacityMode returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasCapacityMode() bool {
	return o != nil && o.CapacityMode != nil
}

// SetCapacityMode gets a reference to the given ApelocalCapacityMode and assigns it to the CapacityMode field.
func (o *ApelocalStorageClassStats) SetCapacityMode(v ApelocalCapacityMode) {
	o.CapacityMode = &v
}

// GetCapacityMultiplier returns the CapacityMultiplier field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetCapacityMultiplier() float64 {
	if o == nil || o.CapacityMultiplier == nil {
		var ret float64
		return ret
	}
	return *o.CapacityMultiplier
}

// GetCapacityMultiplierOk returns a tuple with the CapacityMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetCapacityMultiplierOk() (*float64, bool) {
	if o == nil || o.CapacityMultiplier == nil {
		return nil, false
	}
	return o.CapacityMultiplier, true
}

// HasCapacityMultiplier returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasCapacityMultiplier() bool {
	return o != nil && o.CapacityMultiplier != nil
}

// SetCapacityMultiplier gets a reference to the given float64 and assigns it to the CapacityMultiplier field.
func (o *ApelocalStorageClassStats) SetCapacityMultiplier(v float64) {
	o.CapacityMultiplier = &v
}

// GetFixedCapacity returns the FixedCapacity field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetFixedCapacity() float64 {
	if o == nil || o.FixedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.FixedCapacity
}

// GetFixedCapacityOk returns a tuple with the FixedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetFixedCapacityOk() (*float64, bool) {
	if o == nil || o.FixedCapacity == nil {
		return nil, false
	}
	return o.FixedCapacity, true
}

// HasFixedCapacity returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasFixedCapacity() bool {
	return o != nil && o.FixedCapacity != nil
}

// SetFixedCapacity gets a reference to the given float64 and assigns it to the FixedCapacity field.
func (o *ApelocalStorageClassStats) SetFixedCapacity(v float64) {
	o.FixedCapacity = &v
}

// GetMetricsAvailable returns the MetricsAvailable field value if set, zero value otherwise.
func (o *ApelocalStorageClassStats) GetMetricsAvailable() bool {
	if o == nil || o.MetricsAvailable == nil {
		var ret bool
		return ret
	}
	return *o.MetricsAvailable
}

// GetMetricsAvailableOk returns a tuple with the MetricsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalStorageClassStats) GetMetricsAvailableOk() (*bool, bool) {
	if o == nil || o.MetricsAvailable == nil {
		return nil, false
	}
	return o.MetricsAvailable, true
}

// HasMetricsAvailable returns a boolean if a field has been set.
func (o *ApelocalStorageClassStats) HasMetricsAvailable() bool {
	return o != nil && o.MetricsAvailable != nil
}

// SetMetricsAvailable gets a reference to the given bool and assigns it to the MetricsAvailable field.
func (o *ApelocalStorageClassStats) SetMetricsAvailable(v bool) {
	o.MetricsAvailable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApelocalStorageClassStats) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApelocalStorageClassStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary            *ApelocalCapacitySummary `json:"summary,omitempty"`
		DefinitionName     *string                  `json:"definitionName,omitempty"`
		BackendType        *string                  `json:"backendType,omitempty"`
		CapacityMode       *ApelocalCapacityMode    `json:"capacityMode,omitempty"`
		CapacityMultiplier *float64                 `json:"capacityMultiplier,omitempty"`
		FixedCapacity      *float64                 `json:"fixedCapacity,omitempty"`
		MetricsAvailable   *bool                    `json:"metricsAvailable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "definitionName", "backendType", "capacityMode", "capacityMultiplier", "fixedCapacity", "metricsAvailable"})
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
