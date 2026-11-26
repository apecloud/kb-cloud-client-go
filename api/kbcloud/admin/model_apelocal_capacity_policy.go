// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ApelocalCapacityPolicy Resolved apelocal StorageClass capacity policy populated by environment StorageClass list/get endpoints from the cluster HostPathDefinition. Omitted for non-apelocal storage classes, other endpoints, or when unavailable.
type ApelocalCapacityPolicy struct {
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
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApelocalCapacityPolicy instantiates a new ApelocalCapacityPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApelocalCapacityPolicy() *ApelocalCapacityPolicy {
	this := ApelocalCapacityPolicy{}
	return &this
}

// NewApelocalCapacityPolicyWithDefaults instantiates a new ApelocalCapacityPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApelocalCapacityPolicyWithDefaults() *ApelocalCapacityPolicy {
	this := ApelocalCapacityPolicy{}
	return &this
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise.
func (o *ApelocalCapacityPolicy) GetDefinitionName() string {
	if o == nil || o.DefinitionName == nil {
		var ret string
		return ret
	}
	return *o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityPolicy) GetDefinitionNameOk() (*string, bool) {
	if o == nil || o.DefinitionName == nil {
		return nil, false
	}
	return o.DefinitionName, true
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *ApelocalCapacityPolicy) HasDefinitionName() bool {
	return o != nil && o.DefinitionName != nil
}

// SetDefinitionName gets a reference to the given string and assigns it to the DefinitionName field.
func (o *ApelocalCapacityPolicy) SetDefinitionName(v string) {
	o.DefinitionName = &v
}

// GetBackendType returns the BackendType field value if set, zero value otherwise.
func (o *ApelocalCapacityPolicy) GetBackendType() string {
	if o == nil || o.BackendType == nil {
		var ret string
		return ret
	}
	return *o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityPolicy) GetBackendTypeOk() (*string, bool) {
	if o == nil || o.BackendType == nil {
		return nil, false
	}
	return o.BackendType, true
}

// HasBackendType returns a boolean if a field has been set.
func (o *ApelocalCapacityPolicy) HasBackendType() bool {
	return o != nil && o.BackendType != nil
}

// SetBackendType gets a reference to the given string and assigns it to the BackendType field.
func (o *ApelocalCapacityPolicy) SetBackendType(v string) {
	o.BackendType = &v
}

// GetCapacityMode returns the CapacityMode field value if set, zero value otherwise.
func (o *ApelocalCapacityPolicy) GetCapacityMode() ApelocalCapacityMode {
	if o == nil || o.CapacityMode == nil {
		var ret ApelocalCapacityMode
		return ret
	}
	return *o.CapacityMode
}

// GetCapacityModeOk returns a tuple with the CapacityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityPolicy) GetCapacityModeOk() (*ApelocalCapacityMode, bool) {
	if o == nil || o.CapacityMode == nil {
		return nil, false
	}
	return o.CapacityMode, true
}

// HasCapacityMode returns a boolean if a field has been set.
func (o *ApelocalCapacityPolicy) HasCapacityMode() bool {
	return o != nil && o.CapacityMode != nil
}

// SetCapacityMode gets a reference to the given ApelocalCapacityMode and assigns it to the CapacityMode field.
func (o *ApelocalCapacityPolicy) SetCapacityMode(v ApelocalCapacityMode) {
	o.CapacityMode = &v
}

// GetCapacityMultiplier returns the CapacityMultiplier field value if set, zero value otherwise.
func (o *ApelocalCapacityPolicy) GetCapacityMultiplier() float64 {
	if o == nil || o.CapacityMultiplier == nil {
		var ret float64
		return ret
	}
	return *o.CapacityMultiplier
}

// GetCapacityMultiplierOk returns a tuple with the CapacityMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityPolicy) GetCapacityMultiplierOk() (*float64, bool) {
	if o == nil || o.CapacityMultiplier == nil {
		return nil, false
	}
	return o.CapacityMultiplier, true
}

// HasCapacityMultiplier returns a boolean if a field has been set.
func (o *ApelocalCapacityPolicy) HasCapacityMultiplier() bool {
	return o != nil && o.CapacityMultiplier != nil
}

// SetCapacityMultiplier gets a reference to the given float64 and assigns it to the CapacityMultiplier field.
func (o *ApelocalCapacityPolicy) SetCapacityMultiplier(v float64) {
	o.CapacityMultiplier = &v
}

// GetFixedCapacity returns the FixedCapacity field value if set, zero value otherwise.
func (o *ApelocalCapacityPolicy) GetFixedCapacity() float64 {
	if o == nil || o.FixedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.FixedCapacity
}

// GetFixedCapacityOk returns a tuple with the FixedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityPolicy) GetFixedCapacityOk() (*float64, bool) {
	if o == nil || o.FixedCapacity == nil {
		return nil, false
	}
	return o.FixedCapacity, true
}

// HasFixedCapacity returns a boolean if a field has been set.
func (o *ApelocalCapacityPolicy) HasFixedCapacity() bool {
	return o != nil && o.FixedCapacity != nil
}

// SetFixedCapacity gets a reference to the given float64 and assigns it to the FixedCapacity field.
func (o *ApelocalCapacityPolicy) SetFixedCapacity(v float64) {
	o.FixedCapacity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApelocalCapacityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApelocalCapacityPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefinitionName     *string               `json:"definitionName,omitempty"`
		BackendType        *string               `json:"backendType,omitempty"`
		CapacityMode       *ApelocalCapacityMode `json:"capacityMode,omitempty"`
		CapacityMultiplier *float64              `json:"capacityMultiplier,omitempty"`
		FixedCapacity      *float64              `json:"fixedCapacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"definitionName", "backendType", "capacityMode", "capacityMultiplier", "fixedCapacity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefinitionName = all.DefinitionName
	o.BackendType = all.BackendType
	if all.CapacityMode != nil && !all.CapacityMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CapacityMode = all.CapacityMode
	}
	o.CapacityMultiplier = all.CapacityMultiplier
	o.FixedCapacity = all.FixedCapacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
