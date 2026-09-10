// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchUnassignedInfo struct {
	Reason               common.NullableString `json:"reason,omitempty"`
	At                   common.NullableString `json:"at,omitempty"`
	Details              common.NullableString `json:"details,omitempty"`
	LastAllocationStatus common.NullableString `json:"lastAllocationStatus,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchUnassignedInfo instantiates a new ElasticsearchUnassignedInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchUnassignedInfo() *ElasticsearchUnassignedInfo {
	this := ElasticsearchUnassignedInfo{}
	return &this
}

// NewElasticsearchUnassignedInfoWithDefaults instantiates a new ElasticsearchUnassignedInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchUnassignedInfoWithDefaults() *ElasticsearchUnassignedInfo {
	this := ElasticsearchUnassignedInfo{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchUnassignedInfo) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchUnassignedInfo) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *ElasticsearchUnassignedInfo) HasReason() bool {
	return o != nil && o.Reason.IsSet()
}

// SetReason gets a reference to the given common.NullableString and assigns it to the Reason field.
func (o *ElasticsearchUnassignedInfo) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil.
func (o *ElasticsearchUnassignedInfo) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil.
func (o *ElasticsearchUnassignedInfo) UnsetReason() {
	o.Reason.Unset()
}

// GetAt returns the At field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchUnassignedInfo) GetAt() string {
	if o == nil || o.At.Get() == nil {
		var ret string
		return ret
	}
	return *o.At.Get()
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchUnassignedInfo) GetAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.At.Get(), o.At.IsSet()
}

// HasAt returns a boolean if a field has been set.
func (o *ElasticsearchUnassignedInfo) HasAt() bool {
	return o != nil && o.At.IsSet()
}

// SetAt gets a reference to the given common.NullableString and assigns it to the At field.
func (o *ElasticsearchUnassignedInfo) SetAt(v string) {
	o.At.Set(&v)
}

// SetAtNil sets the value for At to be an explicit nil.
func (o *ElasticsearchUnassignedInfo) SetAtNil() {
	o.At.Set(nil)
}

// UnsetAt ensures that no value is present for At, not even an explicit nil.
func (o *ElasticsearchUnassignedInfo) UnsetAt() {
	o.At.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchUnassignedInfo) GetDetails() string {
	if o == nil || o.Details.Get() == nil {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchUnassignedInfo) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *ElasticsearchUnassignedInfo) HasDetails() bool {
	return o != nil && o.Details.IsSet()
}

// SetDetails gets a reference to the given common.NullableString and assigns it to the Details field.
func (o *ElasticsearchUnassignedInfo) SetDetails(v string) {
	o.Details.Set(&v)
}

// SetDetailsNil sets the value for Details to be an explicit nil.
func (o *ElasticsearchUnassignedInfo) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil.
func (o *ElasticsearchUnassignedInfo) UnsetDetails() {
	o.Details.Unset()
}

// GetLastAllocationStatus returns the LastAllocationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchUnassignedInfo) GetLastAllocationStatus() string {
	if o == nil || o.LastAllocationStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastAllocationStatus.Get()
}

// GetLastAllocationStatusOk returns a tuple with the LastAllocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchUnassignedInfo) GetLastAllocationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAllocationStatus.Get(), o.LastAllocationStatus.IsSet()
}

// HasLastAllocationStatus returns a boolean if a field has been set.
func (o *ElasticsearchUnassignedInfo) HasLastAllocationStatus() bool {
	return o != nil && o.LastAllocationStatus.IsSet()
}

// SetLastAllocationStatus gets a reference to the given common.NullableString and assigns it to the LastAllocationStatus field.
func (o *ElasticsearchUnassignedInfo) SetLastAllocationStatus(v string) {
	o.LastAllocationStatus.Set(&v)
}

// SetLastAllocationStatusNil sets the value for LastAllocationStatus to be an explicit nil.
func (o *ElasticsearchUnassignedInfo) SetLastAllocationStatusNil() {
	o.LastAllocationStatus.Set(nil)
}

// UnsetLastAllocationStatus ensures that no value is present for LastAllocationStatus, not even an explicit nil.
func (o *ElasticsearchUnassignedInfo) UnsetLastAllocationStatus() {
	o.LastAllocationStatus.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchUnassignedInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.At.IsSet() {
		toSerialize["at"] = o.At.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.LastAllocationStatus.IsSet() {
		toSerialize["lastAllocationStatus"] = o.LastAllocationStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchUnassignedInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason               common.NullableString `json:"reason,omitempty"`
		At                   common.NullableString `json:"at,omitempty"`
		Details              common.NullableString `json:"details,omitempty"`
		LastAllocationStatus common.NullableString `json:"lastAllocationStatus,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reason", "at", "details", "lastAllocationStatus"})
	} else {
		return err
	}
	o.Reason = all.Reason
	o.At = all.At
	o.Details = all.Details
	o.LastAllocationStatus = all.LastAllocationStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
