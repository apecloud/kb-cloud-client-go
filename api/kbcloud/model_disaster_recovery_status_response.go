// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// DisasterRecoveryStatusResponse Status of the Disaster Recovery instance
type DisasterRecoveryStatusResponse struct {
	// Unique identifier for the cluster
	ClusterId int32 `json:"clusterId"`
	// Name of the cluster
	ClusterName common.NullableString `json:"clusterName,omitempty"`
	// User-friendly name of the cluster
	DisplayName common.NullableString `json:"displayName,omitempty"`
	// Current operational status of the cluster
	Status common.NullableString `json:"status,omitempty"`
	// ID of the parent cluster, if applicable; can be empty if there is no parent relationship
	ParentId int64 `json:"parentId"`
	// Name of the parent cluster
	ParentName common.NullableString `json:"parentName,omitempty"`
	// User-friendly name of the parent cluster
	ParentDisplayName common.NullableString `json:"parentDisplayName,omitempty"`
	// Current operational status of the parent cluster
	ParentStatus common.NullableString `json:"parentStatus,omitempty"`
	// Delay time in seconds
	Delay common.NullableFloat64 `json:"delay,omitempty"`
	// Current replication point for the disaster recovery instance
	CurrentReplicationPoint common.NullableString `json:"currentReplicationPoint,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryStatusResponse instantiates a new DisasterRecoveryStatusResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryStatusResponse(clusterId int32, parentId int64) *DisasterRecoveryStatusResponse {
	this := DisasterRecoveryStatusResponse{}
	this.ClusterId = clusterId
	this.ParentId = parentId
	return &this
}

// NewDisasterRecoveryStatusResponseWithDefaults instantiates a new DisasterRecoveryStatusResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryStatusResponseWithDefaults() *DisasterRecoveryStatusResponse {
	this := DisasterRecoveryStatusResponse{}
	return &this
}

// GetClusterId returns the ClusterId field value.
func (o *DisasterRecoveryStatusResponse) GetClusterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryStatusResponse) GetClusterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *DisasterRecoveryStatusResponse) SetClusterId(v int32) {
	o.ClusterId = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given common.NullableString and assigns it to the ClusterName field.
func (o *DisasterRecoveryStatusResponse) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasDisplayName() bool {
	return o != nil && o.DisplayName.IsSet()
}

// SetDisplayName gets a reference to the given common.NullableString and assigns it to the DisplayName field.
func (o *DisasterRecoveryStatusResponse) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasStatus() bool {
	return o != nil && o.Status.IsSet()
}

// SetStatus gets a reference to the given common.NullableString and assigns it to the Status field.
func (o *DisasterRecoveryStatusResponse) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetParentId returns the ParentId field value.
func (o *DisasterRecoveryStatusResponse) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryStatusResponse) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value.
func (o *DisasterRecoveryStatusResponse) SetParentId(v int64) {
	o.ParentId = v
}

// GetParentName returns the ParentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetParentName() string {
	if o == nil || o.ParentName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// HasParentName returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasParentName() bool {
	return o != nil && o.ParentName.IsSet()
}

// SetParentName gets a reference to the given common.NullableString and assigns it to the ParentName field.
func (o *DisasterRecoveryStatusResponse) SetParentName(v string) {
	o.ParentName.Set(&v)
}

// SetParentNameNil sets the value for ParentName to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetParentNameNil() {
	o.ParentName.Set(nil)
}

// UnsetParentName ensures that no value is present for ParentName, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetParentName() {
	o.ParentName.Unset()
}

// GetParentDisplayName returns the ParentDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetParentDisplayName() string {
	if o == nil || o.ParentDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentDisplayName.Get()
}

// GetParentDisplayNameOk returns a tuple with the ParentDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetParentDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentDisplayName.Get(), o.ParentDisplayName.IsSet()
}

// HasParentDisplayName returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasParentDisplayName() bool {
	return o != nil && o.ParentDisplayName.IsSet()
}

// SetParentDisplayName gets a reference to the given common.NullableString and assigns it to the ParentDisplayName field.
func (o *DisasterRecoveryStatusResponse) SetParentDisplayName(v string) {
	o.ParentDisplayName.Set(&v)
}

// SetParentDisplayNameNil sets the value for ParentDisplayName to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetParentDisplayNameNil() {
	o.ParentDisplayName.Set(nil)
}

// UnsetParentDisplayName ensures that no value is present for ParentDisplayName, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetParentDisplayName() {
	o.ParentDisplayName.Unset()
}

// GetParentStatus returns the ParentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetParentStatus() string {
	if o == nil || o.ParentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentStatus.Get()
}

// GetParentStatusOk returns a tuple with the ParentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetParentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentStatus.Get(), o.ParentStatus.IsSet()
}

// HasParentStatus returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasParentStatus() bool {
	return o != nil && o.ParentStatus.IsSet()
}

// SetParentStatus gets a reference to the given common.NullableString and assigns it to the ParentStatus field.
func (o *DisasterRecoveryStatusResponse) SetParentStatus(v string) {
	o.ParentStatus.Set(&v)
}

// SetParentStatusNil sets the value for ParentStatus to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetParentStatusNil() {
	o.ParentStatus.Set(nil)
}

// UnsetParentStatus ensures that no value is present for ParentStatus, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetParentStatus() {
	o.ParentStatus.Unset()
}

// GetDelay returns the Delay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetDelay() float64 {
	if o == nil || o.Delay.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Delay.Get()
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetDelayOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delay.Get(), o.Delay.IsSet()
}

// HasDelay returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasDelay() bool {
	return o != nil && o.Delay.IsSet()
}

// SetDelay gets a reference to the given common.NullableFloat64 and assigns it to the Delay field.
func (o *DisasterRecoveryStatusResponse) SetDelay(v float64) {
	o.Delay.Set(&v)
}

// SetDelayNil sets the value for Delay to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetDelayNil() {
	o.Delay.Set(nil)
}

// UnsetDelay ensures that no value is present for Delay, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetDelay() {
	o.Delay.Unset()
}

// GetCurrentReplicationPoint returns the CurrentReplicationPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryStatusResponse) GetCurrentReplicationPoint() string {
	if o == nil || o.CurrentReplicationPoint.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrentReplicationPoint.Get()
}

// GetCurrentReplicationPointOk returns a tuple with the CurrentReplicationPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryStatusResponse) GetCurrentReplicationPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentReplicationPoint.Get(), o.CurrentReplicationPoint.IsSet()
}

// HasCurrentReplicationPoint returns a boolean if a field has been set.
func (o *DisasterRecoveryStatusResponse) HasCurrentReplicationPoint() bool {
	return o != nil && o.CurrentReplicationPoint.IsSet()
}

// SetCurrentReplicationPoint gets a reference to the given common.NullableString and assigns it to the CurrentReplicationPoint field.
func (o *DisasterRecoveryStatusResponse) SetCurrentReplicationPoint(v string) {
	o.CurrentReplicationPoint.Set(&v)
}

// SetCurrentReplicationPointNil sets the value for CurrentReplicationPoint to be an explicit nil.
func (o *DisasterRecoveryStatusResponse) SetCurrentReplicationPointNil() {
	o.CurrentReplicationPoint.Set(nil)
}

// UnsetCurrentReplicationPoint ensures that no value is present for CurrentReplicationPoint, not even an explicit nil.
func (o *DisasterRecoveryStatusResponse) UnsetCurrentReplicationPoint() {
	o.CurrentReplicationPoint.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterId"] = o.ClusterId
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["parentId"] = o.ParentId
	if o.ParentName.IsSet() {
		toSerialize["parentName"] = o.ParentName.Get()
	}
	if o.ParentDisplayName.IsSet() {
		toSerialize["parentDisplayName"] = o.ParentDisplayName.Get()
	}
	if o.ParentStatus.IsSet() {
		toSerialize["parentStatus"] = o.ParentStatus.Get()
	}
	if o.Delay.IsSet() {
		toSerialize["delay"] = o.Delay.Get()
	}
	if o.CurrentReplicationPoint.IsSet() {
		toSerialize["currentReplicationPoint"] = o.CurrentReplicationPoint.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryStatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId               *int32                 `json:"clusterId"`
		ClusterName             common.NullableString  `json:"clusterName,omitempty"`
		DisplayName             common.NullableString  `json:"displayName,omitempty"`
		Status                  common.NullableString  `json:"status,omitempty"`
		ParentId                *int64                 `json:"parentId"`
		ParentName              common.NullableString  `json:"parentName,omitempty"`
		ParentDisplayName       common.NullableString  `json:"parentDisplayName,omitempty"`
		ParentStatus            common.NullableString  `json:"parentStatus,omitempty"`
		Delay                   common.NullableFloat64 `json:"delay,omitempty"`
		CurrentReplicationPoint common.NullableString  `json:"currentReplicationPoint,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterId missing")
	}
	if all.ParentId == nil {
		return fmt.Errorf("required field parentId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterId", "clusterName", "displayName", "status", "parentId", "parentName", "parentDisplayName", "parentStatus", "delay", "currentReplicationPoint"})
	} else {
		return err
	}
	o.ClusterId = *all.ClusterId
	o.ClusterName = all.ClusterName
	o.DisplayName = all.DisplayName
	o.Status = all.Status
	o.ParentId = *all.ParentId
	o.ParentName = all.ParentName
	o.ParentDisplayName = all.ParentDisplayName
	o.ParentStatus = all.ParentStatus
	o.Delay = all.Delay
	o.CurrentReplicationPoint = all.CurrentReplicationPoint

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
