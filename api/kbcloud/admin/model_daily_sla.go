// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "time"

// DailySLA The SLA (Service Level Agreement) for a cluster on a specific date.
type DailySLA struct {
	// The date for which the SLA is calculated.
	Date *time.Time `json:"date,omitempty"`
	// The SLA value for the cluster on the specified date. Null indicates that the cluster was not existing on that date.
	Sla common.NullableFloat `json:"sla,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDailySLA instantiates a new DailySLA object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailySLA() *DailySLA {
	this := DailySLA{}
	return &this
}

// NewDailySLAWithDefaults instantiates a new DailySLA object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDailySLAWithDefaults() *DailySLA {
	this := DailySLA{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DailySLA) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySLA) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DailySLA) HasDate() bool {
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *DailySLA) SetDate(v time.Time) {
	o.Date = &v
}

// GetSla returns the Sla field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DailySLA) GetSla() float {
	if o == nil || o.Sla.Get() == nil {
		var ret float
		return ret
	}
	return *o.Sla.Get()
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DailySLA) GetSlaOk() (*float, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sla.Get(), o.Sla.IsSet()
}

// HasSla returns a boolean if a field has been set.
func (o *DailySLA) HasSla() bool {
	return o != nil && o.Sla.IsSet()
}

// SetSla gets a reference to the given common.NullableFloat and assigns it to the Sla field.
func (o *DailySLA) SetSla(v float) {
	o.Sla.Set(&v)
}

// SetSlaNil sets the value for Sla to be an explicit nil.
func (o *DailySLA) SetSlaNil() {
	o.Sla.Set(nil)
}

// UnsetSla ensures that no value is present for Sla, not even an explicit nil.
func (o *DailySLA) UnsetSla() {
	o.Sla.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DailySLA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Sla.IsSet() {
		toSerialize["sla"] = o.Sla.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DailySLA) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date *time.Time           `json:"date,omitempty"`
		Sla  common.NullableFloat `json:"sla,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"date", "sla"})
	} else {
		return err
	}
	o.Date = all.Date
	o.Sla = all.Sla

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
