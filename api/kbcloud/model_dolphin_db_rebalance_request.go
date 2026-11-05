// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DolphinDBRebalanceRequest struct {
	UpdatedBeforeDays *float64 `json:"updatedBeforeDays,omitempty"`
	// Addon wait timeout. Expiry does not cancel native migrations.
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBRebalanceRequest instantiates a new DolphinDBRebalanceRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBRebalanceRequest() *DolphinDBRebalanceRequest {
	this := DolphinDBRebalanceRequest{}
	var updatedBeforeDays float64 = 7
	this.UpdatedBeforeDays = &updatedBeforeDays
	var timeoutSeconds int32 = 3600
	this.TimeoutSeconds = &timeoutSeconds
	return &this
}

// NewDolphinDBRebalanceRequestWithDefaults instantiates a new DolphinDBRebalanceRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBRebalanceRequestWithDefaults() *DolphinDBRebalanceRequest {
	this := DolphinDBRebalanceRequest{}
	var updatedBeforeDays float64 = 7
	this.UpdatedBeforeDays = &updatedBeforeDays
	var timeoutSeconds int32 = 3600
	this.TimeoutSeconds = &timeoutSeconds
	return &this
}

// GetUpdatedBeforeDays returns the UpdatedBeforeDays field value if set, zero value otherwise.
func (o *DolphinDBRebalanceRequest) GetUpdatedBeforeDays() float64 {
	if o == nil || o.UpdatedBeforeDays == nil {
		var ret float64
		return ret
	}
	return *o.UpdatedBeforeDays
}

// GetUpdatedBeforeDaysOk returns a tuple with the UpdatedBeforeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalanceRequest) GetUpdatedBeforeDaysOk() (*float64, bool) {
	if o == nil || o.UpdatedBeforeDays == nil {
		return nil, false
	}
	return o.UpdatedBeforeDays, true
}

// HasUpdatedBeforeDays returns a boolean if a field has been set.
func (o *DolphinDBRebalanceRequest) HasUpdatedBeforeDays() bool {
	return o != nil && o.UpdatedBeforeDays != nil
}

// SetUpdatedBeforeDays gets a reference to the given float64 and assigns it to the UpdatedBeforeDays field.
func (o *DolphinDBRebalanceRequest) SetUpdatedBeforeDays(v float64) {
	o.UpdatedBeforeDays = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *DolphinDBRebalanceRequest) GetTimeoutSeconds() int32 {
	if o == nil || o.TimeoutSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBRebalanceRequest) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || o.TimeoutSeconds == nil {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *DolphinDBRebalanceRequest) HasTimeoutSeconds() bool {
	return o != nil && o.TimeoutSeconds != nil
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *DolphinDBRebalanceRequest) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBRebalanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.UpdatedBeforeDays != nil {
		toSerialize["updatedBeforeDays"] = o.UpdatedBeforeDays
	}
	if o.TimeoutSeconds != nil {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBRebalanceRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UpdatedBeforeDays *float64 `json:"updatedBeforeDays,omitempty"`
		TimeoutSeconds    *int32   `json:"timeoutSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"updatedBeforeDays", "timeoutSeconds"})
	} else {
		return err
	}
	o.UpdatedBeforeDays = all.UpdatedBeforeDays
	o.TimeoutSeconds = all.TimeoutSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
