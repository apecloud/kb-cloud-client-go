// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ScaleInHealthCheck struct {
	// Whether to perform health checks during scale in
	Enabled *bool `json:"enabled,omitempty"`
	// Number of consecutive successful health checks required
	HealthyThreshold *int32 `json:"healthyThreshold,omitempty"`
	// Number of consecutive failed health checks before marking unhealthy
	UnhealthyThreshold *int32 `json:"unhealthyThreshold,omitempty"`
	// Interval between health checks
	Interval *string `json:"interval,omitempty"`
	// Timeout for each health check
	Timeout *string `json:"timeout,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaleInHealthCheck instantiates a new ScaleInHealthCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaleInHealthCheck() *ScaleInHealthCheck {
	this := ScaleInHealthCheck{}
	var enabled bool = true
	this.Enabled = &enabled
	var healthyThreshold int32 = 3
	this.HealthyThreshold = &healthyThreshold
	var unhealthyThreshold int32 = 3
	this.UnhealthyThreshold = &unhealthyThreshold
	var interval string = "10s"
	this.Interval = &interval
	var timeout string = "5s"
	this.Timeout = &timeout
	return &this
}

// NewScaleInHealthCheckWithDefaults instantiates a new ScaleInHealthCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaleInHealthCheckWithDefaults() *ScaleInHealthCheck {
	this := ScaleInHealthCheck{}
	var enabled bool = true
	this.Enabled = &enabled
	var healthyThreshold int32 = 3
	this.HealthyThreshold = &healthyThreshold
	var unhealthyThreshold int32 = 3
	this.UnhealthyThreshold = &unhealthyThreshold
	var interval string = "10s"
	this.Interval = &interval
	var timeout string = "5s"
	this.Timeout = &timeout
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScaleInHealthCheck) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInHealthCheck) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScaleInHealthCheck) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScaleInHealthCheck) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHealthyThreshold returns the HealthyThreshold field value if set, zero value otherwise.
func (o *ScaleInHealthCheck) GetHealthyThreshold() int32 {
	if o == nil || o.HealthyThreshold == nil {
		var ret int32
		return ret
	}
	return *o.HealthyThreshold
}

// GetHealthyThresholdOk returns a tuple with the HealthyThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInHealthCheck) GetHealthyThresholdOk() (*int32, bool) {
	if o == nil || o.HealthyThreshold == nil {
		return nil, false
	}
	return o.HealthyThreshold, true
}

// HasHealthyThreshold returns a boolean if a field has been set.
func (o *ScaleInHealthCheck) HasHealthyThreshold() bool {
	return o != nil && o.HealthyThreshold != nil
}

// SetHealthyThreshold gets a reference to the given int32 and assigns it to the HealthyThreshold field.
func (o *ScaleInHealthCheck) SetHealthyThreshold(v int32) {
	o.HealthyThreshold = &v
}

// GetUnhealthyThreshold returns the UnhealthyThreshold field value if set, zero value otherwise.
func (o *ScaleInHealthCheck) GetUnhealthyThreshold() int32 {
	if o == nil || o.UnhealthyThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UnhealthyThreshold
}

// GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInHealthCheck) GetUnhealthyThresholdOk() (*int32, bool) {
	if o == nil || o.UnhealthyThreshold == nil {
		return nil, false
	}
	return o.UnhealthyThreshold, true
}

// HasUnhealthyThreshold returns a boolean if a field has been set.
func (o *ScaleInHealthCheck) HasUnhealthyThreshold() bool {
	return o != nil && o.UnhealthyThreshold != nil
}

// SetUnhealthyThreshold gets a reference to the given int32 and assigns it to the UnhealthyThreshold field.
func (o *ScaleInHealthCheck) SetUnhealthyThreshold(v int32) {
	o.UnhealthyThreshold = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ScaleInHealthCheck) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInHealthCheck) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ScaleInHealthCheck) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *ScaleInHealthCheck) SetInterval(v string) {
	o.Interval = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ScaleInHealthCheck) GetTimeout() string {
	if o == nil || o.Timeout == nil {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInHealthCheck) GetTimeoutOk() (*string, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ScaleInHealthCheck) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *ScaleInHealthCheck) SetTimeout(v string) {
	o.Timeout = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaleInHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HealthyThreshold != nil {
		toSerialize["healthyThreshold"] = o.HealthyThreshold
	}
	if o.UnhealthyThreshold != nil {
		toSerialize["unhealthyThreshold"] = o.UnhealthyThreshold
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaleInHealthCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled            *bool   `json:"enabled,omitempty"`
		HealthyThreshold   *int32  `json:"healthyThreshold,omitempty"`
		UnhealthyThreshold *int32  `json:"unhealthyThreshold,omitempty"`
		Interval           *string `json:"interval,omitempty"`
		Timeout            *string `json:"timeout,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "healthyThreshold", "unhealthyThreshold", "interval", "timeout"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.HealthyThreshold = all.HealthyThreshold
	o.UnhealthyThreshold = all.UnhealthyThreshold
	o.Interval = all.Interval
	o.Timeout = all.Timeout

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
