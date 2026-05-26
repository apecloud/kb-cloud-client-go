// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentVisualization struct {
	VisualizationId string                   `json:"visualizationId"`
	Title           string                   `json:"title"`
	Type            AiAgentVisualizationType `json:"type"`
	// References aiAgentMetricSeries.metricSeriesId. Clients hydrate it through aiAgentArtifactBatch.metricSeries.
	MetricSeriesRef *string  `json:"metricSeriesRef,omitempty"`
	XKey            *string  `json:"xKey,omitempty"`
	YKeys           []string `json:"yKeys,omitempty"`
	DisplayUnit     *string  `json:"displayUnit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentVisualization instantiates a new AiAgentVisualization object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentVisualization(visualizationId string, title string, typeVar AiAgentVisualizationType) *AiAgentVisualization {
	this := AiAgentVisualization{}
	this.VisualizationId = visualizationId
	this.Title = title
	this.Type = typeVar
	return &this
}

// NewAiAgentVisualizationWithDefaults instantiates a new AiAgentVisualization object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentVisualizationWithDefaults() *AiAgentVisualization {
	this := AiAgentVisualization{}
	return &this
}

// GetVisualizationId returns the VisualizationId field value.
func (o *AiAgentVisualization) GetVisualizationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VisualizationId
}

// GetVisualizationIdOk returns a tuple with the VisualizationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetVisualizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationId, true
}

// SetVisualizationId sets field value.
func (o *AiAgentVisualization) SetVisualizationId(v string) {
	o.VisualizationId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentVisualization) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentVisualization) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *AiAgentVisualization) GetType() AiAgentVisualizationType {
	if o == nil {
		var ret AiAgentVisualizationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetTypeOk() (*AiAgentVisualizationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentVisualization) SetType(v AiAgentVisualizationType) {
	o.Type = v
}

// GetMetricSeriesRef returns the MetricSeriesRef field value if set, zero value otherwise.
func (o *AiAgentVisualization) GetMetricSeriesRef() string {
	if o == nil || o.MetricSeriesRef == nil {
		var ret string
		return ret
	}
	return *o.MetricSeriesRef
}

// GetMetricSeriesRefOk returns a tuple with the MetricSeriesRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetMetricSeriesRefOk() (*string, bool) {
	if o == nil || o.MetricSeriesRef == nil {
		return nil, false
	}
	return o.MetricSeriesRef, true
}

// HasMetricSeriesRef returns a boolean if a field has been set.
func (o *AiAgentVisualization) HasMetricSeriesRef() bool {
	return o != nil && o.MetricSeriesRef != nil
}

// SetMetricSeriesRef gets a reference to the given string and assigns it to the MetricSeriesRef field.
func (o *AiAgentVisualization) SetMetricSeriesRef(v string) {
	o.MetricSeriesRef = &v
}

// GetXKey returns the XKey field value if set, zero value otherwise.
func (o *AiAgentVisualization) GetXKey() string {
	if o == nil || o.XKey == nil {
		var ret string
		return ret
	}
	return *o.XKey
}

// GetXKeyOk returns a tuple with the XKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetXKeyOk() (*string, bool) {
	if o == nil || o.XKey == nil {
		return nil, false
	}
	return o.XKey, true
}

// HasXKey returns a boolean if a field has been set.
func (o *AiAgentVisualization) HasXKey() bool {
	return o != nil && o.XKey != nil
}

// SetXKey gets a reference to the given string and assigns it to the XKey field.
func (o *AiAgentVisualization) SetXKey(v string) {
	o.XKey = &v
}

// GetYKeys returns the YKeys field value if set, zero value otherwise.
func (o *AiAgentVisualization) GetYKeys() []string {
	if o == nil || o.YKeys == nil {
		var ret []string
		return ret
	}
	return o.YKeys
}

// GetYKeysOk returns a tuple with the YKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetYKeysOk() (*[]string, bool) {
	if o == nil || o.YKeys == nil {
		return nil, false
	}
	return &o.YKeys, true
}

// HasYKeys returns a boolean if a field has been set.
func (o *AiAgentVisualization) HasYKeys() bool {
	return o != nil && o.YKeys != nil
}

// SetYKeys gets a reference to the given []string and assigns it to the YKeys field.
func (o *AiAgentVisualization) SetYKeys(v []string) {
	o.YKeys = v
}

// GetDisplayUnit returns the DisplayUnit field value if set, zero value otherwise.
func (o *AiAgentVisualization) GetDisplayUnit() string {
	if o == nil || o.DisplayUnit == nil {
		var ret string
		return ret
	}
	return *o.DisplayUnit
}

// GetDisplayUnitOk returns a tuple with the DisplayUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentVisualization) GetDisplayUnitOk() (*string, bool) {
	if o == nil || o.DisplayUnit == nil {
		return nil, false
	}
	return o.DisplayUnit, true
}

// HasDisplayUnit returns a boolean if a field has been set.
func (o *AiAgentVisualization) HasDisplayUnit() bool {
	return o != nil && o.DisplayUnit != nil
}

// SetDisplayUnit gets a reference to the given string and assigns it to the DisplayUnit field.
func (o *AiAgentVisualization) SetDisplayUnit(v string) {
	o.DisplayUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["visualizationId"] = o.VisualizationId
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	if o.MetricSeriesRef != nil {
		toSerialize["metricSeriesRef"] = o.MetricSeriesRef
	}
	if o.XKey != nil {
		toSerialize["xKey"] = o.XKey
	}
	if o.YKeys != nil {
		toSerialize["yKeys"] = o.YKeys
	}
	if o.DisplayUnit != nil {
		toSerialize["displayUnit"] = o.DisplayUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentVisualization) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		VisualizationId *string                   `json:"visualizationId"`
		Title           *string                   `json:"title"`
		Type            *AiAgentVisualizationType `json:"type"`
		MetricSeriesRef *string                   `json:"metricSeriesRef,omitempty"`
		XKey            *string                   `json:"xKey,omitempty"`
		YKeys           []string                  `json:"yKeys,omitempty"`
		DisplayUnit     *string                   `json:"displayUnit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.VisualizationId == nil {
		return fmt.Errorf("required field visualizationId missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"visualizationId", "title", "type", "metricSeriesRef", "xKey", "yKeys", "displayUnit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.VisualizationId = *all.VisualizationId
	o.Title = *all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.MetricSeriesRef = all.MetricSeriesRef
	o.XKey = all.XKey
	o.YKeys = all.YKeys
	o.DisplayUnit = all.DisplayUnit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
