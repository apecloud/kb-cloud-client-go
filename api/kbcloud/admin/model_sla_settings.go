// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SLASettings struct {
	// The name of the environment.
	EnvironmentName      *string  `json:"environmentName,omitempty"`
	ProbeInterval        *string  `json:"probeInterval,omitempty"`
	ProbeTimeout         *string  `json:"probeTimeout,omitempty"`
	EnableProbeRetry     *bool    `json:"enableProbeRetry,omitempty"`
	ProbeRetryCount      *int32   `json:"probeRetryCount,omitempty"`
	ProbeRetryInterval   *string  `json:"probeRetryInterval,omitempty"`
	ProbeEngines         []string `json:"probeEngines,omitempty"`
	ProbePoolSize        *int32   `json:"probePoolSize,omitempty"`
	CollectorBatchSize   *int32   `json:"collectorBatchSize,omitempty"`
	CollectorMaxWaitTime *string  `json:"collectorMaxWaitTime,omitempty"`
	// The SLA threshold value for the cluster.
	SlaThreshold *float `json:"slaThreshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLASettings instantiates a new SLASettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLASettings() *SLASettings {
	this := SLASettings{}
	return &this
}

// NewSLASettingsWithDefaults instantiates a new SLASettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLASettingsWithDefaults() *SLASettings {
	this := SLASettings{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *SLASettings) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *SLASettings) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *SLASettings) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetProbeInterval returns the ProbeInterval field value if set, zero value otherwise.
func (o *SLASettings) GetProbeInterval() string {
	if o == nil || o.ProbeInterval == nil {
		var ret string
		return ret
	}
	return *o.ProbeInterval
}

// GetProbeIntervalOk returns a tuple with the ProbeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbeIntervalOk() (*string, bool) {
	if o == nil || o.ProbeInterval == nil {
		return nil, false
	}
	return o.ProbeInterval, true
}

// HasProbeInterval returns a boolean if a field has been set.
func (o *SLASettings) HasProbeInterval() bool {
	return o != nil && o.ProbeInterval != nil
}

// SetProbeInterval gets a reference to the given string and assigns it to the ProbeInterval field.
func (o *SLASettings) SetProbeInterval(v string) {
	o.ProbeInterval = &v
}

// GetProbeTimeout returns the ProbeTimeout field value if set, zero value otherwise.
func (o *SLASettings) GetProbeTimeout() string {
	if o == nil || o.ProbeTimeout == nil {
		var ret string
		return ret
	}
	return *o.ProbeTimeout
}

// GetProbeTimeoutOk returns a tuple with the ProbeTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbeTimeoutOk() (*string, bool) {
	if o == nil || o.ProbeTimeout == nil {
		return nil, false
	}
	return o.ProbeTimeout, true
}

// HasProbeTimeout returns a boolean if a field has been set.
func (o *SLASettings) HasProbeTimeout() bool {
	return o != nil && o.ProbeTimeout != nil
}

// SetProbeTimeout gets a reference to the given string and assigns it to the ProbeTimeout field.
func (o *SLASettings) SetProbeTimeout(v string) {
	o.ProbeTimeout = &v
}

// GetEnableProbeRetry returns the EnableProbeRetry field value if set, zero value otherwise.
func (o *SLASettings) GetEnableProbeRetry() bool {
	if o == nil || o.EnableProbeRetry == nil {
		var ret bool
		return ret
	}
	return *o.EnableProbeRetry
}

// GetEnableProbeRetryOk returns a tuple with the EnableProbeRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetEnableProbeRetryOk() (*bool, bool) {
	if o == nil || o.EnableProbeRetry == nil {
		return nil, false
	}
	return o.EnableProbeRetry, true
}

// HasEnableProbeRetry returns a boolean if a field has been set.
func (o *SLASettings) HasEnableProbeRetry() bool {
	return o != nil && o.EnableProbeRetry != nil
}

// SetEnableProbeRetry gets a reference to the given bool and assigns it to the EnableProbeRetry field.
func (o *SLASettings) SetEnableProbeRetry(v bool) {
	o.EnableProbeRetry = &v
}

// GetProbeRetryCount returns the ProbeRetryCount field value if set, zero value otherwise.
func (o *SLASettings) GetProbeRetryCount() int32 {
	if o == nil || o.ProbeRetryCount == nil {
		var ret int32
		return ret
	}
	return *o.ProbeRetryCount
}

// GetProbeRetryCountOk returns a tuple with the ProbeRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbeRetryCountOk() (*int32, bool) {
	if o == nil || o.ProbeRetryCount == nil {
		return nil, false
	}
	return o.ProbeRetryCount, true
}

// HasProbeRetryCount returns a boolean if a field has been set.
func (o *SLASettings) HasProbeRetryCount() bool {
	return o != nil && o.ProbeRetryCount != nil
}

// SetProbeRetryCount gets a reference to the given int32 and assigns it to the ProbeRetryCount field.
func (o *SLASettings) SetProbeRetryCount(v int32) {
	o.ProbeRetryCount = &v
}

// GetProbeRetryInterval returns the ProbeRetryInterval field value if set, zero value otherwise.
func (o *SLASettings) GetProbeRetryInterval() string {
	if o == nil || o.ProbeRetryInterval == nil {
		var ret string
		return ret
	}
	return *o.ProbeRetryInterval
}

// GetProbeRetryIntervalOk returns a tuple with the ProbeRetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbeRetryIntervalOk() (*string, bool) {
	if o == nil || o.ProbeRetryInterval == nil {
		return nil, false
	}
	return o.ProbeRetryInterval, true
}

// HasProbeRetryInterval returns a boolean if a field has been set.
func (o *SLASettings) HasProbeRetryInterval() bool {
	return o != nil && o.ProbeRetryInterval != nil
}

// SetProbeRetryInterval gets a reference to the given string and assigns it to the ProbeRetryInterval field.
func (o *SLASettings) SetProbeRetryInterval(v string) {
	o.ProbeRetryInterval = &v
}

// GetProbeEngines returns the ProbeEngines field value if set, zero value otherwise.
func (o *SLASettings) GetProbeEngines() []string {
	if o == nil || o.ProbeEngines == nil {
		var ret []string
		return ret
	}
	return o.ProbeEngines
}

// GetProbeEnginesOk returns a tuple with the ProbeEngines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbeEnginesOk() (*[]string, bool) {
	if o == nil || o.ProbeEngines == nil {
		return nil, false
	}
	return &o.ProbeEngines, true
}

// HasProbeEngines returns a boolean if a field has been set.
func (o *SLASettings) HasProbeEngines() bool {
	return o != nil && o.ProbeEngines != nil
}

// SetProbeEngines gets a reference to the given []string and assigns it to the ProbeEngines field.
func (o *SLASettings) SetProbeEngines(v []string) {
	o.ProbeEngines = v
}

// GetProbePoolSize returns the ProbePoolSize field value if set, zero value otherwise.
func (o *SLASettings) GetProbePoolSize() int32 {
	if o == nil || o.ProbePoolSize == nil {
		var ret int32
		return ret
	}
	return *o.ProbePoolSize
}

// GetProbePoolSizeOk returns a tuple with the ProbePoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetProbePoolSizeOk() (*int32, bool) {
	if o == nil || o.ProbePoolSize == nil {
		return nil, false
	}
	return o.ProbePoolSize, true
}

// HasProbePoolSize returns a boolean if a field has been set.
func (o *SLASettings) HasProbePoolSize() bool {
	return o != nil && o.ProbePoolSize != nil
}

// SetProbePoolSize gets a reference to the given int32 and assigns it to the ProbePoolSize field.
func (o *SLASettings) SetProbePoolSize(v int32) {
	o.ProbePoolSize = &v
}

// GetCollectorBatchSize returns the CollectorBatchSize field value if set, zero value otherwise.
func (o *SLASettings) GetCollectorBatchSize() int32 {
	if o == nil || o.CollectorBatchSize == nil {
		var ret int32
		return ret
	}
	return *o.CollectorBatchSize
}

// GetCollectorBatchSizeOk returns a tuple with the CollectorBatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetCollectorBatchSizeOk() (*int32, bool) {
	if o == nil || o.CollectorBatchSize == nil {
		return nil, false
	}
	return o.CollectorBatchSize, true
}

// HasCollectorBatchSize returns a boolean if a field has been set.
func (o *SLASettings) HasCollectorBatchSize() bool {
	return o != nil && o.CollectorBatchSize != nil
}

// SetCollectorBatchSize gets a reference to the given int32 and assigns it to the CollectorBatchSize field.
func (o *SLASettings) SetCollectorBatchSize(v int32) {
	o.CollectorBatchSize = &v
}

// GetCollectorMaxWaitTime returns the CollectorMaxWaitTime field value if set, zero value otherwise.
func (o *SLASettings) GetCollectorMaxWaitTime() string {
	if o == nil || o.CollectorMaxWaitTime == nil {
		var ret string
		return ret
	}
	return *o.CollectorMaxWaitTime
}

// GetCollectorMaxWaitTimeOk returns a tuple with the CollectorMaxWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetCollectorMaxWaitTimeOk() (*string, bool) {
	if o == nil || o.CollectorMaxWaitTime == nil {
		return nil, false
	}
	return o.CollectorMaxWaitTime, true
}

// HasCollectorMaxWaitTime returns a boolean if a field has been set.
func (o *SLASettings) HasCollectorMaxWaitTime() bool {
	return o != nil && o.CollectorMaxWaitTime != nil
}

// SetCollectorMaxWaitTime gets a reference to the given string and assigns it to the CollectorMaxWaitTime field.
func (o *SLASettings) SetCollectorMaxWaitTime(v string) {
	o.CollectorMaxWaitTime = &v
}

// GetSlaThreshold returns the SlaThreshold field value if set, zero value otherwise.
func (o *SLASettings) GetSlaThreshold() float {
	if o == nil || o.SlaThreshold == nil {
		var ret float
		return ret
	}
	return *o.SlaThreshold
}

// GetSlaThresholdOk returns a tuple with the SlaThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLASettings) GetSlaThresholdOk() (*float, bool) {
	if o == nil || o.SlaThreshold == nil {
		return nil, false
	}
	return o.SlaThreshold, true
}

// HasSlaThreshold returns a boolean if a field has been set.
func (o *SLASettings) HasSlaThreshold() bool {
	return o != nil && o.SlaThreshold != nil
}

// SetSlaThreshold gets a reference to the given float and assigns it to the SlaThreshold field.
func (o *SLASettings) SetSlaThreshold(v float) {
	o.SlaThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLASettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.ProbeInterval != nil {
		toSerialize["probeInterval"] = o.ProbeInterval
	}
	if o.ProbeTimeout != nil {
		toSerialize["probeTimeout"] = o.ProbeTimeout
	}
	if o.EnableProbeRetry != nil {
		toSerialize["enableProbeRetry"] = o.EnableProbeRetry
	}
	if o.ProbeRetryCount != nil {
		toSerialize["probeRetryCount"] = o.ProbeRetryCount
	}
	if o.ProbeRetryInterval != nil {
		toSerialize["probeRetryInterval"] = o.ProbeRetryInterval
	}
	if o.ProbeEngines != nil {
		toSerialize["probeEngines"] = o.ProbeEngines
	}
	if o.ProbePoolSize != nil {
		toSerialize["probePoolSize"] = o.ProbePoolSize
	}
	if o.CollectorBatchSize != nil {
		toSerialize["collectorBatchSize"] = o.CollectorBatchSize
	}
	if o.CollectorMaxWaitTime != nil {
		toSerialize["collectorMaxWaitTime"] = o.CollectorMaxWaitTime
	}
	if o.SlaThreshold != nil {
		toSerialize["slaThreshold"] = o.SlaThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLASettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentName      *string  `json:"environmentName,omitempty"`
		ProbeInterval        *string  `json:"probeInterval,omitempty"`
		ProbeTimeout         *string  `json:"probeTimeout,omitempty"`
		EnableProbeRetry     *bool    `json:"enableProbeRetry,omitempty"`
		ProbeRetryCount      *int32   `json:"probeRetryCount,omitempty"`
		ProbeRetryInterval   *string  `json:"probeRetryInterval,omitempty"`
		ProbeEngines         []string `json:"probeEngines,omitempty"`
		ProbePoolSize        *int32   `json:"probePoolSize,omitempty"`
		CollectorBatchSize   *int32   `json:"collectorBatchSize,omitempty"`
		CollectorMaxWaitTime *string  `json:"collectorMaxWaitTime,omitempty"`
		SlaThreshold         *float   `json:"slaThreshold,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environmentName", "probeInterval", "probeTimeout", "enableProbeRetry", "probeRetryCount", "probeRetryInterval", "probeEngines", "probePoolSize", "collectorBatchSize", "collectorMaxWaitTime", "slaThreshold"})
	} else {
		return err
	}
	o.EnvironmentName = all.EnvironmentName
	o.ProbeInterval = all.ProbeInterval
	o.ProbeTimeout = all.ProbeTimeout
	o.EnableProbeRetry = all.EnableProbeRetry
	o.ProbeRetryCount = all.ProbeRetryCount
	o.ProbeRetryInterval = all.ProbeRetryInterval
	o.ProbeEngines = all.ProbeEngines
	o.ProbePoolSize = all.ProbePoolSize
	o.CollectorBatchSize = all.CollectorBatchSize
	o.CollectorMaxWaitTime = all.CollectorMaxWaitTime
	o.SlaThreshold = all.SlaThreshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
