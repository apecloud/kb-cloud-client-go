// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlStorageInstanceUsage struct {
	InstanceName   *string                `json:"instanceName,omitempty"`
	PvcName        *string                `json:"pvcName,omitempty"`
	Role           *string                `json:"role,omitempty"`
	ComponentName  *string                `json:"componentName,omitempty"`
	TotalBytes     *int64                 `json:"totalBytes,omitempty"`
	UsedBytes      *int64                 `json:"usedBytes,omitempty"`
	AvailableBytes *int64                 `json:"availableBytes,omitempty"`
	UsageRatio     common.NullableFloat64 `json:"usageRatio,omitempty"`
	UpdatedAt      *string                `json:"updatedAt,omitempty"`
	Source         *string                `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlStorageInstanceUsage instantiates a new MysqlStorageInstanceUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlStorageInstanceUsage() *MysqlStorageInstanceUsage {
	this := MysqlStorageInstanceUsage{}
	return &this
}

// NewMysqlStorageInstanceUsageWithDefaults instantiates a new MysqlStorageInstanceUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlStorageInstanceUsageWithDefaults() *MysqlStorageInstanceUsage {
	this := MysqlStorageInstanceUsage{}
	return &this
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *MysqlStorageInstanceUsage) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetPvcName returns the PvcName field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetPvcName() string {
	if o == nil || o.PvcName == nil {
		var ret string
		return ret
	}
	return *o.PvcName
}

// GetPvcNameOk returns a tuple with the PvcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetPvcNameOk() (*string, bool) {
	if o == nil || o.PvcName == nil {
		return nil, false
	}
	return o.PvcName, true
}

// HasPvcName returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasPvcName() bool {
	return o != nil && o.PvcName != nil
}

// SetPvcName gets a reference to the given string and assigns it to the PvcName field.
func (o *MysqlStorageInstanceUsage) SetPvcName(v string) {
	o.PvcName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MysqlStorageInstanceUsage) SetRole(v string) {
	o.Role = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *MysqlStorageInstanceUsage) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *MysqlStorageInstanceUsage) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetUsedBytesOk() (*int64, bool) {
	if o == nil || o.UsedBytes == nil {
		return nil, false
	}
	return o.UsedBytes, true
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasUsedBytes() bool {
	return o != nil && o.UsedBytes != nil
}

// SetUsedBytes gets a reference to the given int64 and assigns it to the UsedBytes field.
func (o *MysqlStorageInstanceUsage) SetUsedBytes(v int64) {
	o.UsedBytes = &v
}

// GetAvailableBytes returns the AvailableBytes field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetAvailableBytes() int64 {
	if o == nil || o.AvailableBytes == nil {
		var ret int64
		return ret
	}
	return *o.AvailableBytes
}

// GetAvailableBytesOk returns a tuple with the AvailableBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetAvailableBytesOk() (*int64, bool) {
	if o == nil || o.AvailableBytes == nil {
		return nil, false
	}
	return o.AvailableBytes, true
}

// HasAvailableBytes returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasAvailableBytes() bool {
	return o != nil && o.AvailableBytes != nil
}

// SetAvailableBytes gets a reference to the given int64 and assigns it to the AvailableBytes field.
func (o *MysqlStorageInstanceUsage) SetAvailableBytes(v int64) {
	o.AvailableBytes = &v
}

// GetUsageRatio returns the UsageRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MysqlStorageInstanceUsage) GetUsageRatio() float64 {
	if o == nil || o.UsageRatio.Get() == nil {
		var ret float64
		return ret
	}
	return *o.UsageRatio.Get()
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MysqlStorageInstanceUsage) GetUsageRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageRatio.Get(), o.UsageRatio.IsSet()
}

// HasUsageRatio returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasUsageRatio() bool {
	return o != nil && o.UsageRatio.IsSet()
}

// SetUsageRatio gets a reference to the given common.NullableFloat64 and assigns it to the UsageRatio field.
func (o *MysqlStorageInstanceUsage) SetUsageRatio(v float64) {
	o.UsageRatio.Set(&v)
}

// SetUsageRatioNil sets the value for UsageRatio to be an explicit nil.
func (o *MysqlStorageInstanceUsage) SetUsageRatioNil() {
	o.UsageRatio.Set(nil)
}

// UnsetUsageRatio ensures that no value is present for UsageRatio, not even an explicit nil.
func (o *MysqlStorageInstanceUsage) UnsetUsageRatio() {
	o.UsageRatio.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MysqlStorageInstanceUsage) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MysqlStorageInstanceUsage) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStorageInstanceUsage) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MysqlStorageInstanceUsage) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MysqlStorageInstanceUsage) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlStorageInstanceUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.PvcName != nil {
		toSerialize["pvcName"] = o.PvcName
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.UsedBytes != nil {
		toSerialize["usedBytes"] = o.UsedBytes
	}
	if o.AvailableBytes != nil {
		toSerialize["availableBytes"] = o.AvailableBytes
	}
	if o.UsageRatio.IsSet() {
		toSerialize["usageRatio"] = o.UsageRatio.Get()
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlStorageInstanceUsage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName   *string                `json:"instanceName,omitempty"`
		PvcName        *string                `json:"pvcName,omitempty"`
		Role           *string                `json:"role,omitempty"`
		ComponentName  *string                `json:"componentName,omitempty"`
		TotalBytes     *int64                 `json:"totalBytes,omitempty"`
		UsedBytes      *int64                 `json:"usedBytes,omitempty"`
		AvailableBytes *int64                 `json:"availableBytes,omitempty"`
		UsageRatio     common.NullableFloat64 `json:"usageRatio,omitempty"`
		UpdatedAt      *string                `json:"updatedAt,omitempty"`
		Source         *string                `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "pvcName", "role", "componentName", "totalBytes", "usedBytes", "availableBytes", "usageRatio", "updatedAt", "source"})
	} else {
		return err
	}
	o.InstanceName = all.InstanceName
	o.PvcName = all.PvcName
	o.Role = all.Role
	o.ComponentName = all.ComponentName
	o.TotalBytes = all.TotalBytes
	o.UsedBytes = all.UsedBytes
	o.AvailableBytes = all.AvailableBytes
	o.UsageRatio = all.UsageRatio
	o.UpdatedAt = all.UpdatedAt
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
