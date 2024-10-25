// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION AlertReceiverUserGroup
type AlertReceiverUserGroup struct {
	// NODESCRIPTION EmailEnabled
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// NODESCRIPTION Ids
	Ids []string `json:"ids,omitempty"`
	// NODESCRIPTION SmsEnabled
	SmsEnabled *bool `json:"smsEnabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertReceiverUserGroup instantiates a new AlertReceiverUserGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertReceiverUserGroup() *AlertReceiverUserGroup {
	this := AlertReceiverUserGroup{}
	return &this
}

// NewAlertReceiverUserGroupWithDefaults instantiates a new AlertReceiverUserGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertReceiverUserGroupWithDefaults() *AlertReceiverUserGroup {
	this := AlertReceiverUserGroup{}
	return &this
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *AlertReceiverUserGroup) GetEmailEnabled() bool {
	if o == nil || o.EmailEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiverUserGroup) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || o.EmailEnabled == nil {
		return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *AlertReceiverUserGroup) HasEmailEnabled() bool {
	return o != nil && o.EmailEnabled != nil
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *AlertReceiverUserGroup) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *AlertReceiverUserGroup) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiverUserGroup) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return &o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AlertReceiverUserGroup) HasIds() bool {
	return o != nil && o.Ids != nil
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *AlertReceiverUserGroup) SetIds(v []string) {
	o.Ids = v
}

// GetSmsEnabled returns the SmsEnabled field value if set, zero value otherwise.
func (o *AlertReceiverUserGroup) GetSmsEnabled() bool {
	if o == nil || o.SmsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SmsEnabled
}

// GetSmsEnabledOk returns a tuple with the SmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertReceiverUserGroup) GetSmsEnabledOk() (*bool, bool) {
	if o == nil || o.SmsEnabled == nil {
		return nil, false
	}
	return o.SmsEnabled, true
}

// HasSmsEnabled returns a boolean if a field has been set.
func (o *AlertReceiverUserGroup) HasSmsEnabled() bool {
	return o != nil && o.SmsEnabled != nil
}

// SetSmsEnabled gets a reference to the given bool and assigns it to the SmsEnabled field.
func (o *AlertReceiverUserGroup) SetSmsEnabled(v bool) {
	o.SmsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertReceiverUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EmailEnabled != nil {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.SmsEnabled != nil {
		toSerialize["smsEnabled"] = o.SmsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertReceiverUserGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EmailEnabled *bool    `json:"emailEnabled,omitempty"`
		Ids          []string `json:"ids,omitempty"`
		SmsEnabled   *bool    `json:"smsEnabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"emailEnabled", "ids", "smsEnabled"})
	} else {
		return err
	}
	o.EmailEnabled = all.EmailEnabled
	o.Ids = all.Ids
	o.SmsEnabled = all.SmsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
