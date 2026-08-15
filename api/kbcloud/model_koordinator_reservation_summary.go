// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KoordinatorReservationSummary struct {
	Enabled              bool                                `json:"enabled"`
	SchedulerName        *string                             `json:"schedulerName,omitempty"`
	DefaultSchedulerName *string                             `json:"defaultSchedulerName,omitempty"`
	SchedulerOptions     []SchedulerOption                   `json:"schedulerOptions"`
	Items                []KoordinatorReservationSummaryItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKoordinatorReservationSummary instantiates a new KoordinatorReservationSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorReservationSummary(enabled bool, schedulerOptions []SchedulerOption, items []KoordinatorReservationSummaryItem) *KoordinatorReservationSummary {
	this := KoordinatorReservationSummary{}
	this.Enabled = enabled
	this.SchedulerOptions = schedulerOptions
	this.Items = items
	return &this
}

// NewKoordinatorReservationSummaryWithDefaults instantiates a new KoordinatorReservationSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorReservationSummaryWithDefaults() *KoordinatorReservationSummary {
	this := KoordinatorReservationSummary{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *KoordinatorReservationSummary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *KoordinatorReservationSummary) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchedulerName returns the SchedulerName field value if set, zero value otherwise.
func (o *KoordinatorReservationSummary) GetSchedulerName() string {
	if o == nil || o.SchedulerName == nil {
		var ret string
		return ret
	}
	return *o.SchedulerName
}

// GetSchedulerNameOk returns a tuple with the SchedulerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummary) GetSchedulerNameOk() (*string, bool) {
	if o == nil || o.SchedulerName == nil {
		return nil, false
	}
	return o.SchedulerName, true
}

// HasSchedulerName returns a boolean if a field has been set.
func (o *KoordinatorReservationSummary) HasSchedulerName() bool {
	return o != nil && o.SchedulerName != nil
}

// SetSchedulerName gets a reference to the given string and assigns it to the SchedulerName field.
func (o *KoordinatorReservationSummary) SetSchedulerName(v string) {
	o.SchedulerName = &v
}

// GetDefaultSchedulerName returns the DefaultSchedulerName field value if set, zero value otherwise.
func (o *KoordinatorReservationSummary) GetDefaultSchedulerName() string {
	if o == nil || o.DefaultSchedulerName == nil {
		var ret string
		return ret
	}
	return *o.DefaultSchedulerName
}

// GetDefaultSchedulerNameOk returns a tuple with the DefaultSchedulerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummary) GetDefaultSchedulerNameOk() (*string, bool) {
	if o == nil || o.DefaultSchedulerName == nil {
		return nil, false
	}
	return o.DefaultSchedulerName, true
}

// HasDefaultSchedulerName returns a boolean if a field has been set.
func (o *KoordinatorReservationSummary) HasDefaultSchedulerName() bool {
	return o != nil && o.DefaultSchedulerName != nil
}

// SetDefaultSchedulerName gets a reference to the given string and assigns it to the DefaultSchedulerName field.
func (o *KoordinatorReservationSummary) SetDefaultSchedulerName(v string) {
	o.DefaultSchedulerName = &v
}

// GetSchedulerOptions returns the SchedulerOptions field value.
func (o *KoordinatorReservationSummary) GetSchedulerOptions() []SchedulerOption {
	if o == nil {
		var ret []SchedulerOption
		return ret
	}
	return o.SchedulerOptions
}

// GetSchedulerOptionsOk returns a tuple with the SchedulerOptions field value
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummary) GetSchedulerOptionsOk() (*[]SchedulerOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchedulerOptions, true
}

// SetSchedulerOptions sets field value.
func (o *KoordinatorReservationSummary) SetSchedulerOptions(v []SchedulerOption) {
	o.SchedulerOptions = v
}

// GetItems returns the Items field value.
func (o *KoordinatorReservationSummary) GetItems() []KoordinatorReservationSummaryItem {
	if o == nil {
		var ret []KoordinatorReservationSummaryItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummary) GetItemsOk() (*[]KoordinatorReservationSummaryItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *KoordinatorReservationSummary) SetItems(v []KoordinatorReservationSummaryItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorReservationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.SchedulerName != nil {
		toSerialize["schedulerName"] = o.SchedulerName
	}
	if o.DefaultSchedulerName != nil {
		toSerialize["defaultSchedulerName"] = o.DefaultSchedulerName
	}
	toSerialize["schedulerOptions"] = o.SchedulerOptions
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KoordinatorReservationSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled              *bool                                `json:"enabled"`
		SchedulerName        *string                              `json:"schedulerName,omitempty"`
		DefaultSchedulerName *string                              `json:"defaultSchedulerName,omitempty"`
		SchedulerOptions     *[]SchedulerOption                   `json:"schedulerOptions"`
		Items                *[]KoordinatorReservationSummaryItem `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.SchedulerOptions == nil {
		return fmt.Errorf("required field schedulerOptions missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "schedulerName", "defaultSchedulerName", "schedulerOptions", "items"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.SchedulerName = all.SchedulerName
	o.DefaultSchedulerName = all.DefaultSchedulerName
	o.SchedulerOptions = *all.SchedulerOptions
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
