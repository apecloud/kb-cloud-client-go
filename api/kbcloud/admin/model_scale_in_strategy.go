// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ScaleInStrategy struct {
	// Type of scale in strategy
	Type *ScaleInType `json:"type,omitempty"`
	// Maximum number of nodes that can be unavailable during scale in. Can be an absolute number or a percentage.
	MaxUnavailable *string `json:"maxUnavailable,omitempty"`
	// Maximum number of nodes that can be scheduled above the desired number. Can be an absolute number or a percentage.
	MaxSurge *string `json:"maxSurge,omitempty"`
	// How long to wait for pod eviction when draining a node
	DrainTimeout *string `json:"drainTimeout,omitempty"`
	// Time to wait between pod evictions
	PodEvictionWaitTime *string `json:"podEvictionWaitTime,omitempty"`
	// Minimum time a replacement node should be ready before proceeding
	MinReadySeconds *int32              `json:"minReadySeconds,omitempty"`
	HealthCheck     *ScaleInHealthCheck `json:"healthCheck,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaleInStrategy instantiates a new ScaleInStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaleInStrategy() *ScaleInStrategy {
	this := ScaleInStrategy{}
	var typeVar ScaleInType = ScaleInTypeSerial
	this.Type = &typeVar
	var maxUnavailable string = "25%"
	this.MaxUnavailable = &maxUnavailable
	var maxSurge string = "0%"
	this.MaxSurge = &maxSurge
	var drainTimeout string = "5m"
	this.DrainTimeout = &drainTimeout
	var podEvictionWaitTime string = "30s"
	this.PodEvictionWaitTime = &podEvictionWaitTime
	var minReadySeconds int32 = 30
	this.MinReadySeconds = &minReadySeconds
	return &this
}

// NewScaleInStrategyWithDefaults instantiates a new ScaleInStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaleInStrategyWithDefaults() *ScaleInStrategy {
	this := ScaleInStrategy{}
	var typeVar ScaleInType = ScaleInTypeSerial
	this.Type = &typeVar
	var maxUnavailable string = "25%"
	this.MaxUnavailable = &maxUnavailable
	var maxSurge string = "0%"
	this.MaxSurge = &maxSurge
	var drainTimeout string = "5m"
	this.DrainTimeout = &drainTimeout
	var podEvictionWaitTime string = "30s"
	this.PodEvictionWaitTime = &podEvictionWaitTime
	var minReadySeconds int32 = 30
	this.MinReadySeconds = &minReadySeconds
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetType() ScaleInType {
	if o == nil || o.Type == nil {
		var ret ScaleInType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetTypeOk() (*ScaleInType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ScaleInType and assigns it to the Type field.
func (o *ScaleInStrategy) SetType(v ScaleInType) {
	o.Type = &v
}

// GetMaxUnavailable returns the MaxUnavailable field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetMaxUnavailable() string {
	if o == nil || o.MaxUnavailable == nil {
		var ret string
		return ret
	}
	return *o.MaxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetMaxUnavailableOk() (*string, bool) {
	if o == nil || o.MaxUnavailable == nil {
		return nil, false
	}
	return o.MaxUnavailable, true
}

// HasMaxUnavailable returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasMaxUnavailable() bool {
	return o != nil && o.MaxUnavailable != nil
}

// SetMaxUnavailable gets a reference to the given string and assigns it to the MaxUnavailable field.
func (o *ScaleInStrategy) SetMaxUnavailable(v string) {
	o.MaxUnavailable = &v
}

// GetMaxSurge returns the MaxSurge field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetMaxSurge() string {
	if o == nil || o.MaxSurge == nil {
		var ret string
		return ret
	}
	return *o.MaxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetMaxSurgeOk() (*string, bool) {
	if o == nil || o.MaxSurge == nil {
		return nil, false
	}
	return o.MaxSurge, true
}

// HasMaxSurge returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasMaxSurge() bool {
	return o != nil && o.MaxSurge != nil
}

// SetMaxSurge gets a reference to the given string and assigns it to the MaxSurge field.
func (o *ScaleInStrategy) SetMaxSurge(v string) {
	o.MaxSurge = &v
}

// GetDrainTimeout returns the DrainTimeout field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetDrainTimeout() string {
	if o == nil || o.DrainTimeout == nil {
		var ret string
		return ret
	}
	return *o.DrainTimeout
}

// GetDrainTimeoutOk returns a tuple with the DrainTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetDrainTimeoutOk() (*string, bool) {
	if o == nil || o.DrainTimeout == nil {
		return nil, false
	}
	return o.DrainTimeout, true
}

// HasDrainTimeout returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasDrainTimeout() bool {
	return o != nil && o.DrainTimeout != nil
}

// SetDrainTimeout gets a reference to the given string and assigns it to the DrainTimeout field.
func (o *ScaleInStrategy) SetDrainTimeout(v string) {
	o.DrainTimeout = &v
}

// GetPodEvictionWaitTime returns the PodEvictionWaitTime field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetPodEvictionWaitTime() string {
	if o == nil || o.PodEvictionWaitTime == nil {
		var ret string
		return ret
	}
	return *o.PodEvictionWaitTime
}

// GetPodEvictionWaitTimeOk returns a tuple with the PodEvictionWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetPodEvictionWaitTimeOk() (*string, bool) {
	if o == nil || o.PodEvictionWaitTime == nil {
		return nil, false
	}
	return o.PodEvictionWaitTime, true
}

// HasPodEvictionWaitTime returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasPodEvictionWaitTime() bool {
	return o != nil && o.PodEvictionWaitTime != nil
}

// SetPodEvictionWaitTime gets a reference to the given string and assigns it to the PodEvictionWaitTime field.
func (o *ScaleInStrategy) SetPodEvictionWaitTime(v string) {
	o.PodEvictionWaitTime = &v
}

// GetMinReadySeconds returns the MinReadySeconds field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetMinReadySeconds() int32 {
	if o == nil || o.MinReadySeconds == nil {
		var ret int32
		return ret
	}
	return *o.MinReadySeconds
}

// GetMinReadySecondsOk returns a tuple with the MinReadySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetMinReadySecondsOk() (*int32, bool) {
	if o == nil || o.MinReadySeconds == nil {
		return nil, false
	}
	return o.MinReadySeconds, true
}

// HasMinReadySeconds returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasMinReadySeconds() bool {
	return o != nil && o.MinReadySeconds != nil
}

// SetMinReadySeconds gets a reference to the given int32 and assigns it to the MinReadySeconds field.
func (o *ScaleInStrategy) SetMinReadySeconds(v int32) {
	o.MinReadySeconds = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *ScaleInStrategy) GetHealthCheck() ScaleInHealthCheck {
	if o == nil || o.HealthCheck == nil {
		var ret ScaleInHealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleInStrategy) GetHealthCheckOk() (*ScaleInHealthCheck, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *ScaleInStrategy) HasHealthCheck() bool {
	return o != nil && o.HealthCheck != nil
}

// SetHealthCheck gets a reference to the given ScaleInHealthCheck and assigns it to the HealthCheck field.
func (o *ScaleInStrategy) SetHealthCheck(v ScaleInHealthCheck) {
	o.HealthCheck = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaleInStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MaxUnavailable != nil {
		toSerialize["maxUnavailable"] = o.MaxUnavailable
	}
	if o.MaxSurge != nil {
		toSerialize["maxSurge"] = o.MaxSurge
	}
	if o.DrainTimeout != nil {
		toSerialize["drainTimeout"] = o.DrainTimeout
	}
	if o.PodEvictionWaitTime != nil {
		toSerialize["podEvictionWaitTime"] = o.PodEvictionWaitTime
	}
	if o.MinReadySeconds != nil {
		toSerialize["minReadySeconds"] = o.MinReadySeconds
	}
	if o.HealthCheck != nil {
		toSerialize["healthCheck"] = o.HealthCheck
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaleInStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type                *ScaleInType        `json:"type,omitempty"`
		MaxUnavailable      *string             `json:"maxUnavailable,omitempty"`
		MaxSurge            *string             `json:"maxSurge,omitempty"`
		DrainTimeout        *string             `json:"drainTimeout,omitempty"`
		PodEvictionWaitTime *string             `json:"podEvictionWaitTime,omitempty"`
		MinReadySeconds     *int32              `json:"minReadySeconds,omitempty"`
		HealthCheck         *ScaleInHealthCheck `json:"healthCheck,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "maxUnavailable", "maxSurge", "drainTimeout", "podEvictionWaitTime", "minReadySeconds", "healthCheck"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.MaxUnavailable = all.MaxUnavailable
	o.MaxSurge = all.MaxSurge
	o.DrainTimeout = all.DrainTimeout
	o.PodEvictionWaitTime = all.PodEvictionWaitTime
	o.MinReadySeconds = all.MinReadySeconds
	if all.HealthCheck != nil && all.HealthCheck.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.HealthCheck = all.HealthCheck

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
