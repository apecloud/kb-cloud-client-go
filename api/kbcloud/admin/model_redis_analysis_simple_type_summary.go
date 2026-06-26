// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisAnalysisSimpleTypeSummary struct {
	Type  *string  `json:"type,omitempty"`
	Total *int64   `json:"total,omitempty"`
	Ratio *float64 `json:"ratio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisAnalysisSimpleTypeSummary instantiates a new RedisAnalysisSimpleTypeSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisAnalysisSimpleTypeSummary() *RedisAnalysisSimpleTypeSummary {
	this := RedisAnalysisSimpleTypeSummary{}
	return &this
}

// NewRedisAnalysisSimpleTypeSummaryWithDefaults instantiates a new RedisAnalysisSimpleTypeSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisAnalysisSimpleTypeSummaryWithDefaults() *RedisAnalysisSimpleTypeSummary {
	this := RedisAnalysisSimpleTypeSummary{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RedisAnalysisSimpleTypeSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisSimpleTypeSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RedisAnalysisSimpleTypeSummary) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RedisAnalysisSimpleTypeSummary) SetType(v string) {
	o.Type = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RedisAnalysisSimpleTypeSummary) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisSimpleTypeSummary) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RedisAnalysisSimpleTypeSummary) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RedisAnalysisSimpleTypeSummary) SetTotal(v int64) {
	o.Total = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *RedisAnalysisSimpleTypeSummary) GetRatio() float64 {
	if o == nil || o.Ratio == nil {
		var ret float64
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisAnalysisSimpleTypeSummary) GetRatioOk() (*float64, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *RedisAnalysisSimpleTypeSummary) HasRatio() bool {
	return o != nil && o.Ratio != nil
}

// SetRatio gets a reference to the given float64 and assigns it to the Ratio field.
func (o *RedisAnalysisSimpleTypeSummary) SetRatio(v float64) {
	o.Ratio = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisAnalysisSimpleTypeSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Ratio != nil {
		toSerialize["ratio"] = o.Ratio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisAnalysisSimpleTypeSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *string  `json:"type,omitempty"`
		Total *int64   `json:"total,omitempty"`
		Ratio *float64 `json:"ratio,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "total", "ratio"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Total = all.Total
	o.Ratio = all.Ratio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
