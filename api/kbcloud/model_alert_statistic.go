// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AlertStatistic struct {
	Total    *int32 `json:"total,omitempty"`
	Critical *int32 `json:"critical,omitempty"`
	Warning  *int32 `json:"warning,omitempty"`
	Info     *int32 `json:"info,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertStatistic instantiates a new AlertStatistic object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertStatistic() *AlertStatistic {
	this := AlertStatistic{}
	return &this
}

// NewAlertStatisticWithDefaults instantiates a new AlertStatistic object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertStatisticWithDefaults() *AlertStatistic {
	this := AlertStatistic{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlertStatistic) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStatistic) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlertStatistic) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AlertStatistic) SetTotal(v int32) {
	o.Total = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *AlertStatistic) GetCritical() int32 {
	if o == nil || o.Critical == nil {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStatistic) GetCriticalOk() (*int32, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *AlertStatistic) HasCritical() bool {
	return o != nil && o.Critical != nil
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *AlertStatistic) SetCritical(v int32) {
	o.Critical = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *AlertStatistic) GetWarning() int32 {
	if o == nil || o.Warning == nil {
		var ret int32
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStatistic) GetWarningOk() (*int32, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *AlertStatistic) HasWarning() bool {
	return o != nil && o.Warning != nil
}

// SetWarning gets a reference to the given int32 and assigns it to the Warning field.
func (o *AlertStatistic) SetWarning(v int32) {
	o.Warning = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *AlertStatistic) GetInfo() int32 {
	if o == nil || o.Info == nil {
		var ret int32
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStatistic) GetInfoOk() (*int32, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *AlertStatistic) HasInfo() bool {
	return o != nil && o.Info != nil
}

// SetInfo gets a reference to the given int32 and assigns it to the Info field.
func (o *AlertStatistic) SetInfo(v int32) {
	o.Info = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertStatistic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Critical != nil {
		toSerialize["critical"] = o.Critical
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertStatistic) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total    *int32 `json:"total,omitempty"`
		Critical *int32 `json:"critical,omitempty"`
		Warning  *int32 `json:"warning,omitempty"`
		Info     *int32 `json:"info,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"total", "critical", "warning", "info"})
	} else {
		return err
	}
	o.Total = all.Total
	o.Critical = all.Critical
	o.Warning = all.Warning
	o.Info = all.Info

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
