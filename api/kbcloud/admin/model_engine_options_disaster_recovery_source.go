// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineOptionsDisasterRecoverySource struct {
	MetricSource *MetricsOptionQuery `json:"metricSource,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionsDisasterRecoverySource instantiates a new EngineOptionsDisasterRecoverySource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionsDisasterRecoverySource() *EngineOptionsDisasterRecoverySource {
	this := EngineOptionsDisasterRecoverySource{}
	return &this
}

// NewEngineOptionsDisasterRecoverySourceWithDefaults instantiates a new EngineOptionsDisasterRecoverySource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionsDisasterRecoverySourceWithDefaults() *EngineOptionsDisasterRecoverySource {
	this := EngineOptionsDisasterRecoverySource{}
	return &this
}

// GetMetricSource returns the MetricSource field value if set, zero value otherwise.
func (o *EngineOptionsDisasterRecoverySource) GetMetricSource() MetricsOptionQuery {
	if o == nil || o.MetricSource == nil {
		var ret MetricsOptionQuery
		return ret
	}
	return *o.MetricSource
}

// GetMetricSourceOk returns a tuple with the MetricSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionsDisasterRecoverySource) GetMetricSourceOk() (*MetricsOptionQuery, bool) {
	if o == nil || o.MetricSource == nil {
		return nil, false
	}
	return o.MetricSource, true
}

// HasMetricSource returns a boolean if a field has been set.
func (o *EngineOptionsDisasterRecoverySource) HasMetricSource() bool {
	return o != nil && o.MetricSource != nil
}

// SetMetricSource gets a reference to the given MetricsOptionQuery and assigns it to the MetricSource field.
func (o *EngineOptionsDisasterRecoverySource) SetMetricSource(v MetricsOptionQuery) {
	o.MetricSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionsDisasterRecoverySource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MetricSource != nil {
		toSerialize["metricSource"] = o.MetricSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionsDisasterRecoverySource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricSource *MetricsOptionQuery `json:"metricSource,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"metricSource"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.MetricSource != nil && all.MetricSource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricSource = all.MetricSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
