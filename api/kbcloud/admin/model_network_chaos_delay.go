// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

// NetworkChaosDelay specify the delay in the chaos action
type NetworkChaosDelay struct {
	// specify the latency in the chaos action
	Latency *string `json:"latency,omitempty"`
	// specify the correlation in the chaos action
	Correlation *string `json:"correlation,omitempty"`
	// specify the jitter in the chaos action
	Jitter *string `json:"jitter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkChaosDelay instantiates a new NetworkChaosDelay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkChaosDelay() *NetworkChaosDelay {
	this := NetworkChaosDelay{}
	return &this
}

// NewNetworkChaosDelayWithDefaults instantiates a new NetworkChaosDelay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkChaosDelayWithDefaults() *NetworkChaosDelay {
	this := NetworkChaosDelay{}
	return &this
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *NetworkChaosDelay) GetLatency() string {
	if o == nil || o.Latency == nil {
		var ret string
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaosDelay) GetLatencyOk() (*string, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *NetworkChaosDelay) HasLatency() bool {
	return o != nil && o.Latency != nil
}

// SetLatency gets a reference to the given string and assigns it to the Latency field.
func (o *NetworkChaosDelay) SetLatency(v string) {
	o.Latency = &v
}

// GetCorrelation returns the Correlation field value if set, zero value otherwise.
func (o *NetworkChaosDelay) GetCorrelation() string {
	if o == nil || o.Correlation == nil {
		var ret string
		return ret
	}
	return *o.Correlation
}

// GetCorrelationOk returns a tuple with the Correlation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaosDelay) GetCorrelationOk() (*string, bool) {
	if o == nil || o.Correlation == nil {
		return nil, false
	}
	return o.Correlation, true
}

// HasCorrelation returns a boolean if a field has been set.
func (o *NetworkChaosDelay) HasCorrelation() bool {
	return o != nil && o.Correlation != nil
}

// SetCorrelation gets a reference to the given string and assigns it to the Correlation field.
func (o *NetworkChaosDelay) SetCorrelation(v string) {
	o.Correlation = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *NetworkChaosDelay) GetJitter() string {
	if o == nil || o.Jitter == nil {
		var ret string
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkChaosDelay) GetJitterOk() (*string, bool) {
	if o == nil || o.Jitter == nil {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *NetworkChaosDelay) HasJitter() bool {
	return o != nil && o.Jitter != nil
}

// SetJitter gets a reference to the given string and assigns it to the Jitter field.
func (o *NetworkChaosDelay) SetJitter(v string) {
	o.Jitter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkChaosDelay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Latency != nil {
		toSerialize["latency"] = o.Latency
	}
	if o.Correlation != nil {
		toSerialize["correlation"] = o.Correlation
	}
	if o.Jitter != nil {
		toSerialize["jitter"] = o.Jitter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkChaosDelay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Latency     *string `json:"latency,omitempty"`
		Correlation *string `json:"correlation,omitempty"`
		Jitter      *string `json:"jitter,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"latency", "correlation", "jitter"})
	} else {
		return err
	}
	o.Latency = all.Latency
	o.Correlation = all.Correlation
	o.Jitter = all.Jitter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
